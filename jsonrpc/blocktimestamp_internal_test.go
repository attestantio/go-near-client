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
	"os"
	"testing"
	"time"

	"github.com/attestantio/go-near-client/api"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBlockAtTimestamp(t *testing.T) {
	// Skip if JSONRPC_ADDRESS is not set
	rpcURL := os.Getenv("JSONRPC_ADDRESS")
	if rpcURL == "" {
		t.Skip("JSONRPC_ADDRESS environment variable not set")
	}

	// Test cases with expected block heights and timestamps
	testCases := []struct {
		name                string
		expectedBlockHeight int64
		targetTimestamp     int64
	}{
		{
			name:                "Block 152423388",
			expectedBlockHeight: 152423388,
			targetTimestamp:     1750742496787538908,
		},
		{
			name:                "Block 100000000",
			expectedBlockHeight: 100000000,
			targetTimestamp:     1693394692463909482,
		},
		{
			name:                "Missed block at timestamp",
			expectedBlockHeight: 153107931,
			targetTimestamp:     1751173200000000000,
		},
	}

	// Create JSONRPC client.
	client, err := New(context.Background(),
		WithAddress(rpcURL),
		WithTimeout(30*time.Second),
		WithMonitor(nil),
		WithLogLevel(zerolog.GlobalLevel()))
	require.NoError(t, err)

	// Run tests for each test case.
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			targetTime := time.Unix(0, tc.targetTimestamp) // Convert nanoseconds to time.Time.

			// Test BlockAtTimestamp
			blockResponse, err := client.BlockAtTimestamp(context.Background(), targetTime)
			block := blockResponse.Data
			require.NoError(t, err)
			require.NotNil(t, block)

			// Verify the returned block
			assert.Equal(t, tc.expectedBlockHeight, block.Header.Height)
			assert.GreaterOrEqual(t, tc.targetTimestamp, block.Header.Timestamp)

			// Verify the block timestamp is at or before the target timestamp
			assert.LessOrEqual(t, block.Header.Timestamp, targetTime.UnixNano())

			// Additional verification: check that the next block (if it exists) would be after the target timestamp
			// This ensures we got the last block before the timestamp
			// Try to get the next block after our result
			nextBlock, err := client.Block(context.Background(), &api.BlockOpts{BlockID: tc.expectedBlockHeight + 1})
			if err == nil {
				assert.Greater(t, nextBlock.Data.Header.Timestamp, targetTime.UnixNano())
			}

			t.Logf("Successfully found block at height %d with timestamp %d", block.Header.Height, block.Header.Timestamp)
		})
	}
}

func TestBlockAtTimestampEdgeCases(t *testing.T) {
	// Skip if JSONRPC_ADDRESS is not set
	rpcURL := os.Getenv("JSONRPC_ADDRESS")
	if rpcURL == "" {
		t.Skip("JSONRPC_ADDRESS environment variable not set")
	}

	client, err := New(context.Background(),
		WithAddress(rpcURL),
		WithTimeout(30*time.Second),
		WithMonitor(nil),
		WithLogLevel(zerolog.GlobalLevel()))
	require.NoError(t, err)

	// Test with a timestamp in the future (should return an error)
	futureTime := time.Now().Add(24 * time.Hour)
	blockResponse, err := client.BlockAtTimestamp(context.Background(), futureTime)
	require.Error(t, err, "timestamp is after latest block time")
	require.Nil(t, blockResponse)

	// Test with a timestamp very far in the past
	pastTime := time.Unix(0, 0) // Unix epoch
	blockResponse, err = client.BlockAtTimestamp(context.Background(), pastTime)

	// The error comes from the RPC node, invalid block height.
	require.Error(t, err)
	require.Nil(t, blockResponse)
}

func TestBlockAtTimestampIntegration(t *testing.T) {
	// Skip if JSONRPC_ADDRESS is not set
	rpcURL := os.Getenv("JSONRPC_ADDRESS")
	if rpcURL == "" {
		t.Skip("JSONRPC_ADDRESS environment variable not set")
	}

	client, err := New(context.Background(),
		WithAddress(rpcURL),
		WithTimeout(30*time.Second),
		WithMonitor(nil),
		WithLogLevel(zerolog.GlobalLevel()))
	require.NoError(t, err)

	// Get the latest block first
	latestBlock, err := client.Block(context.Background(), &api.BlockOpts{Finality: "final"})
	require.NoError(t, err)
	require.NotNil(t, latestBlock)

	// Test with the latest block's timestamp
	latestTime := time.Unix(0, latestBlock.Data.Header.Timestamp)
	blockResponse, err := client.BlockAtTimestamp(context.Background(), latestTime)
	block := blockResponse.Data
	require.NoError(t, err)
	require.NotNil(t, block)

	// Should return the same block
	assert.Equal(t, latestBlock.Data.Header.Height, block.Header.Height)
	assert.Equal(t, latestBlock.Data.Header.Timestamp, block.Header.Timestamp)
}
