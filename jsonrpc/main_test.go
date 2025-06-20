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

package jsonrpc_test

import (
	"encoding/hex"
	"encoding/json"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/attestantio/go-near-client/types"
	"github.com/rs/zerolog"
)

// timeout for tests.
var timeout = 60 * time.Second

func TestMain(m *testing.M) {
	zerolog.SetGlobalLevel(zerolog.TraceLevel)
	if os.Getenv("JSONRPC_ADDRESS") != "" {
		os.Exit(m.Run())
	}
}

// strToHash is a helper to create a hash given a string representation.
func strToHash(input string) types.Hash {
	bytes, err := hex.DecodeString(strings.TrimPrefix(input, "0x"))
	if err != nil {
		panic(err)
	}

	var res types.Hash
	copy(res[:], bytes)

	return res
}

// strToAddress is a helper to create an address given a string representation.
func strToAddress(input string) types.Address {
	var res types.Address
	if err := json.Unmarshal([]byte(`"`+input+`"`), &res); err != nil {
		panic(err)
	}

	return res
}

// strToFieldElement is a helper to create a field element given a string representation.
func strToFieldElement(input string) types.FieldElement {
	var res types.FieldElement
	if err := json.Unmarshal([]byte(`"`+input+`"`), &res); err != nil {
		panic(err)
	}

	return res
}

// strToBytes is a helper to create a byte slice given a string representation.
func strToBytes(input string) []byte {
	bytes, err := hex.DecodeString(strings.TrimPrefix(input, "0x"))
	if err != nil {
		panic(err)
	}

	return bytes
}
