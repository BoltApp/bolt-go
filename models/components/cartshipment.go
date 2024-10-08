// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type CartShipment struct {
	Address *AddressReferenceInput `json:"address,omitempty"`
	// A monetary amount, i.e. a base unit amount and a supported currency.
	Cost *Amount `json:"cost,omitempty"`
	// The name of the carrier selected.
	Carrier *string `json:"carrier,omitempty"`
}

func (o *CartShipment) GetAddress() *AddressReferenceInput {
	if o == nil {
		return nil
	}
	return o.Address
}

func (o *CartShipment) GetAddressID() *AddressReferenceID {
	if v := o.GetAddress(); v != nil {
		return v.AddressReferenceID
	}
	return nil
}

func (o *CartShipment) GetAddressExplicit() *AddressReferenceExplicitInput {
	if v := o.GetAddress(); v != nil {
		return v.AddressReferenceExplicitInput
	}
	return nil
}

func (o *CartShipment) GetCost() *Amount {
	if o == nil {
		return nil
	}
	return o.Cost
}

func (o *CartShipment) GetCarrier() *string {
	if o == nil {
		return nil
	}
	return o.Carrier
}
