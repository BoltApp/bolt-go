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
	"github.com/BoltApp/bolt-go/models/components"
	"github.com/BoltApp/bolt-go/models/operations"
	"context"
	"log"
)

func main() {
    s := boltgo.New()


    var xPublishableKey string = "<value>"

    guestPaymentInitializeRequest := components.GuestPaymentInitializeRequest{
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
            Total: components.Amount{
                Currency: components.CurrencyUsd,
                Units: 900,
            },
            Tax: components.Amount{
                Currency: components.CurrencyUsd,
                Units: 900,
            },
        },
        PaymentMethod: components.CreatePaymentMethodInputPaymentMethodPaypal(
                components.PaymentMethodPaypal{
                    DotTag: components.PaymentMethodPaypalTagPaypal,
                    SuccessURL: "https://www.example.com/paypal-callback/success",
                    CancelURL: "https://www.example.com/paypal-callback/cancel",
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
	"github.com/BoltApp/bolt-go/models/components"
	"github.com/BoltApp/bolt-go/models/operations"
	"context"
	"log"
)

func main() {
    s := boltgo.New()


    var id string = "iKv7t5bgt1gg"

    var xPublishableKey string = "<value>"

    paymentUpdateRequest := components.PaymentUpdateRequest{}

    operationSecurity := operations.GuestPaymentsUpdateSecurity{
            APIKey: "<YOUR_API_KEY_HERE>",
        }

    ctx := context.Background()
    res, err := s.Payments.Guest.Update(ctx, operationSecurity, id, xPublishableKey, paymentUpdateRequest)
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
	"github.com/BoltApp/bolt-go/models/components"
	"github.com/BoltApp/bolt-go/models/operations"
	"context"
	"log"
)

func main() {
    s := boltgo.New()


    var id string = "iKv7t5bgt1gg"

    var xPublishableKey string = "<value>"

    paymentActionRequest := components.PaymentActionRequest{
        DotTag: components.PaymentActionRequestTagFinalize,
        RedirectResult: boltgo.String("eyJ0cmFuc"),
    }

    operationSecurity := operations.GuestPaymentsActionSecurity{
            APIKey: "<YOUR_API_KEY_HERE>",
        }

    ctx := context.Background()
    res, err := s.Payments.Guest.PerformAction(ctx, operationSecurity, id, xPublishableKey, paymentActionRequest)
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
