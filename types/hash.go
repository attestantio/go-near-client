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

// HashLength is the length of a hash.
const HashLength = 32

// Hash is a 32-byte hash.
//
//nolint:recvcheck
type Hash [HashLength]byte

var zeroHash = Hash{}

// IsZero returns true if the hash is zero.
func (h Hash) IsZero() bool {
	return bytes.Equal(h[:], zeroHash[:])
}

// String returns the string representation of the hash.
func (h Hash) String() string {
	res := hex.EncodeToString(h[:])
	// Leading 0s not allowed...
	res = strings.TrimLeft(res, "0")
	// ...unless that's all there was.
	if len(res) == 0 {
		res = "0"
	}

	return "0x" + res
}

// Format formats the hash.
func (h Hash) Format(state fmt.State, v rune) {
	format := string(v)
	switch v {
	case 's':
		fmt.Fprint(state, h.String())
	case 'x', 'X':
		if state.Flag('#') {
			format = "#" + format
		}
		fmt.Fprintf(state, "%"+format, h[:])
	default:
		fmt.Fprintf(state, "%"+format, h[:])
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (h *Hash) UnmarshalJSON(input []byte) error {
	if len(input) == 0 {
		return errors.New("hash missing")
	}

	if !bytes.HasPrefix(input, []byte{'"', '0', 'x'}) {
		return errors.New("invalid hash prefix")
	}
	if !bytes.HasSuffix(input, []byte{'"'}) {
		return errors.New("invalid hash suffix")
	}

	// Ensure that there are an even number of characters.
	bytesStr := string(input[3 : len(input)-1])
	if len(bytesStr)%2 == 1 {
		bytesStr = "0" + bytesStr
	}

	val, err := hex.DecodeString(bytesStr)
	if err != nil {
		return errors.New("invalid hash")
	}
	copy(h[len(h)-len(val):], val)

	return nil
}

// MarshalJSON implements json.Marshaler.
func (h Hash) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, h.String())), nil
}

// Parse converts a string to a hash.
func (h *Hash) Parse(input string) (*Hash, error) {
	if err := h.UnmarshalJSON([]byte(fmt.Sprintf("%q", input))); err != nil {
		return h, err
	}

	return h, nil
}

// MustParse converts a string to a hash, panicking on error.
func (h *Hash) MustParse(input string) *Hash {
	if _, err := h.Parse(input); err != nil {
		panic(err)
	}

	return h
}
