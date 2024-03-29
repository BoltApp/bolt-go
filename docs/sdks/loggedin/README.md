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
            Oauth: boltgo.String("Bearer <YOUR_ACCESS_TOKEN_HERE>"),
        }),
    )


    var xPublishableKey string = "<value>"

    paymentInitializeRequest := components.PaymentInitializeRequest{
        Cart: components.Cart{
            OrderReference: "order_100",
            OrderDescription: boltgo.String("Order #1234567890"),
            DisplayID: boltgo.String("215614191"),
            Total: components.Amount{
                Currency: components.CurrencyUsd,
                Units: 900,
            },
            Tax: components.Amount{
                Currency: components.CurrencyUsd,
                Units: 900,
            },
        },
        PaymentMethod: components.CreatePaymentMethodExtendedPaymentMethodCreditCardInput(
                components.PaymentMethodCreditCardInput{
                    DotTag: components.DotTagCreditCard,
                    BillingAddress: components.CreateAddressReferenceInputAddressReferenceID(
                            components.AddressReferenceID{
                                DotTag: components.AddressReferenceIDTagID,
                                ID: "D4g3h5tBuVYK9",
                            },
                    ),
                    Network: components.CreditCardNetworkVisa,
                    Bin: "411111",
                    Last4: "1004",
                    Expiration: "2029-03",
                    Token: "a1B2c3D4e5F6G7H8i9J0k1L2m3N4o5P6Q7r8S9t0",
                },
        ),
    }

    ctx := context.Background()
    res, err := s.Payments.LoggedIn.Initialize(ctx, xPublishableKey, paymentInitializeRequest)
    if err != nil {
        log.Fatal(err)
    }
    if res.PaymentResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `xPublishableKey`                                                                          | *string*                                                                                   | :heavy_check_mark:                                                                         | The publicly viewable identifier used to identify a merchant division.                     |
| `paymentInitializeRequest`                                                                 | [components.PaymentInitializeRequest](../../models/components/paymentinitializerequest.md) | :heavy_check_mark:                                                                         | N/A                                                                                        |


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
            Oauth: boltgo.String("Bearer <YOUR_ACCESS_TOKEN_HERE>"),
        }),
    )


    var id string = "iKv7t5bgt1gg"

    var xPublishableKey string = "<value>"

    paymentUpdateRequest := components.PaymentUpdateRequest{}

    ctx := context.Background()
    res, err := s.Payments.LoggedIn.Update(ctx, id, xPublishableKey, paymentUpdateRequest)
    if err != nil {
        log.Fatal(err)
    }
    if res.PaymentResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        | Example                                                                            |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |                                                                                    |
| `id`                                                                               | *string*                                                                           | :heavy_check_mark:                                                                 | The ID of the payment to update                                                    | iKv7t5bgt1gg                                                                       |
| `xPublishableKey`                                                                  | *string*                                                                           | :heavy_check_mark:                                                                 | The publicly viewable identifier used to identify a merchant division.             |                                                                                    |
| `paymentUpdateRequest`                                                             | [components.PaymentUpdateRequest](../../models/components/paymentupdaterequest.md) | :heavy_check_mark:                                                                 | N/A                                                                                |                                                                                    |


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
            Oauth: boltgo.String("Bearer <YOUR_ACCESS_TOKEN_HERE>"),
        }),
    )


    var id string = "iKv7t5bgt1gg"

    var xPublishableKey string = "<value>"

    paymentActionRequest := components.PaymentActionRequest{
        DotTag: components.PaymentActionRequestTagFinalize,
        RedirectResult: boltgo.String("eyJ0cmFuc"),
    }

    ctx := context.Background()
    res, err := s.Payments.LoggedIn.PerformAction(ctx, id, xPublishableKey, paymentActionRequest)
    if err != nil {
        log.Fatal(err)
    }
    if res.PaymentResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        | Example                                                                            |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |                                                                                    |
| `id`                                                                               | *string*                                                                           | :heavy_check_mark:                                                                 | The ID of the payment to operate on                                                | iKv7t5bgt1gg                                                                       |
| `xPublishableKey`                                                                  | *string*                                                                           | :heavy_check_mark:                                                                 | The publicly viewable identifier used to identify a merchant division.             |                                                                                    |
| `paymentActionRequest`                                                             | [components.PaymentActionRequest](../../models/components/paymentactionrequest.md) | :heavy_check_mark:                                                                 | N/A                                                                                |                                                                                    |


### Response

**[*operations.PaymentsActionResponse](../../models/operations/paymentsactionresponse.md), error**
| Error Object                         | Status Code                          | Content Type                         |
| ------------------------------------ | ------------------------------------ | ------------------------------------ |
| sdkerrors.PaymentsActionResponseBody | 4XX                                  | application/json                     |
| sdkerrors.SDKError                   | 4xx-5xx                              | */*                                  |
