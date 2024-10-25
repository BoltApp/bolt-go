# Bolt Go SDK

<div align="left">
    <a href="https://opensource.org/licenses/MIT">
        <img src="https://img.shields.io/badge/License-MIT-blue.svg" style="width: 100px; height: 28px;" />
    </a>
</div>

<!-- Start Summary [summary] -->
## Summary

Bolt API Reference: A comprehensive Bolt API reference for interacting with Accounts, Payments, Orders and more.
<!-- End Summary [summary] -->

<!-- Start Table of Contents [toc] -->
## Table of Contents

* [SDK Installation](#sdk-installation)
* [SDK Example Usage](#sdk-example-usage)
* [Available Resources and Operations](#available-resources-and-operations)
* [Retries](#retries)
* [Error Handling](#error-handling)
* [Server Selection](#server-selection)
* [Custom HTTP Client](#custom-http-client)
* [Authentication](#authentication)
* [Special Types](#special-types)
<!-- End Table of Contents [toc] -->

<!-- Start SDK Installation [installation] -->
## SDK Installation

To add the SDK as a dependency to your project:
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

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

<details open>
<summary>Available methods</summary>

### [Account](docs/sdks/account/README.md)

* [GetDetails](docs/sdks/account/README.md#getdetails) - Retrieve account details
* [AddAddress](docs/sdks/account/README.md#addaddress) - Add an address
* [UpdateAddress](docs/sdks/account/README.md#updateaddress) - Edit an existing address
* [DeleteAddress](docs/sdks/account/README.md#deleteaddress) - Delete an existing address
* [AddPaymentMethod](docs/sdks/account/README.md#addpaymentmethod) - Add a payment method
* [DeletePaymentMethod](docs/sdks/account/README.md#deletepaymentmethod) - Delete an existing payment method


### [OAuth](docs/sdks/oauth/README.md)

* [GetToken](docs/sdks/oauth/README.md#gettoken) - Get OAuth token

### [Orders](docs/sdks/orders/README.md)

* [OrdersCreate](docs/sdks/orders/README.md#orderscreate) - Create an order that was prepared outside the Bolt ecosystem.

### [Payments](docs/sdks/payments/README.md)


#### [Payments.Guest](docs/sdks/guest/README.md)

* [Initialize](docs/sdks/guest/README.md#initialize) - Initialize a Bolt payment for guest shoppers
* [PerformAction](docs/sdks/guest/README.md#performaction) - Finalize a pending guest payment

#### [Payments.LoggedIn](docs/sdks/loggedin/README.md)

* [Initialize](docs/sdks/loggedin/README.md#initialize) - Initialize a Bolt payment for logged in shoppers
* [PerformAction](docs/sdks/loggedin/README.md#performaction) - Finalize a pending payment

### [Testing](docs/sdks/testing/README.md)

* [CreateAccount](docs/sdks/testing/README.md#createaccount) - Create a test account
* [TestingAccountPhoneGet](docs/sdks/testing/README.md#testingaccountphoneget) - Get a random phone number
* [GetCreditCard](docs/sdks/testing/README.md#getcreditcard) - Retrieve a tokenized test credit card

</details>
<!-- End Available Resources and Operations [operations] -->

<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations. All operations return a response object or an error, they will never return both.

By Default, an API error will return `sdkerrors.SDKError`. When custom error responses are specified for an operation, the SDK may also return their associated error. You can refer to respective *Errors* tables in SDK docs for more details on possible error types for each operation.

For example, the `GetDetails` function may return the following errors:

| Error Type                       | Status Code                      | Content Type                     |
| -------------------------------- | -------------------------------- | -------------------------------- |
| sdkerrors.AccountGetResponseBody | 4XX                              | application/json                 |
| sdkerrors.SDKError               | 5XX                              | \*/\*                            |

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
			Oauth:  boltgo.String("<YOUR_OAUTH_HERE>"),
			APIKey: boltgo.String("<YOUR_API_KEY_HERE>"),
		}),
	)

	ctx := context.Background()
	res, err := s.Account.GetDetails(ctx, "<value>", "<value>")
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

	ctx := context.Background()
	res, err := s.Payments.Guest.Initialize(ctx, operations.GuestPaymentsInitializeSecurity{
		APIKey: "<YOUR_API_KEY_HERE>",
	}, "<value>", "<value>", components.GuestPaymentInitializeRequest{
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
			Shipments: []components.CartShipment{
				components.CartShipment{
					Address: boltgo.Pointer(components.CreateAddressReferenceInputAddressReferenceExplicitInput(
						components.AddressReferenceExplicitInput{
							DotTag:         components.AddressReferenceExplicitTagExplicit,
							FirstName:      "Alice",
							LastName:       "Baker",
							Company:        boltgo.String("ACME Corporation"),
							StreetAddress1: "535 Mission St, Ste 1401",
							StreetAddress2: boltgo.String("c/o Shipping Department"),
							Locality:       "San Francisco",
							PostalCode:     "94105",
							Region:         boltgo.String("CA"),
							CountryCode:    components.CountryCodeUs,
							Email:          boltgo.String("alice@example.com"),
							Phone:          boltgo.String("+14155550199"),
						},
					)),
					Cost: &components.Amount{
						Currency: components.CurrencyUsd,
						Units:    10000,
					},
					Carrier: boltgo.String("FedEx"),
				},
			},
			Discounts: []components.CartDiscount{
				components.CartDiscount{
					Amount: components.Amount{
						Currency: components.CurrencyUsd,
						Units:    10000,
					},
					Code:       boltgo.String("SUMMER10DISCOUNT"),
					DetailsURL: boltgo.String("https://www.example.com/SUMMER-SALE"),
				},
			},
			Items: []components.CartItem{
				components.CartItem{
					Name:        "Bolt Swag Bag",
					Reference:   "item_100",
					Description: boltgo.String("Large tote with Bolt logo."),
					TotalAmount: components.Amount{
						Currency: components.CurrencyUsd,
						Units:    9000,
					},
					UnitPrice: 1000,
					Quantity:  9,
					ImageURL:  boltgo.String("https://www.example.com/products/123456/images/1.png"),
				},
			},
			Total: components.Amount{
				Currency: components.CurrencyUsd,
				Units:    9000,
			},
			Tax: components.Amount{
				Currency: components.CurrencyUsd,
				Units:    100,
			},
		},
		PaymentMethod: components.CreatePaymentMethodInputPaymentMethodCreditCardInput(
			components.PaymentMethodCreditCardInput{
				DotTag: components.DotTagCreditCard,
				BillingAddress: components.CreateAddressReferenceInputAddressReferenceExplicitInput(
					components.AddressReferenceExplicitInput{
						DotTag:         components.AddressReferenceExplicitTagExplicit,
						FirstName:      "Alice",
						LastName:       "Baker",
						Company:        boltgo.String("ACME Corporation"),
						StreetAddress1: "535 Mission St, Ste 1401",
						StreetAddress2: boltgo.String("c/o Shipping Department"),
						Locality:       "San Francisco",
						PostalCode:     "94105",
						Region:         boltgo.String("CA"),
						CountryCode:    components.CountryCodeUs,
						Email:          boltgo.String("alice@example.com"),
						Phone:          boltgo.String("+14155550199"),
					},
				),
				Network:    components.CreditCardNetworkVisa,
				Bin:        "411111",
				Last4:      "1004",
				Expiration: "2025-03",
				Token:      "a1B2c3D4e5F6G7H8i9J0k1L2m3N4o5P6Q7r8S9t0",
			},
		),
	})
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

<!-- Start Retries [retries] -->
## Retries

Some of the endpoints in this SDK support retries. If you use the SDK without any configuration, it will fall back to the default retry strategy provided by the API. However, the default retry strategy can be overridden on a per-operation basis, or across the entire SDK.

To change the default retry strategy for a single API call, simply provide a `retry.Config` object to the call by using the `WithRetries` option:
```go
package main

import (
	"context"
	boltgo "github.com/BoltApp/bolt-go"
	"github.com/BoltApp/bolt-go/models/components"
	"github.com/BoltApp/bolt-go/retry"
	"log"
	"models/operations"
)

func main() {
	s := boltgo.New(
		boltgo.WithSecurity(components.Security{
			Oauth:  boltgo.String("<YOUR_OAUTH_HERE>"),
			APIKey: boltgo.String("<YOUR_API_KEY_HERE>"),
		}),
	)

	ctx := context.Background()
	res, err := s.Account.GetDetails(ctx, "<value>", "<value>", operations.WithRetries(
		retry.Config{
			Strategy: "backoff",
			Backoff: &retry.BackoffStrategy{
				InitialInterval: 1,
				MaxInterval:     50,
				Exponent:        1.1,
				MaxElapsedTime:  100,
			},
			RetryConnectionErrors: false,
		}))
	if err != nil {
		log.Fatal(err)
	}
	if res.Account != nil {
		// handle response
	}
}

```

If you'd like to override the default retry strategy for all operations that support retries, you can use the `WithRetryConfig` option at SDK initialization:
```go
package main

import (
	"context"
	boltgo "github.com/BoltApp/bolt-go"
	"github.com/BoltApp/bolt-go/models/components"
	"github.com/BoltApp/bolt-go/retry"
	"log"
)

func main() {
	s := boltgo.New(
		boltgo.WithRetryConfig(
			retry.Config{
				Strategy: "backoff",
				Backoff: &retry.BackoffStrategy{
					InitialInterval: 1,
					MaxInterval:     50,
					Exponent:        1.1,
					MaxElapsedTime:  100,
				},
				RetryConnectionErrors: false,
			}),
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
<!-- End Retries [retries] -->

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
