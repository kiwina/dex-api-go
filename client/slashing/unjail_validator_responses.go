// Code generated by go-swagger; DO NOT EDIT.

package slashing

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

// UnjailValidatorReader is a Reader for the UnjailValidator structure.
type UnjailValidatorReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UnjailValidatorReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUnjailValidatorOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUnjailValidatorBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUnjailValidatorInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUnjailValidatorOK creates a UnjailValidatorOK with default headers values
func NewUnjailValidatorOK() *UnjailValidatorOK {
	return &UnjailValidatorOK{}
}

/*UnjailValidatorOK handles this case with default header values.

Tx was succesfully generated
*/
type UnjailValidatorOK struct {
	Payload *models.BroadcastTxCommitResult
}

func (o *UnjailValidatorOK) Error() string {
	return fmt.Sprintf("[POST /slashing/validators/{validatorAddr}/unjail][%d] unjailValidatorOK  %+v", 200, o.Payload)
}

func (o *UnjailValidatorOK) GetPayload() *models.BroadcastTxCommitResult {
	return o.Payload
}

func (o *UnjailValidatorOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BroadcastTxCommitResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUnjailValidatorBadRequest creates a UnjailValidatorBadRequest with default headers values
func NewUnjailValidatorBadRequest() *UnjailValidatorBadRequest {
	return &UnjailValidatorBadRequest{}
}

/*UnjailValidatorBadRequest handles this case with default header values.

Invalid validator address or base_req
*/
type UnjailValidatorBadRequest struct {
}

func (o *UnjailValidatorBadRequest) Error() string {
	return fmt.Sprintf("[POST /slashing/validators/{validatorAddr}/unjail][%d] unjailValidatorBadRequest ", 400)
}

func (o *UnjailValidatorBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUnjailValidatorInternalServerError creates a UnjailValidatorInternalServerError with default headers values
func NewUnjailValidatorInternalServerError() *UnjailValidatorInternalServerError {
	return &UnjailValidatorInternalServerError{}
}

/*UnjailValidatorInternalServerError handles this case with default header values.

Internal Server Error
*/
type UnjailValidatorInternalServerError struct {
}

func (o *UnjailValidatorInternalServerError) Error() string {
	return fmt.Sprintf("[POST /slashing/validators/{validatorAddr}/unjail][%d] unjailValidatorInternalServerError ", 500)
}

func (o *UnjailValidatorInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*UnjailValidatorBody unjail validator body
swagger:model UnjailValidatorBody
*/
type UnjailValidatorBody struct {

	// base req
	BaseReq *models.BaseReq `json:"base_req,omitempty"`
}

// Validate validates this unjail validator body
func (o *UnjailValidatorBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateBaseReq(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UnjailValidatorBody) validateBaseReq(formats strfmt.Registry) error {

	if swag.IsZero(o.BaseReq) { // not required
		return nil
	}

	if o.BaseReq != nil {
		if err := o.BaseReq.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("UnjailBody" + "." + "base_req")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UnjailValidatorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UnjailValidatorBody) UnmarshalBinary(b []byte) error {
	var res UnjailValidatorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
