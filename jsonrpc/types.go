package jsonrpc

import "encoding/json"

// RpcRequest represents a JSON-RPC request
type RpcRequest struct {
	Version string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
	ID      string      `json:"id"`
}

// RpcResponse represents a JSON-RPC response
type RpcResponse struct {
	Version string          `json:"jsonrpc"`
	Result  json.RawMessage `json:"result,omitempty"`
	Error   *RpcError       `json:"error,omitempty"`
	ID      string          `json:"id"`
}

// RpcError represents a JSON-RPC error
type RpcError struct {
	Name  string `json:"name"`
	Cause struct {
		Info json.RawMessage `json:"info"`
		Name string          `json:"name"`
	} `json:"cause"`
	Code    int             `json:"code"`
	Data    json.RawMessage `json:"data"`
	Message string          `json:"message"`
}
