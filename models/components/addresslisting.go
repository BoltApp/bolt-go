// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// AddressListing - An address saved on an account, i.e. a physical address plus any additional account-specific metadata.
type AddressListing struct {
	// The address's unique identifier.
	ID *string `json:"id,omitempty"`
	// The first name of the person associated with this address.
	FirstName string `json:"first_name"`
	// The last name of the person associated with this address.
	LastName string `json:"last_name"`
	// The company associated with this address.
	Company *string `json:"company,omitempty"`
	// The street address associated with this address.
	StreetAddress1 string `json:"street_address1"`
	// Any additional, optional, street address information associated with this address.
	StreetAddress2 *string `json:"street_address2,omitempty"`
	// The locality (e.g. city, town, etc...) associated with this address.
	Locality string `json:"locality"`
	// The postal code associated with this address.
	PostalCode string `json:"postal_code"`
	// The region or administrative area (e.g. state, province, county, etc...) associated with this address.
	Region *string `json:"region,omitempty"`
	// The country (in its ISO 3166 alpha-2 format) associated with this address.
	CountryCode CountryCode `json:"country_code"`
	// The email address associated with this address.
	Email *string `json:"email,omitempty"`
	// The phone number associated with this address.
	Phone *string `json:"phone,omitempty"`
	// Whether or not this is the default address saved.
	IsDefault *bool `json:"is_default,omitempty"`
}

func (o *AddressListing) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *AddressListing) GetFirstName() string {
	if o == nil {
		return ""
	}
	return o.FirstName
}

func (o *AddressListing) GetLastName() string {
	if o == nil {
		return ""
	}
	return o.LastName
}

func (o *AddressListing) GetCompany() *string {
	if o == nil {
		return nil
	}
	return o.Company
}

func (o *AddressListing) GetStreetAddress1() string {
	if o == nil {
		return ""
	}
	return o.StreetAddress1
}

func (o *AddressListing) GetStreetAddress2() *string {
	if o == nil {
		return nil
	}
	return o.StreetAddress2
}

func (o *AddressListing) GetLocality() string {
	if o == nil {
		return ""
	}
	return o.Locality
}

func (o *AddressListing) GetPostalCode() string {
	if o == nil {
		return ""
	}
	return o.PostalCode
}

func (o *AddressListing) GetRegion() *string {
	if o == nil {
		return nil
	}
	return o.Region
}

func (o *AddressListing) GetCountryCode() CountryCode {
	if o == nil {
		return CountryCode("")
	}
	return o.CountryCode
}

func (o *AddressListing) GetEmail() *string {
	if o == nil {
		return nil
	}
	return o.Email
}

func (o *AddressListing) GetPhone() *string {
	if o == nil {
		return nil
	}
	return o.Phone
}

func (o *AddressListing) GetIsDefault() *bool {
	if o == nil {
		return nil
	}
	return o.IsDefault
}
