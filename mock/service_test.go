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

package mock_test

import (
	"context"
	"encoding/json"
	"math/big"
	"testing"
	"time"

	"github.com/attestantio/go-near-client/api"
	"github.com/attestantio/go-near-client/mock"
	"github.com/attestantio/go-near-client/spec"
	"github.com/attestantio/go-near-client/types"
	"github.com/stretchr/testify/require"
)

func TestMockService(t *testing.T) {
	s, err := mock.New()
	require.NoError(t, err)
	require.NotNil(t, s)

	// Test basic service methods
	require.Equal(t, "mock", s.Name())
	require.Equal(t, "mock", s.Address())
}

func TestMockBlock(t *testing.T) {
	ctx := context.Background()

	s, err := mock.New()
	require.NoError(t, err)

	now := time.Now()
	block := &spec.Block{
		Header: spec.BlockHeader{
			Height:      100,
			Timestamp:   now.UnixNano(),
			Hash:        "test-hash",
			NextEpochID: "epoch_100",
		},
	}

	// Set the block
	s.SetBlock(100, "test-hash", "final", block)

	// Retrieve by height
	resp, err := s.Block(ctx, &api.BlockOpts{BlockID: 100})
	require.NoError(t, err)
	require.Equal(t, int64(100), resp.Data.Header.Height)

	// Retrieve by hash
	resp, err = s.Block(ctx, &api.BlockOpts{Hash: "test-hash"})
	require.NoError(t, err)
	require.Equal(t, "test-hash", resp.Data.Header.Hash)

	// Retrieve by finality
	resp, err = s.Block(ctx, &api.BlockOpts{Finality: "final"})
	require.NoError(t, err)
	require.Equal(t, int64(100), resp.Data.Header.Height)
}

func TestMockBlockAtTimestamp(t *testing.T) {
	ctx := context.Background()

	s, err := mock.New()
	require.NoError(t, err)

	now := time.Now()

	// Add multiple blocks with different timestamps
	for i := int64(100); i <= 200; i += 10 {
		block := &spec.Block{
			Header: spec.BlockHeader{
				Height:    i,
				Timestamp: now.Add(time.Duration(i-100) * time.Hour).UnixNano(),
			},
		}
		s.SetBlock(i, "", "", block)
	}

	// Find block at a specific timestamp
	targetTime := now.Add(55 * time.Hour)
	resp, err := s.BlockAtTimestamp(ctx, targetTime)
	require.NoError(t, err)
	// Should find block 150 (at 50 hours) as it's the last block before 55 hours
	require.Equal(t, int64(150), resp.Data.Header.Height)
}

func TestMockAccount(t *testing.T) {
	ctx := context.Background()

	s, err := mock.New()
	require.NoError(t, err)

	// Set account data
	accountData := &spec.Account{
		AccountID:       types.AccountID("test.near"),
		StakedBalance:   &types.Balance{Balance: big.NewInt(1000000)},
		UnstakedBalance: &types.Balance{Balance: big.NewInt(500000)},
		CanWithdraw:     true,
	}

	s.SetAccount("contract.near", "test.near", "final", accountData)

	// Retrieve account
	resp, err := s.Account(ctx, &api.AccountOpts{
		ContractID: "contract.near",
		AccountID:  "test.near",
		Block:      &api.BlockOpts{Finality: "final"},
	})
	require.NoError(t, err)
	require.Equal(t, types.AccountID("test.near"), resp.Data.AccountID)
	require.Equal(t, big.NewInt(1000000), resp.Data.StakedBalance.Balance)
}

func TestMockCallQueryFor(t *testing.T) {
	ctx := context.Background()

	s, err := mock.New()
	require.NoError(t, err)

	// Set up mock response
	type HumanReadableAccount struct {
		AccountID     string         `json:"account_id"`
		StakedBalance *types.Balance `json:"staked_balance"`
	}

	mockAccount := &HumanReadableAccount{
		AccountID:     "test.near",
		StakedBalance: &types.Balance{Balance: big.NewInt(2000000)},
	}

	s.SetAccount("stake.pool.near", "test.near", "100", mockAccount)

	// Call function with args
	args := map[string]any{
		"account_id": "test.near",
	}

	var resultAccount HumanReadableAccount
	err = s.CallQueryFor(ctx, &resultAccount, "get_account", "stake.pool.near", &api.BlockOpts{BlockID: 100}, args)
	require.NoError(t, err)
	require.Equal(t, "test.near", resultAccount.AccountID)
	require.Equal(t, big.NewInt(2000000), resultAccount.StakedBalance.Balance)
}

func TestMockMakeRPCQueryCall(t *testing.T) {
	ctx := context.Background()

	s, err := mock.New()
	require.NoError(t, err)

	// Set up mock response
	type HumanReadableAccount struct {
		AccountID     string         `json:"account_id"`
		StakedBalance *types.Balance `json:"staked_balance"`
	}

	mockAccount := &HumanReadableAccount{
		AccountID:     "alice.near",
		StakedBalance: &types.Balance{Balance: big.NewInt(3000000)},
	}

	s.SetAccount("stake.pool.near", "alice.near", "final", mockAccount)

	// Call function with args
	args := map[string]any{
		"account_id": "alice.near",
	}

	result, err := s.MakeRPCQueryCall(ctx, "get_account", "stake.pool.near", &api.BlockOpts{Finality: "final"}, args)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotEmpty(t, result.Result)

	var resultAccount HumanReadableAccount
	err = json.Unmarshal(result.Result, &resultAccount)
	require.NoError(t, err)
	require.Equal(t, "alice.near", resultAccount.AccountID)
	require.Equal(t, big.NewInt(3000000), resultAccount.StakedBalance.Balance)
}

func TestMockStatus(t *testing.T) {
	ctx := context.Background()

	s, err := mock.New()
	require.NoError(t, err)

	// Use default status
	resp, err := s.Status(ctx, &api.StatusOpts{})
	require.NoError(t, err)
	require.Equal(t, "mock-chain", resp.Data.ChainID)
	require.Equal(t, "mock-genesis-hash", resp.Data.GenesisHash)

	// Set custom status
	customStatus := &spec.Status{
		ChainID:     "custom-chain",
		GenesisHash: "custom-genesis",
		SyncInfo: &spec.SyncInfo{
			LatestBlockHeight: 2000,
			LatestBlockTime:   time.Now(),
			Syncing:           false,
		},
	}
	s.SetStatus(customStatus)

	resp, err = s.Status(ctx, &api.StatusOpts{})
	require.NoError(t, err)
	require.Equal(t, "custom-chain", resp.Data.ChainID)
	require.Equal(t, "custom-genesis", resp.Data.GenesisHash)
}

func TestMockSyncing(t *testing.T) {
	ctx := context.Background()

	s, err := mock.New()
	require.NoError(t, err)

	// Test default sync info
	resp, err := s.Syncing(ctx, &api.SyncingOpts{})
	require.NoError(t, err)
	require.False(t, resp.Data.Syncing)
	require.Equal(t, int64(1000), resp.Data.LatestBlockHeight)

	// Set custom sync info
	customSyncInfo := &spec.SyncInfo{
		LatestBlockHeight: 5000,
		LatestBlockTime:   time.Now(),
		Syncing:           true,
	}
	s.SetSyncInfo(customSyncInfo)

	resp, err = s.Syncing(ctx, &api.SyncingOpts{})
	require.NoError(t, err)
	require.True(t, resp.Data.Syncing)
	require.Equal(t, int64(5000), resp.Data.LatestBlockHeight)
}

func TestMockGenesisBlockHeight(t *testing.T) {
	ctx := context.Background()

	s, err := mock.New()
	require.NoError(t, err)

	// Test default genesis height
	height, err := s.GenesisBlockHeight(ctx)
	require.NoError(t, err)
	require.Equal(t, int64(0), height)

	// Set custom genesis height
	s.SetGenesisBlockHeight(1000)

	height, err = s.GenesisBlockHeight(ctx)
	require.NoError(t, err)
	require.Equal(t, int64(1000), height)
}

func TestMockChunk(t *testing.T) {
	ctx := context.Background()

	s, err := mock.New()
	require.NoError(t, err)

	chunk := &spec.Chunk{
		Header: spec.ChunkHeader{
			ChunkHash:      "test-chunk-hash",
			HeightCreated:  100,
			HeightIncluded: 101,
			ShardID:        0,
		},
	}

	// Set chunk by ID
	s.SetChunk("test-chunk-hash", 0, nil, chunk)

	resp, err := s.Chunk(ctx, &api.ChunkOpts{ChunkID: "test-chunk-hash"})
	require.NoError(t, err)
	require.Equal(t, "test-chunk-hash", resp.Data.Header.ChunkHash)

	// Set chunk by block ID and shard ID
	shardID := uint64(0)
	s.SetChunk("", 100, &shardID, chunk)

	resp, err = s.Chunk(ctx, &api.ChunkOpts{BlockID: 100, ShardID: &shardID})
	require.NoError(t, err)
	require.Equal(t, "test-chunk-hash", resp.Data.Header.ChunkHash)
}
