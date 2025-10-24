// Copyright © 2025 Attestant Limited.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package client

import (
	"context"
	"time"

	"github.com/attestantio/go-near-client/api"
	"github.com/attestantio/go-near-client/spec"
)

// Service is the service providing a connection to a NEAR node.
type Service interface {
	// Name returns the name of the node implementation.
	Name() string

	// Address returns the address of the node.
	Address() string
}

// AccountProvider is the interface for providing account details.
type AccountProvider interface {
	// Account returns the account.
	Account(ctx context.Context,
		opts *api.AccountOpts,
	) (
		*api.Response[*spec.Account],
		error,
	)
}

// BlockProvider is the interface for providing block details.
type BlockProvider interface {
	// Block returns the block.
	Block(ctx context.Context,
		opts *api.BlockOpts,
	) (
		*api.Response[*spec.Block],
		error,
	)

	// BlockAtTimestamp returns the block at a given timestamp.
	BlockAtTimestamp(ctx context.Context,
		timestamp time.Time,
	) (
		*api.Response[*spec.Block],
		error,
	)
}

// StatusProvider is the interface for providing status.
type StatusProvider interface {
	// Status returns the status.
	Status(ctx context.Context,
		opts *api.StatusOpts,
	) (
		*api.Response[*spec.Status],
		error,
	)
}

// SyncProvider is the interface for providing sync status.
type SyncProvider interface {
	// Syncing returns the status.
	Syncing(ctx context.Context,
		opts *api.SyncingOpts,
	) (
		*api.Response[*spec.SyncInfo],
		error,
	)
}

// GenesisProvider is the interface for providing genesis block details.
type GenesisProvider interface {
	// Genesis returns the genesis block.
	GenesisBlockHeight(ctx context.Context) (int64, error)
}
