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

package spec

// Validators represents the validators response from NEAR RPC.
type Validators struct {
	CurrentValidators []CurrentValidator `json:"current_validators"`
	NextValidators    []NextValidator    `json:"next_validators"`
	PrevEpochKickout  []PrevEpochKickout `json:"prev_epoch_kickout"`
	EpochStartHeight  int64              `json:"epoch_start_height"`
	EpochHeight       int64              `json:"epoch_height"`
}

// CurrentValidator represents a current validator in the network.
type CurrentValidator struct {
	AccountID                       string  `json:"account_id"`
	PublicKey                       string  `json:"public_key"`
	IsSlashed                       bool    `json:"is_slashed"`
	Stake                           string  `json:"stake"`
	Shards                          []int64 `json:"shards"`
	NumProducedBlocks               int64   `json:"num_produced_blocks"`
	NumExpectedBlocks               int64   `json:"num_expected_blocks"`
	NumProducedChunks               int64   `json:"num_produced_chunks"`
	NumExpectedChunks               int64   `json:"num_expected_chunks"`
	NumProducedChunksPerShard       []int64 `json:"num_produced_chunks_per_shard"`
	NumExpectedChunksPerShard       []int64 `json:"num_expected_chunks_per_shard"`
	NumProducedEndorsementsPerShard []int64 `json:"num_produced_endorsements_per_shard"`
	NumExpectedEndorsementsPerShard []int64 `json:"num_expected_endorsements_per_shard"`
}

// NextValidator represents a validator in the next epoch.
type NextValidator struct {
	AccountID string  `json:"account_id"`
	PublicKey string  `json:"public_key"`
	Stake     string  `json:"stake"`
	Shards    []int64 `json:"shards"`
}

// PrevEpochKickout represents a validator kicked out in the previous epoch.
type PrevEpochKickout struct {
	AccountID string        `json:"account_id"`
	Reason    KickoutReason `json:"reason"`
}

// KickoutReason represents the reason for a validator being kicked out.
type KickoutReason struct {
	NotEnoughBlocks *NotEnoughBlocksProduced `json:"NotEnoughBlocks,omitempty"`
	NotEnoughChunks *NotEnoughChunksProduced `json:"NotEnoughChunks,omitempty"`
	Unstaked        *bool                    `json:"Unstaked,omitempty"`
	NotEnoughStake  *NotEnoughStake          `json:"NotEnoughStake,omitempty"`
	DidNotGetASeat  *bool                    `json:"DidNotGetASeat,omitempty"`
}

// NotEnoughBlocksProduced represents the reason for not producing enough blocks.
type NotEnoughBlocksProduced struct {
	Produced int64 `json:"produced"`
	Expected int64 `json:"expected"`
}

// NotEnoughChunksProduced represents the reason for not producing enough chunks.
type NotEnoughChunksProduced struct {
	Produced int64 `json:"produced"`
	Expected int64 `json:"expected"`
}

// NotEnoughStake represents the reason for not having enough stake.
type NotEnoughStake struct {
	StakeU128 StakeAmount `json:"stake_u128"`
	Threshold StakeAmount `json:"threshold"`
}

// StakeAmount represents a stake amount with low and high parts.
type StakeAmount struct {
	Lo uint64 `json:"lo"`
	Hi uint64 `json:"hi"`
}
