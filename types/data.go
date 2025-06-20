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

// Data is arbitrary-length binary data.
type Data []byte

// String returns the string representation of the data.
func (d *Data) String() string {
	if len(*d) == 0 {
		return "0x"
	}

	res := hex.EncodeToString(*d)
	// Leading 0s not allowed...
	res = strings.TrimLeft(res, "0")
	// ...unless that's all there was.
	if len(res) == 0 {
		res = "0"
	}

	return "0x" + res
}

// Format formats the data.
func (d *Data) Format(state fmt.State, v rune) {
	format := string(v)
	switch v {
	case 's':
		fmt.Fprint(state, d.String())
	case 'x', 'X':
		if state.Flag('#') {
			format = "#" + format
		}
		fmt.Fprintf(state, "%"+format, d)
	default:
		fmt.Fprintf(state, "%"+format, d)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (d *Data) UnmarshalJSON(input []byte) error {
	if len(input) == 0 {
		return errors.New("data missing")
	}

	if !bytes.HasPrefix(input, []byte{'"', '0', 'x'}) {
		return errors.New("invalid data prefix")
	}
	if !bytes.HasSuffix(input, []byte{'"'}) {
		return errors.New("invalid data suffix")
	}

	// Ensure that there are an even number of characters.
	bytesStr := string(input[3 : len(input)-1])
	if len(bytesStr)%2 == 1 {
		bytesStr = "0" + bytesStr
	}

	res, err := hex.DecodeString(bytesStr)
	if err != nil {
		return errors.New("invalid data")
	}

	*d = Data(res)

	return nil
}

// MarshalJSON implements json.Marshaler.
func (d Data) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%q", d.String())), nil
}

// Parse converts a string to data.
func (d *Data) Parse(input string) (*Data, error) {
	if err := d.UnmarshalJSON([]byte(fmt.Sprintf("%q", input))); err != nil {
		return d, err
	}

	return d, nil
}

// MustParse converts a string to data, panicking on error.
func (d *Data) MustParse(input string) *Data {
	if _, err := d.Parse(input); err != nil {
		panic(err)
	}

	return d
}
