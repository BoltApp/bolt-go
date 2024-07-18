# Testing
(*Testing*)

## Overview

Use the Testing API to generate and retrieve test data to verify a subset of flows in non-production environments.

### Available Operations

* [CreateAccount](#createaccount) - Create a test account
* [TestingAccountPhoneGet](#testingaccountphoneget) - Get a random phone number
* [GetCreditCard](#getcreditcard) - Retrieve a tokenized test credit card

## CreateAccount

Create a Bolt shopper account for testing purposes.

### Example Usage

```go
package main

import(
	boltgo "github.com/BoltApp/bolt-go"
	"github.com/BoltApp/bolt-go/models/operations"
	"os"
	"github.com/BoltApp/bolt-go/models/components"
	"context"
	"log"
)

func main() {
    s := boltgo.New()
    security := operations.TestingAccountCreateSecurity{
            APIKey: os.Getenv("API_KEY"),
        }

    var xPublishableKey string = "<value>"

    accountTestCreationData := components.AccountTestCreationData{
        EmailState: components.EmailStateUnverified,
        PhoneState: components.PhoneStateVerified,
        IsMigrated: boltgo.Bool(true),
        HasAddress: boltgo.Bool(true),
        HasCreditCard: boltgo.Bool(true),
    }
    ctx := context.Background()
    res, err := s.Testing.CreateAccount(ctx, security, xPublishableKey, accountTestCreationData)
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountTestCreationData != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `security`                                                                                         | [operations.TestingAccountCreateSecurity](../../models/operations/testingaccountcreatesecurity.md) | :heavy_check_mark:                                                                                 | The security requirements to use for the request.                                                  |
| `xPublishableKey`                                                                                  | *string*                                                                                           | :heavy_check_mark:                                                                                 | The publicly shareable identifier used to identify your Bolt merchant division.                    |
| `accountTestCreationData`                                                                          | [components.AccountTestCreationData](../../models/components/accounttestcreationdata.md)           | :heavy_check_mark:                                                                                 | N/A                                                                                                |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |


### Response

**[*operations.TestingAccountCreateResponse](../../models/operations/testingaccountcreateresponse.md), error**
| Error Object                               | Status Code                                | Content Type                               |
| ------------------------------------------ | ------------------------------------------ | ------------------------------------------ |
| sdkerrors.TestingAccountCreateResponseBody | 4XX                                        | application/json                           |
| sdkerrors.SDKError                         | 4xx-5xx                                    | */*                                        |

## TestingAccountPhoneGet

Get a random, fictitious phone number that is not assigned to any existing Bolt account.

### Example Usage

```go
package main

import(
	boltgo "github.com/BoltApp/bolt-go"
	"github.com/BoltApp/bolt-go/models/operations"
	"os"
	"context"
	"log"
)

func main() {
    s := boltgo.New()
    security := operations.TestingAccountPhoneGetSecurity{
            APIKey: os.Getenv("API_KEY"),
        }

    var xPublishableKey string = "<value>"
    ctx := context.Background()
    res, err := s.Testing.TestingAccountPhoneGet(ctx, security, xPublishableKey)
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountTestPhoneData != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `security`                                                                                             | [operations.TestingAccountPhoneGetSecurity](../../models/operations/testingaccountphonegetsecurity.md) | :heavy_check_mark:                                                                                     | The security requirements to use for the request.                                                      |
| `xPublishableKey`                                                                                      | *string*                                                                                               | :heavy_check_mark:                                                                                     | The publicly shareable identifier used to identify your Bolt merchant division.                        |
| `opts`                                                                                                 | [][operations.Option](../../models/operations/option.md)                                               | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |


### Response

**[*operations.TestingAccountPhoneGetResponse](../../models/operations/testingaccountphonegetresponse.md), error**
| Error Object                                 | Status Code                                  | Content Type                                 |
| -------------------------------------------- | -------------------------------------------- | -------------------------------------------- |
| sdkerrors.TestingAccountPhoneGetResponseBody | 4XX                                          | application/json                             |
| sdkerrors.SDKError                           | 4xx-5xx                                      | */*                                          |

## GetCreditCard

Retrieve a test credit card that can be used to process payments in your Bolt testing environment. The response includes the card's Bolt credit card token.

### Example Usage

```go
package main

import(
	boltgo "github.com/BoltApp/bolt-go"
	"github.com/BoltApp/bolt-go/models/operations"
	"os"
	"context"
	"log"
)

func main() {
    s := boltgo.New()
    request := operations.TestingCreditCardGetRequestBody{
        Type: operations.TypeApprove,
    }

    security := operations.TestingCreditCardGetSecurity{
            APIKey: os.Getenv("API_KEY"),
        }
    ctx := context.Background()
    res, err := s.Testing.GetCreditCard(ctx, request, security)
    if err != nil {
        log.Fatal(err)
    }
    if res.TestCreditCard != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.TestingCreditCardGetRequestBody](../../models/operations/testingcreditcardgetrequestbody.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `security`                                                                                               | [operations.TestingCreditCardGetSecurity](../../models/operations/testingcreditcardgetsecurity.md)       | :heavy_check_mark:                                                                                       | The security requirements to use for the request.                                                        |
| `opts`                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |


### Response

**[*operations.TestingCreditCardGetResponse](../../models/operations/testingcreditcardgetresponse.md), error**
| Error Object                               | Status Code                                | Content Type                               |
| ------------------------------------------ | ------------------------------------------ | ------------------------------------------ |
| sdkerrors.TestingCreditCardGetResponseBody | 4XX                                        | application/json                           |
| sdkerrors.SDKError                         | 4xx-5xx                                    | */*                                        |
