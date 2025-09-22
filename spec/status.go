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

package spec

// Status represents a status.
type Status struct {
	ChainID               string      `json:"chain_id"`
	GenesisHash           string      `json:"genesis_hash"`
	LatestProtocolVersion int         `json:"latest_protocol_version"`
	NodeKey               *string     `json:"node_key"`
	NodePublicKey         string      `json:"node_public_key"`
	ProtocolVersion       int         `json:"protocol_version"`
	RPCAddr               string      `json:"rpc_addr"`
	SyncInfo              *SyncInfo   `json:"sync_info"`
	UptimeSec             int64       `json:"uptime_sec"`
	ValidatorAccountID    *string     `json:"validator_account_id"`
	ValidatorPublicKey    *string     `json:"validator_public_key"`
	Validators            []Validator `json:"validators"`
	Version               Version     `json:"version"`
}

// Validator represents a validator in the network.
type Validator struct {
	AccountID string `json:"account_id"`
	IsSlashed bool   `json:"is_slashed"`
}

// Version contains information about the node's version.
type Version struct {
	Build        string `json:"build"`
	Commit       string `json:"commit"`
	RustcVersion string `json:"rustc_version"`
	Version      string `json:"version"`
}
