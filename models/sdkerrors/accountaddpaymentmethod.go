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

type AccountAddPaymentMethodResponseBodyType string

const (
	AccountAddPaymentMethodResponseBodyTypeGenericError    AccountAddPaymentMethodResponseBodyType = "generic-error"
	AccountAddPaymentMethodResponseBodyTypeFieldError      AccountAddPaymentMethodResponseBodyType = "field-error"
	AccountAddPaymentMethodResponseBodyTypeCreditCardError AccountAddPaymentMethodResponseBodyType = "credit-card-error"
)

// AccountAddPaymentMethodResponseBody - The payment method is invalid and cannot be added, or some other error has occurred
type AccountAddPaymentMethodResponseBody struct {
	GenericError    *components.GenericError    `queryParam:"inline"`
	FieldError      *components.FieldError      `queryParam:"inline"`
	CreditCardError *components.CreditCardError `queryParam:"inline"`

	Type AccountAddPaymentMethodResponseBodyType

	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response `json:"-"`
}

var _ error = &AccountAddPaymentMethodResponseBody{}

func CreateAccountAddPaymentMethodResponseBodyGenericError(genericError components.GenericError) AccountAddPaymentMethodResponseBody {
	typ := AccountAddPaymentMethodResponseBodyTypeGenericError

	return AccountAddPaymentMethodResponseBody{
		GenericError: &genericError,
		Type:         typ,
	}
}

func CreateAccountAddPaymentMethodResponseBodyFieldError(fieldError components.FieldError) AccountAddPaymentMethodResponseBody {
	typ := AccountAddPaymentMethodResponseBodyTypeFieldError

	return AccountAddPaymentMethodResponseBody{
		FieldError: &fieldError,
		Type:       typ,
	}
}

func CreateAccountAddPaymentMethodResponseBodyCreditCardError(creditCardError components.CreditCardError) AccountAddPaymentMethodResponseBody {
	typ := AccountAddPaymentMethodResponseBodyTypeCreditCardError

	return AccountAddPaymentMethodResponseBody{
		CreditCardError: &creditCardError,
		Type:            typ,
	}
}

func (u *AccountAddPaymentMethodResponseBody) UnmarshalJSON(data []byte) error {

	var genericError components.GenericError = components.GenericError{}
	if err := utils.UnmarshalJSON(data, &genericError, "", true, true); err == nil {
		u.GenericError = &genericError
		u.Type = AccountAddPaymentMethodResponseBodyTypeGenericError
		return nil
	}

	var creditCardError components.CreditCardError = components.CreditCardError{}
	if err := utils.UnmarshalJSON(data, &creditCardError, "", true, true); err == nil {
		u.CreditCardError = &creditCardError
		u.Type = AccountAddPaymentMethodResponseBodyTypeCreditCardError
		return nil
	}

	var fieldError components.FieldError = components.FieldError{}
	if err := utils.UnmarshalJSON(data, &fieldError, "", true, true); err == nil {
		u.FieldError = &fieldError
		u.Type = AccountAddPaymentMethodResponseBodyTypeFieldError
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for AccountAddPaymentMethodResponseBody", string(data))
}

func (u AccountAddPaymentMethodResponseBody) MarshalJSON() ([]byte, error) {
	if u.GenericError != nil {
		return utils.MarshalJSON(u.GenericError, "", true)
	}

	if u.FieldError != nil {
		return utils.MarshalJSON(u.FieldError, "", true)
	}

	if u.CreditCardError != nil {
		return utils.MarshalJSON(u.CreditCardError, "", true)
	}

	return nil, errors.New("could not marshal union type AccountAddPaymentMethodResponseBody: all fields are null")
}

func (u AccountAddPaymentMethodResponseBody) Error() string {
	switch u.Type {
	case AccountAddPaymentMethodResponseBodyTypeGenericError:
		data, _ := json.Marshal(u.GenericError)
		return string(data)
	case AccountAddPaymentMethodResponseBodyTypeFieldError:
		data, _ := json.Marshal(u.FieldError)
		return string(data)
	case AccountAddPaymentMethodResponseBodyTypeCreditCardError:
		data, _ := json.Marshal(u.CreditCardError)
		return string(data)
	default:
		return "unknown error"
	}
}
