package jsonrpc

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBlockAtTimestamp(t *testing.T) {
	// Skip if RPC_URL is not set
	rpcURL := os.Getenv("RPC_URL")
	if rpcURL == "" {
		t.Skip("RPC_URL environment variable not set")
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

	// Create client for mainnet
	params := &Parameters{
		Network: "mainnet",
		Address: rpcURL,
		Timeout: 30 * time.Second,
	}

	client, err := New(context.Background(), params)
	require.NoError(t, err)

	// Run tests for each test case
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			targetTime := time.Unix(0, tc.targetTimestamp) // Convert nanoseconds to time.Time

			// Test BlockAtTimestamp
			block, err := client.BlockAtTimestamp(context.Background(), targetTime)
			require.NoError(t, err)
			require.NotNil(t, block)

			// Verify the returned block
			assert.Equal(t, tc.expectedBlockHeight, block.Header.Height)
			assert.GreaterOrEqual(t, tc.targetTimestamp, block.Header.Timestamp)

			// Verify the block timestamp is at or before the target timestamp
			assert.LessOrEqual(t, block.Header.Timestamp, targetTime.UnixNano())

			// Additional verification: check that the next block (if it exists) would be after the target timestamp
			// This ensures we got the last block before the timestamp
			nextBlock, err := client.GetLatestBlock(context.Background())
			if err == nil && nextBlock.Header.Height > block.Header.Height {
				// Try to get the next block after our result
				nextBlockByID, err := client.(*service).BlockByID(context.Background(), tc.expectedBlockHeight+1)
				if err == nil {
					assert.Greater(t, nextBlockByID.Header.Timestamp, targetTime.UnixNano())
				}
			}

			t.Logf("Successfully found block at height %d with timestamp %d", block.Header.Height, block.Header.Timestamp)
		})
	}
}

func TestBlockAtTimestampEdgeCases(t *testing.T) {
	// Skip if RPC_URL is not set
	rpcURL := os.Getenv("RPC_URL")
	if rpcURL == "" {
		t.Skip("RPC_URL environment variable not set")
	}

	params := &Parameters{
		Network: "mainnet",
		Address: rpcURL,
		Timeout: 30 * time.Second,
	}

	client, err := New(context.Background(), params)
	require.NoError(t, err)

	// Test with a timestamp in the future (should return an error)
	futureTime := time.Now().Add(24 * time.Hour)
	block, err := client.BlockAtTimestamp(context.Background(), futureTime)
	require.Error(t, err, "timestamp is after latest block time")
	require.Nil(t, block)

	// Test with a timestamp very far in the past
	pastTime := time.Unix(0, 0) // Unix epoch
	block, err = client.BlockAtTimestamp(context.Background(), pastTime)

	// The error comes from the RPC node, invalid block height.
	require.Error(t, err)
	require.Nil(t, block)
}

func TestBlockAtTimestampIntegration(t *testing.T) {
	// Skip if RPC_URL is not set
	rpcURL := os.Getenv("RPC_URL")
	if rpcURL == "" {
		t.Skip("RPC_URL environment variable not set")
	}

	params := &Parameters{
		Network: "mainnet",
		Address: rpcURL,
		Timeout: 30 * time.Second,
	}

	client, err := New(context.Background(), params)
	require.NoError(t, err)

	// Get the latest block first
	latestBlock, err := client.GetLatestBlock(context.Background())
	require.NoError(t, err)
	require.NotNil(t, latestBlock)

	// Test with the latest block's timestamp
	latestTime := time.Unix(0, latestBlock.Header.Timestamp)
	block, err := client.BlockAtTimestamp(context.Background(), latestTime)
	require.NoError(t, err)
	require.NotNil(t, block)

	// Should return the same block
	assert.Equal(t, latestBlock.Header.Height, block.Header.Height)
	assert.Equal(t, latestBlock.Header.Timestamp, block.Header.Timestamp)
}
