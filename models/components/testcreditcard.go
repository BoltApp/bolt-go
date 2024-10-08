// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"github.com/BoltApp/bolt-go/internal/utils"
	"time"
)

type TestCreditCard struct {
	// The credit card's network.
	Network CreditCardNetwork `json:"network"`
	// The Bank Identification Number (BIN). This is typically the first 4 to 6 digits of the account number.
	Bin string `json:"bin"`
	// The account number's last four digits.
	Last4 string `json:"last4"`
	// The token's expiration date. Tokens used past their expiration will be rejected.
	Expiration time.Time `json:"expiration"`
	// The Bolt token associated with the credit card.
	Token string `json:"token"`
}

func (t TestCreditCard) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TestCreditCard) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *TestCreditCard) GetNetwork() CreditCardNetwork {
	if o == nil {
		return CreditCardNetwork("")
	}
	return o.Network
}

func (o *TestCreditCard) GetBin() string {
	if o == nil {
		return ""
	}
	return o.Bin
}

func (o *TestCreditCard) GetLast4() string {
	if o == nil {
		return ""
	}
	return o.Last4
}

func (o *TestCreditCard) GetExpiration() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.Expiration
}

func (o *TestCreditCard) GetToken() string {
	if o == nil {
		return ""
	}
	return o.Token
}
