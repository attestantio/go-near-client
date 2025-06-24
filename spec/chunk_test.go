package spec_test

import (
	"encoding/json"
	"testing"

	"github.com/attestantio/go-near-client/spec"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const jsonChunk = `{
  "author": "aurora.pool.near",
  "header": {
    "balance_burnt": "0",
    "bandwidth_requests": null,
    "chunk_hash": "8JoY27ADrA534GKWtiU7fkBFRHzgVa6cXJmVjThPzp8u",
    "congestion_info": null,
    "encoded_length": 415,
    "encoded_merkle_root": "8EyxS6SCDScVioygon7sQBqVcfAKFNBZdurnNBjwwf9S",
    "gas_limit": 1000000000000000,
    "gas_used": 223182562500,
    "height_created": 80712125,
    "height_included": 80712125,
    "outcome_root": "B7NuaMuLxHWdxvQhdymaLoh8ksTK1UHHjXBVsbkpcmyc",
    "outgoing_receipts_root": "8s41rye686T2ronWmFE38ji19vgeb6uPxjYMPt8y8pSV",
    "prev_block_hash": "6ZCCmnJGaP6hknzNiFGYzqvoWZzRDkZWQSrvXoH4B7ZL",
    "prev_state_root": "G5o9vYjku2GXy4qobWXnK6uvFLcxPZkmAEcRJQNBYmnm",
    "rent_paid": "0",
    "shard_id": 0,
    "signature": "ed25519:3oLeQC125gsqS5exomq5JiMCJ84Xj1YSbjHVi7mW6XFedx9TpvbuJM85yiTL7yWRT2JxbCsytvonaWpNSLofPNKo",
    "tx_root": "BFm5duguWDn9Axfuazyo6tz1in8H5cNUWSeVaAU7TBp9",
    "validator_proposals": [],
    "validator_reward": "0"
  },
  "receipts": [],
  "transactions": [
    {
      "actions": [
        {
          "FunctionCall": {
            "args": "eyJyZWNlaXZlcl9pZCI6IjI1ZGE5YmYxMzQxYjdlMmFkNzQyMzI1MTdmOTRjYjQ1ODAzNmE1NmU5ZDQ5ZjZkNzU2ZTc5ODM3Mjg3NjAzZWYiLCJhbW91bnQiOiI1NzcwMDAwMDAwMDAwMDAwMDAwIiwibWVtbyI6InN3OnQ6T2o4M200blJ5diJ9",
            "deposit": "1",
            "gas": 14000000000000,
            "method_name": "ft_transfer"
          }
        }
      ],
      "hash": "AmgcGo1ZhJeGWnVtgLhLF5rtbBZMefE4kiBC8P6WTN6k",
      "nonce": 79902626000003,
      "priority_fee": 0,
      "public_key": "ed25519:2wVXmiifSKepdeyR27QtHFjrozuyk1UEoxyLagHXjcsQ",
      "receiver_id": "token.sweat",
      "signature": "ed25519:4KuPzbgfoMmn76CMRoZ95P6PfPwqTpSvionigmSnogKva5khX8YBUncxq3ta1d1os5FDbaHLnwvFHXg7qsdPYMXm",
      "signer_id": "1cd14f68f0db046b2012cbf7e447c576151a6282558816ea063143d9eebd9107"
    }
  ]
}`

const jsonReceipt = `{
  "predecessor_id": "0-relay.hot.tg",
  "priority": 0,
  "receipt": {
    "Action": {
      "actions": [
        {
          "Delegate": {
            "delegate_action": {
              "actions": [
                {
                  "FunctionCall": {
                    "args": "eyJjaGFyZ2VfZ2FzX2ZlZSI6dHJ1ZSwic2lnbmF0dXJlIjoiMmJlNDQ4YTBmZmZlZDYyOWU0ZGIwYmE2ZTY0NzhmOWIwZmM2NjlmY2FiZGQ2YzJmZmVmMzgyZDc4MjY0MTA1NSIsIm1pbmluZ190aW1lIjoiMTAwMDIiLCJtYXhfdHMiOiIxNzQyNDMwMDIyNDY0Mjk2OTYwIn0=",
                    "deposit": "0",
                    "gas": 20000000000000,
                    "method_name": "l2_claim"
                  }
                }
              ],
              "max_block_height": 142647599,
              "nonce": 130374556003260,
              "public_key": "ed25519:DebSnEAXMBuydDvcRz9ExTyXotEQBWAWrSiLFCipRsfn",
              "receiver_id": "game.hot.tg",
              "sender_id": "reinaldoshiroywuoy872.tg"
            },
            "signature": "ed25519:2SB1xjmadSQaTrBZxwHQfWygHZpbAXnR898b7MnJhXnkDzcDXfD3QJFUzDzRr1upDzQWxiHhz55NPpWuVwhh9rT5"
          }
        }
      ],
      "gas_price": "115927408",
      "input_data_ids": [],
      "is_promise_yield": false,
      "output_data_receivers": [],
      "signer_id": "0-relay.hot.tg",
      "signer_public_key": "ed25519:A6jraLRMi8gAU727odbJ9P1t6wniK7w6RToqgojHtE6c"
    }
  },
  "receipt_id": "BD7NGgwoKbyfcGqrvR9WoS5fTqbBH2eV5BuAu79Rbwmi",
  "receiver_id": "reinaldoshiroywuoy872.tg"
}`

func TestChunkUnmarshal(t *testing.T) {
	var chunk spec.Chunk
	err := json.Unmarshal([]byte(jsonChunk), &chunk)
	require.NoError(t, err)

	// Test basic fields
	assert.Equal(t, "aurora.pool.near", chunk.Author)
	assert.Len(t, chunk.Receipts, 0)
	assert.Len(t, chunk.Transactions, 1)

	// Test header fields
	header := chunk.Header
	assert.Equal(t, "0", header.BalanceBurnt)
	assert.Nil(t, header.BandwidthRequests)
	assert.Equal(t, "8JoY27ADrA534GKWtiU7fkBFRHzgVa6cXJmVjThPzp8u", header.ChunkHash)
	assert.Nil(t, header.CongestionInfo)
	assert.Equal(t, 415, header.EncodedLength)
	assert.Equal(t, "8EyxS6SCDScVioygon7sQBqVcfAKFNBZdurnNBjwwf9S", header.EncodedMerkleRoot)
	assert.Equal(t, int64(1000000000000000), header.GasLimit)
	assert.Equal(t, int64(223182562500), header.GasUsed)
	assert.Equal(t, int64(80712125), header.HeightCreated)
	assert.Equal(t, int64(80712125), header.HeightIncluded)
	assert.Equal(t, "B7NuaMuLxHWdxvQhdymaLoh8ksTK1UHHjXBVsbkpcmyc", header.OutcomeRoot)
	assert.Equal(t, "8s41rye686T2ronWmFE38ji19vgeb6uPxjYMPt8y8pSV", header.OutgoingReceiptsRoot)
	assert.Equal(t, "6ZCCmnJGaP6hknzNiFGYzqvoWZzRDkZWQSrvXoH4B7ZL", header.PrevBlockHash)
	assert.Equal(t, "G5o9vYjku2GXy4qobWXnK6uvFLcxPZkmAEcRJQNBYmnm", header.PrevStateRoot)
	assert.Equal(t, "0", header.RentPaid)
	assert.Equal(t, 0, header.ShardID)
	assert.Equal(t, "ed25519:3oLeQC125gsqS5exomq5JiMCJ84Xj1YSbjHVi7mW6XFedx9TpvbuJM85yiTL7yWRT2JxbCsytvonaWpNSLofPNKo", header.Signature)
	assert.Equal(t, "BFm5duguWDn9Axfuazyo6tz1in8H5cNUWSeVaAU7TBp9", header.TxRoot)
	assert.Len(t, header.ValidatorProposals, 0)
	assert.Equal(t, "0", header.ValidatorReward)

	// Test transaction
	tx := chunk.Transactions[0]
	assert.Equal(t, "AmgcGo1ZhJeGWnVtgLhLF5rtbBZMefE4kiBC8P6WTN6k", tx.Hash)
	assert.Equal(t, int64(79902626000003), tx.Nonce)
	assert.Equal(t, int64(0), tx.PriorityFee)
	assert.Equal(t, "ed25519:2wVXmiifSKepdeyR27QtHFjrozuyk1UEoxyLagHXjcsQ", tx.PublicKey)
	assert.Equal(t, "token.sweat", tx.ReceiverID)
	assert.Equal(t, "ed25519:4KuPzbgfoMmn76CMRoZ95P6PfPwqTpSvionigmSnogKva5khX8YBUncxq3ta1d1os5FDbaHLnwvFHXg7qsdPYMXm", tx.Signature)
	assert.Equal(t, "1cd14f68f0db046b2012cbf7e447c576151a6282558816ea063143d9eebd9107", tx.SignerID)
	assert.Len(t, tx.Actions, 1)

	// Test action
	action := tx.Actions[0]
	require.NotNil(t, action.FunctionCall)
	fc := action.FunctionCall
	assert.Equal(t, "eyJyZWNlaXZlcl9pZCI6IjI1ZGE5YmYxMzQxYjdlMmFkNzQyMzI1MTdmOTRjYjQ1ODAzNmE1NmU5ZDQ5ZjZkNzU2ZTc5ODM3Mjg3NjAzZWYiLCJhbW91bnQiOiI1NzcwMDAwMDAwMDAwMDAwMDAwIiwibWVtbyI6InN3OnQ6T2o4M200blJ5diJ9", fc.Args)
	assert.Equal(t, "1", fc.Deposit)
	assert.Equal(t, int64(14000000000000), fc.Gas)
	assert.Equal(t, "ft_transfer", fc.MethodName)
}

func TestChunkMarshal(t *testing.T) {
	// First unmarshal the JSON to get a valid Chunk struct
	var chunk spec.Chunk
	err := json.Unmarshal([]byte(jsonChunk), &chunk)
	require.NoError(t, err)

	// Then marshal it back to JSON
	marshaled, err := json.Marshal(chunk)
	require.NoError(t, err)

	// Unmarshal the marshaled JSON to verify round-trip
	var roundTripChunk spec.Chunk
	err = json.Unmarshal(marshaled, &roundTripChunk)
	require.NoError(t, err)

	// Verify the round-trip preserved all data
	assert.Equal(t, chunk.Author, roundTripChunk.Author)
	assert.Equal(t, chunk.Header, roundTripChunk.Header)
	assert.Equal(t, chunk.Receipts, roundTripChunk.Receipts)
	assert.Equal(t, chunk.Transactions, roundTripChunk.Transactions)
}

func TestReceiptUnmarshal(t *testing.T) {
	var receipt spec.Receipt
	err := json.Unmarshal([]byte(jsonReceipt), &receipt)
	require.NoError(t, err)

	// Test basic fields
	assert.Equal(t, "0-relay.hot.tg", receipt.PredecessorID)
	assert.Equal(t, 0, receipt.Priority)
	assert.Equal(t, "BD7NGgwoKbyfcGqrvR9WoS5fTqbBH2eV5BuAu79Rbwmi", receipt.ReceiptID)
	assert.Equal(t, "reinaldoshiroywuoy872.tg", receipt.ReceiverID)

	// Test receipt data
	require.NotNil(t, receipt.Receipt.Action)
	action := receipt.Receipt.Action
	assert.Equal(t, "115927408", action.GasPrice)
	assert.Len(t, action.InputDataIDs, 0)
	assert.False(t, action.IsPromiseYield)
	assert.Len(t, action.OutputDataReceivers, 0)
	assert.Equal(t, "0-relay.hot.tg", action.SignerID)
	assert.Equal(t, "ed25519:A6jraLRMi8gAU727odbJ9P1t6wniK7w6RToqgojHtE6c", action.SignerPublicKey)
	assert.Len(t, action.Actions, 1)

	// Test delegate action
	actionItem := action.Actions[0]
	require.NotNil(t, actionItem.Delegate)
	delegate := actionItem.Delegate
	assert.Equal(t, "ed25519:2SB1xjmadSQaTrBZxwHQfWygHZpbAXnR898b7MnJhXnkDzcDXfD3QJFUzDzRr1upDzQWxiHhz55NPpWuVwhh9rT5", delegate.Signature)

	// Test delegate action data
	delegateData := delegate.DelegateAction
	assert.Equal(t, int64(142647599), delegateData.MaxBlockHeight)
	assert.Equal(t, int64(130374556003260), delegateData.Nonce)
	assert.Equal(t, "ed25519:DebSnEAXMBuydDvcRz9ExTyXotEQBWAWrSiLFCipRsfn", delegateData.PublicKey)
	assert.Equal(t, "game.hot.tg", delegateData.ReceiverID)
	assert.Equal(t, "reinaldoshiroywuoy872.tg", delegateData.SenderID)
	assert.Len(t, delegateData.Actions, 1)

	// Test function call within delegate action
	delegateActionItem := delegateData.Actions[0]
	require.NotNil(t, delegateActionItem.FunctionCall)
	fc := delegateActionItem.FunctionCall
	assert.Equal(t, "eyJjaGFyZ2VfZ2FzX2ZlZSI6dHJ1ZSwic2lnbmF0dXJlIjoiMmJlNDQ4YTBmZmZlZDYyOWU0ZGIwYmE2ZTY0NzhmOWIwZmM2NjlmY2FiZGQ2YzJmZmVmMzgyZDc4MjY0MTA1NSIsIm1pbmluZ190aW1lIjoiMTAwMDIiLCJtYXhfdHMiOiIxNzQyNDMwMDIyNDY0Mjk2OTYwIn0=", fc.Args)
	assert.Equal(t, "0", fc.Deposit)
	assert.Equal(t, int64(20000000000000), fc.Gas)
	assert.Equal(t, "l2_claim", fc.MethodName)
}

func TestReceiptMarshal(t *testing.T) {
	// First unmarshal the JSON to get a valid Receipt struct
	var receipt spec.Receipt
	err := json.Unmarshal([]byte(jsonReceipt), &receipt)
	require.NoError(t, err)

	// Then marshal it back to JSON
	marshaled, err := json.Marshal(receipt)
	require.NoError(t, err)

	// Unmarshal the marshaled JSON to verify round-trip
	var roundTripReceipt spec.Receipt
	err = json.Unmarshal(marshaled, &roundTripReceipt)
	require.NoError(t, err)

	// Verify the round-trip preserved all data
	assert.Equal(t, receipt.PredecessorID, roundTripReceipt.PredecessorID)
	assert.Equal(t, receipt.Priority, roundTripReceipt.Priority)
	assert.Equal(t, receipt.ReceiptID, roundTripReceipt.ReceiptID)
	assert.Equal(t, receipt.ReceiverID, roundTripReceipt.ReceiverID)
	assert.Equal(t, receipt.Receipt, roundTripReceipt.Receipt)
}

func TestReceiptWithMultipleActions(t *testing.T) {
	// Test receipt with multiple actions
	receiptWithMultipleActions := spec.Receipt{
		PredecessorID: "test.predecessor.near",
		Priority:      1,
		Receipt: spec.ReceiptData{
			Action: &spec.ReceiptAction{
				Actions: []spec.Action{
					{
						Delegate: &spec.DelegateAction{
							DelegateAction: spec.DelegateActionData{
								Actions: []spec.DelegateActionItem{
									{
										FunctionCall: &spec.FunctionCall{
											Args:       "test_args_1",
											Deposit:    "10",
											Gas:        1000000000000,
											MethodName: "test_method_1",
										},
									},
									{
										FunctionCall: &spec.FunctionCall{
											Args:       "test_args_2",
											Deposit:    "20",
											Gas:        2000000000000,
											MethodName: "test_method_2",
										},
									},
								},
								MaxBlockHeight: 1000000,
								Nonce:          123456789,
								PublicKey:      "ed25519:test_public_key",
								ReceiverID:     "test.receiver.near",
								SenderID:       "test.sender.near",
							},
							Signature: "ed25519:test_signature",
						},
					},
				},
				GasPrice:            "100000000",
				InputDataIDs:        []string{"data1", "data2"},
				IsPromiseYield:      true,
				OutputDataReceivers: []interface{}{"receiver1", "receiver2"},
				SignerID:            "test.signer.near",
				SignerPublicKey:     "ed25519:test_signer_public_key",
			},
		},
		ReceiptID:  "test_receipt_id",
		ReceiverID: "test.receiver.near",
	}

	// Marshal to JSON
	marshaled, err := json.Marshal(receiptWithMultipleActions)
	require.NoError(t, err)

	// Unmarshal back
	var unmarshaled spec.Receipt
	err = json.Unmarshal(marshaled, &unmarshaled)
	require.NoError(t, err)

	// Verify all fields are preserved
	assert.Equal(t, receiptWithMultipleActions, unmarshaled)
}

func TestReceiptWithEmptyArrays(t *testing.T) {
	// Test receipt with empty arrays
	emptyReceipt := spec.Receipt{
		PredecessorID: "empty.predecessor.near",
		Priority:      0,
		Receipt: spec.ReceiptData{
			Action: &spec.ReceiptAction{
				Actions:             []spec.Action{},
				GasPrice:            "100000000",
				InputDataIDs:        []string{},
				IsPromiseYield:      false,
				OutputDataReceivers: []interface{}{},
				SignerID:            "empty.signer.near",
				SignerPublicKey:     "ed25519:empty_signer_public_key",
			},
		},
		ReceiptID:  "empty_receipt_id",
		ReceiverID: "empty.receiver.near",
	}

	// Marshal to JSON
	marshaled, err := json.Marshal(emptyReceipt)
	require.NoError(t, err)

	// Unmarshal back
	var unmarshaled spec.Receipt
	err = json.Unmarshal(marshaled, &unmarshaled)
	require.NoError(t, err)

	// Verify all fields are preserved
	assert.Equal(t, emptyReceipt, unmarshaled)
	assert.Len(t, unmarshaled.Receipt.Action.Actions, 0)
	assert.Len(t, unmarshaled.Receipt.Action.InputDataIDs, 0)
	assert.Len(t, unmarshaled.Receipt.Action.OutputDataReceivers, 0)
}

func TestChunkWithReceipts(t *testing.T) {
	// Test chunk with receipts
	chunkWithReceipts := spec.Chunk{
		Author: "test.pool.near",
		Header: spec.ChunkHeader{
			BalanceBurnt:         "1000000000000000000000",
			BandwidthRequests:    nil,
			ChunkHash:            "test_chunk_hash",
			CongestionInfo:       nil,
			EncodedLength:        1000,
			EncodedMerkleRoot:    "test_encoded_merkle_root",
			GasLimit:             1000000000000000,
			GasUsed:              500000000000000,
			HeightCreated:        1000000,
			HeightIncluded:       1000000,
			OutcomeRoot:          "test_outcome_root",
			OutgoingReceiptsRoot: "test_outgoing_receipts_root",
			PrevBlockHash:        "test_prev_block_hash",
			PrevStateRoot:        "test_prev_state_root",
			RentPaid:             "100",
			ShardID:              1,
			Signature:            "ed25519:test_chunk_signature",
			TxRoot:               "test_tx_root",
			ValidatorProposals:   []spec.ValidatorProposal{},
			ValidatorReward:      "200",
		},
		Receipts: []spec.Receipt{
			{
				PredecessorID: "receipt1.predecessor.near",
				Priority:      0,
				Receipt: spec.ReceiptData{
					Action: &spec.ReceiptAction{
						Actions: []spec.Action{
							{
								Delegate: &spec.DelegateAction{
									DelegateAction: spec.DelegateActionData{
										Actions: []spec.DelegateActionItem{
											{
												FunctionCall: &spec.FunctionCall{
													Args:       "test_args",
													Deposit:    "10",
													Gas:        1000000000000,
													MethodName: "test_method",
												},
											},
										},
										MaxBlockHeight: 1000000,
										Nonce:          123456789,
										PublicKey:      "ed25519:test_public_key",
										ReceiverID:     "test.receiver.near",
										SenderID:       "test.sender.near",
									},
									Signature: "ed25519:test_signature",
								},
							},
						},
						GasPrice:            "100000000",
						InputDataIDs:        []string{},
						IsPromiseYield:      false,
						OutputDataReceivers: []interface{}{},
						SignerID:            "test.signer.near",
						SignerPublicKey:     "ed25519:test_signer_public_key",
					},
				},
				ReceiptID:  "test_receipt_id",
				ReceiverID: "test.receiver.near",
			},
		},
		Transactions: []spec.Transaction{},
	}

	// Marshal to JSON
	marshaled, err := json.Marshal(chunkWithReceipts)
	require.NoError(t, err)

	// Unmarshal back
	var unmarshaled spec.Chunk
	err = json.Unmarshal(marshaled, &unmarshaled)
	require.NoError(t, err)

	// Verify all fields are preserved
	assert.Equal(t, chunkWithReceipts, unmarshaled)
	assert.Len(t, unmarshaled.Receipts, 1)
	assert.Len(t, unmarshaled.Transactions, 0)
}
