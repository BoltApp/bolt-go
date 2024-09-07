<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	boltgo "github.com/BoltApp/bolt-go"
	"github.com/BoltApp/bolt-go/models/components"
	"log"
)

func main() {
	s := boltgo.New(
		boltgo.WithSecurity(components.Security{
			Oauth:  boltgo.String("<YOUR_OAUTH_HERE>"),
			APIKey: boltgo.String("<YOUR_API_KEY_HERE>"),
		}),
	)

	ctx := context.Background()
	res, err := s.Account.GetDetails(ctx, "<value>", "<value>")
	if err != nil {
		log.Fatal(err)
	}
	if res.Account != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->