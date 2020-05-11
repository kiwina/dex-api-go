// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CandleStick candle stick
//
// swagger:model CandleStick
type CandleStick struct {

	// close price
	Close string `json:"close,omitempty"`

	// high price
	High string `json:"high,omitempty"`

	// low price
	Low string `json:"low,omitempty"`

	// market
	Market string `json:"market,omitempty"`

	// open price
	Open string `json:"open,omitempty"`

	// Minute:16/Hour:32/Day:48
	TimeSpan uint8 `json:"time_span,omitempty"`

	// total deal
	Total string `json:"total,omitempty"`

	// ending unix time
	UnixTime int64 `json:"unix_time,omitempty"`
}

// Validate validates this candle stick
func (m *CandleStick) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CandleStick) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CandleStick) UnmarshalBinary(b []byte) error {
	var res CandleStick
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
