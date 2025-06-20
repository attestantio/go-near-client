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

// FieldElementLength is the length of a near field element.
const FieldElementLength = 32

// FieldElement is a 32-byte (actually max 252-bit) near field element.
type FieldElement [FieldElementLength]byte

// String returns the string representation of the field element.
func (f *FieldElement) String() string {
	res := hex.EncodeToString(f[:])
	// Leading 0s not allowed...
	res = strings.TrimLeft(res, "0")
	// ...unless that's all there was.
	if len(res) == 0 {
		res = "0"
	}

	return "0x" + res
}

// Format formats the field element.
func (f *FieldElement) Format(state fmt.State, v rune) {
	format := string(v)
	switch v {
	case 's':
		fmt.Fprint(state, f.String())
	case 'x', 'X':
		if state.Flag('#') {
			format = "#" + format
		}
		fmt.Fprintf(state, "%"+format, f[:])
	default:
		fmt.Fprintf(state, "%"+format, f[:])
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (f *FieldElement) UnmarshalJSON(input []byte) error {
	if len(input) == 0 {
		return errors.New("field element missing")
	}

	if !bytes.HasPrefix(input, []byte{'"', '0', 'x'}) {
		return errors.New("invalid field element prefix")
	}
	if !bytes.HasSuffix(input, []byte{'"'}) {
		return errors.New("invalid field element suffix")
	}

	// Ensure that there are an even number of characters.
	bytesStr := string(input[3 : len(input)-1])
	if len(bytesStr)%2 == 1 {
		bytesStr = "0" + bytesStr
	}

	val, err := hex.DecodeString(bytesStr)
	if err != nil {
		return errors.New("invalid field element")
	}
	copy(f[len(f)-len(val):], val)

	return nil
}

// MarshalJSON implements json.Marshaler.
func (f FieldElement) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, f.String())), nil
}

// Parse converts a string to a field element.
func (f *FieldElement) Parse(input string) (*FieldElement, error) {
	if err := f.UnmarshalJSON([]byte(fmt.Sprintf("%q", input))); err != nil {
		return f, err
	}

	return f, nil
}

// MustParse converts a string to a field element, panicking on error.
func (f *FieldElement) MustParse(input string) *FieldElement {
	if _, err := f.Parse(input); err != nil {
		panic(err)
	}

	return f
}
