package telemetry

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/armon/go-metrics"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetrichttp"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/metric/instrument"
	metricsdk "go.opentelemetry.io/otel/sdk/metric"
)

const (
	flushInterval = 100 * time.Millisecond
)

type OTLPSink struct {
	ctx context.Context

	meter         metric.Meter
	meterProvider *metricsdk.MeterProvider

	counters   map[string]*instrument.Float64Counter
	gauges     map[string]*instrument.Float64UpDownCounter
	histograms map[string]*instrument.Float64Histogram
}

func NewOTLPSink() (*OTLPSink, error) {
	ctx := context.Background()

	// TODO: Create a custom exporter (interface)
	// We would create our custom exporter that uses a client that can do HCP auth.
	exp, err := otlpmetrichttp.New(ctx, otlpmetrichttp.WithInsecure(), otlpmetrichttp.WithEndpoint("localhost:9090"), otlpmetrichttp.WithTemporalitySelector(metricsdk.DefaultTemporalitySelector))
	if err != nil {
		return nil, fmt.Errorf("failed to create exporter: %v", err)
	}

	// TODO: Create a custom reader (interface)
	// We would create a custom PeriodicReader that does the batching/queue logic we need within it.
	reader := metricsdk.NewPeriodicReader(exp, metricsdk.WithInterval(10*time.Second))

	// TODO: This is experimental currently? How do we feel about that.
	provider := metricsdk.NewMeterProvider(metricsdk.WithReader(reader))

	// A meter is what registers all the metrics that will be collected.
	meter := provider.Meter("github.com/consul/agent/hcp/telemetry", metric.WithInstrumentationAttributes(attribute.KeyValue{Key: "instance.id", Value: attribute.StringValue("test")}))

	s := &OTLPSink{
		ctx: ctx,

		meter:         meter,
		meterProvider: provider,

		// Currently lazy loading when pushing metrics.
		// Size of these bounded by the number of metrics (n)
		counters:   make(map[string]*instrument.Float64Counter, 0),
		gauges:     make(map[string]*instrument.Float64UpDownCounter, 0),
		histograms: make(map[string]*instrument.Float64Histogram, 0),
	}

	return s, nil
}

func (s *OTLPSink) Shutdown() {
	ctx := context.Background()
	s.meterProvider.Shutdown(ctx)
}

func (s *OTLPSink) SetGauge(key []string, val float32) {
	s.SetGaugeWithLabels(key, val, nil)
}

func (s *OTLPSink) SetGaugeWithLabels(key []string, val float32, labels []metrics.Label) {
	k := s.flattenKeyLabels(key, labels)
	counter, ok := s.gauges[k]
	if !ok {
		c, _ := s.meter.Float64UpDownCounter(k)
		s.gauges[k] = &c
		counter = &c
	}

	(*counter).Add(s.ctx, float64(val))
}

func (s *OTLPSink) EmitKey(key []string, val float32) {
	s.flattenKey(key)
}

func (s *OTLPSink) IncrCounter(key []string, val float32) {
	s.IncrCounterWithLabels(key, val, nil)
}

func (s *OTLPSink) IncrCounterWithLabels(key []string, val float32, labels []metrics.Label) {
	k := s.flattenKeyLabels(key, labels)
	counter, ok := s.counters[k]
	if !ok {
		c, _ := s.meter.Float64Counter(k)
		s.counters[k] = &c
		counter = &c
	}

	(*counter).Add(s.ctx, float64(val))
}

func (s *OTLPSink) AddSample(key []string, val float32) {
	s.AddSampleWithLabels(key, val, nil)
}

func (s *OTLPSink) AddSampleWithLabels(key []string, val float32, labels []metrics.Label) {
	k := s.flattenKeyLabels(key, labels)
	hist, ok := s.histograms[k]
	if !ok {
		c, _ := s.meter.Float64Histogram(k)
		s.histograms[k] = &c
		hist = &c
	}

	(*hist).Record(s.ctx, float64(val))
}

// Flattens the key for formatting, removes spaces
func (s *OTLPSink) flattenKey(parts []string) string {
	joined := strings.Join(parts, ".")
	return strings.Map(func(r rune) rune {
		switch r {
		case ':':
			fallthrough
		case ' ':
			return '_'
		default:
			return r
		}
	}, joined)
}

// Flattens the key along with labels for formatting, removes spaces
func (s *OTLPSink) flattenKeyLabels(parts []string, labels []metrics.Label) string {
	for _, label := range labels {
		parts = append(parts, label.Value)
	}
	return s.flattenKey(parts)
}
