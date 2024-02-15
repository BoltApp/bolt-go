# github.com/BoltApp/bolt-go

<div align="left">
    <a href="https://opensource.org/licenses/MIT">
        <img src="https://img.shields.io/badge/License-MIT-blue.svg" style="width: 100px; height: 28px;" />
    </a>
</div>

<!-- Start SDK Installation [installation] -->
## SDK Installation

```bash
go get github.com/BoltApp/bolt-go
```
<!-- End SDK Installation [installation] -->

<!-- Start SDK Example Usage [usage] -->
## SDK Example Usage

### Example

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
			Oauth: boltgo.String("Bearer <YOUR_ACCESS_TOKEN_HERE>"),
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

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

### [Account](docs/sdks/account/README.md)

* [GetDetails](docs/sdks/account/README.md#getdetails) - Retrieve account details
* [AddAddress](docs/sdks/account/README.md#addaddress) - Add an address
* [UpdateAddress](docs/sdks/account/README.md#updateaddress) - Edit an existing address
* [DeleteAddress](docs/sdks/account/README.md#deleteaddress) - Delete an existing address
* [AddPaymentMethod](docs/sdks/account/README.md#addpaymentmethod) - Add a payment method to a shopper's Bolt account Wallet.
* [DeletePaymentMethod](docs/sdks/account/README.md#deletepaymentmethod) - Delete an existing payment method


### [Payments.Guest](docs/sdks/guest/README.md)

* [Initialize](docs/sdks/guest/README.md#initialize) - Initialize a Bolt payment for guest shoppers
* [Update](docs/sdks/guest/README.md#update) - Update an existing guest payment
* [PerformAction](docs/sdks/guest/README.md#performaction) - Perform an irreversible action (e.g. finalize) on a pending guest payment

### [Payments.LoggedIn](docs/sdks/loggedin/README.md)

* [Initialize](docs/sdks/loggedin/README.md#initialize) - Initialize a Bolt payment for logged in shoppers
* [Update](docs/sdks/loggedin/README.md#update) - Update an existing payment
* [PerformAction](docs/sdks/loggedin/README.md#performaction) - Perform an irreversible action (e.g. finalize) on a pending payment

### [OAuth](docs/sdks/oauth/README.md)

* [GetToken](docs/sdks/oauth/README.md#gettoken) - Get OAuth token

### [Orders](docs/sdks/orders/README.md)

* [OrdersCreate](docs/sdks/orders/README.md#orderscreate) - Create an order that was placed outside the Bolt ecosystem.

### [Testing](docs/sdks/testing/README.md)

* [CreateAccount](docs/sdks/testing/README.md#createaccount) - Create a test account
* [GetCreditCard](docs/sdks/testing/README.md#getcreditcard) - Retrieve a test credit card, including its token
<!-- End Available Resources and Operations [operations] -->

<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations.  All operations return a response object or an error, they will never return both.  When specified by the OpenAPI spec document, the SDK will return the appropriate subclass.

| Error Object                     | Status Code                      | Content Type                     |
| -------------------------------- | -------------------------------- | -------------------------------- |
| sdkerrors.AccountGetResponseBody | 4XX                              | application/json                 |
| sdkerrors.SDKError               | 4xx-5xx                          | */*                              |

### Example

```go
package main

import (
	"context"
	"errors"
	boltgo "github.com/BoltApp/bolt-go"
	"github.com/BoltApp/bolt-go/models/components"
	"github.com/BoltApp/bolt-go/models/sdkerrors"
	"log"
)

func main() {
	s := boltgo.New(
		boltgo.WithSecurity(components.Security{
			Oauth: boltgo.String("Bearer <YOUR_ACCESS_TOKEN_HERE>"),
		}),
	)

	var xPublishableKey string = "<value>"

	ctx := context.Background()
	res, err := s.Account.GetDetails(ctx, xPublishableKey)
	if err != nil {

		var e *sdkerrors.AccountGetResponseBody
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.SDKError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}
	}
}

```
<!-- End Error Handling [errors] -->

<!-- Start Server Selection [server] -->
## Server Selection

### Select Server by Index

You can override the default server globally using the `WithServerIndex` option when initializing the SDK client instance. The selected server will then be used as the default on the operations that use it. This table lists the indexes associated with the available servers:

| # | Server | Variables |
| - | ------ | --------- |
| 0 | `https://{environment}.bolt.com/v3` | `environment` (default is `api-sandbox`) |

#### Example

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
		boltgo.WithServerIndex(0),
		boltgo.WithSecurity(components.Security{
			Oauth: boltgo.String("Bearer <YOUR_ACCESS_TOKEN_HERE>"),
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

#### Variables

Some of the server options above contain variables. If you want to set the values of those variables, the following options are provided for doing so:
 * `WithEnvironment boltgo.ServerEnvironment`

### Override Server URL Per-Client

The default server can also be overridden globally using the `WithServerURL` option when initializing the SDK client instance. For example:
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
		boltgo.WithServerURL("https://{environment}.bolt.com/v3"),
		boltgo.WithSecurity(components.Security{
			Oauth: boltgo.String("Bearer <YOUR_ACCESS_TOKEN_HERE>"),
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
<!-- End Server Selection [server] -->

<!-- Start Custom HTTP Client [http-client] -->
## Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"
	"github.com/myorg/your-go-sdk"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = sdk.New(sdk.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client [http-client] -->

<!-- Start Authentication [security] -->
## Authentication

### Per-Client Security Schemes

This SDK supports the following security schemes globally:

| Name         | Type         | Scheme       |
| ------------ | ------------ | ------------ |
| `Oauth`      | oauth2       | OAuth2 token |
| `APIKey`     | apiKey       | API key      |

You can set the security parameters through the `WithSecurity` option when initializing the SDK client instance. The selected scheme will be used by default to authenticate with the API for all operations that support it. For example:
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
			Oauth: boltgo.String("Bearer <YOUR_ACCESS_TOKEN_HERE>"),
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

### Per-Operation Security Schemes

Some operations in this SDK require the security scheme to be specified at the request level. For example:
```go
package main

import (
	"context"
	boltgo "github.com/BoltApp/bolt-go"
	"github.com/BoltApp/bolt-go/models/components"
	"github.com/BoltApp/bolt-go/models/operations"
	"log"
)

func main() {
	s := boltgo.New()

	var xPublishableKey string = "<value>"

	guestPaymentInitializeRequest := components.GuestPaymentInitializeRequest{
		Profile: components.ProfileCreationData{
			CreateAccount: true,
			FirstName:     "Alice",
			LastName:      "Baker",
			Email:         "alice@example.com",
			Phone:         boltgo.String("+14155550199"),
		},
		Cart: components.Cart{
			OrderReference:   "order_100",
			OrderDescription: boltgo.String("Order #1234567890"),
			DisplayID:        boltgo.String("215614191"),
			Total: components.Amount{
				Currency: components.CurrencyUsd,
				Units:    900,
			},
			Tax: components.Amount{
				Currency: components.CurrencyUsd,
				Units:    900,
			},
		},
		PaymentMethod: components.CreatePaymentMethodInputPaymentMethodPaypal(
			components.PaymentMethodPaypal{
				DotTag:     components.PaymentMethodPaypalTagPaypal,
				SuccessURL: "https://www.example.com/paypal-callback/success",
				CancelURL:  "https://www.example.com/paypal-callback/cancel",
			},
		),
	}

	operationSecurity := operations.GuestPaymentsInitializeSecurity{
		APIKey: "<YOUR_API_KEY_HERE>",
	}

	ctx := context.Background()
	res, err := s.Payments.Guest.Initialize(ctx, operationSecurity, xPublishableKey, guestPaymentInitializeRequest)
	if err != nil {
		log.Fatal(err)
	}

	if res.PaymentResponse != nil {
		// handle response
	}
}

```
<!-- End Authentication [security] -->

<!-- Start Special Types [types] -->
## Special Types


<!-- End Special Types [types] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->

# Development

## Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

## Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release!

<div align="left">
    <a href="https://speakeasyapi.dev/"><img src="https://custom-icon-badges.demolab.com/badge/-Built%20By%20Speakeasy-212015?style=for-the-badge&logoColor=FBE331&logo=speakeasy&labelColor=545454" /></a>
</div>
