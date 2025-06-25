# go-near-client

## Example

```golang
	client, err := nearclient.New(ctx,
		nearclient.WithTimeout(time.Second*600),
		nearclient.WithLogLevel(zerolog.DebugLevel),
		nearclient.WithAddress(address),
		nearclient.WithAllowDelayedStart(true),
	)
    opts := nearapi.AccountOpts{
    AccountID:  accountID,
    ContractID: contractID,
    }
    acc, err := client.Account(ctx, &opts)
    fmt.Println(acc.Data.StakedBalance.PrintNEAR())
```