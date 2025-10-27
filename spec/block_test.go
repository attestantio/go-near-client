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

package spec_test

import (
	"encoding/json"
	"testing"

	"github.com/attestantio/go-near-client/spec"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const jsonBlock = `{
  "author": "figment.poolv1.near",
  "chunks": [
    {
      "balance_burnt": "2217453536878000000000",
      "bandwidth_requests": { "V1": { "requests": [] } },
      "chunk_hash": "ALuXChsM28HTaKgYpvgBqdizmuvPBtAbuxgzx4q7XoTS",
      "congestion_info": {
        "allowed_shard": 9,
        "buffered_receipts_gas": "0",
        "delayed_receipts_gas": "0",
        "receipt_bytes": 0
      },
      "encoded_length": 41714,
      "encoded_merkle_root": "2s1fm25rwGnEg3iURDV8Bc5gWrbMF41pj8eSPQiroHq5",
      "gas_limit": 1000000000000000,
      "gas_used": 31036316868780,
      "height_created": 152423388,
      "height_included": 152423388,
      "outcome_root": "zfFRH1q2Jv6RbF9gdhdgsAQQH2iE6dE4kkJdVpGLBab",
      "outgoing_receipts_root": "GhV4PCKXYGFL6xHHUtezBKc2wTW1T1CE4MyPbq5JUECB",
      "prev_block_hash": "EEANLTSggvcHzJDTZcMA5PmeVK7fWpFbu9P2H4TZqr1n",
      "prev_state_root": "B2KTKWV21p1G9aMaAqA5GrJVnCEeZhUTtwZeBiKvsxjb",
      "rent_paid": "0",
      "shard_id": 0,
      "signature": "ed25519:2AjaQp67MrsGYXDPrXCK5Ui8TXVoNbQDCmmEfEN2kuQGU1JGAjHLxnZvKv66kdgTpYH9wSE4Q6RqMyr1vpM3R26i",
      "tx_root": "8zd5Qp7cnN4TWSrYRKHjJ5hu3UWcgYVYkVth48yjkEvy",
      "validator_proposals": [],
      "validator_reward": "0"
    },
    {
      "balance_burnt": "0",
      "bandwidth_requests": { "V1": { "requests": [] } },
      "chunk_hash": "26tij3sXZ8Jv2TJzVmdDE72Lo1fmgrtyoK1xnc8B9RHQ",
      "congestion_info": {
        "allowed_shard": 6,
        "buffered_receipts_gas": "0",
        "delayed_receipts_gas": "0",
        "receipt_bytes": 0
      },
      "encoded_length": 8,
      "encoded_merkle_root": "9zYue7drR1rhfzEEoc4WUXzaYRnRNihvRoGt1BgK7Lkk",
      "gas_limit": 1000000000000000,
      "gas_used": 0,
      "height_created": 152423388,
      "height_included": 152423388,
      "outcome_root": "11111111111111111111111111111111",
      "outgoing_receipts_root": "27xHDu2bH9gJfna5zxGMaSrcfuH2XZcEnAnvZL6fWMYL",
      "prev_block_hash": "EEANLTSggvcHzJDTZcMA5PmeVK7fWpFbu9P2H4TZqr1n",
      "prev_state_root": "12ttNBBUcRW5TvBhxqZXvc3DBESKPufpj55G59n4aUtC",
      "rent_paid": "0",
      "shard_id": 1,
      "signature": "ed25519:47d5cJGTS3oChaAAxRHf426Q2pWeTFTPSwz5DJ4zaPmX9aVFenoq9LXx6B3C4yQwmzi5WHjtMghFRaqdBTXTRfh1",
      "tx_root": "11111111111111111111111111111111",
      "validator_proposals": [],
      "validator_reward": "0"
    }
  ],
  "header": {
    "approvals": [
      "ed25519:2yAoyNAkTqxRGQBDcxJN4gnXaChnNfkTzjfpMVmBteBazYj5rNnuYo5rz3CsEkxJ1qENfKqsdb333zj9tDW7bHoB",
      "ed25519:5vcGbzKohCQfDtM7uuFUeHHjVedCXSMc8NydHRTxzxDZ6hHQdjUmyyegCwDfuEcepkkFov1cYSz4fVjMWb2pamTz",
      null,
      "ed25519:2os2pEPaNw2Gp4ViwyhNG1bN6sNuZZ9BWDob2krGurakBXqCE89zYLJ6CFYBKzzptjS5ZqQ4wjAYVSEvsn9W6Gcm"
    ],
    "block_body_hash": "TTb6LK7E5k8eGPhfSrA9YtLWyogRbuwgNokT6e99E72",
    "block_merkle_root": "H1YVpWyAN28BFETjnzN88Ut6D5AoKWiYHH6Zq46dFgDj",
    "block_ordinal": 142247608,
    "challenges_result": [],
    "challenges_root": "11111111111111111111111111111111",
    "chunk_endorsements": [
      [251, 238, 247, 102, 123, 247, 247, 221, 3],
      [189, 255, 255, 255, 239, 255, 127, 253, 6]
    ],
    "chunk_headers_root": "39YYs5tn9Y8g1HmkgS3GUR5LWzr5G6uqhJDxRXC23jNG",
    "chunk_mask": [true, true, true, true, true, true, true, true],
    "chunk_receipts_root": "7zYsqovWRr6CqKuyKRz8RnQ9KTKzU1fpQMcJNbLWtvWk",
    "chunk_tx_root": "GSNwMEn7LneQkyLFW9bqQ2EP7reVSj7RfXQeQSGZ527k",
    "chunks_included": 8,
    "epoch_id": "YM6foK5rXbkThBweZ6ZYubCM6WijZdnSgPgyKLKKHoz",
    "epoch_sync_data_hash": null,
    "gas_price": "100000000",
    "hash": "EhrfG2rghZFeedQYWApbpE8swN9c9Rha5mCV1QqJvfXP",
    "height": 152423388,
    "last_ds_final_block": "EEANLTSggvcHzJDTZcMA5PmeVK7fWpFbu9P2H4TZqr1n",
    "last_final_block": "Cj9azMwFbZfqLPa1vGMWqKeLY162dvmGvFysBzEACW1V",
    "latest_protocol_version": 77,
    "next_bp_hash": "5MPayzPnpGuZuFLS9dXyVWMhDjgWDe2tcz1W8L1C9HD2",
    "next_epoch_id": "EywTi5HyZD7tTcfXmwftnvzcf3RC6TuBt6PBXy1S5yfz",
    "outcome_root": "2WtSCchxGL8bDTMFgKWoLrCUZMne66Fg27SPNKZtgu8a",
    "prev_hash": "EEANLTSggvcHzJDTZcMA5PmeVK7fWpFbu9P2H4TZqr1n",
    "prev_height": 152423387,
    "prev_state_root": "4qpMb3QudW4APRr7mx8tMUxJuq9PJvSQFmTvtALTUZEG",
    "random_value": "ALXUva5BM9mju825qa4tDEzTSJpCbMPEWybNKmcv1Lda",
    "rent_paid": "0",
    "signature": "ed25519:2uA5R3eULBL3Xh4SKjtD9NuNmMkvidWUGw7m9aAQyGXPQsQPYcNiyPDyTkkm3GFtQkYQubRgrz5dspVkukC7tPXg",
    "timestamp": 1750742496787538908,
    "timestamp_nanosec": "1750742496787538908",
    "total_supply": "1256905134513715749426049553842780",
    "validator_proposals": [],
    "validator_reward": "0"
  }
}`

func TestBlockUnmarshal(t *testing.T) {
	var block spec.Block
	err := json.Unmarshal([]byte(jsonBlock), &block)
	require.NoError(t, err)

	// Test basic fields
	assert.Equal(t, "figment.poolv1.near", block.Author)
	assert.Len(t, block.Chunks, 2)

	// Test header fields
	header := block.Header
	assert.Len(t, header.Approvals, 4)
	assert.Equal(t, "ed25519:2yAoyNAkTqxRGQBDcxJN4gnXaChnNfkTzjfpMVmBteBazYj5rNnuYo5rz3CsEkxJ1qENfKqsdb333zj9tDW7bHoB", *header.Approvals[0])
	assert.Equal(t, "ed25519:5vcGbzKohCQfDtM7uuFUeHHjVedCXSMc8NydHRTxzxDZ6hHQdjUmyyegCwDfuEcepkkFov1cYSz4fVjMWb2pamTz", *header.Approvals[1])
	assert.Nil(t, header.Approvals[2])
	assert.Equal(t, "ed25519:2os2pEPaNw2Gp4ViwyhNG1bN6sNuZZ9BWDob2krGurakBXqCE89zYLJ6CFYBKzzptjS5ZqQ4wjAYVSEvsn9W6Gcm", *header.Approvals[3])

	assert.Equal(t, "TTb6LK7E5k8eGPhfSrA9YtLWyogRbuwgNokT6e99E72", header.BlockBodyHash)
	assert.Equal(t, "H1YVpWyAN28BFETjnzN88Ut6D5AoKWiYHH6Zq46dFgDj", header.BlockMerkleRoot)
	assert.Equal(t, int64(142247608), header.BlockOrdinal)
	assert.Len(t, header.ChallengesResult, 0)
	assert.Equal(t, "11111111111111111111111111111111", header.ChallengesRoot)
	assert.Len(t, header.ChunkEndorsements, 2)
	assert.Equal(t, []int{251, 238, 247, 102, 123, 247, 247, 221, 3}, header.ChunkEndorsements[0])
	assert.Equal(t, []int{189, 255, 255, 255, 239, 255, 127, 253, 6}, header.ChunkEndorsements[1])
	assert.Equal(t, "39YYs5tn9Y8g1HmkgS3GUR5LWzr5G6uqhJDxRXC23jNG", header.ChunkHeadersRoot)
	assert.Equal(t, []bool{true, true, true, true, true, true, true, true}, header.ChunkMask)
	assert.Equal(t, "7zYsqovWRr6CqKuyKRz8RnQ9KTKzU1fpQMcJNbLWtvWk", header.ChunkReceiptsRoot)
	assert.Equal(t, "GSNwMEn7LneQkyLFW9bqQ2EP7reVSj7RfXQeQSGZ527k", header.ChunkTxRoot)
	assert.Equal(t, 8, header.ChunksIncluded)
	assert.Equal(t, "YM6foK5rXbkThBweZ6ZYubCM6WijZdnSgPgyKLKKHoz", header.EpochID)
	assert.Nil(t, header.EpochSyncDataHash)
	assert.Equal(t, "100000000", header.GasPrice)
	assert.Equal(t, "EhrfG2rghZFeedQYWApbpE8swN9c9Rha5mCV1QqJvfXP", header.Hash)
	assert.Equal(t, int64(152423388), header.Height)
	assert.Equal(t, "EEANLTSggvcHzJDTZcMA5PmeVK7fWpFbu9P2H4TZqr1n", header.LastDsFinalBlock)
	assert.Equal(t, "Cj9azMwFbZfqLPa1vGMWqKeLY162dvmGvFysBzEACW1V", header.LastFinalBlock)
	assert.Equal(t, 77, header.LatestProtocolVersion)
	assert.Equal(t, "5MPayzPnpGuZuFLS9dXyVWMhDjgWDe2tcz1W8L1C9HD2", header.NextBpHash)
	assert.Equal(t, "EywTi5HyZD7tTcfXmwftnvzcf3RC6TuBt6PBXy1S5yfz", header.NextEpochID)
	assert.Equal(t, "2WtSCchxGL8bDTMFgKWoLrCUZMne66Fg27SPNKZtgu8a", header.OutcomeRoot)
	assert.Equal(t, "EEANLTSggvcHzJDTZcMA5PmeVK7fWpFbu9P2H4TZqr1n", header.PrevHash)
	assert.Equal(t, int64(152423387), header.PrevHeight)
	assert.Equal(t, "4qpMb3QudW4APRr7mx8tMUxJuq9PJvSQFmTvtALTUZEG", header.PrevStateRoot)
	assert.Equal(t, "ALXUva5BM9mju825qa4tDEzTSJpCbMPEWybNKmcv1Lda", header.RandomValue)
	assert.Equal(t, "0", header.RentPaid)
	assert.Equal(t, "ed25519:2uA5R3eULBL3Xh4SKjtD9NuNmMkvidWUGw7m9aAQyGXPQsQPYcNiyPDyTkkm3GFtQkYQubRgrz5dspVkukC7tPXg", header.Signature)
	assert.Equal(t, int64(1750742496787538908), header.Timestamp)
	assert.Equal(t, "1750742496787538908", header.TimestampNanosec)
	assert.Equal(t, "1256905134513715749426049553842780", header.TotalSupply)
	assert.Len(t, header.ValidatorProposals, 0)
	assert.Equal(t, "0", header.ValidatorReward)

	// Test first chunk
	chunk1 := block.Chunks[0]
	assert.Equal(t, "2217453536878000000000", chunk1.BalanceBurnt)
	require.NotNil(t, chunk1.BandwidthRequests)
	assert.Len(t, chunk1.BandwidthRequests.V1.Requests, 0)
	assert.Equal(t, "ALuXChsM28HTaKgYpvgBqdizmuvPBtAbuxgzx4q7XoTS", chunk1.ChunkHash)
	require.NotNil(t, chunk1.CongestionInfo)
	assert.Equal(t, 9, chunk1.CongestionInfo.AllowedShard)
	assert.Equal(t, "0", chunk1.CongestionInfo.BufferedReceiptsGas)
	assert.Equal(t, "0", chunk1.CongestionInfo.DelayedReceiptsGas)
	assert.Equal(t, 0, chunk1.CongestionInfo.ReceiptBytes)
	assert.Equal(t, 41714, chunk1.EncodedLength)
	assert.Equal(t, "2s1fm25rwGnEg3iURDV8Bc5gWrbMF41pj8eSPQiroHq5", chunk1.EncodedMerkleRoot)
	assert.Equal(t, int64(1000000000000000), chunk1.GasLimit)
	assert.Equal(t, int64(31036316868780), chunk1.GasUsed)
	assert.Equal(t, int64(152423388), chunk1.HeightCreated)
	assert.Equal(t, int64(152423388), chunk1.HeightIncluded)
	assert.Equal(t, "zfFRH1q2Jv6RbF9gdhdgsAQQH2iE6dE4kkJdVpGLBab", chunk1.OutcomeRoot)
	assert.Equal(t, "GhV4PCKXYGFL6xHHUtezBKc2wTW1T1CE4MyPbq5JUECB", chunk1.OutgoingReceiptsRoot)
	assert.Equal(t, "EEANLTSggvcHzJDTZcMA5PmeVK7fWpFbu9P2H4TZqr1n", chunk1.PrevBlockHash)
	assert.Equal(t, "B2KTKWV21p1G9aMaAqA5GrJVnCEeZhUTtwZeBiKvsxjb", chunk1.PrevStateRoot)
	assert.Equal(t, "0", chunk1.RentPaid)
	assert.Equal(t, 0, chunk1.ShardID)
	assert.Equal(t, "ed25519:2AjaQp67MrsGYXDPrXCK5Ui8TXVoNbQDCmmEfEN2kuQGU1JGAjHLxnZvKv66kdgTpYH9wSE4Q6RqMyr1vpM3R26i", chunk1.Signature)
	assert.Equal(t, "8zd5Qp7cnN4TWSrYRKHjJ5hu3UWcgYVYkVth48yjkEvy", chunk1.TxRoot)
	assert.Len(t, chunk1.ValidatorProposals, 0)
	assert.Equal(t, "0", chunk1.ValidatorReward)

	// Test second chunk
	chunk2 := block.Chunks[1]
	assert.Equal(t, "0", chunk2.BalanceBurnt)
	assert.Equal(t, "26tij3sXZ8Jv2TJzVmdDE72Lo1fmgrtyoK1xnc8B9RHQ", chunk2.ChunkHash)
	assert.Equal(t, 1, chunk2.ShardID)
	assert.Equal(t, int64(0), chunk2.GasUsed)
}

func TestBlockMarshal(t *testing.T) {
	// First unmarshal the JSON to get a valid Block struct
	var block spec.Block
	err := json.Unmarshal([]byte(jsonBlock), &block)
	require.NoError(t, err)

	// Then marshal it back to JSON
	marshaled, err := json.Marshal(block)
	require.NoError(t, err)

	// Unmarshal the marshaled JSON to verify round-trip
	var roundTripBlock spec.Block
	err = json.Unmarshal(marshaled, &roundTripBlock)
	require.NoError(t, err)

	// Verify the round-trip preserved all data
	assert.Equal(t, block.Author, roundTripBlock.Author)
	assert.Equal(t, block.Header, roundTripBlock.Header)
	assert.Equal(t, block.Chunks, roundTripBlock.Chunks)
}

func TestBlockWithNonNullableFields(t *testing.T) {
	// Test with non-null values for nullable fields
	blockWithValues := spec.Block{
		Author: "test.pool.near",
		Header: spec.BlockHeader{
			Approvals: []*string{
				stringPtr("ed25519:test_approval1"),
				stringPtr("ed25519:test_approval2"),
				nil,
				stringPtr("ed25519:test_approval3"),
			},
			BlockBodyHash:         "test_block_body_hash",
			BlockMerkleRoot:       "test_block_merkle_root",
			BlockOrdinal:          1000000,
			ChallengesResult:      []string{"challenge1", "challenge2"},
			ChallengesRoot:        "test_challenges_root",
			ChunkEndorsements:     [][]int{{1, 2, 3}, {4, 5, 6}},
			ChunkHeadersRoot:      "test_chunk_headers_root",
			ChunkMask:             []bool{true, false, true},
			ChunkReceiptsRoot:     "test_chunk_receipts_root",
			ChunkTxRoot:           "test_chunk_tx_root",
			ChunksIncluded:        4,
			EpochID:               "test_epoch_id",
			EpochSyncDataHash:     stringPtr("test_epoch_sync_data_hash"),
			GasPrice:              "100000000",
			Hash:                  "test_hash",
			Height:                1000000,
			LastDsFinalBlock:      "test_last_ds_final_block",
			LastFinalBlock:        "test_last_final_block",
			LatestProtocolVersion: 77,
			NextBpHash:            "test_next_bp_hash",
			NextEpochID:           "test_next_epoch_id",
			OutcomeRoot:           "test_outcome_root",
			PrevHash:              "test_prev_hash",
			PrevHeight:            999999,
			PrevStateRoot:         "test_prev_state_root",
			RandomValue:           "test_random_value",
			RentPaid:              "1000",
			Signature:             "ed25519:test_signature",
			Timestamp:             1750742496787538908,
			TimestampNanosec:      "1750742496787538908",
			TotalSupply:           "1000000000000000000000000000000000",
			ValidatorProposals:    []spec.ValidatorProposal{{AccountID: "test_account_id", PublicKey: "test_public_key", Stake: "100", ValidatorStakeStructVersion: "1.0"}},
			ValidatorReward:       "500",
		},
		Chunks: []spec.ChunkHeader{
			{
				BalanceBurnt: "1000000000000000000000",
				BandwidthRequests: &spec.BandwidthRequests{
					V1: spec.BandwidthRequestsV1{
						Requests: []interface{}{"request1", "request2"},
					},
				},
				ChunkHash: "test_chunk_hash",
				CongestionInfo: &spec.CongestionInfo{
					AllowedShard:        5,
					BufferedReceiptsGas: "1000",
					DelayedReceiptsGas:  "2000",
					ReceiptBytes:        100,
				},
				EncodedLength:        1000,
				EncodedMerkleRoot:    "test_encoded_merkle_root",
				GasLimit:             1000000000000000,
				GasUsed:              500000000000000,
				HeightCreated:        1000000,
				HeightIncluded:       1000000,
				OutcomeRoot:          "test_outcome_root",
				OutgoingReceiptsRoot: "test_outgoing_receipts_root",
				PrevBlockHash:        "test_prev_block_hash",
				PrevStateRoot:        "test_prev_state_root",
				RentPaid:             "100",
				ShardID:              1,
				Signature:            "ed25519:test_chunk_signature",
				TxRoot:               "test_tx_root",
				ValidatorProposals:   []spec.ValidatorProposal{{AccountID: "test_account_id", PublicKey: "test_public_key", Stake: "100", ValidatorStakeStructVersion: "1.0"}},
				ValidatorReward:      "200",
			},
		},
	}

	// Marshal to JSON
	marshaled, err := json.Marshal(blockWithValues)
	require.NoError(t, err)

	// Unmarshal back
	var unmarshaled spec.Block
	err = json.Unmarshal(marshaled, &unmarshaled)
	require.NoError(t, err)

	// Verify all fields are preserved
	assert.Equal(t, blockWithValues, unmarshaled)
}

func stringPtr(s string) *string {
	return &s
}
