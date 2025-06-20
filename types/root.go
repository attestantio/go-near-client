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
	"bytes"
	"encoding/hex"
	"errors"
	"fmt"
	"strings"
)

// RootLength is the length of a root.
const RootLength = 32

// Root is a 32-byte merkle root.
type Root [RootLength]byte

// String returns the string representation of the root.
func (r *Root) String() string {
	res := hex.EncodeToString(r[:])
	// Leading 0s not allowed...
	res = strings.TrimLeft(res, "0")
	// ...unless that's all there was.
	if len(res) == 0 {
		res = "0"
	}

	return "0x" + res
}

// Format formats the root.
func (r *Root) Format(state fmt.State, v rune) {
	format := string(v)
	switch v {
	case 's':
		fmt.Fprint(state, r.String())
	case 'x', 'X':
		if state.Flag('#') {
			format = "#" + format
		}
		fmt.Fprintf(state, "%"+format, r[:])
	default:
		fmt.Fprintf(state, "%"+format, r[:])
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (r *Root) UnmarshalJSON(input []byte) error {
	if len(input) == 0 {
		return errors.New("root missing")
	}

	if !bytes.HasPrefix(input, []byte{'"', '0', 'x'}) {
		return errors.New("invalid root prefix")
	}
	if !bytes.HasSuffix(input, []byte{'"'}) {
		return errors.New("invalid root suffix")
	}

	// Ensure that there are an even number of characters.
	bytesStr := string(input[3 : len(input)-1])
	if len(bytesStr)%2 == 1 {
		bytesStr = "0" + bytesStr
	}

	val, err := hex.DecodeString(bytesStr)
	if err != nil {
		return errors.New("invalid root")
	}
	copy(r[len(r)-len(val):], val)

	return nil
}

// MarshalJSON implements json.Marshaler.
func (r Root) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, r.String())), nil
}

// Parse converts a string to a root.
func (r *Root) Parse(input string) (*Root, error) {
	if err := r.UnmarshalJSON([]byte(fmt.Sprintf("%q", input))); err != nil {
		return r, err
	}

	return r, nil
}

// MustParse converts a string to a root, panicking on error.
func (r *Root) MustParse(input string) *Root {
	if _, err := r.Parse(input); err != nil {
		panic(err)
	}

	return r
}
