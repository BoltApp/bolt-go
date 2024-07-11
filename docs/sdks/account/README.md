# Account
(*Account*)

## Overview

Account endpoints allow you to view and manage shoppers' accounts. For example,
you can add or remove addresses and payment information.


### Available Operations

* [GetDetails](#getdetails) - Retrieve account details
* [AddAddress](#addaddress) - Add an address
* [UpdateAddress](#updateaddress) - Edit an existing address
* [DeleteAddress](#deleteaddress) - Delete an existing address
* [AddPaymentMethod](#addpaymentmethod) - Add a payment method to a shopper's Bolt account Wallet.
* [DeletePaymentMethod](#deletepaymentmethod) - Delete an existing payment method

## GetDetails

Retrieve a shopper's account details, such as addresses and payment information

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
    ctx := context.Background()
    res, err := s.Account.GetDetails(ctx, xPublishableKey, xMerchantClientID)
    if err != nil {
        log.Fatal(err)
    }
    if res.Account != nil {
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
| `opts`                                                                                                                                                                                                              | [][operations.Option](../../models/operations/option.md)                                                                                                                                                            | :heavy_minus_sign:                                                                                                                                                                                                  | The options for this request.                                                                                                                                                                                       |


### Response

**[*operations.AccountGetResponse](../../models/operations/accountgetresponse.md), error**
| Error Object                     | Status Code                      | Content Type                     |
| -------------------------------- | -------------------------------- | -------------------------------- |
| sdkerrors.AccountGetResponseBody | 4XX                              | application/json                 |
| sdkerrors.SDKError               | 4xx-5xx                          | */*                              |

## AddAddress

Add an address to the shopper's account

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

    addressListing := components.AddressListingInput{
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
    }
    ctx := context.Background()
    res, err := s.Account.AddAddress(ctx, xPublishableKey, xMerchantClientID, addressListing)
    if err != nil {
        log.Fatal(err)
    }
    if res.AddressListing != nil {
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
| `addressListing`                                                                                                                                                                                                    | [components.AddressListingInput](../../models/components/addresslistinginput.md)                                                                                                                                    | :heavy_check_mark:                                                                                                                                                                                                  | N/A                                                                                                                                                                                                                 |
| `opts`                                                                                                                                                                                                              | [][operations.Option](../../models/operations/option.md)                                                                                                                                                            | :heavy_minus_sign:                                                                                                                                                                                                  | The options for this request.                                                                                                                                                                                       |


### Response

**[*operations.AccountAddressCreateResponse](../../models/operations/accountaddresscreateresponse.md), error**
| Error Object                               | Status Code                                | Content Type                               |
| ------------------------------------------ | ------------------------------------------ | ------------------------------------------ |
| sdkerrors.AccountAddressCreateResponseBody | 4XX                                        | application/json                           |
| sdkerrors.SDKError                         | 4xx-5xx                                    | */*                                        |

## UpdateAddress

Edit an existing address on the shopper's account. This does not edit addresses
that are already associated with other resources, such as transactions or
shipments.


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
    var id string = "D4g3h5tBuVYK9"

    var xPublishableKey string = "<value>"

    var xMerchantClientID string = "<value>"

    addressListing := components.AddressListingInput{
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
    }
    ctx := context.Background()
    res, err := s.Account.UpdateAddress(ctx, id, xPublishableKey, xMerchantClientID, addressListing)
    if err != nil {
        log.Fatal(err)
    }
    if res.AddressListing != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                           | Type                                                                                                                                                                                                                | Required                                                                                                                                                                                                            | Description                                                                                                                                                                                                         | Example                                                                                                                                                                                                             |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                               | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                               | :heavy_check_mark:                                                                                                                                                                                                  | The context to use for the request.                                                                                                                                                                                 |                                                                                                                                                                                                                     |
| `id`                                                                                                                                                                                                                | *string*                                                                                                                                                                                                            | :heavy_check_mark:                                                                                                                                                                                                  | The ID of the address to edit                                                                                                                                                                                       | D4g3h5tBuVYK9                                                                                                                                                                                                       |
| `xPublishableKey`                                                                                                                                                                                                   | *string*                                                                                                                                                                                                            | :heavy_check_mark:                                                                                                                                                                                                  | The publicly viewable identifier used to identify a merchant division.                                                                                                                                              |                                                                                                                                                                                                                     |
| `xMerchantClientID`                                                                                                                                                                                                 | *string*                                                                                                                                                                                                            | :heavy_check_mark:                                                                                                                                                                                                  | A unique identifier for a shopper's device, generated by Bolt. This header is required for proper attribution of this operation to your analytics reports. Omitting this header may result in incorrect statistics. |                                                                                                                                                                                                                     |
| `addressListing`                                                                                                                                                                                                    | [components.AddressListingInput](../../models/components/addresslistinginput.md)                                                                                                                                    | :heavy_check_mark:                                                                                                                                                                                                  | N/A                                                                                                                                                                                                                 |                                                                                                                                                                                                                     |
| `opts`                                                                                                                                                                                                              | [][operations.Option](../../models/operations/option.md)                                                                                                                                                            | :heavy_minus_sign:                                                                                                                                                                                                  | The options for this request.                                                                                                                                                                                       |                                                                                                                                                                                                                     |


### Response

**[*operations.AccountAddressEditResponse](../../models/operations/accountaddresseditresponse.md), error**
| Error Object                             | Status Code                              | Content Type                             |
| ---------------------------------------- | ---------------------------------------- | ---------------------------------------- |
| sdkerrors.AccountAddressEditResponseBody | 4XX                                      | application/json                         |
| sdkerrors.SDKError                       | 4xx-5xx                                  | */*                                      |

## DeleteAddress

Delete an existing address. Deleting an address does not invalidate transactions or
shipments that are associated with it.


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
    var id string = "D4g3h5tBuVYK9"

    var xPublishableKey string = "<value>"

    var xMerchantClientID string = "<value>"
    ctx := context.Background()
    res, err := s.Account.DeleteAddress(ctx, id, xPublishableKey, xMerchantClientID)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                           | Type                                                                                                                                                                                                                | Required                                                                                                                                                                                                            | Description                                                                                                                                                                                                         | Example                                                                                                                                                                                                             |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                               | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                               | :heavy_check_mark:                                                                                                                                                                                                  | The context to use for the request.                                                                                                                                                                                 |                                                                                                                                                                                                                     |
| `id`                                                                                                                                                                                                                | *string*                                                                                                                                                                                                            | :heavy_check_mark:                                                                                                                                                                                                  | The ID of the address to delete                                                                                                                                                                                     | D4g3h5tBuVYK9                                                                                                                                                                                                       |
| `xPublishableKey`                                                                                                                                                                                                   | *string*                                                                                                                                                                                                            | :heavy_check_mark:                                                                                                                                                                                                  | The publicly viewable identifier used to identify a merchant division.                                                                                                                                              |                                                                                                                                                                                                                     |
| `xMerchantClientID`                                                                                                                                                                                                 | *string*                                                                                                                                                                                                            | :heavy_check_mark:                                                                                                                                                                                                  | A unique identifier for a shopper's device, generated by Bolt. This header is required for proper attribution of this operation to your analytics reports. Omitting this header may result in incorrect statistics. |                                                                                                                                                                                                                     |
| `opts`                                                                                                                                                                                                              | [][operations.Option](../../models/operations/option.md)                                                                                                                                                            | :heavy_minus_sign:                                                                                                                                                                                                  | The options for this request.                                                                                                                                                                                       |                                                                                                                                                                                                                     |


### Response

**[*operations.AccountAddressDeleteResponse](../../models/operations/accountaddressdeleteresponse.md), error**
| Error Object                               | Status Code                                | Content Type                               |
| ------------------------------------------ | ------------------------------------------ | ------------------------------------------ |
| sdkerrors.AccountAddressDeleteResponseBody | 4XX                                        | application/json                           |
| sdkerrors.SDKError                         | 4xx-5xx                                    | */*                                        |

## AddPaymentMethod

Add a payment method to a shopper's Bolt account Wallet. For security purposes, this request must come from
your backend because authentication requires the use of your private key.<br />
**Note**: Before using this API, the credit card details must be tokenized using Bolt's JavaScript library function,
which is documented in [Install the Bolt Tokenizer](https://help.bolt.com/developers/references/bolt-tokenizer).


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

    var paymentMethod components.PaymentMethodInput = components.CreatePaymentMethodInputPaymentMethodCreditCardInput(
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
                Expiration: "2025-03",
                Token: "a1B2c3D4e5F6G7H8i9J0k1L2m3N4o5P6Q7r8S9t0",
            },
    )
    ctx := context.Background()
    res, err := s.Account.AddPaymentMethod(ctx, xPublishableKey, xMerchantClientID, paymentMethod)
    if err != nil {
        log.Fatal(err)
    }
    if res.PaymentMethod != nil {
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
| `paymentMethod`                                                                                                                                                                                                     | [components.PaymentMethodInput](../../models/components/paymentmethodinput.md)                                                                                                                                      | :heavy_check_mark:                                                                                                                                                                                                  | N/A                                                                                                                                                                                                                 |
| `opts`                                                                                                                                                                                                              | [][operations.Option](../../models/operations/option.md)                                                                                                                                                            | :heavy_minus_sign:                                                                                                                                                                                                  | The options for this request.                                                                                                                                                                                       |


### Response

**[*operations.AccountAddPaymentMethodResponse](../../models/operations/accountaddpaymentmethodresponse.md), error**
| Error Object                                  | Status Code                                   | Content Type                                  |
| --------------------------------------------- | --------------------------------------------- | --------------------------------------------- |
| sdkerrors.AccountAddPaymentMethodResponseBody | 4XX                                           | application/json                              |
| sdkerrors.SDKError                            | 4xx-5xx                                       | */*                                           |

## DeletePaymentMethod

Delete an existing payment method. Deleting a payment method does not invalidate transactions or
orders that are associated with it.


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
    var id string = "D4g3h5tBuVYK9"

    var xPublishableKey string = "<value>"

    var xMerchantClientID string = "<value>"
    ctx := context.Background()
    res, err := s.Account.DeletePaymentMethod(ctx, id, xPublishableKey, xMerchantClientID)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                           | Type                                                                                                                                                                                                                | Required                                                                                                                                                                                                            | Description                                                                                                                                                                                                         | Example                                                                                                                                                                                                             |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                               | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                               | :heavy_check_mark:                                                                                                                                                                                                  | The context to use for the request.                                                                                                                                                                                 |                                                                                                                                                                                                                     |
| `id`                                                                                                                                                                                                                | *string*                                                                                                                                                                                                            | :heavy_check_mark:                                                                                                                                                                                                  | The ID of the payment method to delete                                                                                                                                                                              | D4g3h5tBuVYK9                                                                                                                                                                                                       |
| `xPublishableKey`                                                                                                                                                                                                   | *string*                                                                                                                                                                                                            | :heavy_check_mark:                                                                                                                                                                                                  | The publicly viewable identifier used to identify a merchant division.                                                                                                                                              |                                                                                                                                                                                                                     |
| `xMerchantClientID`                                                                                                                                                                                                 | *string*                                                                                                                                                                                                            | :heavy_check_mark:                                                                                                                                                                                                  | A unique identifier for a shopper's device, generated by Bolt. This header is required for proper attribution of this operation to your analytics reports. Omitting this header may result in incorrect statistics. |                                                                                                                                                                                                                     |
| `opts`                                                                                                                                                                                                              | [][operations.Option](../../models/operations/option.md)                                                                                                                                                            | :heavy_minus_sign:                                                                                                                                                                                                  | The options for this request.                                                                                                                                                                                       |                                                                                                                                                                                                                     |


### Response

**[*operations.AccountPaymentMethodDeleteResponse](../../models/operations/accountpaymentmethoddeleteresponse.md), error**
| Error Object                                     | Status Code                                      | Content Type                                     |
| ------------------------------------------------ | ------------------------------------------------ | ------------------------------------------------ |
| sdkerrors.AccountPaymentMethodDeleteResponseBody | 4XX                                              | application/json                                 |
| sdkerrors.SDKError                               | 4xx-5xx                                          | */*                                              |
