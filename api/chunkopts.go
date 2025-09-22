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

package api

// ChunkOpts are the options for chunk calls.
type ChunkOpts struct {
	Common CommonOpts

	// ChunkID gets a chunk by chunk ID.
	ChunkID string
	// BlockID gets a chunk by block ID (used with ShardID).
	BlockID uint64
	// ShardID gets a chunk by shard ID (used with BlockID).
	ShardID uint64
}
