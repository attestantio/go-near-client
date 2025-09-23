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
	"time"

	"github.com/attestantio/go-near-client/api"
	"github.com/attestantio/go-near-client/spec"
	"github.com/pkg/errors"
)

const blockTimeNanoseconds = 600 * 1000 * 1000 // 600ms

// abs returns the absolute value of an int64.
func abs(x int64) int64 {
	if x < 0 {
		return -x
	}

	return x
}

// BlockAtTimestamp returns last block before or at a given timestamp.
func (s *Service) BlockAtTimestamp(ctx context.Context, timestamp time.Time) (*api.Response[*spec.Block], error) {
	// Fetch highest block and timestamp from status
	status, err := s.Status(ctx, nil)
	if err != nil {
		return nil, err
	}

	if status.Data == nil {
		return nil, errors.New("status data is nil")
	}

	latestBlockTime := status.Data.SyncInfo.LatestBlockTime
	if timestamp.After(latestBlockTime) {
		return nil, errors.New("timestamp is after latest block time")
	}

	latestBlockHeight := status.Data.SyncInfo.LatestBlockHeight
	fetchedBlockHeight := latestBlockHeight + (timestamp.UnixNano()-latestBlockTime.UnixNano())/blockTimeNanoseconds

	block, err := s.fetchNextBlock(ctx,
		timestamp.UnixNano(),
		fetchedBlockHeight,
		latestBlockHeight,
		latestBlockTime.UnixNano(),
		latestBlockHeight,
	)
	if err != nil {
		return nil, err
	}

	return &api.Response[*spec.Block]{
		Data:     block,
		Metadata: map[string]any{},
	}, nil
}

func (s *Service) fetchNextBlock(ctx context.Context,
	targetTimestamp int64,
	fetchedBlockHeight int64,
	knownBlockHeight int64,
	knownBlockTimestamp int64,
	latestBlockHeight int64,
) (*spec.Block, error) {
	fetchedBlockResponse, err := s.Block(ctx, &api.BlockOpts{BlockID: fetchedBlockHeight})
	if err != nil {
		if errors.Is(err, ErrBlockNotFound) {
			return s.fetchNextBlock(ctx,
				targetTimestamp,
				fetchedBlockHeight-1,
				knownBlockHeight,
				knownBlockTimestamp,
				latestBlockHeight,
			)
		}

		return nil, err
	}

	fetchedBlockTimestamp := fetchedBlockResponse.Data.Header.Timestamp
	timePerBlock := (fetchedBlockTimestamp - knownBlockTimestamp) / (fetchedBlockHeight - knownBlockHeight)
	nextBlockHeight := fetchedBlockHeight + (targetTimestamp-fetchedBlockTimestamp)/timePerBlock
	if nextBlockHeight > latestBlockHeight {
		return nil, errors.New("timestamp is after latest block time")
	}

	if nextBlockHeight < s.genesisBlockHeight {
		if fetchedBlockHeight == s.genesisBlockHeight {
			return nil, errors.New("timestamp is before first block time")
		}
		nextBlockHeight = s.genesisBlockHeight
	}

	if abs(nextBlockHeight-fetchedBlockHeight) <= 5 {
		block, err := s.findLastBlockBeforeTimestamp(ctx,
			targetTimestamp,
			fetchedBlockHeight,
			fetchedBlockTimestamp,
			latestBlockHeight,
		)
		if err != nil {
			return nil, err
		}

		return block, nil
	}

	return s.fetchNextBlock(ctx,
		targetTimestamp,
		nextBlockHeight,
		fetchedBlockHeight,
		fetchedBlockTimestamp,
		latestBlockHeight,
	)
}

func (s *Service) findLastBlockBeforeTimestamp(
	ctx context.Context,
	targetTimestamp int64,
	startBlockHeight int64,
	startBlockTimestamp int64,
	latestBlockHeight int64,
) (*spec.Block, error) {
	currentHeight := startBlockHeight

	if startBlockTimestamp > targetTimestamp {
		// `startBlockTimestamp`` is after targetTimestamp, so we need to search backwards.
		// Search backwards until we find the last block before the target timestamp.
		for currentHeight > s.genesisBlockHeight {
			blockResponse, err := s.Block(ctx, &api.BlockOpts{BlockID: currentHeight})
			if err != nil {
				if errors.Is(err, ErrBlockNotFound) {
					currentHeight--
					continue
				}

				return nil, err
			}

			if blockResponse.Data.Header.Timestamp <= targetTimestamp {
				// Found the last block before or at the target timestamp.
				return blockResponse.Data, nil
			}

			currentHeight--
		}

		return nil, errors.New("timestamp is before first block time")
	}

	// `startBlockTimestamp`` is before targetTimestamp, so we need to search forwards.
	// Search forwards until we find the first block after the target timestamp.
	previousBlock := &spec.Block{}
	for currentHeight <= latestBlockHeight {
		blockResponse, err := s.Block(ctx, &api.BlockOpts{BlockID: currentHeight})
		if err != nil {
			if errors.Is(err, ErrBlockNotFound) {
				currentHeight++
				continue
			}

			return nil, err
		}

		if blockResponse.Data.Header.Timestamp > targetTimestamp {
			// Found the first block after the target timestamp.
			return previousBlock, nil
		}

		previousBlock = blockResponse.Data
		currentHeight++
	}

	return nil, errors.New("timestamp is after latest block time")
}
