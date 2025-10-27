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
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/attestantio/go-near-client/api"

	client "github.com/attestantio/go-near-client"
	"github.com/attestantio/go-near-client/spec"
)

// MakeRPCQueryCall makes a JSON-RPC query call to the NEAR node.
func (s *Service) MakeRPCQueryCall(
	ctx context.Context,
	method,
	contractID string,
	block *api.BlockOpts,
	args map[string]any,
) (*spec.CallFunctionResult, error) {
	data := &spec.CallFunctionResult{}
	err := s.CallQueryFor(ctx, data, method, contractID, block, args)

	return data, err
}

// CallQueryFor calls a query method and unmarshals the result into the out parameter.
func (s *Service) CallQueryFor(ctx context.Context,
	out any,
	method,
	contractID string,
	block *api.BlockOpts,
	args map[string]any,
) error {
	argsBytes, err := json.Marshal(args)
	if err != nil {
		return errors.Join(fmt.Errorf("failed to marshal %s args to bytes", method), client.ErrInvalidOptions)
	}

	params := map[string]any{
		"request_type": "call_function",
		"account_id":   contractID,
		"method_name":  method,
		"args_base64":  base64.StdEncoding.EncodeToString(argsBytes),
	}

	if block != nil {
		// Check blockOpts are valid
		if err := validateBlockOpts(block); err != nil {
			return errors.Join(err, client.ErrInvalidOptions)
		}

		switch {
		case block.Finality != "":
			params["finality"] = block.Finality
		case block.BlockID != 0:
			params["block_id"] = block.BlockID
		case block.Hash != "":
			params["block_id"] = block.Hash
		default:
			return client.ErrInvalidOptions
		}
	} else {
		params["finality"] = "final"
	}

	response := &spec.CallFunctionResult{}

	err = s.CallFor(ctx, response, "query", params)
	if err != nil {
		return errors.Join(errors.New("failed to call json rpc"), err)
	}

	if err := json.Unmarshal(response.Result, out); err != nil {
		return client.ErrInconsistentResult
	}

	return nil
}
