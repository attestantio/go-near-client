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

// Chunk represents a NEAR chunk with its header, receipts, and transactions
type Chunk struct {
	Author       string        `json:"author"`
	Header       ChunkHeader   `json:"header"`
	Receipts     []Receipt     `json:"receipts"`
	Transactions []Transaction `json:"transactions"`
}

// ChunkHeader contains metadata about a chunk
type ChunkHeader struct {
	BalanceBurnt         string              `json:"balance_burnt"`
	BandwidthRequests    *BandwidthRequests  `json:"bandwidth_requests"`
	ChunkHash            string              `json:"chunk_hash"`
	CongestionInfo       *CongestionInfo     `json:"congestion_info"`
	EncodedLength        int                 `json:"encoded_length"`
	EncodedMerkleRoot    string              `json:"encoded_merkle_root"`
	GasLimit             int64               `json:"gas_limit"`
	GasUsed              int64               `json:"gas_used"`
	HeightCreated        int64               `json:"height_created"`
	HeightIncluded       int64               `json:"height_included"`
	OutcomeRoot          string              `json:"outcome_root"`
	OutgoingReceiptsRoot string              `json:"outgoing_receipts_root"`
	PrevBlockHash        string              `json:"prev_block_hash"`
	PrevStateRoot        string              `json:"prev_state_root"`
	RentPaid             string              `json:"rent_paid"`
	ShardID              int                 `json:"shard_id"`
	Signature            string              `json:"signature"`
	TxRoot               string              `json:"tx_root"`
	ValidatorProposals   []ValidatorProposal `json:"validator_proposals"`
	ValidatorReward      string              `json:"validator_reward"`
}

// ValidatorProposal represents a validator proposal
type ValidatorProposal struct {
	AccountID                   string `json:"account_id"`
	PublicKey                   string `json:"public_key"`
	Stake                       string `json:"stake"`
	ValidatorStakeStructVersion string `json:"validator_stake_struct_version"`
}

// Transaction represents a NEAR transaction
type Transaction struct {
	Actions     []Action `json:"actions"`
	Hash        string   `json:"hash"`
	Nonce       int64    `json:"nonce"`
	PriorityFee int64    `json:"priority_fee"`
	PublicKey   string   `json:"public_key"`
	ReceiverID  string   `json:"receiver_id"`
	Signature   string   `json:"signature"`
	SignerID    string   `json:"signer_id"`
}

// BandwidthRequests represents bandwidth requests information
type BandwidthRequests struct {
	V1 BandwidthRequestsV1 `json:"V1"`
}

// BandwidthRequestsV1 represents V1 bandwidth requests
type BandwidthRequestsV1 struct {
	Requests []interface{} `json:"requests"`
}

// CongestionInfo represents congestion information for a chunk
type CongestionInfo struct {
	AllowedShard        int    `json:"allowed_shard"`
	BufferedReceiptsGas string `json:"buffered_receipts_gas"`
	DelayedReceiptsGas  string `json:"delayed_receipts_gas"`
	ReceiptBytes        int    `json:"receipt_bytes"`
}
