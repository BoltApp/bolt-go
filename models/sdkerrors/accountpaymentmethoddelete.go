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

type AccountPaymentMethodDeleteResponseBodyType string

const (
	AccountPaymentMethodDeleteResponseBodyTypeGenericError AccountPaymentMethodDeleteResponseBodyType = "generic-error"
	AccountPaymentMethodDeleteResponseBodyTypeFieldError   AccountPaymentMethodDeleteResponseBodyType = "field-error"
)

// AccountPaymentMethodDeleteResponseBody - An error has occurred, and further details are contained in the response
type AccountPaymentMethodDeleteResponseBody struct {
	GenericError *components.GenericError `queryParam:"inline"`
	FieldError   *components.FieldError   `queryParam:"inline"`

	Type AccountPaymentMethodDeleteResponseBodyType

	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response `json:"-"`
}

var _ error = &AccountPaymentMethodDeleteResponseBody{}

func CreateAccountPaymentMethodDeleteResponseBodyGenericError(genericError components.GenericError) AccountPaymentMethodDeleteResponseBody {
	typ := AccountPaymentMethodDeleteResponseBodyTypeGenericError

	return AccountPaymentMethodDeleteResponseBody{
		GenericError: &genericError,
		Type:         typ,
	}
}

func CreateAccountPaymentMethodDeleteResponseBodyFieldError(fieldError components.FieldError) AccountPaymentMethodDeleteResponseBody {
	typ := AccountPaymentMethodDeleteResponseBodyTypeFieldError

	return AccountPaymentMethodDeleteResponseBody{
		FieldError: &fieldError,
		Type:       typ,
	}
}

func (u *AccountPaymentMethodDeleteResponseBody) UnmarshalJSON(data []byte) error {

	var genericError components.GenericError = components.GenericError{}
	if err := utils.UnmarshalJSON(data, &genericError, "", true, true); err == nil {
		u.GenericError = &genericError
		u.Type = AccountPaymentMethodDeleteResponseBodyTypeGenericError
		return nil
	}

	var fieldError components.FieldError = components.FieldError{}
	if err := utils.UnmarshalJSON(data, &fieldError, "", true, true); err == nil {
		u.FieldError = &fieldError
		u.Type = AccountPaymentMethodDeleteResponseBodyTypeFieldError
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for AccountPaymentMethodDeleteResponseBody", string(data))
}

func (u AccountPaymentMethodDeleteResponseBody) MarshalJSON() ([]byte, error) {
	if u.GenericError != nil {
		return utils.MarshalJSON(u.GenericError, "", true)
	}

	if u.FieldError != nil {
		return utils.MarshalJSON(u.FieldError, "", true)
	}

	return nil, errors.New("could not marshal union type AccountPaymentMethodDeleteResponseBody: all fields are null")
}

func (u AccountPaymentMethodDeleteResponseBody) Error() string {
	switch u.Type {
	case AccountPaymentMethodDeleteResponseBodyTypeGenericError:
		data, _ := json.Marshal(u.GenericError)
		return string(data)
	case AccountPaymentMethodDeleteResponseBodyTypeFieldError:
		data, _ := json.Marshal(u.FieldError)
		return string(data)
	default:
		return "unknown error"
	}
}
