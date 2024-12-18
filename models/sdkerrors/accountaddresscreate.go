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

type AccountAddressCreateResponseBodyType string

const (
	AccountAddressCreateResponseBodyTypeGenericError AccountAddressCreateResponseBodyType = "generic-error"
	AccountAddressCreateResponseBodyTypeFieldError   AccountAddressCreateResponseBodyType = "field-error"
)

// AccountAddressCreateResponseBody - The address is invalid and cannot be added, or some other error has occurred
type AccountAddressCreateResponseBody struct {
	GenericError *components.GenericError `queryParam:"inline"`
	FieldError   *components.FieldError   `queryParam:"inline"`

	Type AccountAddressCreateResponseBodyType

	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response `json:"-"`
}

var _ error = &AccountAddressCreateResponseBody{}

func CreateAccountAddressCreateResponseBodyGenericError(genericError components.GenericError) AccountAddressCreateResponseBody {
	typ := AccountAddressCreateResponseBodyTypeGenericError

	return AccountAddressCreateResponseBody{
		GenericError: &genericError,
		Type:         typ,
	}
}

func CreateAccountAddressCreateResponseBodyFieldError(fieldError components.FieldError) AccountAddressCreateResponseBody {
	typ := AccountAddressCreateResponseBodyTypeFieldError

	return AccountAddressCreateResponseBody{
		FieldError: &fieldError,
		Type:       typ,
	}
}

func (u *AccountAddressCreateResponseBody) UnmarshalJSON(data []byte) error {

	var genericError components.GenericError = components.GenericError{}
	if err := utils.UnmarshalJSON(data, &genericError, "", true, true); err == nil {
		u.GenericError = &genericError
		u.Type = AccountAddressCreateResponseBodyTypeGenericError
		return nil
	}

	var fieldError components.FieldError = components.FieldError{}
	if err := utils.UnmarshalJSON(data, &fieldError, "", true, true); err == nil {
		u.FieldError = &fieldError
		u.Type = AccountAddressCreateResponseBodyTypeFieldError
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for AccountAddressCreateResponseBody", string(data))
}

func (u AccountAddressCreateResponseBody) MarshalJSON() ([]byte, error) {
	if u.GenericError != nil {
		return utils.MarshalJSON(u.GenericError, "", true)
	}

	if u.FieldError != nil {
		return utils.MarshalJSON(u.FieldError, "", true)
	}

	return nil, errors.New("could not marshal union type AccountAddressCreateResponseBody: all fields are null")
}

func (u AccountAddressCreateResponseBody) Error() string {
	switch u.Type {
	case AccountAddressCreateResponseBodyTypeGenericError:
		data, _ := json.Marshal(u.GenericError)
		return string(data)
	case AccountAddressCreateResponseBodyTypeFieldError:
		data, _ := json.Marshal(u.FieldError)
		return string(data)
	default:
		return "unknown error"
	}
}
