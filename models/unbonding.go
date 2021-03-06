// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Unbonding unbonding
//
// swagger:model Unbonding
type Unbonding struct {

	// amount
	Amount string `json:"amount,omitempty"`

	// completion time
	CompletionTime string `json:"completion_time,omitempty"`

	// delegator
	Delegator Address `json:"delegator,omitempty"`

	// The tx hash
	TxHash string `json:"tx_hash,omitempty"`

	// validator
	Validator Address `json:"validator,omitempty"`
}

// Validate validates this unbonding
func (m *Unbonding) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDelegator(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValidator(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Unbonding) validateDelegator(formats strfmt.Registry) error {

	if swag.IsZero(m.Delegator) { // not required
		return nil
	}

	if err := m.Delegator.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("delegator")
		}
		return err
	}

	return nil
}

func (m *Unbonding) validateValidator(formats strfmt.Registry) error {

	if swag.IsZero(m.Validator) { // not required
		return nil
	}

	if err := m.Validator.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("validator")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Unbonding) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Unbonding) UnmarshalBinary(b []byte) error {
	var res Unbonding
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
