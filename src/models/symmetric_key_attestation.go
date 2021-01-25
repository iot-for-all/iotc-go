// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SymmetricKeyAttestation symmetric key attestation
//
// swagger:model SymmetricKeyAttestation
type SymmetricKeyAttestation struct {

	// The symmetric key credentials for this attestation.
	// Required: true
	SymmetricKey *SymmetricKey `json:"symmetricKey"`
}

// Type gets the type of this subtype
func (m *SymmetricKeyAttestation) Type() string {
	return "SymmetricKeyAttestation"
}

// SetType sets the type of this subtype
func (m *SymmetricKeyAttestation) SetType(val string) {
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *SymmetricKeyAttestation) UnmarshalJSON(raw []byte) error {
	var data struct {

		// The symmetric key credentials for this attestation.
		// Required: true
		SymmetricKey *SymmetricKey `json:"symmetricKey"`
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var base struct {
		/* Just the base type fields. Used for unmashalling polymorphic types.*/

		Type string `json:"type"`
	}
	buf = bytes.NewBuffer(raw)
	dec = json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&base); err != nil {
		return err
	}

	var result SymmetricKeyAttestation

	if base.Type != result.Type() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid type value: %q", base.Type)
	}

	result.SymmetricKey = data.SymmetricKey

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m SymmetricKeyAttestation) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// The symmetric key credentials for this attestation.
		// Required: true
		SymmetricKey *SymmetricKey `json:"symmetricKey"`
	}{

		SymmetricKey: m.SymmetricKey,
	})
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		Type string `json:"type"`
	}{

		Type: m.Type(),
	})
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this symmetric key attestation
func (m *SymmetricKeyAttestation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSymmetricKey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SymmetricKeyAttestation) validateSymmetricKey(formats strfmt.Registry) error {

	if err := validate.Required("symmetricKey", "body", m.SymmetricKey); err != nil {
		return err
	}

	if m.SymmetricKey != nil {
		if err := m.SymmetricKey.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("symmetricKey")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SymmetricKeyAttestation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SymmetricKeyAttestation) UnmarshalBinary(b []byte) error {
	var res SymmetricKeyAttestation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}