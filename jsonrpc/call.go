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
	"encoding/json"
	"errors"
	"fmt"

	client "github.com/attestantio/go-near-client"
	"github.com/attestantio/go-near-client/api"
	"github.com/attestantio/go-near-client/types"
	"github.com/ybbus/jsonrpc/v2"
)

// Call makes a call to the client.
func (s *Service) Call(ctx context.Context,
	opts *api.CallOpts,
) (
	*api.Response[[]types.FieldElement],
	error,
) {
	if err := s.assertIsSynced(ctx); err != nil {
		return nil, err
	}

	if opts == nil {
		return nil, client.ErrNoOptions
	}
	if opts.Block == "" {
		return nil, errors.Join(errors.New("no block specified"), client.ErrInvalidOptions)
	}

	// TODO: Fill in the content.
	rpcOpts := make(map[string]any)
	request := make(map[string]any)
	request["contract_address"] = opts.Contract.String()
	request["entry_point_selector"] = opts.EntryPointSelector.String()

	calldata := make([]string, 0, len(opts.Calldata))
	for i := range opts.Calldata {
		calldata = append(calldata, opts.Calldata[i].String())
	}
	request["calldata"] = calldata
	rpcOpts["request"] = request
	rpcOpts["block_id"] = opts.Block

	var data []types.FieldElement
	// TODO: Get correct call.
	err := s.client.CallFor(&data, "starknet_call", rpcOpts)
	if err != nil {
		return nil, parseJSONRPCError(err)
	}

	return &api.Response[[]types.FieldElement]{
		Data:     data,
		Metadata: map[string]any{},
	}, nil
}

// parseJSONRPCError potentially adds more information to a JSONRPC error.
func parseJSONRPCError(err error) error {
	var jsonrpcErr *jsonrpc.RPCError
	if errors.As(err, &jsonrpcErr) {
		if jsonrpcErr.Data != nil {
			// Include the additional data in the error.
			additional, marshalErr := json.Marshal(jsonrpcErr.Data)
			if marshalErr != nil {
				return errors.Join(err, client.ErrUnsupportedFormat)
			}
			err = fmt.Errorf("%s %s", err.Error(), string(additional))
		}
	}

	return err
}
