# Guest
(*Payments.Guest*)

## Overview

### Available Operations

* [Initialize](#initialize) - Initialize a Bolt payment for guest shoppers
* [PerformAction](#performaction) - Finalize a pending guest payment

## Initialize

Initialize a Bolt guest shopper's intent to pay for a cart, using the specified payment method. Payments must be finalized before indicating the payment result to the shopper. Some payment methods will finalize automatically after initialization. For these payments, they will transition directly to "finalized" and the response from Initialize Payment will contain a finalized payment.

### Example Usage

```go
package main

import(
	boltgo "github.com/BoltApp/bolt-go"
	"context"
	"github.com/BoltApp/bolt-go/models/operations"
	"github.com/BoltApp/bolt-go/models/components"
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
            FirstName: "Alice",
            LastName: "Baker",
            Email: "alice@example.com",
            Phone: boltgo.String("+14155550199"),
        },
        Cart: components.Cart{
            OrderReference: "order_100",
            OrderDescription: boltgo.String("Order #1234567890"),
            DisplayID: boltgo.String("215614191"),
            Shipments: []components.CartShipment{
                components.CartShipment{
                    Address: boltgo.Pointer(components.CreateAddressReferenceInputAddressReferenceExplicitInput(
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
                    )),
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
        PaymentMethod: components.CreatePaymentMethodInputPaymentMethodCreditCardInput(
            components.PaymentMethodCreditCardInput{
                DotTag: components.DotTagCreditCard,
                BillingAddress: components.CreateAddressReferenceInputAddressReferenceExplicitInput(
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
                Network: components.CreditCardNetworkVisa,
                Bin: "411111",
                Last4: "1004",
                Expiration: "2025-03",
                Token: "a1B2c3D4e5F6G7H8i9J0k1L2m3N4o5P6Q7r8S9t0",
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

### Parameters

| Parameter                                                                                                                                                                                                           | Type                                                                                                                                                                                                                | Required                                                                                                                                                                                                            | Description                                                                                                                                                                                                         |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                               | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                               | :heavy_check_mark:                                                                                                                                                                                                  | The context to use for the request.                                                                                                                                                                                 |
| `security`                                                                                                                                                                                                          | [operations.GuestPaymentsInitializeSecurity](../../models/operations/guestpaymentsinitializesecurity.md)                                                                                                            | :heavy_check_mark:                                                                                                                                                                                                  | The security requirements to use for the request.                                                                                                                                                                   |
| `xPublishableKey`                                                                                                                                                                                                   | *string*                                                                                                                                                                                                            | :heavy_check_mark:                                                                                                                                                                                                  | The publicly shareable identifier used to identify your Bolt merchant division.                                                                                                                                     |
| `xMerchantClientID`                                                                                                                                                                                                 | *string*                                                                                                                                                                                                            | :heavy_check_mark:                                                                                                                                                                                                  | A unique identifier for a shopper's device, generated by Bolt. This header is required for proper attribution of this operation to your analytics reports. Omitting this header may result in incorrect statistics. |
| `guestPaymentInitializeRequest`                                                                                                                                                                                     | [components.GuestPaymentInitializeRequest](../../models/components/guestpaymentinitializerequest.md)                                                                                                                | :heavy_check_mark:                                                                                                                                                                                                  | N/A                                                                                                                                                                                                                 |
| `opts`                                                                                                                                                                                                              | [][operations.Option](../../models/operations/option.md)                                                                                                                                                            | :heavy_minus_sign:                                                                                                                                                                                                  | The options for this request.                                                                                                                                                                                       |

### Response

**[*operations.GuestPaymentsInitializeResponse](../../models/operations/guestpaymentsinitializeresponse.md), error**

### Errors

| Error Type                                    | Status Code                                   | Content Type                                  |
| --------------------------------------------- | --------------------------------------------- | --------------------------------------------- |
| sdkerrors.GuestPaymentsInitializeResponseBody | 4XX                                           | application/json                              |
| sdkerrors.SDKError                            | 5XX                                           | \*/\*                                         |

## PerformAction

Finalize a pending payment being made by a Bolt guest shopper. Upon receipt of a finalized payment result, payment success should be communicated to the shopper.

### Example Usage

```go
package main

import(
	boltgo "github.com/BoltApp/bolt-go"
	"context"
	"github.com/BoltApp/bolt-go/models/operations"
	"github.com/BoltApp/bolt-go/models/components"
	"log"
)

func main() {
    s := boltgo.New()

    ctx := context.Background()
    res, err := s.Payments.Guest.PerformAction(ctx, operations.GuestPaymentsActionSecurity{
        APIKey: "<YOUR_API_KEY_HERE>",
    }, "iKv7t5bgt1gg", "<value>", "<value>", components.PaymentActionRequest{
        DotTag: components.PaymentActionRequestTagFinalize,
        RedirectResult: boltgo.String("eyJ0cmFuc"),
    })
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
| `security`                                                                                                                                                                                                          | [operations.GuestPaymentsActionSecurity](../../models/operations/guestpaymentsactionsecurity.md)                                                                                                                    | :heavy_check_mark:                                                                                                                                                                                                  | The security requirements to use for the request.                                                                                                                                                                   |                                                                                                                                                                                                                     |
| `id`                                                                                                                                                                                                                | *string*                                                                                                                                                                                                            | :heavy_check_mark:                                                                                                                                                                                                  | The ID of the guest payment to operate on                                                                                                                                                                           | iKv7t5bgt1gg                                                                                                                                                                                                        |
| `xPublishableKey`                                                                                                                                                                                                   | *string*                                                                                                                                                                                                            | :heavy_check_mark:                                                                                                                                                                                                  | The publicly shareable identifier used to identify your Bolt merchant division.                                                                                                                                     |                                                                                                                                                                                                                     |
| `xMerchantClientID`                                                                                                                                                                                                 | *string*                                                                                                                                                                                                            | :heavy_check_mark:                                                                                                                                                                                                  | A unique identifier for a shopper's device, generated by Bolt. This header is required for proper attribution of this operation to your analytics reports. Omitting this header may result in incorrect statistics. |                                                                                                                                                                                                                     |
| `paymentActionRequest`                                                                                                                                                                                              | [components.PaymentActionRequest](../../models/components/paymentactionrequest.md)                                                                                                                                  | :heavy_check_mark:                                                                                                                                                                                                  | N/A                                                                                                                                                                                                                 |                                                                                                                                                                                                                     |
| `opts`                                                                                                                                                                                                              | [][operations.Option](../../models/operations/option.md)                                                                                                                                                            | :heavy_minus_sign:                                                                                                                                                                                                  | The options for this request.                                                                                                                                                                                       |                                                                                                                                                                                                                     |

### Response

**[*operations.GuestPaymentsActionResponse](../../models/operations/guestpaymentsactionresponse.md), error**

### Errors

| Error Type                                | Status Code                               | Content Type                              |
| ----------------------------------------- | ----------------------------------------- | ----------------------------------------- |
| sdkerrors.GuestPaymentsActionResponseBody | 4XX                                       | application/json                          |
| sdkerrors.SDKError                        | 5XX                                       | \*/\*                                     |