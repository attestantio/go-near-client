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

package jsonrpc_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/attestantio/go-near-client/api"
	"github.com/attestantio/go-near-client/jsonrpc"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestChunk(t *testing.T) {
	// Skip if JSONRPC_ADDRESS is not set
	rpcURL := os.Getenv("JSONRPC_ADDRESS")
	if rpcURL == "" {
		t.Skip("JSONRPC_ADDRESS environment variable not set")
	}

	// Create JSONRPC client
	client, err := jsonrpc.New(context.Background(),
		jsonrpc.WithAddress(rpcURL),
		jsonrpc.WithTimeout(30*time.Second),
		jsonrpc.WithMonitor(nil),
		jsonrpc.WithLogLevel(zerolog.GlobalLevel()))
	require.NoError(t, err)

	// First, get a known block to extract chunk information from it
	blockResponse, err := client.Block(context.Background(), &api.BlockOpts{
		BlockID: 80712125,
	})
	require.NoError(t, err)
	require.NotNil(t, blockResponse)
	require.NotNil(t, blockResponse.Data)
	require.Greater(t, len(blockResponse.Data.Chunks), 0, "Block should have at least one chunk")

	// Use the first chunk from the known block
	firstChunk := blockResponse.Data.Chunks[0]
	knownChunkHash := firstChunk.ChunkHash
	knownBlockID := uint64(80712125)
	knownShardID := uint64(firstChunk.ShardID)
	t.Logf("Known chunk hash: %s", knownChunkHash)
	t.Logf("Known block ID: %d", knownBlockID)
	t.Logf("Known shard ID: %d", knownShardID)

	nonExistentBlockID := uint64(999999999999)
	nonExistentShardID := uint64(999)

	// Test cases with different ChunkOpts configurations
	testCases := []struct {
		name        string
		opts        *api.ChunkOpts
		expectError bool
		description string
	}{
		{
			name: "Chunk by chunk ID - known chunk",
			opts: &api.ChunkOpts{
				ChunkID: knownChunkHash,
			},
			expectError: false,
			description: "Should get chunk by known chunk ID from block 80712125",
		},
		{
			name: "Chunk by block ID and shard ID - known chunk",
			opts: &api.ChunkOpts{
				BlockID: knownBlockID,
				ShardID: &knownShardID,
			},
			expectError: false,
			description: "Should get chunk by known block ID and shard ID",
		},
		{
			name: "Chunk by invalid chunk ID",
			opts: &api.ChunkOpts{
				ChunkID: "InvalidChunkHashThatDoesNotExist123456789",
			},
			expectError: true,
			description: "Should fail for invalid chunk ID",
		},
		{
			name: "Chunk by non-existent block ID and shard ID",
			opts: &api.ChunkOpts{
				BlockID: nonExistentBlockID, // Very high, likely non-existent block
				ShardID: &knownShardID,
			},
			expectError: true,
			description: "Should fail for non-existent block ID",
		},
		{
			name: "Chunk by valid block ID and invalid shard ID",
			opts: &api.ChunkOpts{
				BlockID: knownBlockID,
				ShardID: &nonExistentShardID, // Very high shard ID that likely doesn't exist
			},
			expectError: true,
			description: "Should fail for invalid shard ID",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			response, err := client.Chunk(context.Background(), tc.opts)

			if tc.expectError {
				require.Error(t, err, tc.description)
				require.Nil(t, response)
				t.Logf("Correctly failed with error: %v", err)
				return
			}

			require.NoError(t, err, tc.description)
			require.NotNil(t, response)
			require.NotNil(t, response.Data)

			chunk := response.Data

			// Verify basic chunk structure
			assert.NotEmpty(t, chunk.Author, "Chunk should have an author")
			assert.NotEmpty(t, chunk.Header.ChunkHash, "Chunk should have a hash")
			assert.Greater(t, chunk.Header.HeightCreated, int64(0), "Chunk should have a positive height created")
			assert.Greater(t, chunk.Header.HeightIncluded, int64(0), "Chunk should have a positive height included")
			assert.GreaterOrEqual(t, chunk.Header.ShardID, 0, "Chunk should have a valid shard ID")

			// Test specific expectations based on the test case
			switch tc.name {
			case "Chunk by chunk ID - known chunk":
				assert.Equal(t, knownChunkHash, chunk.Header.ChunkHash, "Should return chunk with correct hash")
				assert.Equal(t, int64(80712125), chunk.Header.HeightCreated, "Should return chunk with correct height created")
				assert.Equal(t, int64(80712125), chunk.Header.HeightIncluded, "Should return chunk with correct height included")

			case "Chunk by block ID and shard ID - known chunk":
				assert.Equal(t, int64(80712125), chunk.Header.HeightCreated, "Should return chunk with correct height created")
				assert.Equal(t, int64(80712125), chunk.Header.HeightIncluded, "Should return chunk with correct height included")
				assert.Equal(t, int(knownShardID), chunk.Header.ShardID, "Should return chunk with correct shard ID")

			case "Chunk by block ID and shard ID - different shard":
				assert.Equal(t, int64(80712125), chunk.Header.HeightCreated, "Should return chunk with correct height created")
				assert.Equal(t, int64(80712125), chunk.Header.HeightIncluded, "Should return chunk with correct height included")
				assert.Equal(t, 1, chunk.Header.ShardID, "Should return chunk with correct shard ID")
			}

			// Verify response metadata
			assert.NotNil(t, response.Metadata, "Response should have metadata")

			// Verify chunk header structure
			assert.NotEmpty(t, chunk.Header.PrevBlockHash, "Chunk should have a previous block hash")
			assert.NotEmpty(t, chunk.Header.PrevStateRoot, "Chunk should have a previous state root")
			assert.GreaterOrEqual(t, chunk.Header.GasLimit, int64(0), "Chunk should have a valid gas limit")
			assert.GreaterOrEqual(t, chunk.Header.GasUsed, int64(0), "Chunk should have a valid gas used")

			// Verify chunk data structures
			assert.NotNil(t, chunk.Receipts, "Chunk should have receipts (even if empty)")
			assert.NotNil(t, chunk.Transactions, "Chunk should have transactions (even if empty)")

			t.Logf("Successfully retrieved chunk with hash %s at height %d (shard %d)",
				chunk.Header.ChunkHash, chunk.Header.HeightCreated, chunk.Header.ShardID)
		})
	}
}

func TestChunkInvalidOptions(t *testing.T) {
	// Skip if JSONRPC_ADDRESS is not set
	rpcURL := os.Getenv("JSONRPC_ADDRESS")
	if rpcURL == "" {
		t.Skip("JSONRPC_ADDRESS environment variable not set")
	}

	client, err := jsonrpc.New(context.Background(),
		jsonrpc.WithAddress(rpcURL),
		jsonrpc.WithTimeout(30*time.Second),
		jsonrpc.WithMonitor(nil),
		jsonrpc.WithLogLevel(zerolog.GlobalLevel()))
	require.NoError(t, err)

	testShardID := uint64(0)
	// Test invalid option combinations
	invalidOptionTests := []struct {
		name        string
		opts        *api.ChunkOpts
		description string
	}{
		{
			name:        "No options provided",
			opts:        nil,
			description: "Should fail when no options are provided",
		},
		{
			name:        "Empty options",
			opts:        &api.ChunkOpts{},
			description: "Should fail when no specific options are set",
		},
		{
			name: "Only block ID provided",
			opts: &api.ChunkOpts{
				BlockID: 80712125,
			},
			description: "Should fail when only block ID is provided without shard ID",
		},
		{
			name: "Only shard ID provided",
			opts: &api.ChunkOpts{
				ShardID: &testShardID,
			},
			description: "Should fail when only shard ID is provided without block ID",
		},
		{
			name: "Chunk ID with block ID",
			opts: &api.ChunkOpts{
				ChunkID: "8JoY27ADrA534GKWtiU7fkBFRHzgVa6cXJmVjThPzp8u",
				BlockID: 80712125,
			},
			description: "Should fail when chunk ID is provided with block ID",
		},
		{
			name: "Chunk ID with shard ID",
			opts: &api.ChunkOpts{
				ChunkID: "8JoY27ADrA534GKWtiU7fkBFRHzgVa6cXJmVjThPzp8u",
				ShardID: &testShardID,
			},
			description: "Should fail when chunk ID is provided with shard ID",
		},
		{
			name: "All options provided",
			opts: &api.ChunkOpts{
				ChunkID: "8JoY27ADrA534GKWtiU7fkBFRHzgVa6cXJmVjThPzp8u",
				BlockID: 80712125,
				ShardID: &testShardID,
			},
			description: "Should fail when all options are provided",
		},
	}

	for _, tc := range invalidOptionTests {
		t.Run(tc.name, func(t *testing.T) {
			response, err := client.Chunk(context.Background(), tc.opts)
			require.Error(t, err, tc.description)
			require.Nil(t, response)
			t.Logf("Correctly failed with error: %v", err)
		})
	}
}

func TestChunkWithTimeout(t *testing.T) {
	// Skip if JSONRPC_ADDRESS is not set
	rpcURL := os.Getenv("JSONRPC_ADDRESS")
	if rpcURL == "" {
		t.Skip("JSONRPC_ADDRESS environment variable not set")
	}

	client, err := jsonrpc.New(context.Background(),
		jsonrpc.WithAddress(rpcURL),
		jsonrpc.WithTimeout(30*time.Second),
		jsonrpc.WithMonitor(nil),
		jsonrpc.WithLogLevel(zerolog.GlobalLevel()))
	require.NoError(t, err)

	// First get a known block to extract chunk info
	blockResponse, err := client.Block(context.Background(), &api.BlockOpts{
		BlockID: 80712125,
	})
	require.NoError(t, err)
	require.Greater(t, len(blockResponse.Data.Chunks), 0, "Block should have chunks")

	firstChunk := blockResponse.Data.Chunks[0]

	// Test with custom timeout in options
	opts := &api.ChunkOpts{
		Common: api.CommonOpts{
			Timeout: 5 * time.Second,
		},
		ChunkID: firstChunk.ChunkHash,
	}

	response, err := client.Chunk(context.Background(), opts)
	require.NoError(t, err)
	require.NotNil(t, response)
	require.NotNil(t, response.Data)

	chunk := response.Data
	assert.NotEmpty(t, chunk.Author, "Chunk should have an author")
	assert.Equal(t, firstChunk.ChunkHash, chunk.Header.ChunkHash, "Should return correct chunk")

	t.Logf("Successfully retrieved chunk with custom timeout: %s", chunk.Header.ChunkHash)
}

func TestChunkFromRecentBlock(t *testing.T) {
	// Skip if JSONRPC_ADDRESS is not set
	rpcURL := os.Getenv("JSONRPC_ADDRESS")
	if rpcURL == "" {
		t.Skip("JSONRPC_ADDRESS environment variable not set")
	}

	client, err := jsonrpc.New(context.Background(),
		jsonrpc.WithAddress(rpcURL),
		jsonrpc.WithTimeout(30*time.Second),
		jsonrpc.WithMonitor(nil),
		jsonrpc.WithLogLevel(zerolog.GlobalLevel()))
	require.NoError(t, err)

	// Get the latest final block
	blockResponse, err := client.Block(context.Background(), &api.BlockOpts{
		Finality: "final",
	})
	require.NoError(t, err)
	require.NotNil(t, blockResponse)
	require.Greater(t, len(blockResponse.Data.Chunks), 0, "Block should have chunks")

	// Get chunk from the latest final block
	firstChunk := blockResponse.Data.Chunks[0]

	response, err := client.Chunk(context.Background(), &api.ChunkOpts{
		ChunkID: firstChunk.ChunkHash,
	})
	require.NoError(t, err)
	require.NotNil(t, response)
	require.NotNil(t, response.Data)

	chunk := response.Data

	// Verify the chunk matches what we expected from the block
	assert.Equal(t, firstChunk.ChunkHash, chunk.Header.ChunkHash, "Chunk hash should match")
	assert.Equal(t, firstChunk.ShardID, chunk.Header.ShardID, "Shard ID should match")
	assert.Equal(t, firstChunk.HeightCreated, chunk.Header.HeightCreated, "Height created should match")
	assert.Equal(t, firstChunk.HeightIncluded, chunk.Header.HeightIncluded, "Height included should match")

	// Verify chunk structure
	assert.NotEmpty(t, chunk.Author, "Chunk should have an author")
	assert.Greater(t, chunk.Header.HeightCreated, int64(0), "Chunk should have a positive height")
	assert.NotNil(t, chunk.Receipts, "Chunk should have receipts")
	assert.NotNil(t, chunk.Transactions, "Chunk should have transactions")

	t.Logf("Successfully retrieved recent chunk: hash=%s, height=%d, shard=%d",
		chunk.Header.ChunkHash, chunk.Header.HeightCreated, chunk.Header.ShardID)
}

func TestChunkCompareByIDvsBlockShard(t *testing.T) {
	// Skip if JSONRPC_ADDRESS is not set
	rpcURL := os.Getenv("JSONRPC_ADDRESS")
	if rpcURL == "" {
		t.Skip("JSONRPC_ADDRESS environment variable not set")
	}

	client, err := jsonrpc.New(context.Background(),
		jsonrpc.WithAddress(rpcURL),
		jsonrpc.WithTimeout(30*time.Second),
		jsonrpc.WithMonitor(nil),
		jsonrpc.WithLogLevel(zerolog.GlobalLevel()))
	require.NoError(t, err)

	// Get chunk by ID
	knownBlockID := uint64(80712125)
	knownShardID := uint64(0)

	// First, get chunk by block ID and shard ID
	chunkByBlockShard, err := client.Chunk(context.Background(), &api.ChunkOpts{
		BlockID: knownBlockID,
		ShardID: &knownShardID,
	})
	require.NoError(t, err)
	require.NotNil(t, chunkByBlockShard)

	// Then get the same chunk by its chunk ID
	chunkByID, err := client.Chunk(context.Background(), &api.ChunkOpts{
		ChunkID: chunkByBlockShard.Data.Header.ChunkHash,
	})
	require.NoError(t, err)
	require.NotNil(t, chunkByID)

	// Both methods should return the exact same chunk
	assert.Equal(t, chunkByBlockShard.Data.Header.ChunkHash, chunkByID.Data.Header.ChunkHash, "Chunk hashes should match")
	assert.Equal(t, chunkByBlockShard.Data.Author, chunkByID.Data.Author, "Authors should match")
	assert.Equal(t, chunkByBlockShard.Data.Header.HeightCreated, chunkByID.Data.Header.HeightCreated, "Heights should match")
	assert.Equal(t, chunkByBlockShard.Data.Header.ShardID, chunkByID.Data.Header.ShardID, "Shard IDs should match")
	assert.Equal(t, len(chunkByBlockShard.Data.Transactions), len(chunkByID.Data.Transactions), "Transaction counts should match")
	assert.Equal(t, len(chunkByBlockShard.Data.Receipts), len(chunkByID.Data.Receipts), "Receipt counts should match")

	t.Logf("Successfully verified both methods return the same chunk: %s", chunkByID.Data.Header.ChunkHash)
}
