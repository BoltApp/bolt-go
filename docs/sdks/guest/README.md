# Guest
(*Payments.Guest*)

### Available Operations

* [Initialize](#initialize) - Initialize a Bolt payment for guest shoppers
* [Update](#update) - Update an existing guest payment
* [PerformAction](#performaction) - Perform an irreversible action (e.g. finalize) on a pending guest payment

## Initialize

Initialize a Bolt payment token that will be used to reference this payment to
Bolt when it is updated or finalized for guest shoppers.


### Example Usage

```go
package main

import(
	boltgo "github.com/BoltApp/bolt-go"
	"github.com/BoltApp/bolt-go/models/operations"
	"github.com/BoltApp/bolt-go/models/components"
	"context"
	"log"
)

func main() {
    s := boltgo.New()
    security := operations.GuestPaymentsInitializeSecurity{
            APIKey: "<YOUR_API_KEY_HERE>",
        }

    var xPublishableKey string = "<value>"

    guestPaymentInitializeRequest := components.GuestPaymentInitializeRequest{
        Cart: components.Cart{
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
            DisplayID: boltgo.String("215614191"),
            Items: []components.CartItem{
                components.CartItem{
                    Description: boltgo.String("Large tote with Bolt logo."),
                    ImageURL: boltgo.String("https://www.example.com/products/123456/images/1.png"),
                    Name: "Bolt Swag Bag",
                    Quantity: 9,
                    Reference: "item_100",
                    TotalAmount: components.Amount{
                        Currency: components.CurrencyUsd,
                        Units: 9000,
                    },
                    UnitPrice: 1000,
                },
            },
            OrderDescription: boltgo.String("Order #1234567890"),
            OrderReference: "order_100",
            Shipments: []components.CartShipment{
                components.CartShipment{
                    Address: components.CreateAddressReferenceInputAddressReferenceExplicitInput(
                            components.AddressReferenceExplicitInput{
                                DotTag: components.AddressReferenceExplicitTagExplicit,
                                Company: boltgo.String("ACME Corporation"),
                                CountryCode: components.CountryCodeUs,
                                Email: boltgo.String("alice@example.com"),
                                FirstName: "Alice",
                                LastName: "Baker",
                                Locality: "San Francisco",
                                Phone: boltgo.String("+14155550199"),
                                PostalCode: "94105",
                                Region: boltgo.String("CA"),
                                StreetAddress1: "535 Mission St, Ste 1401",
                                StreetAddress2: boltgo.String("c/o Shipping Department"),
                            },
                    ),
                    Carrier: boltgo.String("FedEx"),
                    Cost: &components.Amount{
                        Currency: components.CurrencyUsd,
                        Units: 10000,
                    },
                },
            },
            Tax: components.Amount{
                Currency: components.CurrencyUsd,
                Units: 100,
            },
            Total: components.Amount{
                Currency: components.CurrencyUsd,
                Units: 9000,
            },
        },
        PaymentMethod: components.CreatePaymentMethodInputPaymentMethodAffirm(
                components.PaymentMethodAffirm{
                    DotTag: components.PaymentMethodAffirmTagAffirm,
                    ReturnURL: "www.example.com/handle_affirm_success",
                },
        ),
        Profile: components.ProfileCreationData{
            CreateAccount: true,
            Email: "alice@example.com",
            FirstName: "Alice",
            LastName: "Baker",
            Phone: boltgo.String("+14155550199"),
        },
    }
    ctx := context.Background()
    res, err := s.Payments.Guest.Initialize(ctx, security, xPublishableKey, guestPaymentInitializeRequest)
    if err != nil {
        log.Fatal(err)
    }
    if res.PaymentResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `security`                                                                                               | [operations.GuestPaymentsInitializeSecurity](../../models/operations/guestpaymentsinitializesecurity.md) | :heavy_check_mark:                                                                                       | The security requirements to use for the request.                                                        |
| `xPublishableKey`                                                                                        | *string*                                                                                                 | :heavy_check_mark:                                                                                       | The publicly viewable identifier used to identify a merchant division.                                   |
| `guestPaymentInitializeRequest`                                                                          | [components.GuestPaymentInitializeRequest](../../models/components/guestpaymentinitializerequest.md)     | :heavy_check_mark:                                                                                       | N/A                                                                                                      |


### Response

**[*operations.GuestPaymentsInitializeResponse](../../models/operations/guestpaymentsinitializeresponse.md), error**
| Error Object                                  | Status Code                                   | Content Type                                  |
| --------------------------------------------- | --------------------------------------------- | --------------------------------------------- |
| sdkerrors.GuestPaymentsInitializeResponseBody | 4XX                                           | application/json                              |
| sdkerrors.SDKError                            | 4xx-5xx                                       | */*                                           |

## Update

Update a pending guest payment


### Example Usage

```go
package main

import(
	boltgo "github.com/BoltApp/bolt-go"
	"github.com/BoltApp/bolt-go/models/operations"
	"github.com/BoltApp/bolt-go/models/components"
	"context"
	"log"
)

func main() {
    s := boltgo.New()
    security := operations.GuestPaymentsUpdateSecurity{
            APIKey: "<YOUR_API_KEY_HERE>",
        }

    var id string = "iKv7t5bgt1gg"

    var xPublishableKey string = "<value>"

    paymentUpdateRequest := components.PaymentUpdateRequest{
        Cart: &components.Cart{
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
            DisplayID: boltgo.String("215614191"),
            Items: []components.CartItem{
                components.CartItem{
                    Description: boltgo.String("Large tote with Bolt logo."),
                    ImageURL: boltgo.String("https://www.example.com/products/123456/images/1.png"),
                    Name: "Bolt Swag Bag",
                    Quantity: 1,
                    Reference: "item_100",
                    TotalAmount: components.Amount{
                        Currency: components.CurrencyUsd,
                        Units: 900,
                    },
                    UnitPrice: 1000,
                },
            },
            OrderDescription: boltgo.String("Order #1234567890"),
            OrderReference: "order_100",
            Shipments: []components.CartShipment{
                components.CartShipment{
                    Address: components.CreateAddressReferenceInputAddressReferenceID(
                            components.AddressReferenceID{
                                DotTag: components.AddressReferenceIDTagID,
                                ID: "D4g3h5tBuVYK9",
                            },
                    ),
                    Carrier: boltgo.String("FedEx"),
                    Cost: &components.Amount{
                        Currency: components.CurrencyUsd,
                        Units: 900,
                    },
                },
            },
            Tax: components.Amount{
                Currency: components.CurrencyUsd,
                Units: 900,
            },
            Total: components.Amount{
                Currency: components.CurrencyUsd,
                Units: 900,
            },
        },
    }
    ctx := context.Background()
    res, err := s.Payments.Guest.Update(ctx, security, id, xPublishableKey, paymentUpdateRequest)
    if err != nil {
        log.Fatal(err)
    }
    if res.PaymentResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      | Example                                                                                          |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |                                                                                                  |
| `security`                                                                                       | [operations.GuestPaymentsUpdateSecurity](../../models/operations/guestpaymentsupdatesecurity.md) | :heavy_check_mark:                                                                               | The security requirements to use for the request.                                                |                                                                                                  |
| `id`                                                                                             | *string*                                                                                         | :heavy_check_mark:                                                                               | The ID of the guest payment to update                                                            | iKv7t5bgt1gg                                                                                     |
| `xPublishableKey`                                                                                | *string*                                                                                         | :heavy_check_mark:                                                                               | The publicly viewable identifier used to identify a merchant division.                           |                                                                                                  |
| `paymentUpdateRequest`                                                                           | [components.PaymentUpdateRequest](../../models/components/paymentupdaterequest.md)               | :heavy_check_mark:                                                                               | N/A                                                                                              |                                                                                                  |


### Response

**[*operations.GuestPaymentsUpdateResponse](../../models/operations/guestpaymentsupdateresponse.md), error**
| Error Object                              | Status Code                               | Content Type                              |
| ----------------------------------------- | ----------------------------------------- | ----------------------------------------- |
| sdkerrors.GuestPaymentsUpdateResponseBody | 4XX                                       | application/json                          |
| sdkerrors.SDKError                        | 4xx-5xx                                   | */*                                       |

## PerformAction

Perform an irreversible action on a pending guest payment, such as finalizing it.


### Example Usage

```go
package main

import(
	boltgo "github.com/BoltApp/bolt-go"
	"github.com/BoltApp/bolt-go/models/operations"
	"github.com/BoltApp/bolt-go/models/components"
	"context"
	"log"
)

func main() {
    s := boltgo.New()
    security := operations.GuestPaymentsActionSecurity{
            APIKey: "<YOUR_API_KEY_HERE>",
        }

    var id string = "iKv7t5bgt1gg"

    var xPublishableKey string = "<value>"

    paymentActionRequest := components.PaymentActionRequest{
        DotTag: components.PaymentActionRequestTagFinalize,
        RedirectResult: boltgo.String("eyJ0cmFuc"),
    }
    ctx := context.Background()
    res, err := s.Payments.Guest.PerformAction(ctx, security, id, xPublishableKey, paymentActionRequest)
    if err != nil {
        log.Fatal(err)
    }
    if res.PaymentResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      | Example                                                                                          |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |                                                                                                  |
| `security`                                                                                       | [operations.GuestPaymentsActionSecurity](../../models/operations/guestpaymentsactionsecurity.md) | :heavy_check_mark:                                                                               | The security requirements to use for the request.                                                |                                                                                                  |
| `id`                                                                                             | *string*                                                                                         | :heavy_check_mark:                                                                               | The ID of the guest payment to operate on                                                        | iKv7t5bgt1gg                                                                                     |
| `xPublishableKey`                                                                                | *string*                                                                                         | :heavy_check_mark:                                                                               | The publicly viewable identifier used to identify a merchant division.                           |                                                                                                  |
| `paymentActionRequest`                                                                           | [components.PaymentActionRequest](../../models/components/paymentactionrequest.md)               | :heavy_check_mark:                                                                               | N/A                                                                                              |                                                                                                  |


### Response

**[*operations.GuestPaymentsActionResponse](../../models/operations/guestpaymentsactionresponse.md), error**
| Error Object                              | Status Code                               | Content Type                              |
| ----------------------------------------- | ----------------------------------------- | ----------------------------------------- |
| sdkerrors.GuestPaymentsActionResponseBody | 4XX                                       | application/json                          |
| sdkerrors.SDKError                        | 4xx-5xx                                   | */*                                       |
