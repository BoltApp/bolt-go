// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/BoltApp/bolt-go/internal/utils"
	"github.com/BoltApp/bolt-go/models/components"
	"net/http"
)

type PaymentsActionResponseBodyType string

const (
	PaymentsActionResponseBodyTypeGenericError PaymentsActionResponseBodyType = "generic-error"
	PaymentsActionResponseBodyTypeFieldError   PaymentsActionResponseBodyType = "field-error"
)

// PaymentsActionResponseBody - An error has occurred, and further details are contained in the response
type PaymentsActionResponseBody struct {
	GenericError *components.GenericError `queryParam:"inline"`
	FieldError   *components.FieldError   `queryParam:"inline"`

	Type PaymentsActionResponseBodyType

	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response `json:"-"`
}

var _ error = &PaymentsActionResponseBody{}

func CreatePaymentsActionResponseBodyGenericError(genericError components.GenericError) PaymentsActionResponseBody {
	typ := PaymentsActionResponseBodyTypeGenericError

	return PaymentsActionResponseBody{
		GenericError: &genericError,
		Type:         typ,
	}
}

func CreatePaymentsActionResponseBodyFieldError(fieldError components.FieldError) PaymentsActionResponseBody {
	typ := PaymentsActionResponseBodyTypeFieldError

	return PaymentsActionResponseBody{
		FieldError: &fieldError,
		Type:       typ,
	}
}

func (u *PaymentsActionResponseBody) UnmarshalJSON(data []byte) error {

	var genericError components.GenericError = components.GenericError{}
	if err := utils.UnmarshalJSON(data, &genericError, "", true, true); err == nil {
		u.GenericError = &genericError
		u.Type = PaymentsActionResponseBodyTypeGenericError
		return nil
	}

	var fieldError components.FieldError = components.FieldError{}
	if err := utils.UnmarshalJSON(data, &fieldError, "", true, true); err == nil {
		u.FieldError = &fieldError
		u.Type = PaymentsActionResponseBodyTypeFieldError
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for PaymentsActionResponseBody", string(data))
}

func (u PaymentsActionResponseBody) MarshalJSON() ([]byte, error) {
	if u.GenericError != nil {
		return utils.MarshalJSON(u.GenericError, "", true)
	}

	if u.FieldError != nil {
		return utils.MarshalJSON(u.FieldError, "", true)
	}

	return nil, errors.New("could not marshal union type PaymentsActionResponseBody: all fields are null")
}

func (u PaymentsActionResponseBody) Error() string {
	switch u.Type {
	case PaymentsActionResponseBodyTypeGenericError:
		data, _ := json.Marshal(u.GenericError)
		return string(data)
	case PaymentsActionResponseBodyTypeFieldError:
		data, _ := json.Marshal(u.FieldError)
		return string(data)
	default:
		return "unknown error"
	}
}
