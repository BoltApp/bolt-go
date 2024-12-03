// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"net/http"
)

type AccountAddressDeleteRequest struct {
	// The ID of the address to delete
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// The publicly shareable identifier used to identify your Bolt merchant division.
	XPublishableKey string `header:"style=simple,explode=false,name=X-Publishable-Key"`
	// A unique identifier for a shopper's device, generated by Bolt. The value is retrieved with `Bolt.state.merchantClientId` in your frontend context, per-shopper. This header is required for proper attribution of this operation to your analytics reports. Omitting this header may result in incorrect statistics.
	XMerchantClientID *string `header:"style=simple,explode=false,name=X-Merchant-Client-Id"`
}

func (o *AccountAddressDeleteRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *AccountAddressDeleteRequest) GetXPublishableKey() string {
	if o == nil {
		return ""
	}
	return o.XPublishableKey
}

func (o *AccountAddressDeleteRequest) GetXMerchantClientID() *string {
	if o == nil {
		return nil
	}
	return o.XMerchantClientID
}

type AccountAddressDeleteResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *AccountAddressDeleteResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *AccountAddressDeleteResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *AccountAddressDeleteResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
