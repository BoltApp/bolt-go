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
			APIKey: boltgo.String("<YOUR_API_KEY_HERE>"),
		}),
	)
	var xPublishableKey string = "<value>"
	ctx := context.Background()
	res, err := s.Account.GetDetails(ctx, xPublishableKey)
	if err != nil {
		log.Fatal(err)
	}
	if res.Account != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->