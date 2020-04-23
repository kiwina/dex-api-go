// Code generated by go-swagger; DO NOT EDIT.

package bancorlite

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/coinexchain/dex-api-go/models"
)

// BancorInitReader is a Reader for the BancorInit structure.
type BancorInitReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BancorInitReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBancorInitOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewBancorInitBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewBancorInitInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewBancorInitOK creates a BancorInitOK with default headers values
func NewBancorInitOK() *BancorInitOK {
	return &BancorInitOK{}
}

/*BancorInitOK handles this case with default header values.

Tx was succesfully generated
*/
type BancorInitOK struct {
	Payload *models.StdTx
}

func (o *BancorInitOK) Error() string {
	return fmt.Sprintf("[POST /bancorlite/bancor-init][%d] bancorInitOK  %+v", 200, o.Payload)
}

func (o *BancorInitOK) GetPayload() *models.StdTx {
	return o.Payload
}

func (o *BancorInitOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StdTx)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBancorInitBadRequest creates a BancorInitBadRequest with default headers values
func NewBancorInitBadRequest() *BancorInitBadRequest {
	return &BancorInitBadRequest{}
}

/*BancorInitBadRequest handles this case with default header values.

Invalid request
*/
type BancorInitBadRequest struct {
}

func (o *BancorInitBadRequest) Error() string {
	return fmt.Sprintf("[POST /bancorlite/bancor-init][%d] bancorInitBadRequest ", 400)
}

func (o *BancorInitBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewBancorInitInternalServerError creates a BancorInitInternalServerError with default headers values
func NewBancorInitInternalServerError() *BancorInitInternalServerError {
	return &BancorInitInternalServerError{}
}

/*BancorInitInternalServerError handles this case with default header values.

Server internal error
*/
type BancorInitInternalServerError struct {
}

func (o *BancorInitInternalServerError) Error() string {
	return fmt.Sprintf("[POST /bancorlite/bancor-init][%d] bancorInitInternalServerError ", 500)
}

func (o *BancorInitInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*BancorInitBody bancor init body
swagger:model BancorInitBody
*/
type BancorInitBody struct {

	// base req
	// Required: true
	BaseReq *models.BaseReq `json:"base_req"`

	// earliest cancel time
	// Required: true
	EarliestCancelTime *string `json:"earliest_cancel_time"`

	// init price
	// Required: true
	InitPrice *string `json:"init_price"`

	// max money
	// Required: true
	MaxMoney *string `json:"max_money"`

	// max price
	// Required: true
	MaxPrice *string `json:"max_price"`

	// max supply
	// Required: true
	MaxSupply *string `json:"max_supply"`

	// money
	// Required: true
	Money *string `json:"money"`

	// stock
	// Required: true
	Stock *string `json:"stock"`

	// stock precision
	StockPrecision string `json:"stock_precision,omitempty"`
}

// Validate validates this bancor init body
func (o *BancorInitBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateBaseReq(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateEarliestCancelTime(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateInitPrice(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateMaxMoney(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateMaxPrice(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateMaxSupply(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateMoney(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateStock(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *BancorInitBody) validateBaseReq(formats strfmt.Registry) error {

	if err := validate.Required("bancorInit"+"."+"base_req", "body", o.BaseReq); err != nil {
		return err
	}

	if o.BaseReq != nil {
		if err := o.BaseReq.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bancorInit" + "." + "base_req")
			}
			return err
		}
	}

	return nil
}

func (o *BancorInitBody) validateEarliestCancelTime(formats strfmt.Registry) error {

	if err := validate.Required("bancorInit"+"."+"earliest_cancel_time", "body", o.EarliestCancelTime); err != nil {
		return err
	}

	return nil
}

func (o *BancorInitBody) validateInitPrice(formats strfmt.Registry) error {

	if err := validate.Required("bancorInit"+"."+"init_price", "body", o.InitPrice); err != nil {
		return err
	}

	return nil
}

func (o *BancorInitBody) validateMaxMoney(formats strfmt.Registry) error {

	if err := validate.Required("bancorInit"+"."+"max_money", "body", o.MaxMoney); err != nil {
		return err
	}

	return nil
}

func (o *BancorInitBody) validateMaxPrice(formats strfmt.Registry) error {

	if err := validate.Required("bancorInit"+"."+"max_price", "body", o.MaxPrice); err != nil {
		return err
	}

	return nil
}

func (o *BancorInitBody) validateMaxSupply(formats strfmt.Registry) error {

	if err := validate.Required("bancorInit"+"."+"max_supply", "body", o.MaxSupply); err != nil {
		return err
	}

	return nil
}

func (o *BancorInitBody) validateMoney(formats strfmt.Registry) error {

	if err := validate.Required("bancorInit"+"."+"money", "body", o.Money); err != nil {
		return err
	}

	return nil
}

func (o *BancorInitBody) validateStock(formats strfmt.Registry) error {

	if err := validate.Required("bancorInit"+"."+"stock", "body", o.Stock); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *BancorInitBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *BancorInitBody) UnmarshalBinary(b []byte) error {
	var res BancorInitBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}