// Code generated by go-swagger; DO NOT EDIT.

package market

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/coinexchain/dex-api-go/models"
)

// CreateTradingPairReader is a Reader for the CreateTradingPair structure.
type CreateTradingPairReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateTradingPairReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateTradingPairOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateTradingPairBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateTradingPairInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateTradingPairOK creates a CreateTradingPairOK with default headers values
func NewCreateTradingPairOK() *CreateTradingPairOK {
	return &CreateTradingPairOK{}
}

/*CreateTradingPairOK handles this case with default header values.

Tx was succesfully generated
*/
type CreateTradingPairOK struct {
	Payload *models.StdTx
}

func (o *CreateTradingPairOK) Error() string {
	return fmt.Sprintf("[POST /market/trading-pairs][%d] createTradingPairOK  %+v", 200, o.Payload)
}

func (o *CreateTradingPairOK) GetPayload() *models.StdTx {
	return o.Payload
}

func (o *CreateTradingPairOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StdTx)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateTradingPairBadRequest creates a CreateTradingPairBadRequest with default headers values
func NewCreateTradingPairBadRequest() *CreateTradingPairBadRequest {
	return &CreateTradingPairBadRequest{}
}

/*CreateTradingPairBadRequest handles this case with default header values.

Invalid request
*/
type CreateTradingPairBadRequest struct {
}

func (o *CreateTradingPairBadRequest) Error() string {
	return fmt.Sprintf("[POST /market/trading-pairs][%d] createTradingPairBadRequest ", 400)
}

func (o *CreateTradingPairBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateTradingPairInternalServerError creates a CreateTradingPairInternalServerError with default headers values
func NewCreateTradingPairInternalServerError() *CreateTradingPairInternalServerError {
	return &CreateTradingPairInternalServerError{}
}

/*CreateTradingPairInternalServerError handles this case with default header values.

Server internal error
*/
type CreateTradingPairInternalServerError struct {
}

func (o *CreateTradingPairInternalServerError) Error() string {
	return fmt.Sprintf("[POST /market/trading-pairs][%d] createTradingPairInternalServerError ", 500)
}

func (o *CreateTradingPairInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*CreateTradingPairBody create trading pair body
swagger:model CreateTradingPairBody
*/
type CreateTradingPairBody struct {

	// base req
	BaseReq *models.BaseReq `json:"base_req,omitempty"`

	models.BaseMarket
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *CreateTradingPairBody) UnmarshalJSON(raw []byte) error {
	// CreateTradingPairParamsBodyAO0
	var dataCreateTradingPairParamsBodyAO0 struct {
		BaseReq *models.BaseReq `json:"base_req,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataCreateTradingPairParamsBodyAO0); err != nil {
		return err
	}

	o.BaseReq = dataCreateTradingPairParamsBodyAO0.BaseReq

	// CreateTradingPairParamsBodyAO1
	var createTradingPairParamsBodyAO1 models.BaseMarket
	if err := swag.ReadJSON(raw, &createTradingPairParamsBodyAO1); err != nil {
		return err
	}
	o.BaseMarket = createTradingPairParamsBodyAO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o CreateTradingPairBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	var dataCreateTradingPairParamsBodyAO0 struct {
		BaseReq *models.BaseReq `json:"base_req,omitempty"`
	}

	dataCreateTradingPairParamsBodyAO0.BaseReq = o.BaseReq

	jsonDataCreateTradingPairParamsBodyAO0, errCreateTradingPairParamsBodyAO0 := swag.WriteJSON(dataCreateTradingPairParamsBodyAO0)
	if errCreateTradingPairParamsBodyAO0 != nil {
		return nil, errCreateTradingPairParamsBodyAO0
	}
	_parts = append(_parts, jsonDataCreateTradingPairParamsBodyAO0)

	createTradingPairParamsBodyAO1, err := swag.WriteJSON(o.BaseMarket)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, createTradingPairParamsBodyAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this create trading pair body
func (o *CreateTradingPairBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateBaseReq(formats); err != nil {
		res = append(res, err)
	}

	// validation for a type composition with models.BaseMarket
	if err := o.BaseMarket.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateTradingPairBody) validateBaseReq(formats strfmt.Registry) error {

	if swag.IsZero(o.BaseReq) { // not required
		return nil
	}

	if o.BaseReq != nil {
		if err := o.BaseReq.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "base_req")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreateTradingPairBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateTradingPairBody) UnmarshalBinary(b []byte) error {
	var res CreateTradingPairBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}