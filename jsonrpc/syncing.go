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

package jsonrpc

import (
	"context"

	client "github.com/attestantio/go-near-client"
	"github.com/attestantio/go-near-client/api"
)

// Syncing obtains information about the sync state of the node.
func (s *Service) Syncing(_ context.Context,
	opts *api.SyncingOpts,
) (
	*api.Response[*api.SyncState],
	error,
) {
	// We do not run assertIsActive here as it calls this function, and so it would cause a loop.

	if opts == nil {
		return nil, client.ErrNoOptions
	}

	var data api.SyncState
	// TODO: Get correct call
	if err := s.client.CallFor(&data, "starknet_syncing"); err != nil {
		return nil, err
	}

	return &api.Response[*api.SyncState]{
		Data:     &data,
		Metadata: map[string]any{},
	}, nil
}
