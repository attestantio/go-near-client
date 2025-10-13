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

package types

import (
	"encoding/json"
	"math/big"
	"strings"

	"github.com/pkg/errors"
)

const (
	yoctoNEARPrecision = 24
	milliNEARPrecision = 4
)

// Balance is a NEAR balance.
type Balance struct {
	Balance *big.Int
}

// String returns the string representation of the balance.
func (b *Balance) String() string {
	return b.Balance.String()
}

// Add returns the addition of the two balances.
func (b *Balance) Add(other *Balance) *Balance {
	return &Balance{Balance: new(big.Int).Add(b.Balance, other.Balance)}
}

// Sub subtracts the provided balance from this balance.
func (b *Balance) Sub(other *Balance) *Balance {
	return &Balance{Balance: new(big.Int).Sub(b.Balance, other.Balance)}
}

// Mul multiplies the two balances together.
func (b *Balance) Mul(other *Balance) *Balance {
	return &Balance{Balance: new(big.Int).Mul(b.Balance, other.Balance)}
}

// Div implements Euclidean division (unlike Go). Divides this balance by the other balance.
func (b *Balance) Div(other *Balance) *Balance {
	return &Balance{Balance: new(big.Int).Div(b.Balance, other.Balance)}
}

// Quo implements truncated division (like Go). Divides this balance by the other balance.
func (b *Balance) Quo(other *Balance) *Balance {
	return &Balance{Balance: new(big.Int).Quo(b.Balance, other.Balance)}
}

// Mod returns the modulus of this balance by the other balance.
func (b *Balance) Mod(other *Balance) *Balance {
	return &Balance{Balance: new(big.Int).Mod(b.Balance, other.Balance)}
}

// Exp returns the value of the exponent of the balance.
func (b *Balance) Exp(other *Balance) *Balance {
	return &Balance{Balance: new(big.Int).Exp(b.Balance, other.Balance, nil)}
}

// Cmp returns the comparison value of two balances.
func (b *Balance) Cmp(other *Balance) int {
	return b.Balance.Cmp(other.Balance)
}

// Abs returns the absolute value of the balance.
func (b *Balance) Abs() *Balance {
	return &Balance{Balance: new(big.Int).Abs(b.Balance)}
}

// Sign returns whether the sign of the balance.
func (b *Balance) Sign() int {
	return b.Balance.Sign()
}

// IsZero returns whether the balance is zero.
func (b *Balance) IsZero() bool {
	return b.Balance.Sign() == 0
}

// UnmarshalJSON implements json.Unmarshaler.
func (b *Balance) UnmarshalJSON(input []byte) error {
	if len(input) == 0 {
		return errors.New("balance missing")
	}

	balanceBigInt := big.NewInt(0)

	_, ok := balanceBigInt.SetString(strings.Trim(string(input), "\""), 10)
	if !ok {
		return errors.New("failed to unmarshal balance")
	}

	b.Balance = balanceBigInt

	return nil
}

// MarshalJSON implements json.Marshaler.
func (b *Balance) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.Balance.String())
}

// PrintNEAR returns the string representation of NEAR.MilliNEAR.
func (b *Balance) PrintNEAR() string {
	var precision big.Int
	precision.Exp(big.NewInt(10), big.NewInt(yoctoNEARPrecision), nil)

	near := new(big.Float).Quo(new(big.Float).SetInt(b.Balance), new(big.Float).SetInt(&precision))

	return near.Text('f', milliNEARPrecision)
}
