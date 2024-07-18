<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	boltgo "github.com/BoltApp/bolt-go"
	"github.com/BoltApp/bolt-go/models/components"
	"log"
	"os"
)

func main() {
	s := boltgo.New(
		boltgo.WithSecurity(components.Security{
			Oauth: boltgo.String(os.Getenv("OAUTH")),
		}),
	)
	var xPublishableKey string = "<value>"

	var xMerchantClientID string = "<value>"
	ctx := context.Background()
	res, err := s.Account.GetDetails(ctx, xPublishableKey, xMerchantClientID)
	if err != nil {
		log.Fatal(err)
	}
	if res.Account != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->