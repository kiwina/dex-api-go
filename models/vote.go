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

// Vote vote
//
// swagger:model Vote
type Vote struct {

	// option
	// Required: true
	Option *string `json:"option"`

	// proposal id
	// Required: true
	ProposalID *string `json:"proposal_id"`

	// voter
	// Required: true
	Voter *string `json:"voter"`
}

// Validate validates this vote
func (m *Vote) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOption(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProposalID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVoter(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Vote) validateOption(formats strfmt.Registry) error {

	if err := validate.Required("option", "body", m.Option); err != nil {
		return err
	}

	return nil
}

func (m *Vote) validateProposalID(formats strfmt.Registry) error {

	if err := validate.Required("proposal_id", "body", m.ProposalID); err != nil {
		return err
	}

	return nil
}

func (m *Vote) validateVoter(formats strfmt.Registry) error {

	if err := validate.Required("voter", "body", m.Voter); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Vote) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Vote) UnmarshalBinary(b []byte) error {
	var res Vote
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}