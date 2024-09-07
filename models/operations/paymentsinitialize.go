// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/BoltApp/bolt-go/models/components"
	"net/http"
)

type PaymentsInitializeRequest struct {
	// The publicly shareable identifier used to identify your Bolt merchant division.
	XPublishableKey string `header:"style=simple,explode=false,name=X-Publishable-Key"`
	// A unique identifier for a shopper's device, generated by Bolt. This header is required for proper attribution of this operation to your analytics reports. Omitting this header may result in incorrect statistics.
	XMerchantClientID        string                              `header:"style=simple,explode=false,name=X-Merchant-Client-Id"`
	PaymentInitializeRequest components.PaymentInitializeRequest `request:"mediaType=application/json"`
}

func (o *PaymentsInitializeRequest) GetXPublishableKey() string {
	if o == nil {
		return ""
	}
	return o.XPublishableKey
}

func (o *PaymentsInitializeRequest) GetXMerchantClientID() string {
	if o == nil {
		return ""
	}
	return o.XMerchantClientID
}

func (o *PaymentsInitializeRequest) GetPaymentInitializeRequest() components.PaymentInitializeRequest {
	if o == nil {
		return components.PaymentInitializeRequest{}
	}
	return o.PaymentInitializeRequest
}

type PaymentsInitializeResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// The payment was successfully initialized, and was either immediately finalized or is pending
	PaymentResponse *components.PaymentResponse
}

func (o *PaymentsInitializeResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PaymentsInitializeResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PaymentsInitializeResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *PaymentsInitializeResponse) GetPaymentResponse() *components.PaymentResponse {
	if o == nil {
		return nil
	}
	return o.PaymentResponse
}

func (o *PaymentsInitializeResponse) GetPaymentResponseFinalized() *components.PaymentResponseFinalized {
	if v := o.GetPaymentResponse(); v != nil {
		return v.PaymentResponseFinalized
	}
	return nil
}

func (o *PaymentsInitializeResponse) GetPaymentResponsePending() *components.PaymentResponsePending {
	if v := o.GetPaymentResponse(); v != nil {
		return v.PaymentResponsePending
	}
	return nil
}
