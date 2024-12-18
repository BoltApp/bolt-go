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

type AccountAddressEditResponseBodyType string

const (
	AccountAddressEditResponseBodyTypeGenericError AccountAddressEditResponseBodyType = "generic-error"
	AccountAddressEditResponseBodyTypeFieldError   AccountAddressEditResponseBodyType = "field-error"
)

// AccountAddressEditResponseBody - The address is invalid and cannot be added, or some other error has occurred
type AccountAddressEditResponseBody struct {
	GenericError *components.GenericError `queryParam:"inline"`
	FieldError   *components.FieldError   `queryParam:"inline"`

	Type AccountAddressEditResponseBodyType

	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response `json:"-"`
}

var _ error = &AccountAddressEditResponseBody{}

func CreateAccountAddressEditResponseBodyGenericError(genericError components.GenericError) AccountAddressEditResponseBody {
	typ := AccountAddressEditResponseBodyTypeGenericError

	return AccountAddressEditResponseBody{
		GenericError: &genericError,
		Type:         typ,
	}
}

func CreateAccountAddressEditResponseBodyFieldError(fieldError components.FieldError) AccountAddressEditResponseBody {
	typ := AccountAddressEditResponseBodyTypeFieldError

	return AccountAddressEditResponseBody{
		FieldError: &fieldError,
		Type:       typ,
	}
}

func (u *AccountAddressEditResponseBody) UnmarshalJSON(data []byte) error {

	var genericError components.GenericError = components.GenericError{}
	if err := utils.UnmarshalJSON(data, &genericError, "", true, true); err == nil {
		u.GenericError = &genericError
		u.Type = AccountAddressEditResponseBodyTypeGenericError
		return nil
	}

	var fieldError components.FieldError = components.FieldError{}
	if err := utils.UnmarshalJSON(data, &fieldError, "", true, true); err == nil {
		u.FieldError = &fieldError
		u.Type = AccountAddressEditResponseBodyTypeFieldError
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for AccountAddressEditResponseBody", string(data))
}

func (u AccountAddressEditResponseBody) MarshalJSON() ([]byte, error) {
	if u.GenericError != nil {
		return utils.MarshalJSON(u.GenericError, "", true)
	}

	if u.FieldError != nil {
		return utils.MarshalJSON(u.FieldError, "", true)
	}

	return nil, errors.New("could not marshal union type AccountAddressEditResponseBody: all fields are null")
}

func (u AccountAddressEditResponseBody) Error() string {
	switch u.Type {
	case AccountAddressEditResponseBodyTypeGenericError:
		data, _ := json.Marshal(u.GenericError)
		return string(data)
	case AccountAddressEditResponseBodyTypeFieldError:
		data, _ := json.Marshal(u.FieldError)
		return string(data)
	default:
		return "unknown error"
	}
}
