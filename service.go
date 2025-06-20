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

package client

import (
	"context"
	"github.com/attestantio/go-near-client/api"
	"github.com/attestantio/go-near-client/types"
)

// Service is the service providing a connection to a NEAR node.
type Service interface {
	// Name returns the name of the node implementation.
	Name() string

	// Address returns the address of the node.
	Address() string
}

// CallProvider is the interface for making calls to the client.
type CallProvider interface {
	// Call makes a call to the client.
	Call(ctx context.Context,
		opts *api.CallOpts,
	) (
		*api.Response[[]types.FieldElement],
		error,
	)
}
