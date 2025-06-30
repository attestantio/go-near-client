package jsonrpc

import (
	"context"
	"time"

	"github.com/attestantio/go-near-client/spec"
	"github.com/pkg/errors"
)

const BLOCKTIME_NANOSECONDS = 600 * 1000 * 1000 // 600ms

// abs returns the absolute value of an int64
func abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}

func (s *service) BlockAtTimestamp(ctx context.Context, timestamp time.Time) (*spec.Block, error) {
	// Fetch highest block and timestamp from status
	status, err := s.Status(ctx)
	if err != nil {
		return nil, err
	}

	latestBlockTime := status.SyncInfo.LatestBlockTime
	if timestamp.After(latestBlockTime) {
		return nil, errors.New("timestamp is after latest block time")
	}

	latestBlockHeight := status.SyncInfo.LatestBlockHeight
	fetchedBlockHeight := latestBlockHeight + (timestamp.UnixNano()-latestBlockTime.UnixNano())/BLOCKTIME_NANOSECONDS

	return s.fetchNextBlock(ctx, timestamp.UnixNano(), fetchedBlockHeight, latestBlockHeight, latestBlockTime.UnixNano(), latestBlockHeight)
}

func (s *service) fetchNextBlock(ctx context.Context, targetTimestamp int64, fetchedBlockHeight int64, knownBlockHeight int64, knownBlockTimestamp int64, latestBlockHeight int64) (*spec.Block, error) {
	fetchedBlock, err := s.BlockByID(ctx, fetchedBlockHeight)
	if err != nil {
		if err == ErrBlockNotFound {
			return s.fetchNextBlock(ctx, targetTimestamp, fetchedBlockHeight-1, knownBlockHeight, knownBlockTimestamp, latestBlockHeight)
		}
		return nil, err
	}

	fetchedBlockTimestamp := fetchedBlock.Header.Timestamp
	timePerBlock := (fetchedBlockTimestamp - knownBlockTimestamp) / (fetchedBlockHeight - knownBlockHeight)
	nextBlockHeight := fetchedBlockHeight + (targetTimestamp-fetchedBlockTimestamp)/timePerBlock
	if nextBlockHeight > latestBlockHeight {
		return nil, errors.New("timestamp is after latest block time")
	}

	// TODO: RPCs and archival nodes do not store blocks all the way to zero.
	// First suupport support seems to be around block 10000000
	if nextBlockHeight < 0 {
		if fetchedBlockHeight == 0 {
			return nil, errors.New("timestamp is before first block time")
		}
		nextBlockHeight = 0
	}

	if abs(nextBlockHeight-fetchedBlockHeight) <= 5 {
		block, err := s.findLastBlockBeforeTimestamp(ctx, targetTimestamp, nextBlockHeight, fetchedBlockTimestamp, latestBlockHeight)
		if err != nil {
			return nil, err
		}
		return block, nil
	}
	return s.fetchNextBlock(ctx, targetTimestamp, nextBlockHeight, fetchedBlockHeight, fetchedBlockTimestamp, latestBlockHeight)
}

func (s *service) findLastBlockBeforeTimestamp(ctx context.Context, targetTimestamp int64, startBlockHeight int64, startBlockTimestamp int64, latestBlockHeight int64) (*spec.Block, error) {
	currentHeight := startBlockHeight

	if startBlockTimestamp > targetTimestamp {
		// `startBlockTimestamp`` is after targetTimestamp, so we need to search backwards.
		// Search backwards until we find the last block before the target timestamp.
		for currentHeight > 0 {
			block, err := s.BlockByID(ctx, currentHeight)
			if err != nil {
				if err == ErrBlockNotFound {
					currentHeight--
					continue
				}
				return nil, err
			}

			if block.Header.Timestamp <= targetTimestamp {
				// Found the last block before or at the target timestamp.
				return block, nil
			}

			currentHeight--
		}
		return nil, errors.New("timestamp is before first block time")
	} else {
		// `startBlockTimestamp`` is before targetTimestamp, so we need to search forwards.
		// Search forwards until we find the first block after the target timestamp.
		previousBlock := &spec.Block{}
		for currentHeight <= latestBlockHeight {
			block, err := s.BlockByID(ctx, currentHeight)
			if err != nil {
				if err == ErrBlockNotFound {
					currentHeight++
					continue
				}
				return nil, err
			}

			if block.Header.Timestamp >= targetTimestamp {
				// Found the first block after the target timestamp.
				return previousBlock, nil
			}

			previousBlock = block
			currentHeight++
		}
		return nil, errors.New("timestamp is after latest block time")
	}
}
