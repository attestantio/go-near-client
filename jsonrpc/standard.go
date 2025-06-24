// Copyright © 2024 Attestant Limited.
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
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"time"

	"github.com/okx/go-wallet-sdk/crypto/base58"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

// service implements Service.
type service struct {
	network string
	address string
	log     zerolog.Logger
	client  *http.Client
}

// GetAccountStakedBalance implements Service.GetAccountStakedBalance.
func (s *service) GetAccountStakedBalance(ctx context.Context, accountID string) (*big.Int, error) {
	if accountID == "" {
		return nil, errors.New("account ID not specified")
	}

	// TODO: Implement RPC call to get staked balance
	return nil, errors.New("not implemented")
}

// CallFunction implements Service.CallFunction.
func (s *service) CallFunction(ctx context.Context, contractID string, methodName string, args []byte) ([]byte, error) {
	if contractID == "" {
		return nil, errors.New("contract ID not specified")
	}
	if methodName == "" {
		return nil, errors.New("method name not specified")
	}

	params := map[string]interface{}{
		"request_type": "call_function",
		"finality":     "final",
		"account_id":   contractID,
		"method_name":  methodName,
		"args_base64":  base64.StdEncoding.EncodeToString(args),
	}

	s.log.Debug().
		Str("contract_id", contractID).
		Str("method_name", methodName).
		Bytes("args", args).
		Interface("params", params).
		Msg("calling contract function")

	result, err := s.makeRPCCall(ctx, "query", params)
	if err != nil {
		return nil, errors.Wrap(err, "failed to call function")
	}

	// Parse the result which contains the actual function result in result.result
	var response struct {
		Result []byte `json:"result"`
	}
	if err := json.Unmarshal(result, &response); err != nil {
		return nil, errors.Wrap(err, "failed to parse function response")
	}

	return response.Result, nil
}

// makeRPCCall makes a JSON-RPC call to the NEAR node
func (s *service) makeRPCCall(ctx context.Context, method string, params interface{}) (json.RawMessage, error) {
	// If params is a string (base64 tx), wrap it in an array
	if base64Str, ok := params.(string); ok {
		params = []string{base64Str}
	}

	request := RpcRequest{
		Version: "2.0",
		Method:  method,
		Params:  params,
		ID:      fmt.Sprintf("%d", time.Now().UnixNano()),
	}

	requestBody, err := json.Marshal(request)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal request")
	}

	req, err := http.NewRequestWithContext(ctx, "POST", s.address, bytes.NewReader(requestBody))
	if err != nil {
		return nil, errors.Wrap(err, "failed to create request")
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to send request")
	}
	defer resp.Body.Close()

	var rpcResp RpcResponse
	if err := json.NewDecoder(resp.Body).Decode(&rpcResp); err != nil {
		return nil, errors.Wrap(err, "failed to decode response")
	}

	if rpcResp.Error != nil {
		s.log.Error().
			Str("error_name", rpcResp.Error.Name).
			Str("error_cause_name", rpcResp.Error.Cause.Name).
			RawJSON("error_cause_info", rpcResp.Error.Cause.Info).
			Int("error_code", rpcResp.Error.Code).
			RawJSON("error_data", rpcResp.Error.Data).
			Str("error_message", rpcResp.Error.Message).
			Msg("RPC error occurred")
		return nil, fmt.Errorf("RPC error: %s - %s (cause: %s)", rpcResp.Error.Name, rpcResp.Error.Message, rpcResp.Error.Cause.Name)
	}

	return rpcResp.Result, nil
}

// GetAccountNonce gets the current nonce for an account
func (s *service) GetAccountNonce(ctx context.Context, accountID string, publicKey string) (uint64, error) {
	params := map[string]interface{}{
		"request_type": "view_access_key",
		"finality":     "final",
		"account_id":   accountID,
		"public_key":   publicKey,
	}

	s.log.Debug().
		Str("account_id", accountID).
		Interface("params", params).
		Msg("getting account nonce")

	result, err := s.makeRPCCall(ctx, "query", params)
	if err != nil {
		return 0, errors.Wrap(err, "failed to get account")
	}

	// Debug log the raw response
	s.log.Debug().
		RawJSON("raw_response", result).
		Msg("received raw response")

	var response struct {
		BlockHash   string      `json:"block_hash"`
		BlockHeight uint64      `json:"block_height"`
		Nonce       uint64      `json:"nonce"`
		Permission  interface{} `json:"permission"` // Can be string "FullAccess" or object
	}

	if err := json.Unmarshal(result, &response); err != nil {
		s.log.Error().
			Err(err).
			Str("result", string(result)).
			Msg("failed to parse account response")
		return 0, errors.Wrap(err, "failed to parse account response")
	}

	nonce := response.Nonce

	s.log.Debug().
		Uint64("nonce", nonce).
		Uint64("block_height", response.BlockHeight).
		Str("block_hash", response.BlockHash).
		Interface("permission", response.Permission).
		Msg("got account nonce")

	return nonce + 1, nil
}

// GetTxStatus gets detailed transaction status
func (s *service) GetTxStatus(ctx context.Context, txHash string, accountID string) (json.RawMessage, error) {
	params := []interface{}{
		txHash,
		accountID,
	}

	s.log.Debug().
		Str("tx_hash", txHash).
		Str("account_id", accountID).
		Interface("params", params).
		Msg("getting transaction status")

	result, err := s.makeRPCCall(ctx, "EXPERIMENTAL_tx_status", params)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get transaction status")
	}

	s.log.Debug().
		RawJSON("result", result).
		Msg("got transaction status")

	return result, nil
}

// SubmitTransaction implements Service.SubmitTransaction.
func (s *service) SubmitTransaction(ctx context.Context, signedTxBase58 string) (string, error) {
	if signedTxBase58 == "" {
		return "", errors.New("signed transaction not specified")
	}

	// Convert base58 to base64
	txBytes := base58.Decode(signedTxBase58)
	signedTxBase64 := base64.StdEncoding.EncodeToString(txBytes)

	s.log.Debug().
		Str("base58", signedTxBase58).
		Str("base64", signedTxBase64).
		Bytes("raw_bytes", txBytes).
		Msg("submitting transaction")

	// Use send_tx with structured params
	params := map[string]interface{}{
		"signed_tx_base64": signedTxBase64,
	}

	result, err := s.makeRPCCall(ctx, "send_tx", params)
	if err != nil {
		// The error is already properly formatted by makeRPCCall
		return "", errors.Wrap(err, "failed to submit transaction")
	}

	// Parse the transaction hash from the result
	var response struct {
		Transaction struct {
			Hash string `json:"hash"`
		} `json:"transaction"`
	}
	if err := json.Unmarshal(result, &response); err != nil {
		return "", errors.Wrap(err, "failed to parse transaction response")
	}

	return response.Transaction.Hash, nil
}
