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
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Number is a generic number.
type Number uint64

// String returns the string representation of the number.
func (n *Number) String() string {
	res := strconv.FormatUint(uint64(*n), 16)

	// Leading 0s not allowed...
	res = strings.TrimLeft(res, "0")
	// ...unless that's all there was.
	if len(res) == 0 {
		res = "0"
	}

	return "0x" + res
}

// Format formats the number.
func (n *Number) Format(state fmt.State, v rune) {
	format := string(v)
	switch v {
	case 's':
		fmt.Fprint(state, n.String())
	case 'x', 'X':
		if state.Flag('#') {
			format = "#" + format
		}
		fmt.Fprintf(state, "%"+format, n)
	default:
		fmt.Fprintf(state, "%"+format, n)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (n *Number) UnmarshalJSON(input []byte) error {
	if len(input) == 0 {
		return errors.New("number missing")
	}

	if !bytes.HasPrefix(input, []byte{'"', '0', 'x'}) {
		return errors.New("invalid number prefix")
	}
	if !bytes.HasSuffix(input, []byte{'"'}) {
		return errors.New("invalid number suffix")
	}

	// Ensure that there are an even number of characters.
	bytesStr := string(input[3 : len(input)-1])
	if len(bytesStr)%2 == 1 {
		bytesStr = "0" + bytesStr
	}

	val, err := strconv.ParseUint(bytesStr, 16, 64)
	if err != nil {
		return errors.New("invalid number")
	}

	*n = Number(val)

	return nil
}

// MarshalJSON implements json.Marshaler.
func (n Number) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%q", n.String())), nil
}

// Parse converts a string to a number.
func (n *Number) Parse(input string) (*Number, error) {
	if err := n.UnmarshalJSON([]byte(fmt.Sprintf("%q", input))); err != nil {
		return n, err
	}

	return n, nil
}

// MustParse converts a string to a number, panicking on error.
func (n *Number) MustParse(input string) *Number {
	if _, err := n.Parse(input); err != nil {
		panic(err)
	}

	return n
}
