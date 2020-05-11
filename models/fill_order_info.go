// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// FillOrderInfo fill order info
//
// swagger:model FillOrderInfo
type FillOrderInfo struct {

	// Current block deal money amount
	CurrMoney int64 `json:"curr_money,omitempty"`

	// Current block deal stock amount
	CurrStock int64 `json:"curr_stock,omitempty"`

	// Order accumulated deal money amount
	DealMoney int64 `json:"deal_money,omitempty"`

	// Order accumulated deal stock amount
	DealStock int64 `json:"deal_stock,omitempty"`

	// Freeze sato.CET amount
	Freeze int64 `json:"freeze,omitempty"`

	// height
	Height int64 `json:"height,omitempty"`

	// Order left stock
	LeftStock int64 `json:"left_stock,omitempty"`

	// order id
	OrderID string `json:"order_id,omitempty"`

	// price
	Price string `json:"price,omitempty"`

	// BUY:1/SELL:2
	Side int64 `json:"side,omitempty"`

	// trading pair
	TradingPair string `json:"trading_pair,omitempty"`
}

// Validate validates this fill order info
func (m *FillOrderInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FillOrderInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FillOrderInfo) UnmarshalBinary(b []byte) error {
	var res FillOrderInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
