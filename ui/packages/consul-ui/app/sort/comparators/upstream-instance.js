/**
 * Copyright (c) HashiCorp, Inc.
 * SPDX-License-Identifier: MPL-2.0
 */

export default ({ properties }) =>
  (key = 'DestinationName:asc') => {
    return properties(['DestinationName'])(key);
  };
