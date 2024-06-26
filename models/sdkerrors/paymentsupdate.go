// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/BoltApp/bolt-go/internal/utils"
	"github.com/BoltApp/bolt-go/models/components"
	"net/http"
)

type PaymentsUpdateResponseBodyType string

const (
	PaymentsUpdateResponseBodyTypeGenericError PaymentsUpdateResponseBodyType = "generic-error"
	PaymentsUpdateResponseBodyTypeFieldError   PaymentsUpdateResponseBodyType = "field-error"
)

// PaymentsUpdateResponseBody - An error has occurred, and further details are contained in the response
type PaymentsUpdateResponseBody struct {
	GenericError *components.GenericError
	FieldError   *components.FieldError

	Type PaymentsUpdateResponseBodyType

	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response `json:"-"`
}

var _ error = &PaymentsUpdateResponseBody{}

func CreatePaymentsUpdateResponseBodyGenericError(genericError components.GenericError) PaymentsUpdateResponseBody {
	typ := PaymentsUpdateResponseBodyTypeGenericError

	return PaymentsUpdateResponseBody{
		GenericError: &genericError,
		Type:         typ,
	}
}

func CreatePaymentsUpdateResponseBodyFieldError(fieldError components.FieldError) PaymentsUpdateResponseBody {
	typ := PaymentsUpdateResponseBodyTypeFieldError

	return PaymentsUpdateResponseBody{
		FieldError: &fieldError,
		Type:       typ,
	}
}

func (u *PaymentsUpdateResponseBody) UnmarshalJSON(data []byte) error {

	var genericError components.GenericError = components.GenericError{}
	if err := utils.UnmarshalJSON(data, &genericError, "", true, true); err == nil {
		u.GenericError = &genericError
		u.Type = PaymentsUpdateResponseBodyTypeGenericError
		return nil
	}

	var fieldError components.FieldError = components.FieldError{}
	if err := utils.UnmarshalJSON(data, &fieldError, "", true, true); err == nil {
		u.FieldError = &fieldError
		u.Type = PaymentsUpdateResponseBodyTypeFieldError
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for PaymentsUpdateResponseBody", string(data))
}

func (u PaymentsUpdateResponseBody) MarshalJSON() ([]byte, error) {
	if u.GenericError != nil {
		return utils.MarshalJSON(u.GenericError, "", true)
	}

	if u.FieldError != nil {
		return utils.MarshalJSON(u.FieldError, "", true)
	}

	return nil, errors.New("could not marshal union type PaymentsUpdateResponseBody: all fields are null")
}

func (u PaymentsUpdateResponseBody) Error() string {
	switch u.Type {
	case PaymentsUpdateResponseBodyTypeGenericError:
		data, _ := json.Marshal(u.GenericError)
		return string(data)
	case PaymentsUpdateResponseBodyTypeFieldError:
		data, _ := json.Marshal(u.FieldError)
		return string(data)
	default:
		return "unknown error"
	}
}
