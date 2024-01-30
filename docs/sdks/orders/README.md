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
	"github.com/BoltApp/bolt-go/models/components"
	"github.com/BoltApp/bolt-go/models/operations"
	"context"
	"log"
)

func main() {
    s := boltgo.New()


    var xPublishableKey string = "string"

    order := components.Order{
        Profile: components.Profile{
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

    operationSecurity := operations.OrdersCreateSecurity{
            APIKey: "<YOUR_API_KEY_HERE>",
        }

    ctx := context.Background()
    res, err := s.Orders.OrdersCreate(ctx, operationSecurity, xPublishableKey, order)
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
