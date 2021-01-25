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

// PowerAutomateAction power automate action
//
// swagger:model PowerAutomateAction
type PowerAutomateAction struct {
	displayNameField string

	idField string

	rulesField []string
}

// DisplayName gets the display name of this subtype
func (m *PowerAutomateAction) DisplayName() string {
	return m.displayNameField
}

// SetDisplayName sets the display name of this subtype
func (m *PowerAutomateAction) SetDisplayName(val string) {
	m.displayNameField = val
}

// ID gets the id of this subtype
func (m *PowerAutomateAction) ID() string {
	return m.idField
}

// SetID sets the id of this subtype
func (m *PowerAutomateAction) SetID(val string) {
	m.idField = val
}

// Rules gets the rules of this subtype
func (m *PowerAutomateAction) Rules() []string {
	return m.rulesField
}

// SetRules sets the rules of this subtype
func (m *PowerAutomateAction) SetRules(val []string) {
	m.rulesField = val
}

// Type gets the type of this subtype
func (m *PowerAutomateAction) Type() string {
	return "PowerAutomateAction"
}

// SetType sets the type of this subtype
func (m *PowerAutomateAction) SetType(val string) {
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *PowerAutomateAction) UnmarshalJSON(raw []byte) error {
	var data struct {
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var base struct {
		/* Just the base type fields. Used for unmashalling polymorphic types.*/

		DisplayName string `json:"displayName,omitempty"`

		ID string `json:"id,omitempty"`

		Rules []string `json:"rules"`

		Type string `json:"type"`
	}
	buf = bytes.NewBuffer(raw)
	dec = json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&base); err != nil {
		return err
	}

	var result PowerAutomateAction

	result.displayNameField = base.DisplayName

	result.idField = base.ID

	result.rulesField = base.Rules

	if base.Type != result.Type() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid type value: %q", base.Type)
	}

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m PowerAutomateAction) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {
	}{})
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		DisplayName string `json:"displayName,omitempty"`

		ID string `json:"id,omitempty"`

		Rules []string `json:"rules"`

		Type string `json:"type"`
	}{

		DisplayName: m.DisplayName(),

		ID: m.ID(),

		Rules: m.Rules(),

		Type: m.Type(),
	})
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this power automate action
func (m *PowerAutomateAction) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRules(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PowerAutomateAction) validateRules(formats strfmt.Registry) error {

	if err := validate.Required("rules", "body", m.Rules()); err != nil {
		return err
	}

	iRulesSize := int64(len(m.Rules()))

	if err := validate.MinItems("rules", "body", iRulesSize, 1); err != nil {
		return err
	}

	if err := validate.MaxItems("rules", "body", iRulesSize, 1); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PowerAutomateAction) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PowerAutomateAction) UnmarshalBinary(b []byte) error {
	var res PowerAutomateAction
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}