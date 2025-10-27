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

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	client "github.com/attestantio/go-near-client"
	"github.com/pkg/errors"
)

// MakeRPCCall makes a JSON-RPC call to the NEAR node.
func (s *Service) MakeRPCCall(ctx context.Context, method string, params any) (json.RawMessage, error) {
	// If params is a string (base64 tx), wrap it in an array
	if base64Str, ok := params.(string); ok {
		params = []string{base64Str}
	}

	request := RPCRequest{
		Version: "2.0",
		Method:  method,
		Params:  params,
		ID:      fmt.Sprintf("%d", time.Now().UnixNano()),
	}

	requestBody, err := json.Marshal(request)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal request")
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, s.address, bytes.NewReader(requestBody))
	if err != nil {
		return nil, errors.Wrap(err, "failed to create request")
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to send request")
	}
	defer resp.Body.Close()

	var rpcResp RPCResponse
	if err := json.NewDecoder(resp.Body).Decode(&rpcResp); err != nil {
		// The RPC didnot return a valid JSON response. Just return the status code.
		if resp.StatusCode != http.StatusOK {
			s.log.Debug().
				Int("status_code", resp.StatusCode).
				Str("method", method).
				Interface("params", params).Msg("json rpc request failed")

			return nil, fmt.Errorf("json rpc request failed with status code %d", resp.StatusCode)
		}

		return nil, errors.Wrap(err, "failed to decode response")
	}

	if rpcResp.Error != nil {
		return nil, formatJSONRPCError(rpcResp.Error)
	}

	return rpcResp.Result, nil
}

// CallFor makes a JSON-RPC call to the NEAR node and unmarshals the result into the out parameter.
func (s *Service) CallFor(ctx context.Context, out any, method string, params any) error {
	rpcResp, err := s.MakeRPCCall(ctx, method, params)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(rpcResp, out); err != nil {
		return client.ErrInconsistentResult
	}

	return nil
}

func formatJSONRPCError(err *RPCError) error {
	if err.Name == "HANDLER_ERROR" && err.Cause.Name == "UNKNOWN_BLOCK" {
		return ErrBlockNotFound
	}

	return fmt.Errorf("RPC error: %s - %s (cause: %s)",
		err.Name, err.Message, err.Cause.Name)
}
