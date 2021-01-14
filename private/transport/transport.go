// Copyright (C) 2020 Storj Labs, Inc.
// See LICENSE for copying information.

package transport

import (
	"context"

	"storj.io/common/rpc/rpcpool"
	"storj.io/uplink"
	"storj.io/uplink/internal/expose"
)

// SetConnectionPool configures connection pool on the passed in config. If
// argument pool is nil, it will clear the pool on the config.
func SetConnectionPool(ctx context.Context, config *uplink.Config, pool *rpcpool.Pool) error {
	expose.ConfigSetConnectionPool(config, pool)
	return nil
}
