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

type OrdersCreateResponseBodyType string

const (
	OrdersCreateResponseBodyTypeGenericError OrdersCreateResponseBodyType = "generic-error"
	OrdersCreateResponseBodyTypeFieldError   OrdersCreateResponseBodyType = "field-error"
)

// OrdersCreateResponseBody - An error has occurred, and further details are contained in the response
type OrdersCreateResponseBody struct {
	GenericError *components.GenericError `queryParam:"inline"`
	FieldError   *components.FieldError   `queryParam:"inline"`

	Type OrdersCreateResponseBodyType

	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response `json:"-"`
}

var _ error = &OrdersCreateResponseBody{}

func CreateOrdersCreateResponseBodyGenericError(genericError components.GenericError) OrdersCreateResponseBody {
	typ := OrdersCreateResponseBodyTypeGenericError

	return OrdersCreateResponseBody{
		GenericError: &genericError,
		Type:         typ,
	}
}

func CreateOrdersCreateResponseBodyFieldError(fieldError components.FieldError) OrdersCreateResponseBody {
	typ := OrdersCreateResponseBodyTypeFieldError

	return OrdersCreateResponseBody{
		FieldError: &fieldError,
		Type:       typ,
	}
}

func (u *OrdersCreateResponseBody) UnmarshalJSON(data []byte) error {

	var genericError components.GenericError = components.GenericError{}
	if err := utils.UnmarshalJSON(data, &genericError, "", true, true); err == nil {
		u.GenericError = &genericError
		u.Type = OrdersCreateResponseBodyTypeGenericError
		return nil
	}

	var fieldError components.FieldError = components.FieldError{}
	if err := utils.UnmarshalJSON(data, &fieldError, "", true, true); err == nil {
		u.FieldError = &fieldError
		u.Type = OrdersCreateResponseBodyTypeFieldError
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for OrdersCreateResponseBody", string(data))
}

func (u OrdersCreateResponseBody) MarshalJSON() ([]byte, error) {
	if u.GenericError != nil {
		return utils.MarshalJSON(u.GenericError, "", true)
	}

	if u.FieldError != nil {
		return utils.MarshalJSON(u.FieldError, "", true)
	}

	return nil, errors.New("could not marshal union type OrdersCreateResponseBody: all fields are null")
}

func (u OrdersCreateResponseBody) Error() string {
	switch u.Type {
	case OrdersCreateResponseBodyTypeGenericError:
		data, _ := json.Marshal(u.GenericError)
		return string(data)
	case OrdersCreateResponseBodyTypeFieldError:
		data, _ := json.Marshal(u.FieldError)
		return string(data)
	default:
		return "unknown error"
	}
}
