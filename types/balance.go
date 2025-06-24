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

// Balance is a NEAR balance.
type Balance struct {
	NEAR      uint64
	MilliNEAR uint64
	Raw       []byte
}

// String returns the string representation of the balance.
func (b *Balance) String() string {
	return strings.ReplaceAll(string(b.Raw), `"`, "")
}

// PrintNEAR returns the string representation of NEAR.MilliNEAR.
func (b *Balance) PrintNEAR() string {
	return fmt.Sprintf(`%d.%d`, b.NEAR, b.MilliNEAR)
}

// UnmarshalJSON implements json.Unmarshaler.
func (b *Balance) UnmarshalJSON(input []byte) error {
	if len(input) == 0 {
		return errors.New("balance missing")
	}
	balance := Balance{}
	bytesStr := strings.ReplaceAll(string(input), `"`, "")
	bytesArr := []byte(bytesStr)
	yoctoIndex := 24
	if len(bytesArr) <= yoctoIndex {
		balance.NEAR = 0
	} else {
		shortVal := bytesArr[:len(bytesArr)-yoctoIndex]
		val, err := strconv.ParseUint(string(shortVal), 10, 64)
		if err != nil {
			return errors.New("invalid NEAR number")
		}
		balance.NEAR = val
	}

	milliIndex := 20
	if len(bytesArr) <= milliIndex {
		balance.MilliNEAR = 0
	} else {
		shortVal := bytesArr[len(bytesArr)-yoctoIndex : len(bytesArr)-milliIndex]
		val, err := strconv.ParseUint(string(shortVal), 10, 64)
		if err != nil {
			return errors.New("invalid MilliNEAR number")
		}
		balance.MilliNEAR = val
	}

	balance.Raw = input

	*b = balance

	return nil
}

// MarshalJSON implements json.Marshaler.
func (b Balance) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%q", b.String())), nil
}

// Parse converts a string to a number.
func (b *Balance) Parse(input string) (*Balance, error) {
	if err := b.UnmarshalJSON([]byte(fmt.Sprintf("%q", input))); err != nil {
		return b, err
	}

	return b, nil
}

// MustParse converts a string to a number, panicking on error.
func (b *Balance) MustParse(input string) *Balance {
	if _, err := b.Parse(input); err != nil {
		panic(err)
	}

	return b
}
