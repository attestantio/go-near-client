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

package client_test

import (
	"testing"

	client "github.com/attestantio/go-near-client"
	"github.com/attestantio/go-near-client/jsonrpc"
	"github.com/attestantio/go-near-client/mock"
	"github.com/stretchr/testify/assert"
)

// TestMockImplementsInterfaces verifies that the mock service implements all required interfaces.
func TestMockImplementsInterfaces(t *testing.T) {
	s, err := mock.New()
	if err != nil {
		t.Fatal(err)
	}

	// Test that mock implements all interfaces
	assert.Implements(t, (*client.Service)(nil), s, "mock should implement Service")
	assert.Implements(t, (*client.AccountProvider)(nil), s, "mock should implement AccountProvider")
	assert.Implements(t, (*client.BlockProvider)(nil), s, "mock should implement BlockProvider")
	assert.Implements(t, (*client.StatusProvider)(nil), s, "mock should implement StatusProvider")
	assert.Implements(t, (*client.SyncProvider)(nil), s, "mock should implement SyncProvider")
	assert.Implements(t, (*client.GenesisProvider)(nil), s, "mock should implement GenesisProvider")
	assert.Implements(t, (*client.ChunkProvider)(nil), s, "mock should implement ChunkProvider")
	assert.Implements(t, (*client.QueryProvider)(nil), s, "mock should implement QueryProvider")
}

// TestJSONRPCImplementsInterfaces verifies that the jsonrpc service implements all required interfaces.
func TestJSONRPCImplementsInterfaces(t *testing.T) {
	// We can't create a real jsonrpc service without a valid endpoint,
	// so we just check the types implement the interfaces at compile time.
	var s *jsonrpc.Service

	// Test that jsonrpc implements all interfaces
	assert.Implements(t, (*client.Service)(nil), s, "jsonrpc should implement Service")
	assert.Implements(t, (*client.AccountProvider)(nil), s, "jsonrpc should implement AccountProvider")
	assert.Implements(t, (*client.BlockProvider)(nil), s, "jsonrpc should implement BlockProvider")
	assert.Implements(t, (*client.StatusProvider)(nil), s, "jsonrpc should implement StatusProvider")
	assert.Implements(t, (*client.SyncProvider)(nil), s, "jsonrpc should implement SyncProvider")
	assert.Implements(t, (*client.GenesisProvider)(nil), s, "jsonrpc should implement GenesisProvider")
	assert.Implements(t, (*client.ChunkProvider)(nil), s, "jsonrpc should implement ChunkProvider")
	assert.Implements(t, (*client.QueryProvider)(nil), s, "jsonrpc should implement QueryProvider")
}

// TestClientCompatibility ensures both implementations can be used interchangeably.
func TestClientCompatibility(t *testing.T) {
	// This function demonstrates that both mock and jsonrpc services
	// can be used through the same interfaces.
	mockService, err := mock.New()
	if err != nil {
		t.Fatal(err)
	}

	// Test using interface types
	var service client.Service = mockService
	assert.Equal(t, "mock", service.Name())
	assert.Equal(t, "mock", service.Address())

	var accountProvider client.AccountProvider = mockService
	assert.NotNil(t, accountProvider)

	var blockProvider client.BlockProvider = mockService
	assert.NotNil(t, blockProvider)

	var statusProvider client.StatusProvider = mockService
	assert.NotNil(t, statusProvider)

	var syncProvider client.SyncProvider = mockService
	assert.NotNil(t, syncProvider)

	var genesisProvider client.GenesisProvider = mockService
	assert.NotNil(t, genesisProvider)

	var chunkProvider client.ChunkProvider = mockService
	assert.NotNil(t, chunkProvider)

	var queryProvider client.QueryProvider = mockService
	assert.NotNil(t, queryProvider)
}
