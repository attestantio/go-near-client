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

func TestAccountInfo(t *testing.T) {
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

	// Test account info for bitwise_1.near
	// Note: The current Account function doesn't support block-specific queries,
	// so this will return the current state rather than at block 143,828,695
	opts := &api.AccountOpts{
		AccountID:  "bitwise_1.near",
		ContractID: "bitwise_1.poolv1.near",
		Block: &api.BlockOpts{
			BlockID: 143828695,
		},
	}

	response, err := client.Account(context.Background(), opts)
	require.NoError(t, err, "Should successfully retrieve account info")
	require.NotNil(t, response, "Response should not be nil")
	require.NotNil(t, response.Data, "Response data should not be nil")

	account := response.Data
	require.NotEmpty(t, string(account.AccountID), "Account ID should not be empty")

	assert.Equal(t, "bitwise_1.near", string(account.AccountID), "Account ID should match")
	assert.Equal(t, "0", account.UnstakedBalance.Balance.String(), "Unstaked balance should be 0")
	assert.Equal(t, "3160288881354999999", account.StakedBalance.Balance.String(), "Staked balance should match expected value")
	assert.True(t, account.CanWithdraw, "Should be able to withdraw")

	t.Logf("✓ All account values validated successfully")
}
