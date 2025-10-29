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

package mock

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/attestantio/go-near-client/api"
	"github.com/attestantio/go-near-client/spec"
)

// Service is a mock Near client service.
type Service struct {
	// Mock data stores
	blocks   map[string]*spec.Block
	accounts map[string]any
	chunks   map[string]*spec.Chunk
	status   *spec.Status
	syncInfo *spec.SyncInfo

	// Genesis block height
	genesisBlockHeight int64
}

// New creates a new mock.
func New() (*Service, error) {
	return &Service{
		blocks:             make(map[string]*spec.Block),
		accounts:           make(map[string]any),
		chunks:             make(map[string]*spec.Chunk),
		genesisBlockHeight: 0,
		status: &spec.Status{
			ChainID:               "mock-chain",
			GenesisHash:           "mock-genesis-hash",
			LatestProtocolVersion: 1,
			ProtocolVersion:       1,
			RPCAddr:               "mock-rpc",
			SyncInfo: &spec.SyncInfo{
				LatestBlockHeight: 1000,
				LatestBlockTime:   time.Now(),
				Syncing:           false,
			},
		},
		syncInfo: &spec.SyncInfo{
			LatestBlockHeight: 1000,
			LatestBlockTime:   time.Now(),
			Syncing:           false,
		},
	}, nil
}

// Name returns the name of the client implementation.
func (*Service) Name() string { return "mock" }

// Address returns the address of the client.
func (*Service) Address() string { return "mock" }

// SetBlock sets a block in the mock store by height, hash, and finality.
func (s *Service) SetBlock(height int64, hash string, finality string, block *spec.Block) {
	if height > 0 {
		s.blocks[fmt.Sprintf("height:%d", height)] = block
	}

	if hash != "" {
		s.blocks[fmt.Sprintf("hash:%s", hash)] = block
	}

	if finality != "" {
		s.blocks[fmt.Sprintf("finality:%s", finality)] = block
	}
}

// SetAccount sets account data in the mock store.
func (s *Service) SetAccount(contractID, accountID, blockID string, data any) {
	key := fmt.Sprintf("%s:%s:%s", contractID, accountID, blockID)
	s.accounts[key] = data
}

// SetChunk sets a chunk in the mock store.
func (s *Service) SetChunk(chunkID string, blockID uint64, shardID *uint64, chunk *spec.Chunk) {
	if chunkID != "" {
		s.chunks[fmt.Sprintf("chunkid:%s", chunkID)] = chunk
	}

	if blockID > 0 && shardID != nil {
		s.chunks[fmt.Sprintf("block:%d:shard:%d", blockID, *shardID)] = chunk
	}
}

// SetStatus sets the status in the mock store.
func (s *Service) SetStatus(status *spec.Status) {
	s.status = status
	if status.SyncInfo != nil {
		s.syncInfo = status.SyncInfo
	}
}

// SetSyncInfo sets the sync info in the mock store.
func (s *Service) SetSyncInfo(syncInfo *spec.SyncInfo) {
	s.syncInfo = syncInfo
	if s.status != nil {
		s.status.SyncInfo = syncInfo
	}
}

// SetGenesisBlockHeight sets the genesis block height.
func (s *Service) SetGenesisBlockHeight(height int64) {
	s.genesisBlockHeight = height
}

// Account retrieves account info from the mock.
func (s *Service) Account(_ context.Context, opts *api.AccountOpts) (*api.Response[*spec.Account], error) {
	if opts == nil {
		return nil, errors.New("no options provided")
	}

	blockID := "final"

	if opts.Block != nil {
		switch {
		case opts.Block.Finality != "":
			blockID = opts.Block.Finality
		case opts.Block.BlockID != 0:
			blockID = fmt.Sprintf("%d", opts.Block.BlockID)
		case opts.Block.Hash != "":
			blockID = opts.Block.Hash
		default:
			// blockID remains "final"
		}
	}

	key := fmt.Sprintf("%s:%s:%s", opts.ContractID, opts.AccountID, blockID)
	if data, exists := s.accounts[key]; exists {
		account, ok := data.(*spec.Account)
		if !ok {
			return nil, errors.New("account data is not of type *spec.Account")
		}

		return &api.Response[*spec.Account]{
			Data:     account,
			Metadata: map[string]any{},
		}, nil
	}

	return nil, fmt.Errorf("account not found: %s", key)
}

// Block returns the block.
func (s *Service) Block(_ context.Context, opts *api.BlockOpts) (*api.Response[*spec.Block], error) {
	if opts == nil {
		return nil, errors.New("no options provided")
	}

	var key string

	switch {
	case opts.Finality != "":
		key = fmt.Sprintf("finality:%s", opts.Finality)
	case opts.BlockID != 0:
		key = fmt.Sprintf("height:%d", opts.BlockID)
	case opts.Hash != "":
		key = fmt.Sprintf("hash:%s", opts.Hash)
	default:
		return nil, errors.New("invalid block options")
	}

	if block, exists := s.blocks[key]; exists {
		return &api.Response[*spec.Block]{
			Data:     block,
			Metadata: map[string]any{},
		}, nil
	}

	return nil, fmt.Errorf("block not found: %s", key)
}

// BlockAtTimestamp returns the last block before or at a given timestamp.
func (s *Service) BlockAtTimestamp(_ context.Context, timestamp time.Time) (*api.Response[*spec.Block], error) {
	// Simple mock implementation - find the block with the closest timestamp
	var closestBlock *spec.Block

	var closestDiff int64 = 1<<63 - 1

	for _, block := range s.blocks {
		if block.Header.Timestamp <= timestamp.UnixNano() {
			diff := timestamp.UnixNano() - block.Header.Timestamp
			if diff < closestDiff {
				closestDiff = diff
				closestBlock = block
			}
		}
	}

	if closestBlock == nil {
		return nil, fmt.Errorf("no block found before timestamp %v", timestamp)
	}

	return &api.Response[*spec.Block]{
		Data:     closestBlock,
		Metadata: map[string]any{},
	}, nil
}

// Status returns the status.
func (s *Service) Status(_ context.Context, _ *api.StatusOpts) (*api.Response[*spec.Status], error) {
	if s.status == nil {
		return nil, errors.New("status not set in mock")
	}

	return &api.Response[*spec.Status]{
		Data:     s.status,
		Metadata: map[string]any{},
	}, nil
}

// Syncing returns the sync info.
func (s *Service) Syncing(_ context.Context, opts *api.SyncingOpts) (*api.Response[*spec.SyncInfo], error) {
	if opts == nil {
		return nil, errors.New("no options provided")
	}

	if s.syncInfo == nil {
		return nil, errors.New("sync info not set in mock")
	}

	return &api.Response[*spec.SyncInfo]{
		Data:     s.syncInfo,
		Metadata: map[string]any{},
	}, nil
}

// GenesisBlockHeight returns the height of the genesis block.
func (s *Service) GenesisBlockHeight(_ context.Context) (int64, error) {
	return s.genesisBlockHeight, nil
}

// Chunk returns the chunk.
func (s *Service) Chunk(_ context.Context, opts *api.ChunkOpts) (*api.Response[*spec.Chunk], error) {
	if opts == nil {
		return nil, errors.New("no options provided")
	}

	var key string

	switch {
	case opts.ChunkID != "":
		key = fmt.Sprintf("chunkid:%s", opts.ChunkID)
	case opts.BlockID > 0 && opts.ShardID != nil:
		key = fmt.Sprintf("block:%d:shard:%d", opts.BlockID, *opts.ShardID)
	default:
		return nil, errors.New("invalid chunk options")
	}

	if chunk, exists := s.chunks[key]; exists {
		return &api.Response[*spec.Chunk]{
			Data:     chunk,
			Metadata: map[string]any{},
		}, nil
	}

	return nil, fmt.Errorf("chunk not found: %s", key)
}

// CallQueryFor calls a query method and unmarshals the result into the out parameter.
func (s *Service) CallQueryFor(_ context.Context,
	out any,
	method string,
	contractID string,
	block *api.BlockOpts,
	args map[string]any,
) error {
	blockID := "final"

	if block != nil {
		switch {
		case block.Finality != "":
			blockID = block.Finality
		case block.BlockID != 0:
			blockID = fmt.Sprintf("%d", block.BlockID)
		case block.Hash != "":
			blockID = block.Hash
		default:
			// blockID remains "final"
		}
	}

	// Extract account_id from args if present
	accountID := ""

	if args != nil {
		if id, ok := args["account_id"].(string); ok {
			accountID = id
		}
	}

	key := fmt.Sprintf("%s:%s:%s", contractID, accountID, blockID)
	if data, exists := s.accounts[key]; exists {
		// Marshal and unmarshal to convert to the output type
		jsonData, err := json.Marshal(data)
		if err != nil {
			return fmt.Errorf("failed to marshal mock data: %w", err)
		}

		if err := json.Unmarshal(jsonData, out); err != nil {
			return fmt.Errorf("failed to unmarshal to output type: %w", err)
		}

		return nil
	}

	return fmt.Errorf("query result not found: method=%s, contract=%s, key=%s", method, contractID, key)
}

// MakeRPCQueryCall makes a JSON-RPC query call to the NEAR node.
func (s *Service) MakeRPCQueryCall(
	_ context.Context,
	method string,
	contractID string,
	block *api.BlockOpts,
	args map[string]any,
) (*spec.CallFunctionResult, error) {
	blockID := "final"
	blockHeight := int64(0)
	blockHash := ""

	if block != nil {
		switch {
		case block.Finality != "":
			blockID = block.Finality
		case block.BlockID != 0:
			blockID = fmt.Sprintf("%d", block.BlockID)
			blockHeight = int64(block.BlockID) //nolint:unconvert // Conversion from uint64 to int64 is necessary
		case block.Hash != "":
			blockID = block.Hash
			blockHash = block.Hash
		default:
			// blockID remains "final"
		}
	}

	// Extract account_id from args if present
	accountID := ""

	if args != nil {
		if id, ok := args["account_id"].(string); ok {
			accountID = id
		}
	}

	key := fmt.Sprintf("%s:%s:%s", contractID, accountID, blockID)
	if data, exists := s.accounts[key]; exists {
		// Marshal the result
		jsonData, err := json.Marshal(data)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal mock data: %w", err)
		}

		return &spec.CallFunctionResult{
			BlockHeight: blockHeight,
			BlockHash:   blockHash,
			Logs:        []string{},
			Result:      jsonData,
		}, nil
	}

	return nil, fmt.Errorf("query result not found: method=%s, contract=%s, key=%s", method, contractID, key)
}
