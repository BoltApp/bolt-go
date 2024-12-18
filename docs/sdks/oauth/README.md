# OAuth
(*OAuth*)

## Overview

Use the OAuth API to enable your ecommerce server to make API calls on behalf of a Bolt logged-in shopper.
<https://help.bolt.com/products/accounts/direct-api/oauth-guide/>

### Available Operations

* [GetToken](#gettoken) - Get OAuth token

## GetToken

Retrieve a new or refresh an existing OAuth token.

### Example Usage

```go
package main

import(
	"context"
	boltgo "github.com/BoltApp/bolt-go"
	"github.com/BoltApp/bolt-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := boltgo.New()

    res, err := s.OAuth.GetToken(ctx, components.CreateTokenRequestAuthorizationCodeRequest(
        components.AuthorizationCodeRequest{
            GrantType: components.GrantTypeAuthorizationCode,
            Code: "7GSjMRSHs6Ak7C_zvVW6P2IhZOHxMK7HZKW1fMX85ms",
            ClientID: "8fd9diIy59sj.IraJdeIgmdsO.fd233434fg2c616cgo932aa6e1e4fc627a9385045gr395222a127gi93c595rg4",
            ClientSecret: "23ee7ec7301779eaff451d7c6f6cba322499e3c0ec752f800c72a8f99217e3a8",
            Scope: []components.Scope{
                components.ScopeBoltAccountManage,
                components.ScopeBoltAccountView,
                components.ScopeOpenid,
            },
            State: boltgo.String("xyzABC123"),
        },
    ), boltgo.String("<value>"))
    if err != nil {
        log.Fatal(err)
    }
    if res.GetAccessTokenResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                                                                                                            | Type                                                                                                                                                                                                                                                                                                                 | Required                                                                                                                                                                                                                                                                                                             | Description                                                                                                                                                                                                                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                                                                                                | :heavy_check_mark:                                                                                                                                                                                                                                                                                                   | The context to use for the request.                                                                                                                                                                                                                                                                                  |
| `tokenRequest`                                                                                                                                                                                                                                                                                                       | [components.TokenRequest](../../models/components/tokenrequest.md)                                                                                                                                                                                                                                                   | :heavy_check_mark:                                                                                                                                                                                                                                                                                                   | N/A                                                                                                                                                                                                                                                                                                                  |
| `xMerchantClientID`                                                                                                                                                                                                                                                                                                  | **string*                                                                                                                                                                                                                                                                                                            | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                   | A unique identifier for a shopper's device, generated by Bolt. The value is retrieved with `Bolt.state.merchantClientId` in your frontend context, per-shopper. This header is required for proper attribution of this operation to your analytics reports. Omitting this header may result in incorrect statistics. |
| `opts`                                                                                                                                                                                                                                                                                                               | [][operations.Option](../../models/operations/option.md)                                                                                                                                                                                                                                                             | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                   | The options for this request.                                                                                                                                                                                                                                                                                        |

### Response

**[*operations.OauthGetTokenResponse](../../models/operations/oauthgettokenresponse.md), error**

### Errors

| Error Type                          | Status Code                         | Content Type                        |
| ----------------------------------- | ----------------------------------- | ----------------------------------- |
| sdkerrors.OauthGetTokenResponseBody | 4XX                                 | application/json                    |
| sdkerrors.SDKError                  | 5XX                                 | \*/\*                               |