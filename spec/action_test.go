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

package spec_test

import (
	"encoding/json"
	"testing"

	"github.com/attestantio/go-near-client/spec"
)

func TestFunctionCallActionUnmarshal(t *testing.T) {
	var action spec.Action
	err := json.Unmarshal([]byte(FunctionCallActionExample), &action)
	if err != nil {
		t.Fatalf("Failed to unmarshal FunctionCall action: %v", err)
	}
}

func TestTransferActionUnmarshal(t *testing.T) {
	var action spec.Action
	err := json.Unmarshal([]byte(TransferActionExample), &action)
	if err != nil {
		t.Fatalf("Failed to unmarshal Transfer action: %v", err)
	}
}

func TestDelegateActionUnmarshal(t *testing.T) {
	var action spec.Action
	err := json.Unmarshal([]byte(DelegateActionExample), &action)
	if err != nil {
		t.Fatalf("Failed to unmarshal Delegate action: %v", err)
	}
}

// FunctionCall action example
const FunctionCallActionExample = `{
  "FunctionCall": {
    "args": "eyJjaGFyZ2VfZ2FzX2ZlZSI6dHJ1ZSwic2lnbmF0dXJlIjoiMmJlNDQ4YTBmZmZlZDYyOWU0ZGIwYmE2ZTY0NzhmOWIwZmM2NjlmY2FiZGQ2YzJmZmVmMzgyZDc4MjY0MTA1NSIsIm1pbmluZ190aW1lIjoiMTAwMDIiLCJtYXhfdHMiOiIxNzQyNDMwMDIyNDY0Mjk2OTYwIn0=",
    "deposit": "0",
    "gas": 20000000000000,
    "method_name": "l2_claim"
  }
}`

// Transfer action example
const TransferActionExample = `{
  "Transfer": {
    "deposit": "1574371140228467194032"
  }
}`

// Delegate action example
const DelegateActionExample = `{
  "Delegate": {
    "delegate_action": {
      "actions": [
        {
          "FunctionCall": {
            "args": "eyJjaGFyZ2VfZ2FzX2ZlZSI6dHJ1ZSwic2lnbmF0dXJlIjoiMmJlNDQ4YTBmZmZlZDYyOWU0ZGIwYmE2ZTY0NzhmOWIwZmM2NjlmY2FiZGQ2YzJmZmVmMzgyZDc4MjY0MTA1NSIsIm1pbmluZ190aW1lIjoiMTAwMDIiLCJtYXhfdHMiOiIxNzQyNDMwMDIyNDY0Mjk2OTYwIn0=",
            "deposit": "0",
            "gas": 20000000000000,
            "method_name": "l2_claim"
          }
        }
      ],
      "max_block_height": 142647599,
      "nonce": 130374556003260,
      "public_key": "ed25519:DebSnEAXMBuydDvcRz9ExTyXotEQBWAWrSiLFCipRsfn",
      "receiver_id": "game.hot.tg",
      "sender_id": "reinaldoshiroywuoy872.tg"
    },
    "signature": "ed25519:2SB1xjmadSQaTrBZxwHQfWygHZpbAXnR898b7MnJhXnkDzcDXfD3QJFUzDzRr1upDzQWxiHhz55NPpWuVwhh9rT5"
  }
}`
