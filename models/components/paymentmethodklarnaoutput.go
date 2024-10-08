// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

type PaymentMethodKlarnaTag string

const (
	PaymentMethodKlarnaTagKlarna PaymentMethodKlarnaTag = "klarna"
)

func (e PaymentMethodKlarnaTag) ToPointer() *PaymentMethodKlarnaTag {
	return &e
}
func (e *PaymentMethodKlarnaTag) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "klarna":
		*e = PaymentMethodKlarnaTag(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PaymentMethodKlarnaTag: %v", v)
	}
}

type PaymentMethodKlarnaOutput struct {
	DotTag PaymentMethodKlarnaTag `json:".tag"`
}

func (o *PaymentMethodKlarnaOutput) GetDotTag() PaymentMethodKlarnaTag {
	if o == nil {
		return PaymentMethodKlarnaTag("")
	}
	return o.DotTag
}

type PaymentMethodKlarna struct {
	DotTag PaymentMethodKlarnaTag `json:".tag"`
	// Return URL to return to after payment completion in Klarna.
	ReturnURL string `json:"return_url"`
}

func (o *PaymentMethodKlarna) GetDotTag() PaymentMethodKlarnaTag {
	if o == nil {
		return PaymentMethodKlarnaTag("")
	}
	return o.DotTag
}

func (o *PaymentMethodKlarna) GetReturnURL() string {
	if o == nil {
		return ""
	}
	return o.ReturnURL
}
