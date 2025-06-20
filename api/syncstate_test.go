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

package api_test

import (
	"encoding/json"
	"testing"

	"github.com/attestantio/go-near-client/api"
	"github.com/stretchr/testify/require"
)

func TestSyncState(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		expected []byte
		err      string
	}{
		{
			name: "Empty",
			err:  "unexpected end of JSON input",
		},
		{
			name:  "JSONBad",
			input: []byte("[]"),
			err:   "invalid sync state JSON \"[]\"",
		},
		{
			name:     "GoodSynced",
			input:    []byte(`false`),
			expected: []byte(`{"syncing":false}`),
		},
		{
			name:  "StartingBlockHashMissing",
			input: []byte(`{"starting_block_num":198300,"current_block_hash":"0x60d543771bbf5712e8a93b584c4d9287a40344b6fe3ca3e1fb04e3b73733209","current_block_num":198323,"highest_block_hash":"0x54f1d4f9ad2fc60ddd84a9defe58de0273c1261f5caddf261789a5a9a570ad3","highest_block_num":198335}`),
			err:   "starting block hash missing",
		},
		{
			name:  "StartingBlockHashWrongType",
			input: []byte(`{"starting_block_hash":true,"starting_block_num":198300,"current_block_hash":"0x60d543771bbf5712e8a93b584c4d9287a40344b6fe3ca3e1fb04e3b73733209","current_block_num":198323,"highest_block_hash":"0x54f1d4f9ad2fc60ddd84a9defe58de0273c1261f5caddf261789a5a9a570ad3","highest_block_num":198335}`),
			err:   "invalid JSON\njson: cannot unmarshal bool into Go struct field syncStateJSON.starting_block_hash of type string",
		},
		{
			name:  "StartingBlockHashInvalid",
			input: []byte(`{"starting_block_hash":"invalid","starting_block_num":198300,"current_block_hash":"0x60d543771bbf5712e8a93b584c4d9287a40344b6fe3ca3e1fb04e3b73733209","current_block_num":198323,"highest_block_hash":"0x54f1d4f9ad2fc60ddd84a9defe58de0273c1261f5caddf261789a5a9a570ad3","highest_block_num":198335}`),
			err:   "starting block hash invalid\nencoding/hex: invalid byte: U+0069 'i'",
		},
		{
			name:  "CurrentBlockHashMissing",
			input: []byte(`{"starting_block_hash":"0x39056d179abe29cdbdee010d04c3a33cb10f86e53cbf49537d41afefa367e56","starting_block_num":198300,"current_block_num":198323,"highest_block_hash":"0x54f1d4f9ad2fc60ddd84a9defe58de0273c1261f5caddf261789a5a9a570ad3","highest_block_num":198335}`),
			err:   "current block hash missing",
		},
		{
			name:  "CurrentBlockHashWrongType",
			input: []byte(`{"starting_block_hash":"0x39056d179abe29cdbdee010d04c3a33cb10f86e53cbf49537d41afefa367e56","starting_block_num":198300,"current_block_hash":true,"current_block_num":198323,"highest_block_hash":"0x54f1d4f9ad2fc60ddd84a9defe58de0273c1261f5caddf261789a5a9a570ad3","highest_block_num":198335}`),
			err:   "invalid JSON\njson: cannot unmarshal bool into Go struct field syncStateJSON.current_block_hash of type string",
		},
		{
			name:  "CurrentBlockHashInvalid",
			input: []byte(`{"starting_block_hash":"0x39056d179abe29cdbdee010d04c3a33cb10f86e53cbf49537d41afefa367e56","starting_block_num":198300,"current_block_hash":"invalid","current_block_num":198323,"highest_block_hash":"0x54f1d4f9ad2fc60ddd84a9defe58de0273c1261f5caddf261789a5a9a570ad3","highest_block_num":198335}`),
			err:   "current block hash invalid\nencoding/hex: invalid byte: U+0069 'i'",
		},
		{
			name:  "HighestBlockHashMissing",
			input: []byte(`{"starting_block_hash":"0x39056d179abe29cdbdee010d04c3a33cb10f86e53cbf49537d41afefa367e56","starting_block_num":198300,"current_block_hash":"0x60d543771bbf5712e8a93b584c4d9287a40344b6fe3ca3e1fb04e3b73733209","current_block_num":198323,"highest_block_num":198335}`),
			err:   "highest block hash missing",
		},
		{
			name:  "HighestBlockHashWrongType",
			input: []byte(`{"starting_block_hash":"0x39056d179abe29cdbdee010d04c3a33cb10f86e53cbf49537d41afefa367e56","starting_block_num":198300,"current_block_hash":"0x60d543771bbf5712e8a93b584c4d9287a40344b6fe3ca3e1fb04e3b73733209","current_block_num":198323,"highest_block_hash":true,"highest_block_num":198335}`),
			err:   "invalid JSON\njson: cannot unmarshal bool into Go struct field syncStateJSON.highest_block_hash of type string",
		},
		{
			name:  "HighestBlockHashInvalid",
			input: []byte(`{"starting_block_hash":"0x39056d179abe29cdbdee010d04c3a33cb10f86e53cbf49537d41afefa367e56","starting_block_num":198300,"current_block_hash":"0x60d543771bbf5712e8a93b584c4d9287a40344b6fe3ca3e1fb04e3b73733209","current_block_num":198323,"highest_block_hash":"invalid","highest_block_num":198335}`),
			err:   "highest block hash invalid\nencoding/hex: invalid byte: U+0069 'i'",
		},
		{
			name:     "GoodSyncing",
			input:    []byte(`{"starting_block_hash":"0x39056d179abe29cdbdee010d04c3a33cb10f86e53cbf49537d41afefa367e56","starting_block_num":198300,"current_block_hash":"0x60d543771bbf5712e8a93b584c4d9287a40344b6fe3ca3e1fb04e3b73733209","current_block_num":198323,"highest_block_hash":"0x54f1d4f9ad2fc60ddd84a9defe58de0273c1261f5caddf261789a5a9a570ad3","highest_block_num":198335}`),
			expected: []byte(`{"syncing":true,"starting_block_hash":"0x39056d179abe29cdbdee010d04c3a33cb10f86e53cbf49537d41afefa367e56","starting_block_num":198300,"current_block_hash":"0x60d543771bbf5712e8a93b584c4d9287a40344b6fe3ca3e1fb04e3b73733209","current_block_num":198323,"highest_block_hash":"0x54f1d4f9ad2fc60ddd84a9defe58de0273c1261f5caddf261789a5a9a570ad3","highest_block_num":198335}`),
		},
		{
			name:     "GoodSyncing",
			input:    []byte(`{"current_block_hash":"0x3b7efb91e3ac1ddb612fe5b8cc85cd47f34348a3e4b4d6caf20234808fe102c","current_block_num":7137,"highest_block_hash":"0x6c73a0012b790387dd11b9250e30671e27eb5ab136addc97becb60d429c4f8e","highest_block_num":196447,"starting_block_hash":"0x23a86829b1d007e593c500c5cbce1993eea0c9ec3b0d2dfbd7ae455ad5a6ddc","starting_block_num":966}`),
			expected: []byte(`{"syncing":true,"starting_block_hash":"0x23a86829b1d007e593c500c5cbce1993eea0c9ec3b0d2dfbd7ae455ad5a6ddc","starting_block_num":966,"current_block_hash":"0x3b7efb91e3ac1ddb612fe5b8cc85cd47f34348a3e4b4d6caf20234808fe102c","current_block_num":7137,"highest_block_hash":"0x6c73a0012b790387dd11b9250e30671e27eb5ab136addc97becb60d429c4f8e","highest_block_num":196447}`),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var res api.SyncState
			err := json.Unmarshal(test.input, &res)
			if test.err != "" {
				require.EqualError(t, err, test.err)
			} else {
				require.NoError(t, err)
				rt, err := json.Marshal(&res)
				require.NoError(t, err)
				if test.expected != nil {
					require.Equal(t, string(test.expected), string(rt))
				} else {
					require.Equal(t, string(test.input), string(rt))
				}
				require.Equal(t, string(rt), res.String())
			}
		})
	}
}
