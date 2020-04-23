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

// Coin coin
//
// swagger:model Coin
type Coin struct {

	// amount
	// Required: true
	Amount *string `json:"amount"`

	// denom
	// Required: true
	Denom *string `json:"denom"`
}

// Validate validates this coin
func (m *Coin) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDenom(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Coin) validateAmount(formats strfmt.Registry) error {

	if err := validate.Required("amount", "body", m.Amount); err != nil {
		return err
	}

	return nil
}

func (m *Coin) validateDenom(formats strfmt.Registry) error {

	if err := validate.Required("denom", "body", m.Denom); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Coin) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Coin) UnmarshalBinary(b []byte) error {
	var res Coin
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}