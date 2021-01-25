// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// APIToken Api token
//
// swagger:model ApiToken
type APIToken struct {
	Permission

	// String-formatted date representing the time when the token expires.
	// Format: date-time
	Expiry strfmt.DateTime `json:"expiry,omitempty"`

	// Unique ID of the API token.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// Value of the API token.
	// Read Only: true
	Token string `json:"token,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *APIToken) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 Permission
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.Permission = aO0

	// AO1
	var dataAO1 struct {
		Expiry strfmt.DateTime `json:"expiry,omitempty"`

		ID string `json:"id,omitempty"`

		Token string `json:"token,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Expiry = dataAO1.Expiry

	m.ID = dataAO1.ID

	m.Token = dataAO1.Token

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m APIToken) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.Permission)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		Expiry strfmt.DateTime `json:"expiry,omitempty"`

		ID string `json:"id,omitempty"`

		Token string `json:"token,omitempty"`
	}

	dataAO1.Expiry = m.Expiry

	dataAO1.ID = m.ID

	dataAO1.Token = m.Token

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this Api token
func (m *APIToken) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with Permission
	if err := m.Permission.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExpiry(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APIToken) validateExpiry(formats strfmt.Registry) error {

	if swag.IsZero(m.Expiry) { // not required
		return nil
	}

	if err := validate.FormatOf("expiry", "body", "date-time", m.Expiry.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *APIToken) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIToken) UnmarshalBinary(b []byte) error {
	var res APIToken
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
