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

package jsonrpc

import "encoding/json"

// RPCRequest represents a JSON-RPC request.
type RPCRequest struct {
	Version string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  any    `json:"params"`
	ID      string `json:"id"`
}

// RPCResponse represents a JSON-RPC response.
type RPCResponse struct {
	Version string          `json:"jsonrpc"`
	Result  json.RawMessage `json:"result,omitempty"`
	Error   *RPCError       `json:"error,omitempty"`
	ID      string          `json:"id"`
}

// RPCError represents a JSON-RPC error.
type RPCError struct {
	Name    string          `json:"name"`
	Cause   Cause           `json:"cause"`
	Code    int             `json:"code"`
	Data    json.RawMessage `json:"data"`
	Message string          `json:"message"`
}

// Cause represents a JSON-RPC error cause.
type Cause struct {
	Info json.RawMessage `json:"info"`
	Name string          `json:"name"`
}
