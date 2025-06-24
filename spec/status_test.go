package spec_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/attestantio/go-near-client/spec"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const jsonStatus = `{
  "chain_id": "mainnet",
  "genesis_hash": "EPnLgE7iEq9s7yTkos96M3cWymH5avBAPm3qx3NXqR8H",
  "latest_protocol_version": 77,
  "node_key": null,
  "node_public_key": "ed25519:4V9pPpYNcFDBh3EbdMEjwiajxgLPQjDn2Gvvos1UqKCj",
  "protocol_version": 77,
  "rpc_addr": "0.0.0.0:3030",
  "sync_info": {
    "earliest_block_hash": "CJGN6UEPsY7xMji5PmWwhU5Y5NQZiqRyXEZcYE8kUyMm",
    "earliest_block_height": 152207510,
    "earliest_block_time": "2025-06-22T15:43:15.846327398Z",
    "epoch_id": "YM6foK5rXbkThBweZ6ZYubCM6WijZdnSgPgyKLKKHoz",
    "epoch_start_height": 152380311,
    "latest_block_hash": "FSezxCQkZw3MqUJmh25m5TkcMei167jjRAyAofRkU189",
    "latest_block_height": 152416780,
    "latest_block_time": "2025-06-24T04:13:36.579430453Z",
    "latest_state_root": "BmsxonQr78hKGbULNv3uZSjSwoEL4tpWPdERBQq6fwBK",
    "syncing": false
  },
  "uptime_sec": 2473536,
  "validator_account_id": null,
  "validator_public_key": null,
  "validators": [
    { "account_id": "astro-stakers.poolv1.near", "is_slashed": false },
    { "account_id": "zavodil.poolv1.near", "is_slashed": false },
    { "account_id": "figment.poolv1.near", "is_slashed": false },
    { "account_id": "ledgerbyfigment.poolv1.near", "is_slashed": false },
    { "account_id": "luganodes.pool.near", "is_slashed": false },
    { "account_id": "sumerian.poolv1.near", "is_slashed": false },
    { "account_id": "sofarsonear.poolv1.near", "is_slashed": false },
    { "account_id": "kiln-1.poolv1.near", "is_slashed": false },
    { "account_id": "pinnacle1.poolv1.near", "is_slashed": false },
    { "account_id": "twinstake.poolv1.near", "is_slashed": false },
    {
      "account_id": "staking_yes_protocol1.poolv1.near",
      "is_slashed": false
    },
    { "account_id": "liver.pool.near", "is_slashed": false },
    { "account_id": "stake1.poolv1.near", "is_slashed": false },
    { "account_id": "epic.poolv1.near", "is_slashed": false },
    { "account_id": "foundry.poolv1.near", "is_slashed": false },
    { "account_id": "bisontrails2.poolv1.near", "is_slashed": false },
    { "account_id": "binancenode1.poolv1.near", "is_slashed": false },
    { "account_id": "nearone.pool.near", "is_slashed": false },
    {
      "account_id": "aca87218e28c41f5a693dee3dff12238.poolv1.near",
      "is_slashed": false
    },
    { "account_id": "p2p-org.poolv1.near", "is_slashed": false },
    { "account_id": "cosmose.poolv1.near", "is_slashed": false },
    { "account_id": "kiln.poolv1.near", "is_slashed": false },
    { "account_id": "here.poolv1.near", "is_slashed": false },
    { "account_id": "stakin.poolv1.near", "is_slashed": false },
    { "account_id": "republic.poolv1.near", "is_slashed": false },
    { "account_id": "marcus.pool.near", "is_slashed": false },
    { "account_id": "flipside.pool.near", "is_slashed": false },
    { "account_id": "galaxydigital.poolv1.near", "is_slashed": false },
    { "account_id": "falcon.pool.near", "is_slashed": false },
    { "account_id": "bisontrails.poolv1.near", "is_slashed": false },
    { "account_id": "rekt.poolv1.near", "is_slashed": false },
    { "account_id": "d1.poolv1.near", "is_slashed": false },
    { "account_id": "macrodatarefinement.poolv1.near", "is_slashed": false },
    { "account_id": "near-fans.poolv1.near", "is_slashed": false },
    { "account_id": "everstake.poolv1.near", "is_slashed": false },
    { "account_id": "future_is_near.poolv1.near", "is_slashed": false },
    { "account_id": "blockdaemon.poolv1.near", "is_slashed": false },
    { "account_id": "nansen.poolv1.near", "is_slashed": false },
    { "account_id": "pandora.poolv1.near", "is_slashed": false },
    { "account_id": "allnodes.poolv1.near", "is_slashed": false },
    { "account_id": "northernlights.poolv1.near", "is_slashed": false },
    { "account_id": "chorusone.poolv1.near", "is_slashed": false },
    { "account_id": "sweat_validator.poolv1.near", "is_slashed": false },
    { "account_id": "nearfans.poolv1.near", "is_slashed": false },
    { "account_id": "dragonfly.poolv1.near", "is_slashed": false },
    { "account_id": "bitcoinsuisse.poolv1.near", "is_slashed": false },
    { "account_id": "aurora.pool.near", "is_slashed": false },
    { "account_id": "okx-earn.poolv1.near", "is_slashed": false },
    { "account_id": "x.poolv1.near", "is_slashed": false },
    { "account_id": "trust-nodes.poolv1.near", "is_slashed": false },
    { "account_id": "lux.poolv1.near", "is_slashed": false },
    { "account_id": "erm.poolv1.near", "is_slashed": false },
    { "account_id": "buildlinks.poolv1.near", "is_slashed": false },
    { "account_id": "anonymous.poolv1.near", "is_slashed": false },
    { "account_id": "dsrvlabs.poolv1.near", "is_slashed": false },
    { "account_id": "baziliknear.poolv1.near", "is_slashed": false },
    { "account_id": "openshards.poolv1.near", "is_slashed": false },
    { "account_id": "cryptium.poolv1.near", "is_slashed": false },
    { "account_id": "staking4all.poolv1.near", "is_slashed": false },
    { "account_id": "brea.poolv1.near", "is_slashed": false },
    { "account_id": "stakesabai.poolv1.near", "is_slashed": false },
    { "account_id": "meteor.poolv1.near", "is_slashed": false },
    { "account_id": "moonlet.poolv1.near", "is_slashed": false },
    { "account_id": "staked.poolv1.near", "is_slashed": false },
    { "account_id": "kaiching.poolv1.near", "is_slashed": false },
    { "account_id": "lunanova.poolv1.near", "is_slashed": false },
    { "account_id": "masternode24.poolv1.near", "is_slashed": false },
    { "account_id": "colossus.poolv1.near", "is_slashed": false },
    {
      "account_id": "readylayerone_staking.poolv1.near",
      "is_slashed": false
    },
    { "account_id": "smart-stake.poolv1.near", "is_slashed": false },
    { "account_id": "hapi.poolv1.near", "is_slashed": false },
    { "account_id": "nearkoreahub.poolv1.near", "is_slashed": false },
    { "account_id": "stardust.poolv1.near", "is_slashed": false },
    { "account_id": "polkachu.poolv1.near", "is_slashed": false },
    { "account_id": "dexagon.poolv1.near", "is_slashed": false },
    { "account_id": "stakely_io.poolv1.near", "is_slashed": false },
    { "account_id": "qbit.poolv1.near", "is_slashed": false },
    { "account_id": "hb436_pool.poolv1.near", "is_slashed": false },
    { "account_id": "fresh.poolv1.near", "is_slashed": false },
    { "account_id": "avado.poolv1.near", "is_slashed": false },
    { "account_id": "01node.poolv1.near", "is_slashed": false },
    { "account_id": "lavenderfive.poolv1.near", "is_slashed": false },
    { "account_id": "autostake.poolv1.near", "is_slashed": false },
    { "account_id": "pandateam.poolv1.near", "is_slashed": false },
    { "account_id": "delightlabs.pool.near", "is_slashed": false },
    { "account_id": "modernlion.poolv1.near", "is_slashed": false },
    { "account_id": "cryptogarik.poolv1.near", "is_slashed": false },
    { "account_id": "atomic-nodes.poolv1.near", "is_slashed": false },
    { "account_id": "gfi-validator.poolv1.near", "is_slashed": false },
    { "account_id": "hashquark.poolv1.near", "is_slashed": false },
    { "account_id": "lionstake.poolv1.near", "is_slashed": false },
    { "account_id": "oe.poolv1.near", "is_slashed": false },
    { "account_id": "pangdao.poolv1.near", "is_slashed": false },
    { "account_id": "galactic.poolv1.near", "is_slashed": false },
    { "account_id": "wackazong.poolv1.near", "is_slashed": false },
    { "account_id": "namdokmai.poolv1.near", "is_slashed": false },
    {
      "account_id": "optimusvalidatornetwork.poolv1.near",
      "is_slashed": false
    },
    { "account_id": "oharanodes.poolv1.near", "is_slashed": false },
    { "account_id": "2pilot.poolv1.near", "is_slashed": false },
    { "account_id": "rhea-validator.poolv1.near", "is_slashed": false }
  ],
  "version": {
    "build": "2.6.3",
    "commit": "680b27eb0105655345535a20623aaeb3b49b82e0",
    "rustc_version": "1.85.0",
    "version": "2.6.3"
  }
}`

func TestStatusUnmarshal(t *testing.T) {
	var status spec.Status
	err := json.Unmarshal([]byte(jsonStatus), &status)
	require.NoError(t, err)

	// Test basic fields
	assert.Equal(t, "mainnet", status.ChainID)
	assert.Equal(t, "EPnLgE7iEq9s7yTkos96M3cWymH5avBAPm3qx3NXqR8H", status.GenesisHash)
	assert.Equal(t, 77, status.LatestProtocolVersion)
	assert.Equal(t, "ed25519:4V9pPpYNcFDBh3EbdMEjwiajxgLPQjDn2Gvvos1UqKCj", status.NodePublicKey)
	assert.Equal(t, 77, status.ProtocolVersion)
	assert.Equal(t, "0.0.0.0:3030", status.RPCAddr)
	assert.Equal(t, int64(2473536), status.UptimeSec)

	// Test nullable fields
	assert.Nil(t, status.NodeKey)
	assert.Nil(t, status.ValidatorAccountID)
	assert.Nil(t, status.ValidatorPublicKey)

	// Test SyncInfo
	assert.Equal(t, "CJGN6UEPsY7xMji5PmWwhU5Y5NQZiqRyXEZcYE8kUyMm", status.SyncInfo.EarliestBlockHash)
	assert.Equal(t, int64(152207510), status.SyncInfo.EarliestBlockHeight)
	assert.Equal(t, "YM6foK5rXbkThBweZ6ZYubCM6WijZdnSgPgyKLKKHoz", status.SyncInfo.EpochID)
	assert.Equal(t, int64(152380311), status.SyncInfo.EpochStartHeight)
	assert.Equal(t, "FSezxCQkZw3MqUJmh25m5TkcMei167jjRAyAofRkU189", status.SyncInfo.LatestBlockHash)
	assert.Equal(t, int64(152416780), status.SyncInfo.LatestBlockHeight)
	assert.Equal(t, "BmsxonQr78hKGbULNv3uZSjSwoEL4tpWPdERBQq6fwBK", status.SyncInfo.LatestStateRoot)
	assert.False(t, status.SyncInfo.Syncing)

	earliestBlockTime, err := time.Parse(time.RFC3339, "2025-06-22T15:43:15.846327398Z")
	require.NoError(t, err)
	assert.Equal(t, earliestBlockTime, status.SyncInfo.EarliestBlockTime)

	latestBlockTime, err := time.Parse(time.RFC3339, "2025-06-24T04:13:36.579430453Z")
	require.NoError(t, err)
	assert.Equal(t, latestBlockTime, status.SyncInfo.LatestBlockTime)

	// Test validators
	assert.Len(t, status.Validators, 100)
	assert.Equal(t, "astro-stakers.poolv1.near", status.Validators[0].AccountID)
	assert.False(t, status.Validators[0].IsSlashed)
	assert.Equal(t, "rhea-validator.poolv1.near", status.Validators[99].AccountID)
	assert.False(t, status.Validators[99].IsSlashed)

	// Test version
	assert.Equal(t, "2.6.3", status.Version.Build)
	assert.Equal(t, "680b27eb0105655345535a20623aaeb3b49b82e0", status.Version.Commit)
	assert.Equal(t, "1.85.0", status.Version.RustcVersion)
	assert.Equal(t, "2.6.3", status.Version.Version)
}

func TestStatusMarshal(t *testing.T) {
	// First unmarshal the JSON to get a valid Status struct
	var status spec.Status
	err := json.Unmarshal([]byte(jsonStatus), &status)
	require.NoError(t, err)

	// Then marshal it back to JSON
	marshaled, err := json.Marshal(status)
	require.NoError(t, err)

	// Unmarshal the marshaled JSON to verify round-trip
	var roundTripStatus spec.Status
	err = json.Unmarshal(marshaled, &roundTripStatus)
	require.NoError(t, err)

	// Verify the round-trip preserved all data
	assert.Equal(t, status.ChainID, roundTripStatus.ChainID)
	assert.Equal(t, status.GenesisHash, roundTripStatus.GenesisHash)
	assert.Equal(t, status.LatestProtocolVersion, roundTripStatus.LatestProtocolVersion)
	assert.Equal(t, status.NodePublicKey, roundTripStatus.NodePublicKey)
	assert.Equal(t, status.ProtocolVersion, roundTripStatus.ProtocolVersion)
	assert.Equal(t, status.RPCAddr, roundTripStatus.RPCAddr)
	assert.Equal(t, status.UptimeSec, roundTripStatus.UptimeSec)
	assert.Equal(t, status.NodeKey, roundTripStatus.NodeKey)
	assert.Equal(t, status.ValidatorAccountID, roundTripStatus.ValidatorAccountID)
	assert.Equal(t, status.ValidatorPublicKey, roundTripStatus.ValidatorPublicKey)
	assert.Equal(t, status.SyncInfo, roundTripStatus.SyncInfo)
	assert.Equal(t, status.Validators, roundTripStatus.Validators)
	assert.Equal(t, status.Version, roundTripStatus.Version)
}

func TestStatusWithNonNullableFields(t *testing.T) {
	// Test with non-null values for nullable fields
	statusWithValues := spec.Status{
		ChainID:               "testnet",
		GenesisHash:           "test_hash",
		LatestProtocolVersion: 78,
		NodeKey:               stringPtr("test_node_key"),
		NodePublicKey:         "ed25519:test_key",
		ProtocolVersion:       78,
		RPCAddr:               "127.0.0.1:3030",
		SyncInfo: spec.SyncInfo{
			EarliestBlockHash:   "earliest_hash",
			EarliestBlockHeight: 1000,
			EpochID:             "epoch_id",
			EpochStartHeight:    2000,
			LatestBlockHash:     "latest_hash",
			LatestBlockHeight:   3000,
			LatestStateRoot:     "state_root",
			Syncing:             true,
		},
		UptimeSec:          1000,
		ValidatorAccountID: stringPtr("test.validator.near"),
		ValidatorPublicKey: stringPtr("ed25519:validator_key"),
		Validators: []spec.Validator{
			{AccountID: "validator1.near", IsSlashed: false},
			{AccountID: "validator2.near", IsSlashed: true},
		},
		Version: spec.Version{
			Build:        "test_build",
			Commit:       "test_commit",
			RustcVersion: "test_rustc",
			Version:      "test_version",
		},
	}

	// Marshal to JSON
	marshaled, err := json.Marshal(statusWithValues)
	require.NoError(t, err)

	// Unmarshal back
	var unmarshaled spec.Status
	err = json.Unmarshal(marshaled, &unmarshaled)
	require.NoError(t, err)

	// Verify all fields are preserved
	assert.Equal(t, statusWithValues, unmarshaled)
}
