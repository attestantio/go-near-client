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

package types

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// BlockID is a block identifier.
// It can be a block number, block hash, or one of the special values "latest" or "pending".
type BlockID string

// String returns the string representation of the block ID.
func (b BlockID) String() string {
	return string(b)
}

// Format formats the address.
func (b BlockID) Format(state fmt.State, v rune) {
	format := string(v)
	switch v {
	case 'q':
		fmt.Fprint(state, `"`+string(b)+`"`)
	case 's':
		fmt.Fprint(state, string(b))
	case 'x', 'X':
		if state.Flag('#') {
			format = "#" + format
		}
		fmt.Fprintf(state, "%"+format, b[:])
	default:
		fmt.Fprintf(state, "%"+format, b[:])
	}
}

// MarshalJSON implements json.Marshaler.
func (b BlockID) MarshalJSON() ([]byte, error) {
	switch {
	case b == "latest" || b == "pending":
		return []byte(fmt.Sprintf("%q", b)), nil
	case strings.HasPrefix(string(b), "0x"):
		return []byte(fmt.Sprintf(`{"block_hash":"%s"}`, b)), nil
	default:
		block, err := strconv.ParseUint(string(b), 10, 64)
		if err != nil {
			return nil, errors.New("invalid from block")
		}

		return []byte(fmt.Sprintf(`{"block_number":%d}`, block)), nil
	}
}
