// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/BoltApp/bolt-go/internal/utils"
)

type PaymentMethodInputType string

const (
	PaymentMethodInputTypeCreditCard    PaymentMethodInputType = "credit_card"
	PaymentMethodInputTypePaypal        PaymentMethodInputType = "paypal"
	PaymentMethodInputTypeAffirm        PaymentMethodInputType = "affirm"
	PaymentMethodInputTypeAfterpay      PaymentMethodInputType = "afterpay"
	PaymentMethodInputTypeKlarna        PaymentMethodInputType = "klarna"
	PaymentMethodInputTypeKlarnaAccount PaymentMethodInputType = "klarna_account"
	PaymentMethodInputTypeKlarnaPaynow  PaymentMethodInputType = "klarna_paynow"
)

type PaymentMethodInput struct {
	PaymentMethodCreditCardInput *PaymentMethodCreditCardInput `queryParam:"inline"`
	PaymentMethodPaypal          *PaymentMethodPaypal          `queryParam:"inline"`
	PaymentMethodAffirm          *PaymentMethodAffirm          `queryParam:"inline"`
	PaymentMethodAfterpay        *PaymentMethodAfterpay        `queryParam:"inline"`
	PaymentMethodKlarna          *PaymentMethodKlarna          `queryParam:"inline"`
	PaymentMethodKlarnaAccount   *PaymentMethodKlarnaAccount   `queryParam:"inline"`
	PaymentMethodKlarnaPaynow    *PaymentMethodKlarnaPaynow    `queryParam:"inline"`

	Type PaymentMethodInputType
}

func CreatePaymentMethodInputCreditCard(creditCard PaymentMethodCreditCardInput) PaymentMethodInput {
	typ := PaymentMethodInputTypeCreditCard

	typStr := DotTag(typ)
	creditCard.DotTag = typStr

	return PaymentMethodInput{
		PaymentMethodCreditCardInput: &creditCard,
		Type:                         typ,
	}
}

func CreatePaymentMethodInputPaypal(paypal PaymentMethodPaypal) PaymentMethodInput {
	typ := PaymentMethodInputTypePaypal

	typStr := PaymentMethodPaypalTag(typ)
	paypal.DotTag = typStr

	return PaymentMethodInput{
		PaymentMethodPaypal: &paypal,
		Type:                typ,
	}
}

func CreatePaymentMethodInputAffirm(affirm PaymentMethodAffirm) PaymentMethodInput {
	typ := PaymentMethodInputTypeAffirm

	typStr := PaymentMethodAffirmTag(typ)
	affirm.DotTag = typStr

	return PaymentMethodInput{
		PaymentMethodAffirm: &affirm,
		Type:                typ,
	}
}

func CreatePaymentMethodInputAfterpay(afterpay PaymentMethodAfterpay) PaymentMethodInput {
	typ := PaymentMethodInputTypeAfterpay

	typStr := PaymentMethodAfterpayTag(typ)
	afterpay.DotTag = typStr

	return PaymentMethodInput{
		PaymentMethodAfterpay: &afterpay,
		Type:                  typ,
	}
}

func CreatePaymentMethodInputKlarna(klarna PaymentMethodKlarna) PaymentMethodInput {
	typ := PaymentMethodInputTypeKlarna

	typStr := PaymentMethodKlarnaTag(typ)
	klarna.DotTag = typStr

	return PaymentMethodInput{
		PaymentMethodKlarna: &klarna,
		Type:                typ,
	}
}

func CreatePaymentMethodInputKlarnaAccount(klarnaAccount PaymentMethodKlarnaAccount) PaymentMethodInput {
	typ := PaymentMethodInputTypeKlarnaAccount

	typStr := PaymentMethodKlarnaAccountTag(typ)
	klarnaAccount.DotTag = typStr

	return PaymentMethodInput{
		PaymentMethodKlarnaAccount: &klarnaAccount,
		Type:                       typ,
	}
}

func CreatePaymentMethodInputKlarnaPaynow(klarnaPaynow PaymentMethodKlarnaPaynow) PaymentMethodInput {
	typ := PaymentMethodInputTypeKlarnaPaynow

	typStr := PaymentMethodKlarnaPaynowTag(typ)
	klarnaPaynow.DotTag = typStr

	return PaymentMethodInput{
		PaymentMethodKlarnaPaynow: &klarnaPaynow,
		Type:                      typ,
	}
}

func (u *PaymentMethodInput) UnmarshalJSON(data []byte) error {

	type discriminator struct {
		DotTag string `json:".tag"`
	}

	dis := new(discriminator)
	if err := json.Unmarshal(data, &dis); err != nil {
		return fmt.Errorf("could not unmarshal discriminator: %w", err)
	}

	switch dis.DotTag {
	case "credit_card":
		paymentMethodCreditCardInput := new(PaymentMethodCreditCardInput)
		if err := utils.UnmarshalJSON(data, &paymentMethodCreditCardInput, "", true, false); err != nil {
			return fmt.Errorf("could not unmarshal `%s` into expected (DotTag == credit_card) type PaymentMethodCreditCardInput within PaymentMethodInput: %w", string(data), err)
		}

		u.PaymentMethodCreditCardInput = paymentMethodCreditCardInput
		u.Type = PaymentMethodInputTypeCreditCard
		return nil
	case "paypal":
		paymentMethodPaypal := new(PaymentMethodPaypal)
		if err := utils.UnmarshalJSON(data, &paymentMethodPaypal, "", true, false); err != nil {
			return fmt.Errorf("could not unmarshal `%s` into expected (DotTag == paypal) type PaymentMethodPaypal within PaymentMethodInput: %w", string(data), err)
		}

		u.PaymentMethodPaypal = paymentMethodPaypal
		u.Type = PaymentMethodInputTypePaypal
		return nil
	case "affirm":
		paymentMethodAffirm := new(PaymentMethodAffirm)
		if err := utils.UnmarshalJSON(data, &paymentMethodAffirm, "", true, false); err != nil {
			return fmt.Errorf("could not unmarshal `%s` into expected (DotTag == affirm) type PaymentMethodAffirm within PaymentMethodInput: %w", string(data), err)
		}

		u.PaymentMethodAffirm = paymentMethodAffirm
		u.Type = PaymentMethodInputTypeAffirm
		return nil
	case "afterpay":
		paymentMethodAfterpay := new(PaymentMethodAfterpay)
		if err := utils.UnmarshalJSON(data, &paymentMethodAfterpay, "", true, false); err != nil {
			return fmt.Errorf("could not unmarshal `%s` into expected (DotTag == afterpay) type PaymentMethodAfterpay within PaymentMethodInput: %w", string(data), err)
		}

		u.PaymentMethodAfterpay = paymentMethodAfterpay
		u.Type = PaymentMethodInputTypeAfterpay
		return nil
	case "klarna":
		paymentMethodKlarna := new(PaymentMethodKlarna)
		if err := utils.UnmarshalJSON(data, &paymentMethodKlarna, "", true, false); err != nil {
			return fmt.Errorf("could not unmarshal `%s` into expected (DotTag == klarna) type PaymentMethodKlarna within PaymentMethodInput: %w", string(data), err)
		}

		u.PaymentMethodKlarna = paymentMethodKlarna
		u.Type = PaymentMethodInputTypeKlarna
		return nil
	case "klarna_account":
		paymentMethodKlarnaAccount := new(PaymentMethodKlarnaAccount)
		if err := utils.UnmarshalJSON(data, &paymentMethodKlarnaAccount, "", true, false); err != nil {
			return fmt.Errorf("could not unmarshal `%s` into expected (DotTag == klarna_account) type PaymentMethodKlarnaAccount within PaymentMethodInput: %w", string(data), err)
		}

		u.PaymentMethodKlarnaAccount = paymentMethodKlarnaAccount
		u.Type = PaymentMethodInputTypeKlarnaAccount
		return nil
	case "klarna_paynow":
		paymentMethodKlarnaPaynow := new(PaymentMethodKlarnaPaynow)
		if err := utils.UnmarshalJSON(data, &paymentMethodKlarnaPaynow, "", true, false); err != nil {
			return fmt.Errorf("could not unmarshal `%s` into expected (DotTag == klarna_paynow) type PaymentMethodKlarnaPaynow within PaymentMethodInput: %w", string(data), err)
		}

		u.PaymentMethodKlarnaPaynow = paymentMethodKlarnaPaynow
		u.Type = PaymentMethodInputTypeKlarnaPaynow
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for PaymentMethodInput", string(data))
}

func (u PaymentMethodInput) MarshalJSON() ([]byte, error) {
	if u.PaymentMethodCreditCardInput != nil {
		return utils.MarshalJSON(u.PaymentMethodCreditCardInput, "", true)
	}

	if u.PaymentMethodPaypal != nil {
		return utils.MarshalJSON(u.PaymentMethodPaypal, "", true)
	}

	if u.PaymentMethodAffirm != nil {
		return utils.MarshalJSON(u.PaymentMethodAffirm, "", true)
	}

	if u.PaymentMethodAfterpay != nil {
		return utils.MarshalJSON(u.PaymentMethodAfterpay, "", true)
	}

	if u.PaymentMethodKlarna != nil {
		return utils.MarshalJSON(u.PaymentMethodKlarna, "", true)
	}

	if u.PaymentMethodKlarnaAccount != nil {
		return utils.MarshalJSON(u.PaymentMethodKlarnaAccount, "", true)
	}

	if u.PaymentMethodKlarnaPaynow != nil {
		return utils.MarshalJSON(u.PaymentMethodKlarnaPaynow, "", true)
	}

	return nil, errors.New("could not marshal union type PaymentMethodInput: all fields are null")
}
