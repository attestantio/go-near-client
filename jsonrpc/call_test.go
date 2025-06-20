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

package jsonrpc_test

import (
	"context"
	"os"
	"testing"

	"github.com/attestantio/go-near-client/api"
	"github.com/attestantio/go-near-client/jsonrpc"
	"github.com/attestantio/go-near-client/types"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
)

func TestCall(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	tests := []struct {
		name     string
		opts     *api.CallOpts
		expected []types.FieldElement
		err      string
		errCode  int
	}{
		{
			name: "SingleReturnElement",
			opts: &api.CallOpts{
				Block:              "latest",
				Contract:           strToAddress("0x0028c3ac8a0d8e8505486cd2857c309f1557cab0f93d9bb3686704d3cd26af96"),
				EntryPointSelector: strToFieldElement("0x1a35984e05126dbecb7c3bb9929e7dd9106d460c59b1633739a5c733a5fb13b"),
			},
			expected: []types.FieldElement{
				strToFieldElement("0x1ff755e135eac251b4a10bc7aadd11e37a4cab7a552b52d99827c85605ba333"),
			},
		},
		{
			name: "MultipleReturnElements",
			opts: &api.CallOpts{
				Block:              "273028",
				Contract:           strToAddress("0x0590e76a2e65435b7288bf3526cfa5c3ec7748d2f3433a934c931cce62460fc5"),
				EntryPointSelector: strToFieldElement("0x36735aa694184cd8116c479c296d9431cc04a470e0467c07067e4586f647ece"),
			},
			expected: []types.FieldElement{
				strToFieldElement("0xde0b6b3a7640000"),
				strToFieldElement("0x4718f5a0fc34cc1af16a1cdee98ffb20c31f5cd61d6ab07201858f4287c938d"),
				strToFieldElement("0x2264e064ae"),
				strToFieldElement("0x3660fc6334c9485065394f6432933c2f04b4716c67511ba174384c65faebc19"),
				strToFieldElement("0x52bf6ec001452dbace0ca4c7db1b232b8031b5a0ccb117bb47a05569df435de"),
				strToFieldElement("0x12c"),
			},
		},
		{
			name: "CallData",
			opts: &api.CallOpts{
				Block:              "272995",
				Contract:           strToAddress("0x1d3ec552f7875a008783a1a2706c387ef96067469534458dbfb41003ebecb4f"),
				EntryPointSelector: strToFieldElement("0x1af63b1cdc061704af767c81df4fd9eb5a3d989868ff970dd79d5431a89db44"),
				Calldata: []types.FieldElement{
					strToFieldElement("0x714792a41f3651e171c46c93fc53adeb922a414a891cc36d73029d23e99a6ec"),
				},
			},
			expected: []types.FieldElement{
				strToFieldElement("0x391d69afc1b49f01ad8d2e0e8a03756b694dd056fb6645781eb00f33dbd8caf"),
				strToFieldElement("0x4a7876e03402cf253efdb3b17c760ee7349c7ec2876059b83ec2c92ca451e16"),
				strToFieldElement("0x1"),
				strToFieldElement("0xde0b6b3a7640000"),
				strToFieldElement("0x2264e064ae"),
				strToFieldElement("0x2f4d1ee7e060"),
				strToFieldElement("0x0"),
				strToFieldElement("0xca4ef197ec48699c2d4c9f9802770fd08743256759f6f0e500c92e2d397cd7"),
				strToFieldElement("0x4563918244f40000"),
				strToFieldElement("0x0"),
				strToFieldElement("0x9c4"),
			},
		},
		{
			name: "Revert",
			opts: &api.CallOpts{
				Block:              "327585",
				Contract:           strToAddress("0x288b53071caf5a66460cef5673da0e0057e81f2f22c27a69ef1e0222f7a86ba"),
				EntryPointSelector: strToFieldElement("0xcf37a862e5bf34bd0e858865ea02d4ba6db9cc722f3424eb452c94d4ea567f"),
				Calldata: []types.FieldElement{
					strToFieldElement("0x6743648963f634c530c426fe7938007719ccc68169717ad286a9e1ad8753955"),
				},
			},
			err: `40:Contract error {"revert_error":"Execution failed. Failure reason: \"Pool member does not exist\"."}`,
		},
	}

	s, err := jsonrpc.New(ctx,
		jsonrpc.WithLogLevel(zerolog.Disabled),
		jsonrpc.WithAddress(os.Getenv("JSONRPC_ADDRESS")),
		jsonrpc.WithTimeout(timeout),
	)
	require.NoError(t, err)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			response, err := s.Call(ctx, test.opts)
			switch {
			case test.err != "":
				require.ErrorContains(t, err, test.err)
			// case test.errCode != 0:
			//	var apiErr *api.Error
			//	if errors.As(err, &apiErr) {
			//		require.Equal(t, test.errCode, apiErr.StatusCode)
			//	}
			default:
				require.NoError(t, err)
				require.NotNil(t, response)
				require.Equal(t, test.expected, response.Data)
			}
		})
	}
}
