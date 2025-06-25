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

type Action struct {
	Delegate     *DelegateAction `json:"Delegate,omitempty"`
	Transfer     *TransferAction `json:"Transfer,omitempty"`
	FunctionCall *FunctionCall   `json:"FunctionCall,omitempty"`
}

// FunctionCall represents a function call action
type FunctionCall struct {
	Args       string `json:"args"`
	Deposit    string `json:"deposit"`
	Gas        int64  `json:"gas"`
	MethodName string `json:"method_name"`
}

// TransferAction represents a transfer action
type TransferAction struct {
	Deposit string `json:"deposit"`
}

// DelegateAction represents a delegate action
type DelegateAction struct {
	DelegateAction DelegateActionData `json:"delegate_action"`
	Signature      string             `json:"signature"`
}

// DelegateActionData represents the data within a delegate action
type DelegateActionData struct {
	Actions        []DelegateActionItem `json:"actions"`
	MaxBlockHeight int64                `json:"max_block_height"`
	Nonce          int64                `json:"nonce"`
	PublicKey      string               `json:"public_key"`
	ReceiverID     string               `json:"receiver_id"`
	SenderID       string               `json:"sender_id"`
}

// DelegateActionItem represents an individual action within a delegate action
type DelegateActionItem struct {
	FunctionCall *FunctionCall `json:"FunctionCall,omitempty"`
}
