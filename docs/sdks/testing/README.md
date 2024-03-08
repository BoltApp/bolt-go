# Testing
(*Testing*)

## Overview

Endpoints that allow you to generate and retrieve test data to verify certain
flows in non-production environments.


### Available Operations

* [CreateAccount](#createaccount) - Create a test account
* [GetCreditCard](#getcreditcard) - Retrieve a test credit card, including its token

## CreateAccount

Create a Bolt shopper account for testing purposes.


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

    accountTestCreationData := components.AccountTestCreationData{
        EmailState: components.EmailStateUnverified,
        PhoneState: components.PhoneStateVerified,
        IsMigrated: boltgo.Bool(true),
        HasAddress: boltgo.Bool(true),
        HasCreditCard: boltgo.Bool(true),
    }

    operationSecurity := operations.TestingAccountCreateSecurity{
            APIKey: "<YOUR_API_KEY_HERE>",
        }

    ctx := context.Background()
    res, err := s.Testing.CreateAccount(ctx, operationSecurity, xPublishableKey, accountTestCreationData)
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
| `xPublishableKey`                                                                                  | *string*                                                                                           | :heavy_check_mark:                                                                                 | The publicly viewable identifier used to identify a merchant division.                             |
| `accountTestCreationData`                                                                          | [components.AccountTestCreationData](../../models/components/accounttestcreationdata.md)           | :heavy_check_mark:                                                                                 | N/A                                                                                                |


### Response

**[*operations.TestingAccountCreateResponse](../../models/operations/testingaccountcreateresponse.md), error**
| Error Object                               | Status Code                                | Content Type                               |
| ------------------------------------------ | ------------------------------------------ | ------------------------------------------ |
| sdkerrors.TestingAccountCreateResponseBody | 4XX                                        | application/json                           |
| sdkerrors.SDKError                         | 4xx-5xx                                    | */*                                        |

## GetCreditCard

Retrieve test credit card information. This includes its token, which can be used to process payments.


### Example Usage

```go
package main

import(
	boltgo "github.com/BoltApp/bolt-go"
	"github.com/BoltApp/bolt-go/models/operations"
	"context"
	"log"
)

func main() {
    s := boltgo.New()


    operationSecurity := operations.TestingCreditCardGetSecurity{
            APIKey: "<YOUR_API_KEY_HERE>",
        }

    ctx := context.Background()
    res, err := s.Testing.GetCreditCard(ctx, operations.TestingCreditCardGetRequestBody{
        Type: operations.TypeApprove,
    }, operationSecurity)
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


### Response

**[*operations.TestingCreditCardGetResponse](../../models/operations/testingcreditcardgetresponse.md), error**
| Error Object                               | Status Code                                | Content Type                               |
| ------------------------------------------ | ------------------------------------------ | ------------------------------------------ |
| sdkerrors.TestingCreditCardGetResponseBody | 4XX                                        | application/json                           |
| sdkerrors.SDKError                         | 4xx-5xx                                    | */*                                        |
