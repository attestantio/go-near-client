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

package mock

// Service is a mock Near client service.
type Service struct{}

// New creates a new mock.
func New() (*Service, error) {
	return &Service{}, nil
}

// Name returns the name of the client implementation.
func (*Service) Name() string { return "mock" }

// Address returns the address of the client.
func (*Service) Address() string { return "mock" }
