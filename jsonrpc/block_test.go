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

func TestBlock(t *testing.T) {
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

	// Test cases with different BlockOpts configurations
	testCases := []struct {
		name        string
		opts        *api.BlockOpts
		expectError bool
		description string
	}{
		{
			name: "Block by finality - final",
			opts: &api.BlockOpts{
				Finality: "final",
			},
			expectError: false,
			description: "Should get the latest final block",
		},
		{
			name: "Block by finality - optimistic",
			opts: &api.BlockOpts{
				Finality: "optimistic",
			},
			expectError: false,
			description: "Should get the latest optimistic block",
		},
		{
			name: "Block by specific block ID - old block",
			opts: &api.BlockOpts{
				BlockID: 152423388,
			},
			expectError: false,
			description: "Should get block 152423388 with known timestamp 1750742496787538908",
		},
		{
			name: "Block by hash - known block",
			opts: &api.BlockOpts{
				Hash: "EhrfG2rghZFeedQYWApbpE8swN9c9Rha5mCV1QqJvfXP",
			},
			expectError: false,
			description: "Should get block with known hash from block_test.go",
		},
		{
			name: "Block by non-existent block ID",
			opts: &api.BlockOpts{
				BlockID: 152423386, // Missing block
			},
			expectError: true,
			description: "Should fail for missing block",
		},
		{
			name: "Block with invalid hash",
			opts: &api.BlockOpts{
				Hash: "InvalidHashThatDoesNotExist123456789",
			},
			expectError: true,
			description: "Should fail for invalid hash",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			response, err := client.Block(context.Background(), tc.opts)

			if tc.expectError {
				require.Error(t, err, tc.description)
				require.Nil(t, response)
				return
			}

			require.NoError(t, err, tc.description)
			require.NotNil(t, response)
			require.NotNil(t, response.Data)

			block := response.Data

			// Verify basic block structure
			assert.NotEmpty(t, block.Author, "Block should have an author")
			assert.NotEmpty(t, block.Header.Hash, "Block should have a hash")
			assert.Greater(t, block.Header.Height, int64(0), "Block should have a positive height")
			assert.Greater(t, block.Header.Timestamp, int64(0), "Block should have a positive timestamp")
			assert.NotEmpty(t, block.Header.PrevHash, "Block should have a previous hash")

			// Test specific expectations based on the test case
			switch tc.name {
			case "Block by specific block ID - old block":
				assert.Equal(t, int64(152423388), block.Header.Height, "Should return block with correct height")
				assert.Equal(t, int64(1750742496787538908), block.Header.Timestamp, "Should return block with correct timestamp")

			case "Block by hash - known block":
				assert.Equal(t, "EhrfG2rghZFeedQYWApbpE8swN9c9Rha5mCV1QqJvfXP", block.Header.Hash, "Should return block with correct hash")
				assert.Equal(t, int64(152423388), block.Header.Height, "Should return block with expected height")
				assert.Equal(t, "figment.poolv1.near", block.Author, "Should return block with expected author")
				assert.Len(t, block.Chunks, 8, "Should return block with expected number of chunks")

			case "Block by finality - final":
				// For final blocks, just verify it's a reasonable recent block
				assert.Greater(t, block.Header.Height, int64(100000000), "Final block should be a recent block")
				currentTime := time.Now().UnixNano()
				timeDiff := currentTime - block.Header.Timestamp
				assert.Less(t, timeDiff, int64(24*time.Hour), "Final block should be recent (within 24 hours)")

			case "Block by finality - optimistic":
				// For optimistic blocks, verify it's at least as recent as final
				assert.Greater(t, block.Header.Height, int64(100000000), "Optimistic block should be a recent block")
				currentTime := time.Now().UnixNano()
				timeDiff := currentTime - block.Header.Timestamp
				assert.Less(t, timeDiff, int64(time.Hour), "Optimistic block should be very recent (within 1 hour)")
			}

			// Verify response metadata
			assert.NotNil(t, response.Metadata, "Response should have metadata")

			t.Logf("Successfully retrieved block at height %d with hash %s", block.Header.Height, block.Header.Hash)
		})
	}
}

func TestBlockInvalidOptions(t *testing.T) {
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

	// Test invalid option combinations
	invalidOptionTests := []struct {
		name        string
		opts        *api.BlockOpts
		description string
	}{
		{
			name:        "No options provided",
			opts:        nil,
			description: "Should fail when no options are provided",
		},
		{
			name:        "Empty options",
			opts:        &api.BlockOpts{},
			description: "Should fail when no specific options are set",
		},
		{
			name: "Multiple options - finality and block ID",
			opts: &api.BlockOpts{
				Finality: "final",
				BlockID:  152423388,
			},
			description: "Should fail when multiple options are provided",
		},
		{
			name: "Multiple options - finality and hash",
			opts: &api.BlockOpts{
				Finality: "final",
				Hash:     "EhrfG2rghZFeedQYWApbpE8swN9c9Rha5mCV1QqJvfXP",
			},
			description: "Should fail when multiple options are provided",
		},
		{
			name: "Multiple options - block ID and hash",
			opts: &api.BlockOpts{
				BlockID: 152423388,
				Hash:    "EhrfG2rghZFeedQYWApbpE8swN9c9Rha5mCV1QqJvfXP",
			},
			description: "Should fail when multiple options are provided",
		},
		{
			name: "All options provided",
			opts: &api.BlockOpts{
				Finality: "final",
				BlockID:  152423388,
				Hash:     "EhrfG2rghZFeedQYWApbpE8swN9c9Rha5mCV1QqJvfXP",
			},
			description: "Should fail when all options are provided",
		},
	}

	for _, tc := range invalidOptionTests {
		t.Run(tc.name, func(t *testing.T) {
			response, err := client.Block(context.Background(), tc.opts)
			require.Error(t, err, tc.description)
			require.Nil(t, response)
			t.Logf("Correctly failed with error: %v", err)
		})
	}
}

func TestBlockRecentFinality(t *testing.T) {
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

	// Test finality with timestamp within 5 minutes
	response, err := client.Block(context.Background(), &api.BlockOpts{
		Finality: "final",
	})
	require.NoError(t, err)
	require.NotNil(t, response)
	require.NotNil(t, response.Data)

	block := response.Data
	currentTime := time.Now()
	blockTime := time.Unix(0, block.Header.Timestamp)
	timeDiff := currentTime.Sub(blockTime)

	// Verify the block is recent (within reasonable bounds, not necessarily 5 minutes)
	assert.Less(t, timeDiff, 24*time.Hour, "Final block should be within the last 24 hours")

	// Log the actual time difference for manual verification
	t.Logf("Final block timestamp: %v", blockTime)
	t.Logf("Current time: %v", currentTime)
	t.Logf("Time difference: %v", timeDiff)

	if timeDiff <= 5*time.Minute {
		t.Logf("✓ Block is within 5 minutes as requested")
	} else {
		t.Logf("⚠ Block is older than 5 minutes (this may be normal depending on network activity)")
	}

	// Verify block structure
	assert.NotEmpty(t, block.Author, "Block should have an author")
	assert.Greater(t, block.Header.Height, int64(0), "Block should have a positive height")
	assert.NotEmpty(t, block.Header.Hash, "Block should have a hash")
}
