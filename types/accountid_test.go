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

func TestAccountIDUnmarshalJSON(t *testing.T) {
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
			name:  "Short",
			input: []byte(`"1"`),
			err:   "account id too short",
		},
		{
			name:  "Long",
			input: []byte(`"01234567890123456789012345678901234567890123456789012345678901234"`),
			err:   "account id too long",
		},
		{
			name:   "Minimal",
			input:  []byte(`"7a"`),
			output: []byte(`"7a"`),
		},
		{
			name:   "HexString",
			input:  []byte(`"658a1905ba0ce2eb5c2db802f8adc2c05c371c96174b2b90a04d6f2aaada1391"`),
			output: []byte(`"658a1905ba0ce2eb5c2db802f8adc2c05c371c96174b2b90a04d6f2aaada1391"`),
		},
		{
			name:   "FQDN",
			input:  []byte(`"apy-tracer.eco.linear-protocol.near"`),
			output: []byte(`"apy-tracer.eco.linear-protocol.near"`),
		},
		{
			name:   "ShortDomain",
			input:  []byte(`"meta-pool.near"`),
			output: []byte(`"meta-pool.near"`),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var res types.AccountID
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
