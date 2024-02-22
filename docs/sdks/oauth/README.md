# OAuth
(*OAuth*)

## Overview

Use this endpoint to retrieve an OAuth token. Use the token to allow your ecommerce server to make calls to the Account
endpoint and create a one-click checkout experience for shoppers.


<https://help.bolt.com/products/accounts/direct-api/oauth-guide/>
### Available Operations

* [GetToken](#gettoken) - Get OAuth token

## GetToken

Retrieve a new or refresh an existing OAuth token.

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

    ctx := context.Background()
    res, err := s.OAuth.GetToken(ctx, components.CreateTokenRequestAuthorizationCodeRequest(
            components.AuthorizationCodeRequest{
                GrantType: components.GrantTypeAuthorizationCode,
                Code: "7GSjMRSHs6Ak7C_zvVW6P2IhZOHxMK7HZKW1fMX85ms",
                ClientID: "8fd9diIy59sj.IraJdeIgmdsO.fd233434fg2c616cgo932aa6e1e4fc627a9385045gr395222a127gi93c595rg4",
                ClientSecret: "23ee7ec7301779eaff451d7c6f6cba322499e3c0ec752f800c72a8f99217e3a8",
                Scope: []components.Scope{
                    components.ScopeBoltAccountManage,
                },
                State: boltgo.String("xyzABC123"),
            },
    ))
    if err != nil {
        log.Fatal(err)
    }

    if res.GetAccessTokenResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                          | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `ctx`                                                              | [context.Context](https://pkg.go.dev/context#Context)              | :heavy_check_mark:                                                 | The context to use for the request.                                |
| `request`                                                          | [components.TokenRequest](../../models/components/tokenrequest.md) | :heavy_check_mark:                                                 | The request object to use for the request.                         |


### Response

**[*operations.OauthGetTokenResponse](../../models/operations/oauthgettokenresponse.md), error**
| Error Object                        | Status Code                         | Content Type                        |
| ----------------------------------- | ----------------------------------- | ----------------------------------- |
| sdkerrors.OauthGetTokenResponseBody | 4XX                                 | application/json                    |
| sdkerrors.SDKError                  | 4xx-5xx                             | */*                                 |
