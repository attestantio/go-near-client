package jsonrpc

import (
	"context"
	"encoding/json"

	"github.com/attestantio/go-near-client/spec"
)

type ChunkByIDRequest struct {
	ChunkID string `json:"chunk_id"`
}

type ChunkByBlockAndShardRequest struct {
	BlockID uint64 `json:"block_id"`
	ShardID uint64 `json:"shard_id"`
}

func (s *service) ChunkByID(ctx context.Context, chunkID string) (*spec.Chunk, error) {
	result, err := s.makeRPCCall(ctx, "chunk", ChunkByIDRequest{ChunkID: chunkID})
	if err != nil {
		return nil, err
	}

	chunk, err := s.parseChunkResponse(result)
	if err != nil {
		return nil, err
	}

	return chunk, nil
}

func (s *service) ChunkByBlockAndShard(ctx context.Context, blockID uint64, shardID uint64) (*spec.Chunk, error) {
	result, err := s.makeRPCCall(ctx, "chunk", ChunkByBlockAndShardRequest{BlockID: blockID, ShardID: shardID})
	if err != nil {
		return nil, err
	}

	chunk, err := s.parseChunkResponse(result)
	if err != nil {
		return nil, err
	}

	return chunk, nil
}

func (s *service) parseChunkResponse(result json.RawMessage) (*spec.Chunk, error) {
	var chunk spec.Chunk
	if err := json.Unmarshal(result, &chunk); err != nil {
		return nil, err
	}

	return &chunk, nil
}
