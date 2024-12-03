# Orders
(*Orders*)

## Overview

Use the Orders API to create and manage orders, including orders that have been placed outside the Bolt ecosystem.

### Available Operations

* [OrdersCreate](#orderscreate) - Create an order that was prepared outside the Bolt ecosystem.

## OrdersCreate

Create an order that was prepared outside the Bolt ecosystem. Some Bolt-powered flows automatically manage order creation - in those flows the order ID will be provided separately and not through this API.

### Example Usage

```go
package main

import(
	boltgo "github.com/BoltApp/bolt-go"
	"github.com/BoltApp/bolt-go/models/components"
	"context"
	"github.com/BoltApp/bolt-go/models/operations"
	"log"
)

func main() {
    s := boltgo.New()

    ctx := context.Background()
    res, err := s.Orders.OrdersCreate(ctx, operations.OrdersCreateSecurity{
        APIKey: "<YOUR_API_KEY_HERE>",
    }, "<value>", components.Order{
        Profile: components.Profile{
            FirstName: "Charlie",
            LastName: "Dunn",
            Email: "charlie@example.com",
            Phone: boltgo.String("+14085551111"),
        },
        Cart: components.Cart{
            OrderReference: "instore_20240116-878",
            OrderDescription: boltgo.String("Order #878"),
            DisplayID: boltgo.String("20240116-878"),
            Shipments: []components.CartShipment{
                components.CartShipment{
                    Address: boltgo.Pointer(components.CreateAddressReferenceInputAddressReferenceExplicitInput(
                        components.AddressReferenceExplicitInput{
                            DotTag: components.AddressReferenceExplicitTagExplicit,
                            FirstName: "Charlie",
                            LastName: "Dunn",
                            StreetAddress1: "535 Mission St",
                            Locality: "San Francisco",
                            PostalCode: "94105",
                            Region: boltgo.String("CA"),
                            CountryCode: components.CountryCodeUs,
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
                        Units: 900,
                    },
                    Code: boltgo.String("SUMMER10DISCOUNT"),
                    DetailsURL: boltgo.String("https://www.example.com/SUMMER-SALE"),
                },
            },
            Items: []components.CartItem{
                components.CartItem{
                    Name: "Red Fidget Spinner",
                    Reference: "sku-984",
                    Description: boltgo.String("Single-packed fidget spinner, red"),
                    TotalAmount: components.Amount{
                        Currency: components.CurrencyUsd,
                        Units: 1000,
                    },
                    UnitPrice: 1000,
                    Quantity: 1,
                    ImageURL: boltgo.String("https://www.example.com/products/984/image.png"),
                },
            },
            Total: components.Amount{
                Currency: components.CurrencyUsd,
                Units: 1000,
            },
            Tax: components.Amount{
                Currency: components.CurrencyUsd,
                Units: 100,
            },
        },
    }, boltgo.String("<value>"))
    if err != nil {
        log.Fatal(err)
    }
    if res.OrderResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                                                                                                            | Type                                                                                                                                                                                                                                                                                                                 | Required                                                                                                                                                                                                                                                                                                             | Description                                                                                                                                                                                                                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                                                                                                | :heavy_check_mark:                                                                                                                                                                                                                                                                                                   | The context to use for the request.                                                                                                                                                                                                                                                                                  |
| `security`                                                                                                                                                                                                                                                                                                           | [operations.OrdersCreateSecurity](../../models/operations/orderscreatesecurity.md)                                                                                                                                                                                                                                   | :heavy_check_mark:                                                                                                                                                                                                                                                                                                   | The security requirements to use for the request.                                                                                                                                                                                                                                                                    |
| `xPublishableKey`                                                                                                                                                                                                                                                                                                    | *string*                                                                                                                                                                                                                                                                                                             | :heavy_check_mark:                                                                                                                                                                                                                                                                                                   | The publicly shareable identifier used to identify your Bolt merchant division.                                                                                                                                                                                                                                      |
| `order`                                                                                                                                                                                                                                                                                                              | [components.Order](../../models/components/order.md)                                                                                                                                                                                                                                                                 | :heavy_check_mark:                                                                                                                                                                                                                                                                                                   | N/A                                                                                                                                                                                                                                                                                                                  |
| `xMerchantClientID`                                                                                                                                                                                                                                                                                                  | **string*                                                                                                                                                                                                                                                                                                            | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                   | A unique identifier for a shopper's device, generated by Bolt. The value is retrieved with `Bolt.state.merchantClientId` in your frontend context, per-shopper. This header is required for proper attribution of this operation to your analytics reports. Omitting this header may result in incorrect statistics. |
| `opts`                                                                                                                                                                                                                                                                                                               | [][operations.Option](../../models/operations/option.md)                                                                                                                                                                                                                                                             | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                   | The options for this request.                                                                                                                                                                                                                                                                                        |

### Response

**[*operations.OrdersCreateResponse](../../models/operations/orderscreateresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.OrdersCreateResponseBody | 4XX                                | application/json                   |
| sdkerrors.SDKError                 | 5XX                                | \*/\*                              |