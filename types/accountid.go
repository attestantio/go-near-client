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
	"fmt"
	"regexp"
	"strings"

	"github.com/pkg/errors"
)

// AccountID is a string representation of a NEAR account.
//
//nolint:recvcheck
type AccountID string

// String returns the string representation of the Account ID.
func (a AccountID) String() string {
	return string(a)
}

// UnmarshalJSON implements json.Unmarshaler.
func (a *AccountID) UnmarshalJSON(input []byte) error {
	if len(input) == 0 {
		return errors.New("account id missing")
	}

	bytesStr := strings.ReplaceAll(string(input), `"`, ``)

	// Check lengths.
	if len(bytesStr) < 2 {
		return errors.New("account id too short")
	}

	if len(bytesStr) > 64 {
		return errors.New("account id too long")
	}

	// Check regex.
	matchString, err := regexp.MatchString(`^(([a-z\d]+[\-_])*[a-z\d]+\.)*([a-z\d]+[\-_])*[a-z\d]+$`, bytesStr)
	if err != nil {
		return errors.Wrap(err, "account id failed to compile regexp")
	}

	if !matchString {
		return errors.New("account id did not match regex")
	}

	*a = AccountID(bytesStr)

	return nil
}

// MarshalJSON implements json.Marshaler.
func (a AccountID) MarshalJSON() ([]byte, error) {
	if len(a) == 0 {
		return nil, errors.New("account id missing")
	}

	// Check lengths.
	if len(a) < 2 {
		return nil, errors.New("account id too short")
	}

	if len(a) > 64 {
		return nil, errors.New("account id too long")
	}

	// Check regex.
	matchString, err := regexp.MatchString(`^(([a-z\d]+[\-_])*[a-z\d]+\.)*([a-z\d]+[\-_])*[a-z\d]+$`, string(a))
	if err != nil {
		return nil, errors.Wrap(err, "account id failed to compile regexp")
	}

	if !matchString {
		return nil, errors.New("account id did not match regex")
	}

	return []byte(fmt.Sprintf(`"%s"`, a)), nil
}
