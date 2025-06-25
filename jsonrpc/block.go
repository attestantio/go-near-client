// Copyright © 2025 Attestant Limited.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
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

// Block returns the block.
func (s *Service) Block(ctx context.Context,
	opts *api.BlockOpts,
) (
	*api.Response[*spec.Block],
	error,
) {
	if err := s.assertIsSynced(ctx); err != nil {
		return nil, err
	}

	if opts == nil {
		return nil, client.ErrNoOptions
	}

	if opts.Finality == "" {
		return nil, client.ErrInvalidOptions
	}

	args := map[string]any{
		"finality": opts.Finality,
	}
	data, err := s.makeRPCCall(ctx, "block", args)
	if err != nil {
		return nil, parseJSONRPCError(err)
	}

	block := &spec.Block{}
	err = json.Unmarshal(data, block)
	if err != nil {
		return nil, errors.Join(errors.New("failed to unmarshal result to block"), client.ErrInconsistentResult)
	}

	return &api.Response[*spec.Block]{
		Data:     block,
		Metadata: map[string]any{},
	}, nil
}
