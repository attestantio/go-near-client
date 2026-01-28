// Copyright © 2026 Attestant Limited.
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

func TestValidators(t *testing.T) {
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

	// Test validators with Latest option
	opts := &api.ValidatorsOpts{
		Latest: true,
	}

	response, err := client.Validators(context.Background(), opts)
	require.NoError(t, err, "Should successfully retrieve validators")
	require.NotNil(t, response, "Response should not be nil")
	require.NotNil(t, response.Data, "Response data should not be nil")

	validators := response.Data

	// Verify basic validators structure
	require.NotEmpty(t, validators.CurrentValidators, "Current validators should not be empty")
	assert.Greater(t, validators.EpochHeight, int64(0), "Epoch height should be positive")
	assert.Greater(t, validators.EpochStartHeight, int64(0), "Epoch start height should be positive")

	// Verify first current validator has required fields
	firstValidator := validators.CurrentValidators[0]
	assert.NotEmpty(t, firstValidator.AccountID, "First validator should have an account ID")
	assert.NotEmpty(t, firstValidator.PublicKey, "First validator should have a public key")

	t.Logf("✓ All validators checks passed successfully")
}
