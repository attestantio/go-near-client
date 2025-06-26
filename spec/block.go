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

package spec

// Block represents a block.
// Block represents a NEAR block with its header and chunks.
type Block struct {
	Author string        `json:"author"`
	Chunks []ChunkHeader `json:"chunks"`
	Header BlockHeader   `json:"header"`
}

// BlockHeader contains metadata about a block.
type BlockHeader struct {
	Approvals             []*string           `json:"approvals"`
	BlockBodyHash         string              `json:"block_body_hash"`
	BlockMerkleRoot       string              `json:"block_merkle_root"`
	BlockOrdinal          int64               `json:"block_ordinal"`
	ChallengesResult      []string            `json:"challenges_result"`
	ChallengesRoot        string              `json:"challenges_root"`
	ChunkEndorsements     [][]int             `json:"chunk_endorsements"`
	ChunkHeadersRoot      string              `json:"chunk_headers_root"`
	ChunkMask             []bool              `json:"chunk_mask"`
	ChunkReceiptsRoot     string              `json:"chunk_receipts_root"`
	ChunkTxRoot           string              `json:"chunk_tx_root"`
	ChunksIncluded        int                 `json:"chunks_included"`
	EpochID               string              `json:"epoch_id"`
	EpochSyncDataHash     *string             `json:"epoch_sync_data_hash"`
	GasPrice              string              `json:"gas_price"`
	Hash                  string              `json:"hash"`
	Height                int64               `json:"height"`
	LastDsFinalBlock      string              `json:"last_ds_final_block"`
	LastFinalBlock        string              `json:"last_final_block"`
	LatestProtocolVersion int                 `json:"latest_protocol_version"`
	NextBpHash            string              `json:"next_bp_hash"`
	NextEpochID           string              `json:"next_epoch_id"`
	OutcomeRoot           string              `json:"outcome_root"`
	PrevHash              string              `json:"prev_hash"`
	PrevHeight            int64               `json:"prev_height"`
	PrevStateRoot         string              `json:"prev_state_root"`
	RandomValue           string              `json:"random_value"`
	RentPaid              string              `json:"rent_paid"`
	Signature             string              `json:"signature"`
	Timestamp             int64               `json:"timestamp"`
	TimestampNanosec      string              `json:"timestamp_nanosec"`
	TotalSupply           string              `json:"total_supply"`
	ValidatorProposals    []ValidatorProposal `json:"validator_proposals"`
	ValidatorReward       string              `json:"validator_reward"`
}
