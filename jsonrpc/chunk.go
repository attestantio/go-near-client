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
	"github.com/attestantio/go-near-client/spec"
)

// Chunk returns the chunk.
func (s *Service) Chunk(ctx context.Context,
	opts *api.ChunkOpts,
) (
	*api.Response[*spec.Chunk],
	error,
) {
	if err := s.assertIsSynced(ctx); err != nil {
		return nil, err
	}

	if opts == nil {
		return nil, client.ErrNoOptions
	}

	// Chunk ID should not be provided with block ID or shard ID.
	if opts.ChunkID != "" && (opts.BlockID != 0 || opts.ShardID != nil) {
		return nil, client.ErrInvalidOptions
	}

	var args map[string]any

	switch {
	case opts.ChunkID != "":
		args = map[string]any{
			"chunk_id": opts.ChunkID,
		}
	case opts.BlockID != 0 && opts.ShardID != nil:
		// We've already checked that ChunkID is not provided.
		args = map[string]any{
			"block_id": opts.BlockID,
			"shard_id": opts.ShardID,
		}
	default:
		return nil, client.ErrInvalidOptions
	}

	chunk := &spec.Chunk{}
	err := s.client.CallFor(chunk, "chunk", args)
	if err != nil {
		return nil, parseJSONRPCError(err)
	}

	return &api.Response[*spec.Chunk]{
		Data:     chunk,
		Metadata: map[string]any{},
	}, nil
}
