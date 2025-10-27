# go-near-client

[![Tag](https://img.shields.io/github/tag/attestantio/go-near-client.svg)](https://github.com/attestantio/go-near-client/releases/)
[![License](https://img.shields.io/github/license/attestantio/go-near-client.svg)](LICENSE)
[![GoDoc](https://godoc.org/github.com/attestantio/go-near-client?status.svg)](https://godoc.org/github.com/attestantio/go-near-client)
![Lint](https://github.com/attestantio/go-near-client/workflows/golangci-lint/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/attestantio/go-near-client)](https://goreportcard.com/report/github.com/attestantio/go-near-client)

Go library providing an abstraction to NEAR Protocol nodes.

This library is under development; expect APIs and data structures to change until it reaches version 1.0. In addition, clients' implementations of both their own and the standard API are themselves under development so implementation of the full API can be incomplete.

## Table of Contents

- [go-near-client](#go-near-client)
  - [Table of Contents](#table-of-contents)
  - [Install](#install)
  - [Support](#support)
  - [Usage](#usage)
  - [Example](#example)
  - [Maintainers](#maintainers)
  - [Contribute](#contribute)
  - [License](#license)

## Install

`go-near-client` is a standard Go module which can be installed with:

```sh
go get github.com/attestantio/go-near-client
```

## Support

`go-near-client` supports NEAR nodes that comply with the standard NEAR JSON-RPC API.

## Usage

Please read the [Go documentation for this library](https://godoc.org/github.com/attestantio/go-near-client) for interface information.

## Example

Below is a complete annotated example to access a NEAR node.

```go
package main

import (
    "context"
    "fmt"
    "time"

    nearclient "github.com/attestantio/go-near-client"
    "github.com/attestantio/go-near-client/api"
    "github.com/attestantio/go-near-client/jsonrpc"
    "github.com/rs/zerolog"
)

func main() {
    // Provide a cancellable context to the creation function.
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

    client, err := jsonrpc.New(ctx,
        // WithAddress supplies the address of the NEAR node, as a URL.
        jsonrpc.WithAddress("https://rpc.mainnet.near.org"),
        // WithTimeout sets the maximum duration for all requests.
        jsonrpc.WithTimeout(time.Second*600),
        // WithLogLevel supplies the level of logging to carry out.
        jsonrpc.WithLogLevel(zerolog.DebugLevel),
        // WithAllowDelayedStart allows the service to start even if the client is unavailable.
        jsonrpc.WithAllowDelayedStart(true),
    )
    if err != nil {
        panic(err)
    }

    fmt.Printf("Connected to %s\n", client.Name())

    // Query account information
    opts := api.AccountOpts{
        AccountID:  "near",
        ContractID: "near",
    }
    acc, err := client.Account(ctx, &opts)
    if err != nil {
        panic(err)
    }
    fmt.Printf("Account balance: %s\n", acc.Data.StakedBalance.PrintNEAR())

    // Query latest block
    blockOpts := api.BlockOpts{}
    block, err := client.Block(ctx, &blockOpts)
    if err != nil {
        panic(err)
    }
    fmt.Printf("Latest block height: %d\n", block.Data.Header.Height)

    // Check node status
    status, err := client.Status(ctx, &api.StatusOpts{})
    if err != nil {
        panic(err)
    }
    fmt.Printf("Node version: %s\n", status.Data.Version.Version)

    // Cancelling the context passed to New() frees up resources held by the
    // client, closes connections, clears handlers, etc.
    cancel()
}
```

## Maintainers

Chris Berry: [@bez625](https://github.com/Bez625).
Ahmed Mohamed: [@ahmohamed](https://github.com/ahmohamed).

## Contribute

Contributions welcome. Please check out [the issues](https://github.com/attestantio/go-near-client/issues).

## License

[Apache-2.0](LICENSE) © 2025 Attestant Limited