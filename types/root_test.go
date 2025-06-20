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

package types_test

import (
	"encoding/json"
	"testing"

	"github.com/attestantio/go-near-client/types"
	"github.com/stretchr/testify/require"
)

func TestRootUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name   string
		input  []byte
		output []byte
		err    string
	}{
		{
			name:  "Empty",
			input: nil,
			err:   "unexpected end of JSON input",
		},
		{
			name:  "Minimal",
			input: []byte(`"0x1"`),
		},
		{
			name:   "Short",
			input:  []byte(`"0x01"`),
			output: []byte(`"0x1"`),
		},
		{
			name:   "NotTruncated",
			input:  []byte(`"0x0f102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f"`),
			output: []byte(`"0xf102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f"`),
		},
		{
			name:  "Truncated",
			input: []byte(`"0xf102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f"`),
		},
		{
			name:  "Full",
			input: []byte(`"0xff102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f"`),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var res types.Root
			err := json.Unmarshal(test.input, &res)
			if test.err != "" {
				require.EqualError(t, err, test.err)
			} else {
				require.NoError(t, err)
				rt, err := json.Marshal(&res)
				require.NoError(t, err)
				if len(test.output) == 0 {
					require.Equal(t, string(test.input), string(rt))
					require.Equal(t, string(test.input), `"`+res.String()+`"`)
				} else {
					require.Equal(t, string(test.output), string(rt))
					require.Equal(t, string(test.output), `"`+res.String()+`"`)
				}
			}
		})
	}
}
