// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/BoltApp/bolt-go/internal/utils"
)

type PaymentMethodType string

const (
	PaymentMethodTypeCreditCard    PaymentMethodType = "credit_card"
	PaymentMethodTypePaypal        PaymentMethodType = "paypal"
	PaymentMethodTypeAffirm        PaymentMethodType = "affirm"
	PaymentMethodTypeAfterpay      PaymentMethodType = "afterpay"
	PaymentMethodTypeKlarna        PaymentMethodType = "klarna"
	PaymentMethodTypeKlarnaAccount PaymentMethodType = "klarna_account"
	PaymentMethodTypeKlarnaPaynow  PaymentMethodType = "klarna_paynow"
)

type PaymentMethod struct {
	PaymentMethodCreditCard          *PaymentMethodCreditCard          `queryParam:"inline"`
	PaymentMethodPaypalOutput        *PaymentMethodPaypalOutput        `queryParam:"inline"`
	PaymentMethodAffirmOutput        *PaymentMethodAffirmOutput        `queryParam:"inline"`
	PaymentMethodAfterpayOutput      *PaymentMethodAfterpayOutput      `queryParam:"inline"`
	PaymentMethodKlarnaOutput        *PaymentMethodKlarnaOutput        `queryParam:"inline"`
	PaymentMethodKlarnaAccountOutput *PaymentMethodKlarnaAccountOutput `queryParam:"inline"`
	PaymentMethodKlarnaPaynowOutput  *PaymentMethodKlarnaPaynowOutput  `queryParam:"inline"`

	Type PaymentMethodType
}

func CreatePaymentMethodCreditCard(creditCard PaymentMethodCreditCard) PaymentMethod {
	typ := PaymentMethodTypeCreditCard

	typStr := DotTag(typ)
	creditCard.DotTag = typStr

	return PaymentMethod{
		PaymentMethodCreditCard: &creditCard,
		Type:                    typ,
	}
}

func CreatePaymentMethodPaypal(paypal PaymentMethodPaypalOutput) PaymentMethod {
	typ := PaymentMethodTypePaypal

	typStr := PaymentMethodPaypalTag(typ)
	paypal.DotTag = typStr

	return PaymentMethod{
		PaymentMethodPaypalOutput: &paypal,
		Type:                      typ,
	}
}

func CreatePaymentMethodAffirm(affirm PaymentMethodAffirmOutput) PaymentMethod {
	typ := PaymentMethodTypeAffirm

	typStr := PaymentMethodAffirmTag(typ)
	affirm.DotTag = typStr

	return PaymentMethod{
		PaymentMethodAffirmOutput: &affirm,
		Type:                      typ,
	}
}

func CreatePaymentMethodAfterpay(afterpay PaymentMethodAfterpayOutput) PaymentMethod {
	typ := PaymentMethodTypeAfterpay

	typStr := PaymentMethodAfterpayTag(typ)
	afterpay.DotTag = typStr

	return PaymentMethod{
		PaymentMethodAfterpayOutput: &afterpay,
		Type:                        typ,
	}
}

func CreatePaymentMethodKlarna(klarna PaymentMethodKlarnaOutput) PaymentMethod {
	typ := PaymentMethodTypeKlarna

	typStr := PaymentMethodKlarnaTag(typ)
	klarna.DotTag = typStr

	return PaymentMethod{
		PaymentMethodKlarnaOutput: &klarna,
		Type:                      typ,
	}
}

func CreatePaymentMethodKlarnaAccount(klarnaAccount PaymentMethodKlarnaAccountOutput) PaymentMethod {
	typ := PaymentMethodTypeKlarnaAccount

	typStr := PaymentMethodKlarnaAccountTag(typ)
	klarnaAccount.DotTag = typStr

	return PaymentMethod{
		PaymentMethodKlarnaAccountOutput: &klarnaAccount,
		Type:                             typ,
	}
}

func CreatePaymentMethodKlarnaPaynow(klarnaPaynow PaymentMethodKlarnaPaynowOutput) PaymentMethod {
	typ := PaymentMethodTypeKlarnaPaynow

	typStr := PaymentMethodKlarnaPaynowTag(typ)
	klarnaPaynow.DotTag = typStr

	return PaymentMethod{
		PaymentMethodKlarnaPaynowOutput: &klarnaPaynow,
		Type:                            typ,
	}
}

func (u *PaymentMethod) UnmarshalJSON(data []byte) error {

	type discriminator struct {
		DotTag string `json:".tag"`
	}

	dis := new(discriminator)
	if err := json.Unmarshal(data, &dis); err != nil {
		return fmt.Errorf("could not unmarshal discriminator: %w", err)
	}

	switch dis.DotTag {
	case "credit_card":
		paymentMethodCreditCard := new(PaymentMethodCreditCard)
		if err := utils.UnmarshalJSON(data, &paymentMethodCreditCard, "", true, false); err != nil {
			return fmt.Errorf("could not unmarshal `%s` into expected (DotTag == credit_card) type PaymentMethodCreditCard within PaymentMethod: %w", string(data), err)
		}

		u.PaymentMethodCreditCard = paymentMethodCreditCard
		u.Type = PaymentMethodTypeCreditCard
		return nil
	case "paypal":
		paymentMethodPaypalOutput := new(PaymentMethodPaypalOutput)
		if err := utils.UnmarshalJSON(data, &paymentMethodPaypalOutput, "", true, false); err != nil {
			return fmt.Errorf("could not unmarshal `%s` into expected (DotTag == paypal) type PaymentMethodPaypalOutput within PaymentMethod: %w", string(data), err)
		}

		u.PaymentMethodPaypalOutput = paymentMethodPaypalOutput
		u.Type = PaymentMethodTypePaypal
		return nil
	case "affirm":
		paymentMethodAffirmOutput := new(PaymentMethodAffirmOutput)
		if err := utils.UnmarshalJSON(data, &paymentMethodAffirmOutput, "", true, false); err != nil {
			return fmt.Errorf("could not unmarshal `%s` into expected (DotTag == affirm) type PaymentMethodAffirmOutput within PaymentMethod: %w", string(data), err)
		}

		u.PaymentMethodAffirmOutput = paymentMethodAffirmOutput
		u.Type = PaymentMethodTypeAffirm
		return nil
	case "afterpay":
		paymentMethodAfterpayOutput := new(PaymentMethodAfterpayOutput)
		if err := utils.UnmarshalJSON(data, &paymentMethodAfterpayOutput, "", true, false); err != nil {
			return fmt.Errorf("could not unmarshal `%s` into expected (DotTag == afterpay) type PaymentMethodAfterpayOutput within PaymentMethod: %w", string(data), err)
		}

		u.PaymentMethodAfterpayOutput = paymentMethodAfterpayOutput
		u.Type = PaymentMethodTypeAfterpay
		return nil
	case "klarna":
		paymentMethodKlarnaOutput := new(PaymentMethodKlarnaOutput)
		if err := utils.UnmarshalJSON(data, &paymentMethodKlarnaOutput, "", true, false); err != nil {
			return fmt.Errorf("could not unmarshal `%s` into expected (DotTag == klarna) type PaymentMethodKlarnaOutput within PaymentMethod: %w", string(data), err)
		}

		u.PaymentMethodKlarnaOutput = paymentMethodKlarnaOutput
		u.Type = PaymentMethodTypeKlarna
		return nil
	case "klarna_account":
		paymentMethodKlarnaAccountOutput := new(PaymentMethodKlarnaAccountOutput)
		if err := utils.UnmarshalJSON(data, &paymentMethodKlarnaAccountOutput, "", true, false); err != nil {
			return fmt.Errorf("could not unmarshal `%s` into expected (DotTag == klarna_account) type PaymentMethodKlarnaAccountOutput within PaymentMethod: %w", string(data), err)
		}

		u.PaymentMethodKlarnaAccountOutput = paymentMethodKlarnaAccountOutput
		u.Type = PaymentMethodTypeKlarnaAccount
		return nil
	case "klarna_paynow":
		paymentMethodKlarnaPaynowOutput := new(PaymentMethodKlarnaPaynowOutput)
		if err := utils.UnmarshalJSON(data, &paymentMethodKlarnaPaynowOutput, "", true, false); err != nil {
			return fmt.Errorf("could not unmarshal `%s` into expected (DotTag == klarna_paynow) type PaymentMethodKlarnaPaynowOutput within PaymentMethod: %w", string(data), err)
		}

		u.PaymentMethodKlarnaPaynowOutput = paymentMethodKlarnaPaynowOutput
		u.Type = PaymentMethodTypeKlarnaPaynow
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for PaymentMethod", string(data))
}

func (u PaymentMethod) MarshalJSON() ([]byte, error) {
	if u.PaymentMethodCreditCard != nil {
		return utils.MarshalJSON(u.PaymentMethodCreditCard, "", true)
	}

	if u.PaymentMethodPaypalOutput != nil {
		return utils.MarshalJSON(u.PaymentMethodPaypalOutput, "", true)
	}

	if u.PaymentMethodAffirmOutput != nil {
		return utils.MarshalJSON(u.PaymentMethodAffirmOutput, "", true)
	}

	if u.PaymentMethodAfterpayOutput != nil {
		return utils.MarshalJSON(u.PaymentMethodAfterpayOutput, "", true)
	}

	if u.PaymentMethodKlarnaOutput != nil {
		return utils.MarshalJSON(u.PaymentMethodKlarnaOutput, "", true)
	}

	if u.PaymentMethodKlarnaAccountOutput != nil {
		return utils.MarshalJSON(u.PaymentMethodKlarnaAccountOutput, "", true)
	}

	if u.PaymentMethodKlarnaPaynowOutput != nil {
		return utils.MarshalJSON(u.PaymentMethodKlarnaPaynowOutput, "", true)
	}

	return nil, errors.New("could not marshal union type PaymentMethod: all fields are null")
}
