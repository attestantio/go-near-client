// Copyright © 2025 Attestant Limited.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package types_test

import (
	"encoding/json"
	"math/big"
	"testing"

	"github.com/attestantio/go-near-client/types"
	"github.com/stretchr/testify/require"
)

func getBalance(input string) *big.Int {
	newBigInt := big.NewInt(0)
	newBigInt.SetString(input, 10)
	return newBigInt
}

func TestBalanceUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		output  []byte
		balance *big.Int
		err     string
	}{
		{
			name:  "Empty",
			input: nil,
			err:   "unexpected end of JSON input",
		},
		{
			name:    "FullNEAR",
			input:   []byte(`"30246689787546400881670512075"`),
			output:  []byte(`"30246689787546400881670512075"`),
			balance: getBalance("30246689787546400881670512075"),
		},
		{
			name:    "MilliNEAR",
			input:   []byte(`"689787546400881670512075"`),
			output:  []byte(`"689787546400881670512075"`),
			balance: getBalance("689787546400881670512075"),
		},
		{
			name:    "yoctoNEAR",
			input:   []byte(`"87546400881670512075"`),
			output:  []byte(`"87546400881670512075"`),
			balance: getBalance("87546400881670512075"),
		},
		{
			name:    "small yoctoNEAR",
			input:   []byte(`"75"`),
			output:  []byte(`"75"`),
			balance: getBalance("75"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var res types.Balance
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
					require.Equal(t, test.balance, res.Balance)
				}
			}
		})
	}
}
