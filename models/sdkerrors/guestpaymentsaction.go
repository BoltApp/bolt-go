// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
	"net/http"
)

// GuestPaymentsActionResponseBody - An error has occurred, and further details are contained in the response
type GuestPaymentsActionResponseBody struct {
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response `json:"-"`
}

var _ error = &GuestPaymentsActionResponseBody{}

func (e *GuestPaymentsActionResponseBody) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
