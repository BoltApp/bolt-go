# LoggedIn
(*Payments.LoggedIn*)

### Available Operations

* [Initialize](#initialize) - Initialize a Bolt payment for logged in shoppers
* [Update](#update) - Update an existing payment
* [PerformAction](#performaction) - Perform an irreversible action (e.g. finalize) on a pending payment

## Initialize

Initialize a Bolt payment token that will be used to reference this payment to
Bolt when it is updated or finalized for logged in shoppers.


### Example Usage

```go
package main

import(
	"github.com/BoltApp/bolt-go/models/components"
	boltgo "github.com/BoltApp/bolt-go"
	"context"
	"log"
)

func main() {
    s := boltgo.New(
        boltgo.WithSecurity(components.Security{
            Oauth: boltgo.String("<YOUR_OAUTH_HERE>"),
        }),
    )
    var xPublishableKey string = "<value>"

    var xMerchantClientID string = "<value>"

    paymentInitializeRequest := components.PaymentInitializeRequest{
        Cart: components.Cart{
            OrderReference: "order_100",
            OrderDescription: boltgo.String("Order #1234567890"),
            DisplayID: boltgo.String("215614191"),
            Shipments: []components.CartShipment{
                components.CartShipment{
                    Address: components.CreateAddressReferenceInputAddressReferenceID(
                            components.AddressReferenceID{
                                DotTag: components.AddressReferenceIDTagID,
                                ID: "D4g3h5tBuVYK9",
                            },
                    ),
                    Cost: &components.Amount{
                        Currency: components.CurrencyUsd,
                        Units: 10000,
                    },
                    Carrier: boltgo.String("FedEx"),
                },
            },
            Discounts: []components.CartDiscount{
                components.CartDiscount{
                    Amount: components.Amount{
                        Currency: components.CurrencyUsd,
                        Units: 10000,
                    },
                    Code: boltgo.String("SUMMER10DISCOUNT"),
                    DetailsURL: boltgo.String("https://www.example.com/SUMMER-SALE"),
                },
            },
            Items: []components.CartItem{
                components.CartItem{
                    Name: "Bolt Swag Bag",
                    Reference: "item_100",
                    Description: boltgo.String("Large tote with Bolt logo."),
                    TotalAmount: components.Amount{
                        Currency: components.CurrencyUsd,
                        Units: 9000,
                    },
                    UnitPrice: 1000,
                    Quantity: 9,
                    ImageURL: boltgo.String("https://www.example.com/products/123456/images/1.png"),
                },
            },
            Total: components.Amount{
                Currency: components.CurrencyUsd,
                Units: 9000,
            },
            Tax: components.Amount{
                Currency: components.CurrencyUsd,
                Units: 100,
            },
        },
        PaymentMethod: components.CreatePaymentMethodExtendedPaymentMethodReference(
                components.PaymentMethodReference{
                    DotTag: components.PaymentMethodReferenceTagID,
                    ID: "X5h6j8uLpVGK",
                },
        ),
    }
    ctx := context.Background()
    res, err := s.Payments.LoggedIn.Initialize(ctx, xPublishableKey, xMerchantClientID, paymentInitializeRequest)
    if err != nil {
        log.Fatal(err)
    }
    if res.PaymentResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                           | Type                                                                                                                                                                                                                | Required                                                                                                                                                                                                            | Description                                                                                                                                                                                                         |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                               | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                               | :heavy_check_mark:                                                                                                                                                                                                  | The context to use for the request.                                                                                                                                                                                 |
| `xPublishableKey`                                                                                                                                                                                                   | *string*                                                                                                                                                                                                            | :heavy_check_mark:                                                                                                                                                                                                  | The publicly viewable identifier used to identify a merchant division.                                                                                                                                              |
| `xMerchantClientID`                                                                                                                                                                                                 | *string*                                                                                                                                                                                                            | :heavy_check_mark:                                                                                                                                                                                                  | A unique identifier for a shopper's device, generated by Bolt. This header is required for proper attribution of this operation to your analytics reports. Omitting this header may result in incorrect statistics. |
| `paymentInitializeRequest`                                                                                                                                                                                          | [components.PaymentInitializeRequest](../../models/components/paymentinitializerequest.md)                                                                                                                          | :heavy_check_mark:                                                                                                                                                                                                  | N/A                                                                                                                                                                                                                 |
| `opts`                                                                                                                                                                                                              | [][operations.Option](../../models/operations/option.md)                                                                                                                                                            | :heavy_minus_sign:                                                                                                                                                                                                  | The options for this request.                                                                                                                                                                                       |


### Response

**[*operations.PaymentsInitializeResponse](../../models/operations/paymentsinitializeresponse.md), error**
| Error Object                             | Status Code                              | Content Type                             |
| ---------------------------------------- | ---------------------------------------- | ---------------------------------------- |
| sdkerrors.PaymentsInitializeResponseBody | 4XX                                      | application/json                         |
| sdkerrors.SDKError                       | 4xx-5xx                                  | */*                                      |

## Update

Update a pending payment


### Example Usage

```go
package main

import(
	"github.com/BoltApp/bolt-go/models/components"
	boltgo "github.com/BoltApp/bolt-go"
	"context"
	"log"
)

func main() {
    s := boltgo.New(
        boltgo.WithSecurity(components.Security{
            Oauth: boltgo.String("<YOUR_OAUTH_HERE>"),
        }),
    )
    var id string = "iKv7t5bgt1gg"

    var xPublishableKey string = "<value>"

    var xMerchantClientID string = "<value>"

    paymentUpdateRequest := components.PaymentUpdateRequest{
        Cart: &components.Cart{
            OrderReference: "order_100",
            OrderDescription: boltgo.String("Order #1234567890"),
            DisplayID: boltgo.String("215614191"),
            Shipments: []components.CartShipment{
                components.CartShipment{
                    Address: components.CreateAddressReferenceInputAddressReferenceExplicitInput(
                            components.AddressReferenceExplicitInput{
                                DotTag: components.AddressReferenceExplicitTagExplicit,
                                FirstName: "Alice",
                                LastName: "Baker",
                                Company: boltgo.String("ACME Corporation"),
                                StreetAddress1: "535 Mission St, Ste 1401",
                                StreetAddress2: boltgo.String("c/o Shipping Department"),
                                Locality: "San Francisco",
                                PostalCode: "94105",
                                Region: boltgo.String("CA"),
                                CountryCode: components.CountryCodeUs,
                                Email: boltgo.String("alice@example.com"),
                                Phone: boltgo.String("+14155550199"),
                            },
                    ),
                    Cost: &components.Amount{
                        Currency: components.CurrencyUsd,
                        Units: 900,
                    },
                    Carrier: boltgo.String("FedEx"),
                },
            },
            Discounts: []components.CartDiscount{
                components.CartDiscount{
                    Amount: components.Amount{
                        Currency: components.CurrencyUsd,
                        Units: 900,
                    },
                    Code: boltgo.String("SUMMER10DISCOUNT"),
                    DetailsURL: boltgo.String("https://www.example.com/SUMMER-SALE"),
                },
            },
            Items: []components.CartItem{
                components.CartItem{
                    Name: "Bolt Swag Bag",
                    Reference: "item_100",
                    Description: boltgo.String("Large tote with Bolt logo."),
                    TotalAmount: components.Amount{
                        Currency: components.CurrencyUsd,
                        Units: 900,
                    },
                    UnitPrice: 1000,
                    Quantity: 1,
                    ImageURL: boltgo.String("https://www.example.com/products/123456/images/1.png"),
                },
            },
            Total: components.Amount{
                Currency: components.CurrencyUsd,
                Units: 900,
            },
            Tax: components.Amount{
                Currency: components.CurrencyUsd,
                Units: 900,
            },
        },
    }
    ctx := context.Background()
    res, err := s.Payments.LoggedIn.Update(ctx, id, xPublishableKey, xMerchantClientID, paymentUpdateRequest)
    if err != nil {
        log.Fatal(err)
    }
    if res.PaymentResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                           | Type                                                                                                                                                                                                                | Required                                                                                                                                                                                                            | Description                                                                                                                                                                                                         | Example                                                                                                                                                                                                             |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                               | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                               | :heavy_check_mark:                                                                                                                                                                                                  | The context to use for the request.                                                                                                                                                                                 |                                                                                                                                                                                                                     |
| `id`                                                                                                                                                                                                                | *string*                                                                                                                                                                                                            | :heavy_check_mark:                                                                                                                                                                                                  | The ID of the payment to update                                                                                                                                                                                     | iKv7t5bgt1gg                                                                                                                                                                                                        |
| `xPublishableKey`                                                                                                                                                                                                   | *string*                                                                                                                                                                                                            | :heavy_check_mark:                                                                                                                                                                                                  | The publicly viewable identifier used to identify a merchant division.                                                                                                                                              |                                                                                                                                                                                                                     |
| `xMerchantClientID`                                                                                                                                                                                                 | *string*                                                                                                                                                                                                            | :heavy_check_mark:                                                                                                                                                                                                  | A unique identifier for a shopper's device, generated by Bolt. This header is required for proper attribution of this operation to your analytics reports. Omitting this header may result in incorrect statistics. |                                                                                                                                                                                                                     |
| `paymentUpdateRequest`                                                                                                                                                                                              | [components.PaymentUpdateRequest](../../models/components/paymentupdaterequest.md)                                                                                                                                  | :heavy_check_mark:                                                                                                                                                                                                  | N/A                                                                                                                                                                                                                 |                                                                                                                                                                                                                     |
| `opts`                                                                                                                                                                                                              | [][operations.Option](../../models/operations/option.md)                                                                                                                                                            | :heavy_minus_sign:                                                                                                                                                                                                  | The options for this request.                                                                                                                                                                                       |                                                                                                                                                                                                                     |


### Response

**[*operations.PaymentsUpdateResponse](../../models/operations/paymentsupdateresponse.md), error**
| Error Object                         | Status Code                          | Content Type                         |
| ------------------------------------ | ------------------------------------ | ------------------------------------ |
| sdkerrors.PaymentsUpdateResponseBody | 4XX                                  | application/json                     |
| sdkerrors.SDKError                   | 4xx-5xx                              | */*                                  |

## PerformAction

Perform an irreversible action on a pending payment, such as finalizing it.


### Example Usage

```go
package main

import(
	"github.com/BoltApp/bolt-go/models/components"
	boltgo "github.com/BoltApp/bolt-go"
	"context"
	"log"
)

func main() {
    s := boltgo.New(
        boltgo.WithSecurity(components.Security{
            Oauth: boltgo.String("<YOUR_OAUTH_HERE>"),
        }),
    )
    var id string = "iKv7t5bgt1gg"

    var xPublishableKey string = "<value>"

    var xMerchantClientID string = "<value>"

    paymentActionRequest := components.PaymentActionRequest{
        DotTag: components.PaymentActionRequestTagFinalize,
        RedirectResult: boltgo.String("eyJ0cmFuc"),
    }
    ctx := context.Background()
    res, err := s.Payments.LoggedIn.PerformAction(ctx, id, xPublishableKey, xMerchantClientID, paymentActionRequest)
    if err != nil {
        log.Fatal(err)
    }
    if res.PaymentResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                           | Type                                                                                                                                                                                                                | Required                                                                                                                                                                                                            | Description                                                                                                                                                                                                         | Example                                                                                                                                                                                                             |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                               | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                               | :heavy_check_mark:                                                                                                                                                                                                  | The context to use for the request.                                                                                                                                                                                 |                                                                                                                                                                                                                     |
| `id`                                                                                                                                                                                                                | *string*                                                                                                                                                                                                            | :heavy_check_mark:                                                                                                                                                                                                  | The ID of the payment to operate on                                                                                                                                                                                 | iKv7t5bgt1gg                                                                                                                                                                                                        |
| `xPublishableKey`                                                                                                                                                                                                   | *string*                                                                                                                                                                                                            | :heavy_check_mark:                                                                                                                                                                                                  | The publicly viewable identifier used to identify a merchant division.                                                                                                                                              |                                                                                                                                                                                                                     |
| `xMerchantClientID`                                                                                                                                                                                                 | *string*                                                                                                                                                                                                            | :heavy_check_mark:                                                                                                                                                                                                  | A unique identifier for a shopper's device, generated by Bolt. This header is required for proper attribution of this operation to your analytics reports. Omitting this header may result in incorrect statistics. |                                                                                                                                                                                                                     |
| `paymentActionRequest`                                                                                                                                                                                              | [components.PaymentActionRequest](../../models/components/paymentactionrequest.md)                                                                                                                                  | :heavy_check_mark:                                                                                                                                                                                                  | N/A                                                                                                                                                                                                                 |                                                                                                                                                                                                                     |
| `opts`                                                                                                                                                                                                              | [][operations.Option](../../models/operations/option.md)                                                                                                                                                            | :heavy_minus_sign:                                                                                                                                                                                                  | The options for this request.                                                                                                                                                                                       |                                                                                                                                                                                                                     |


### Response

**[*operations.PaymentsActionResponse](../../models/operations/paymentsactionresponse.md), error**
| Error Object                         | Status Code                          | Content Type                         |
| ------------------------------------ | ------------------------------------ | ------------------------------------ |
| sdkerrors.PaymentsActionResponseBody | 4XX                                  | application/json                     |
| sdkerrors.SDKError                   | 4xx-5xx                              | */*                                  |
