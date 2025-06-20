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

// AddressLength is the length of a near address.
const AddressLength = 32

// Address is a 20-byte near address.
//
//nolint:recvcheck
type Address [AddressLength]byte

var zeroAddress = Address{}

// IsZero returns true if the address is zero.
func (a Address) IsZero() bool {
	return bytes.Equal(a[:], zeroAddress[:])
}

// String returns the string representation of the address.
func (a Address) String() string {
	res := hex.EncodeToString(a[:])
	// Leading 0s not allowed...
	res = strings.TrimLeft(res, "0")
	// ...unless that's all there was.
	if len(res) == 0 {
		res = "0"
	}

	return "0x" + res
}

// Format formats the address.
func (a Address) Format(state fmt.State, v rune) {
	format := string(v)
	switch v {
	case 's':
		fmt.Fprint(state, a.String())
	case 'x', 'X':
		if state.Flag('#') {
			format = "#" + format
		}
		fmt.Fprintf(state, "%"+format, a[:])
	default:
		fmt.Fprintf(state, "%"+format, a[:])
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (a *Address) UnmarshalJSON(input []byte) error {
	if len(input) == 0 {
		return errors.New("address missing")
	}

	if !bytes.HasPrefix(input, []byte{'"', '0', 'x'}) {
		return errors.New("invalid address prefix")
	}
	if !bytes.HasSuffix(input, []byte{'"'}) {
		return errors.New("invalid address suffix")
	}

	// Ensure that there are an even number of characters.
	bytesStr := string(input[3 : len(input)-1])
	if len(bytesStr)%2 == 1 {
		bytesStr = "0" + bytesStr
	}

	val, err := hex.DecodeString(bytesStr)
	if err != nil {
		return errors.New("invalid address")
	}
	copy(a[len(a)-len(val):], val)

	return nil
}

// MarshalJSON implements json.Marshaler.
func (a Address) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%q", a.String())), nil
}

// Parse converts a string to an address.
func (a *Address) Parse(input string) (*Address, error) {
	if err := a.UnmarshalJSON([]byte(fmt.Sprintf("%q", input))); err != nil {
		return a, err
	}

	return a, nil
}

// MustParse converts a string to an address , panicking on error.
func (a *Address) MustParse(input string) *Address {
	if _, err := a.Parse(input); err != nil {
		panic(err)
	}

	return a
}
