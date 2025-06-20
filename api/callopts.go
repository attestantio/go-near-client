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

package api

import "github.com/attestantio/go-near-client/types"

// CallOpts are the options for transaction calls.
type CallOpts struct {
	Common CommonOpts

	// Block is the block for which the data is obtained.
	// It can be a block number, block hash, or one of the special values "latest" or "pending".
	Block types.BlockID

	// Contact is the contract for which the data is obtained.
	Contract types.Address

	// EntryPointSelector defines the entry point for the call.
	EntryPointSelector types.FieldElement

	// Calldata is the data passed to the call.
	Calldata []types.FieldElement
}
