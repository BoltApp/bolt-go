// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

type PaymentMethodKlarnaPaynowTag string

const (
	PaymentMethodKlarnaPaynowTagKlarnaPaynow PaymentMethodKlarnaPaynowTag = "klarna_paynow"
)

func (e PaymentMethodKlarnaPaynowTag) ToPointer() *PaymentMethodKlarnaPaynowTag {
	return &e
}
func (e *PaymentMethodKlarnaPaynowTag) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "klarna_paynow":
		*e = PaymentMethodKlarnaPaynowTag(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PaymentMethodKlarnaPaynowTag: %v", v)
	}
}

type PaymentMethodKlarnaPaynowOutput struct {
	DotTag PaymentMethodKlarnaPaynowTag `json:".tag"`
}

func (o *PaymentMethodKlarnaPaynowOutput) GetDotTag() PaymentMethodKlarnaPaynowTag {
	if o == nil {
		return PaymentMethodKlarnaPaynowTag("")
	}
	return o.DotTag
}

type PaymentMethodKlarnaPaynow struct {
	DotTag PaymentMethodKlarnaPaynowTag `json:".tag"`
	// Return URL to return to after payment completion in Klarna.
	ReturnURL string `json:"return_url"`
}

func (o *PaymentMethodKlarnaPaynow) GetDotTag() PaymentMethodKlarnaPaynowTag {
	if o == nil {
		return PaymentMethodKlarnaPaynowTag("")
	}
	return o.DotTag
}

func (o *PaymentMethodKlarnaPaynow) GetReturnURL() string {
	if o == nil {
		return ""
	}
	return o.ReturnURL
}
