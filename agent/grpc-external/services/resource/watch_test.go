package resource

import (
	context "context"
	"errors"
	"fmt"
	"io"
	"testing"
	"time"

	"github.com/hashicorp/consul/internal/resource"
	"github.com/hashicorp/consul/internal/storage/inmem"
	pbresource "github.com/hashicorp/consul/proto-public/pbresource"
	"github.com/hashicorp/consul/proto/private/prototest"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

func TestWatchList_TypeNotFound(t *testing.T) {
	server := NewServer(Config{registry: resource.NewRegistry()})
	client := testClient(t, server)

	stream, err := client.WatchList(context.Background(), &pbresource.WatchListRequest{
		Type: &pbresource.Type{
			Group:        "mesh",
			GroupVersion: "v1",
			Kind:         "service",
		},
		Tenancy: &pbresource.Tenancy{
			Partition: "default",
			Namespace: "default",
			PeerName:  "",
		},
		NamePrefix: "",
	})
	require.NoError(t, err)
	rspCh := handleResourceStream(t, stream)

	err = mustGetError(t, rspCh)
	require.Equal(t, codes.InvalidArgument.String(), status.Code(err).String())
	require.Contains(t, err.Error(), "resource type mesh/v1/service not registered")
}

func TestWatchList_Upsert(t *testing.T) {
	t1 := &pbresource.Type{Group: "mesh", GroupVersion: "v1", Kind: "service"}
	ten1 := &pbresource.Tenancy{
		Partition: "default",
		Namespace: "default",
		PeerName:  "local",
	}
	registry := resource.NewRegistry()
	registry.Register(resource.Registration{Type: t1})

	backend, err := inmem.NewBackend()
	go backend.Run(testContext(t))

	require.NoError(t, err)
	server := NewServer(Config{registry: registry, backend: backend})
	client := testClient(t, server)

	ctx := context.Background()

	// create resource r1 with is of type t1
	r1 := &pbresource.Resource{
		Id: &pbresource.ID{
			Uid:     "someUid",
			Name:    "someName",
			Type:    t1,
			Tenancy: ten1,
		},
		Version: "1",
	}
	r1, err = backend.WriteCAS(ctx, r1, "")
	require.NoError(t, err)

	// watch t1
	stream, err := client.WatchList(ctx, &pbresource.WatchListRequest{
		Type:       t1,
		Tenancy:    ten1,
		NamePrefix: "",
	})
	require.NoError(t, err)

	// verify upsert event received for creation
	rspCh := handleResourceStream(t, stream)
	rsp := mustGetResource(t, rspCh)
	require.Equal(t, pbresource.WatchListResponse_OPERATION_UPSERT, rsp.Operation)
	prototest.AssertDeepEqual(t, r1, rsp.Resource)

	// rsp2 := mustGetResource(t, rspCh)
	// fmt.Println(rsp2)

	//mutate and write and v2
	r2 := clone(r1)
	r2.Version = "2"
	r2, err = backend.WriteCAS(ctx, r2, r1.Version)
	require.NoError(t, err)

	// verify upsert event received for update
	rsp = mustGetResource(t, rspCh)
	// rsp2 := mustGetResource(t, rspCh)
	// fmt.Println(rsp2)
	require.Equal(t, pbresource.WatchListResponse_OPERATION_UPSERT, rsp.Operation)
	prototest.AssertDeepEqual(t, r2, rsp.Resource)

}

func TestWatchList_Loop(t *testing.T) {
	t1 := &pbresource.Type{Group: "mesh", GroupVersion: "v1", Kind: "service"}
	ten1 := &pbresource.Tenancy{
		Partition: "default",
		Namespace: "default",
		PeerName:  "local",
	}
	registry := resource.NewRegistry()
	registry.Register(resource.Registration{Type: t1})
	backend, err := inmem.NewBackend()
	require.NoError(t, err)
	server := NewServer(Config{registry: registry, backend: backend})
	client := testClient(t, server)

	ctx := context.Background()

	// watch t1
	stream, err := client.WatchList(ctx, &pbresource.WatchListRequest{
		Type:       t1,
		Tenancy:    ten1,
		NamePrefix: "",
	})
	require.NoError(t, err)

	for i := 0; i < 10; i++ {
		r := &pbresource.Resource{
			Id: &pbresource.ID{
				Uid:     fmt.Sprintf("someUid %d", i),
				Name:    fmt.Sprintf("someName %d", i),
				Type:    t1,
				Tenancy: ten1,
			},
			Version: "1",
		}
		r, err = backend.WriteCAS(ctx, r, "")
		require.NoError(t, err)
	}

	// verify upsert event received for creation
	rspCh := handleResourceStream(t, stream)

	for i := 0; i < 10; i++ {
		rsp := mustGetResource(t, rspCh)
		require.Equal(t, pbresource.WatchListResponse_OPERATION_UPSERT, rsp.Operation)
		//prototest.AssertDeepEqual(t, r1, rsp.Resource)
	}
}

func mustGetResource(t *testing.T, ch <-chan resourceOrError) *pbresource.WatchListResponse {
	t.Helper()

	select {
	case rsp := <-ch:
		require.NoError(t, rsp.err)
		return rsp.rsp
	case <-time.After(1 * time.Second):
		t.Fatal("timeout waiting for WatchListResponse")
		return nil
	}
}

func mustGetError(t *testing.T, ch <-chan resourceOrError) error {
	t.Helper()

	select {
	case rsp := <-ch:
		require.Error(t, rsp.err)
		return rsp.err
	case <-time.After(2 * time.Second):
		t.Fatal("timeout waiting for WatchListResponse")
		return nil
	}
}

func handleResourceStream(t *testing.T, stream pbresource.ResourceService_WatchListClient) <-chan resourceOrError {
	t.Helper()

	rspCh := make(chan resourceOrError)
	go func() {
		for {
			rsp, err := stream.Recv()
			if errors.Is(err, io.EOF) ||
				errors.Is(err, context.Canceled) ||
				errors.Is(err, context.DeadlineExceeded) {
				return
			}
			rspCh <- resourceOrError{
				rsp: rsp,
				err: err,
			}
		}
	}()
	return rspCh
}

type resourceOrError struct {
	rsp *pbresource.WatchListResponse
	err error
}

func clone[T proto.Message](v T) T { return proto.Clone(v).(T) }

func testContext(t *testing.T) context.Context {
	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)
	return ctx
}
