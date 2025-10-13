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

	client "github.com/attestantio/go-near-client"
	"github.com/attestantio/go-near-client/api"
	"github.com/attestantio/go-near-client/spec"
)

// Account retrieves account info from the client.
func (s *Service) Account(ctx context.Context,
	opts *api.AccountOpts,
) (
	*api.Response[*spec.Account],
	error,
) {
	if err := s.assertIsSynced(ctx); err != nil {
		return nil, err
	}

	if opts == nil {
		return nil, client.ErrNoOptions
	}

	if opts.AccountID == "" {
		return nil, errors.Join(errors.New("no account id specified"), client.ErrInvalidOptions)
	}

	if opts.ContractID == "" {
		return nil, errors.Join(errors.New("no contract id specified"), client.ErrInvalidOptions)
	}

	args := map[string]any{
		"account_id": opts.AccountID,
	}

	data, err := s.makeRPCQueryCall("get_account", opts.ContractID, args)
	if err != nil {
		return nil, parseJSONRPCError(err)
	}

	account := &spec.Account{}

	err = json.Unmarshal(data.Result, account)
	if err != nil {
		return nil, errors.Join(errors.New("failed to unmarshal result to account"), client.ErrInconsistentResult)
	}

	return &api.Response[*spec.Account]{
		Data:     account,
		Metadata: map[string]any{},
	}, nil
}
