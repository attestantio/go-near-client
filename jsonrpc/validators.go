// Copyright © 2026 Attestant Limited.
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

	client "github.com/attestantio/go-near-client"
	"github.com/attestantio/go-near-client/api"
	"github.com/attestantio/go-near-client/spec"
)

// Validators returns the validators.
func (s *Service) Validators(ctx context.Context,
	opts *api.ValidatorsOpts,
) (
	*api.Response[*spec.Validators],
	error,
) {
	if err := s.assertIsSynced(ctx); err != nil {
		return nil, err
	}

	if opts == nil {
		return nil, client.ErrNoOptions
	}

	if err := validateValidatorsOpts(opts); err != nil {
		return nil, err
	}

	var args any

	switch {
	case opts.Latest:
		args = []any{nil}
	case opts.BlockID != 0:
		args = []any{opts.BlockID}
	case opts.EpochID != "":
		args = map[string]any{
			"epoch_id": opts.EpochID,
		}
	default:
		return nil, client.ErrInvalidOptions
	}

	validators := &spec.Validators{}

	err := s.CallFor(ctx, validators, "validators", args)
	if err != nil {
		return nil, parseJSONRPCError(err)
	}

	return &api.Response[*spec.Validators]{
		Data:     validators,
		Metadata: map[string]any{},
	}, nil
}

// validateValidatorsOpts validates the validators options.
func validateValidatorsOpts(opts *api.ValidatorsOpts) error {
	// Check only one of the options is set.
	nparams := 0
	if opts.Latest {
		nparams++
	}

	if opts.BlockID != 0 {
		nparams++
	}

	if opts.EpochID != "" {
		nparams++
	}

	if nparams != 1 {
		return client.ErrInvalidOptions
	}

	return nil
}
