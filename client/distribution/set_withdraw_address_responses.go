// Code generated by go-swagger; DO NOT EDIT.

package distribution

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

// SetWithdrawAddressReader is a Reader for the SetWithdrawAddress structure.
type SetWithdrawAddressReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetWithdrawAddressReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSetWithdrawAddressOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSetWithdrawAddressBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewSetWithdrawAddressUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSetWithdrawAddressInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSetWithdrawAddressOK creates a SetWithdrawAddressOK with default headers values
func NewSetWithdrawAddressOK() *SetWithdrawAddressOK {
	return &SetWithdrawAddressOK{}
}

/*SetWithdrawAddressOK handles this case with default header values.

OK
*/
type SetWithdrawAddressOK struct {
	Payload *models.BroadcastTxCommitResult
}

func (o *SetWithdrawAddressOK) Error() string {
	return fmt.Sprintf("[POST /distribution/delegators/{delegatorAddr}/withdraw_address][%d] setWithdrawAddressOK  %+v", 200, o.Payload)
}

func (o *SetWithdrawAddressOK) GetPayload() *models.BroadcastTxCommitResult {
	return o.Payload
}

func (o *SetWithdrawAddressOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BroadcastTxCommitResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetWithdrawAddressBadRequest creates a SetWithdrawAddressBadRequest with default headers values
func NewSetWithdrawAddressBadRequest() *SetWithdrawAddressBadRequest {
	return &SetWithdrawAddressBadRequest{}
}

/*SetWithdrawAddressBadRequest handles this case with default header values.

Invalid delegator or withdraw address
*/
type SetWithdrawAddressBadRequest struct {
}

func (o *SetWithdrawAddressBadRequest) Error() string {
	return fmt.Sprintf("[POST /distribution/delegators/{delegatorAddr}/withdraw_address][%d] setWithdrawAddressBadRequest ", 400)
}

func (o *SetWithdrawAddressBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSetWithdrawAddressUnauthorized creates a SetWithdrawAddressUnauthorized with default headers values
func NewSetWithdrawAddressUnauthorized() *SetWithdrawAddressUnauthorized {
	return &SetWithdrawAddressUnauthorized{}
}

/*SetWithdrawAddressUnauthorized handles this case with default header values.

Key password is wrong
*/
type SetWithdrawAddressUnauthorized struct {
}

func (o *SetWithdrawAddressUnauthorized) Error() string {
	return fmt.Sprintf("[POST /distribution/delegators/{delegatorAddr}/withdraw_address][%d] setWithdrawAddressUnauthorized ", 401)
}

func (o *SetWithdrawAddressUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSetWithdrawAddressInternalServerError creates a SetWithdrawAddressInternalServerError with default headers values
func NewSetWithdrawAddressInternalServerError() *SetWithdrawAddressInternalServerError {
	return &SetWithdrawAddressInternalServerError{}
}

/*SetWithdrawAddressInternalServerError handles this case with default header values.

Internal Server Error
*/
type SetWithdrawAddressInternalServerError struct {
}

func (o *SetWithdrawAddressInternalServerError) Error() string {
	return fmt.Sprintf("[POST /distribution/delegators/{delegatorAddr}/withdraw_address][%d] setWithdrawAddressInternalServerError ", 500)
}

func (o *SetWithdrawAddressInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*SetWithdrawAddressBody set withdraw address body
swagger:model SetWithdrawAddressBody
*/
type SetWithdrawAddressBody struct {

	// base req
	BaseReq *models.BaseReq `json:"base_req,omitempty"`

	// withdraw address
	WithdrawAddress models.Address `json:"withdraw_address,omitempty"`
}

// Validate validates this set withdraw address body
func (o *SetWithdrawAddressBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateBaseReq(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateWithdrawAddress(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SetWithdrawAddressBody) validateBaseReq(formats strfmt.Registry) error {

	if swag.IsZero(o.BaseReq) { // not required
		return nil
	}

	if o.BaseReq != nil {
		if err := o.BaseReq.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Withdraw request body" + "." + "base_req")
			}
			return err
		}
	}

	return nil
}

func (o *SetWithdrawAddressBody) validateWithdrawAddress(formats strfmt.Registry) error {

	if swag.IsZero(o.WithdrawAddress) { // not required
		return nil
	}

	if err := o.WithdrawAddress.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Withdraw request body" + "." + "withdraw_address")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *SetWithdrawAddressBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SetWithdrawAddressBody) UnmarshalBinary(b []byte) error {
	var res SetWithdrawAddressBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
