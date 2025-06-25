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

package spec

// Receipt represents a NEAR receipt
type Receipt struct {
	PredecessorID string      `json:"predecessor_id"`
	Priority      int         `json:"priority"`
	Receipt       ReceiptData `json:"receipt"`
	ReceiptID     string      `json:"receipt_id"`
	ReceiverID    string      `json:"receiver_id"`
}

// ReceiptData represents the actual receipt data
type ReceiptData struct {
	Action *ReceiptAction `json:"Action,omitempty"`
	// Add other receipt types as needed
}

// ReceiptAction represents an action receipt
type ReceiptAction struct {
	Actions             []Action      `json:"actions"`
	GasPrice            string        `json:"gas_price"`
	InputDataIDs        []string      `json:"input_data_ids"`
	IsPromiseYield      bool          `json:"is_promise_yield"`
	OutputDataReceivers []interface{} `json:"output_data_receivers"`
	SignerID            string        `json:"signer_id"`
	SignerPublicKey     string        `json:"signer_public_key"`
}
