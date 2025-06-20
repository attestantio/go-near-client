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

package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/attestantio/go-near-client/types"
	"github.com/attestantio/go-near-client/util"
)

// SyncState contains the sync state.
type SyncState struct {
	Syncing           bool
	StartingBlockHash types.Hash
	StartingBlockNum  uint32
	CurrentBlockHash  types.Hash
	CurrentBlockNum   uint32
	HighestBlockHash  types.Hash
	HighestBlockNum   uint32
}

// syncStateJSON is the spec representation of the struct.
type syncStateJSON struct {
	Syncing           bool   `json:"syncing"`
	StartingBlockHash string `json:"starting_block_hash"`
	StartingBlockNum  uint32 `json:"starting_block_num"`
	CurrentBlockHash  string `json:"current_block_hash"`
	CurrentBlockNum   uint32 `json:"current_block_num"`
	HighestBlockHash  string `json:"highest_block_hash"`
	HighestBlockNum   uint32 `json:"highest_block_num"`
}

// MarshalJSON implements json.Marshaler.
func (s *SyncState) MarshalJSON() ([]byte, error) {
	if !s.Syncing {
		return json.Marshal(map[string]bool{"syncing": false})
	}

	return json.Marshal(&syncStateJSON{
		Syncing:           s.Syncing,
		StartingBlockHash: s.StartingBlockHash.String(),
		StartingBlockNum:  s.StartingBlockNum,
		CurrentBlockHash:  s.CurrentBlockHash.String(),
		CurrentBlockNum:   s.CurrentBlockNum,
		HighestBlockHash:  s.HighestBlockHash.String(),
		HighestBlockNum:   s.HighestBlockNum,
	})
}

// UnmarshalJSON implements json.Unmarshaler.
func (s *SyncState) UnmarshalJSON(input []byte) error {
	// May be a simple bool, or an object.
	if bytes.HasPrefix(input, []byte("{")) {
		// It's an object.
		s.Syncing = true
		var data syncStateJSON
		if err := json.Unmarshal(input, &data); err != nil {
			return errors.Join(errors.New("invalid JSON"), err)
		}

		if err := s.unpack(&data); err != nil {
			return err
		}

		// Patch for pathfinder; remove when pathfinder 0.15 is released.
		s.Syncing = s.CurrentBlockNum < s.HighestBlockNum

		return nil
	}

	// It's a simple bool.
	if string(input) == "false" {
		s.StartingBlockNum = 0
		s.StartingBlockHash = types.Hash{}
		s.CurrentBlockNum = 0
		s.CurrentBlockHash = types.Hash{}
		s.HighestBlockNum = 0
		s.HighestBlockHash = types.Hash{}
		s.Syncing = false

		return nil
	}

	return fmt.Errorf("invalid sync state JSON %q", string(input))
}

func (s *SyncState) unpack(data *syncStateJSON) error {
	var err error

	s.StartingBlockNum = data.StartingBlockNum

	s.StartingBlockHash, err = util.StrToHash("starting block hash", data.StartingBlockHash)
	if err != nil {
		return err
	}

	s.CurrentBlockNum = data.CurrentBlockNum

	s.CurrentBlockHash, err = util.StrToHash("current block hash", data.CurrentBlockHash)
	if err != nil {
		return err
	}

	s.HighestBlockNum = data.HighestBlockNum

	s.HighestBlockHash, err = util.StrToHash("highest block hash", data.HighestBlockHash)
	if err != nil {
		return err
	}

	return nil
}

// String returns a string version of the structure.
func (s *SyncState) String() string {
	data, err := json.Marshal(s)
	if err != nil {
		return fmt.Sprintf("ERR: %v", err)
	}

	return string(data)
}
