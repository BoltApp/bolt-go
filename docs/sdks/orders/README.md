# Orders
(*Orders*)

## Overview

Use the Orders API to create and manage orders, including orders that have been placed outside the Bolt ecosystem.


### Available Operations

* [OrdersCreate](#orderscreate) - Create an order that was placed outside the Bolt ecosystem.

## OrdersCreate

Create an order that was placed outside the Bolt ecosystem.


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
    security := operations.OrdersCreateSecurity{
            APIKey: "<YOUR_API_KEY_HERE>",
        }

    var xPublishableKey string = "<value>"

    order := components.Order{
        Cart: components.Cart{
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
            DisplayID: boltgo.String("20240116-878"),
            Items: []components.CartItem{
                components.CartItem{
                    Description: boltgo.String("Single-packed fidget spinner, red"),
                    ImageURL: boltgo.String("https://www.example.com/products/984/image.png"),
                    Name: "Red Fidget Spinner",
                    Quantity: 1,
                    Reference: "sku-984",
                    TotalAmount: components.Amount{
                        Currency: components.CurrencyUsd,
                        Units: 1000,
                    },
                    UnitPrice: 1000,
                },
            },
            OrderDescription: boltgo.String("Order #878"),
            OrderReference: "instore_20240116-878",
            Shipments: []components.CartShipment{
                components.CartShipment{
                    Address: components.CreateAddressReferenceInputAddressReferenceExplicitInput(
                            components.AddressReferenceExplicitInput{
                                DotTag: components.AddressReferenceExplicitTagExplicit,
                                Company: boltgo.String("ACME Corporation"),
                                CountryCode: components.CountryCodeUs,
                                Email: boltgo.String("alice@example.com"),
                                FirstName: "Charlie",
                                LastName: "Dunn",
                                Locality: "San Francisco",
                                Phone: boltgo.String("+14155550199"),
                                PostalCode: "94105",
                                Region: boltgo.String("CA"),
                                StreetAddress1: "535 Mission St",
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
                Units: 1000,
            },
        },
        Profile: components.Profile{
            Email: "charlie@example.com",
            FirstName: "Charlie",
            LastName: "Dunn",
            Phone: boltgo.String("+14085551111"),
        },
    }
    ctx := context.Background()
    res, err := s.Orders.OrdersCreate(ctx, security, xPublishableKey, order)
    if err != nil {
        log.Fatal(err)
    }
    if res.OrderResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `security`                                                                         | [operations.OrdersCreateSecurity](../../models/operations/orderscreatesecurity.md) | :heavy_check_mark:                                                                 | The security requirements to use for the request.                                  |
| `xPublishableKey`                                                                  | *string*                                                                           | :heavy_check_mark:                                                                 | The publicly viewable identifier used to identify a merchant division.             |
| `order`                                                                            | [components.Order](../../models/components/order.md)                               | :heavy_check_mark:                                                                 | N/A                                                                                |


### Response

**[*operations.OrdersCreateResponse](../../models/operations/orderscreateresponse.md), error**
| Error Object                       | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.OrdersCreateResponseBody | 4XX                                | application/json                   |
| sdkerrors.SDKError                 | 4xx-5xx                            | */*                                |
