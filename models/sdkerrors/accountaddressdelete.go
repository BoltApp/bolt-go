// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
	"net/http"
)

// AccountAddressDeleteResponseBody - An error has occurred, and further details are contained in the response
type AccountAddressDeleteResponseBody struct {
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response `json:"-"`
}

var _ error = &AccountAddressDeleteResponseBody{}

func (e *AccountAddressDeleteResponseBody) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
