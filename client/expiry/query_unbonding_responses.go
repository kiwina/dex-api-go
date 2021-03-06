// Code generated by go-swagger; DO NOT EDIT.

package expiry

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/coinexchain/dex-api-go/models"
)

// QueryUnbondingReader is a Reader for the QueryUnbonding structure.
type QueryUnbondingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QueryUnbondingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQueryUnbondingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewQueryUnbondingInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewQueryUnbondingOK creates a QueryUnbondingOK with default headers values
func NewQueryUnbondingOK() *QueryUnbondingOK {
	return &QueryUnbondingOK{}
}

/*QueryUnbondingOK handles this case with default header values.

OK
*/
type QueryUnbondingOK struct {
	Payload *QueryUnbondingOKBody
}

func (o *QueryUnbondingOK) Error() string {
	return fmt.Sprintf("[GET /expiry/unbondings][%d] queryUnbondingOK  %+v", 200, o.Payload)
}

func (o *QueryUnbondingOK) GetPayload() *QueryUnbondingOKBody {
	return o.Payload
}

func (o *QueryUnbondingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(QueryUnbondingOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryUnbondingInternalServerError creates a QueryUnbondingInternalServerError with default headers values
func NewQueryUnbondingInternalServerError() *QueryUnbondingInternalServerError {
	return &QueryUnbondingInternalServerError{}
}

/*QueryUnbondingInternalServerError handles this case with default header values.

Server internal error
*/
type QueryUnbondingInternalServerError struct {
}

func (o *QueryUnbondingInternalServerError) Error() string {
	return fmt.Sprintf("[GET /expiry/unbondings][%d] queryUnbondingInternalServerError ", 500)
}

func (o *QueryUnbondingInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*QueryUnbondingOKBody query unbonding o k body
swagger:model QueryUnbondingOKBody
*/
type QueryUnbondingOKBody struct {

	// data
	Data []*models.Unbonding `json:"data"`

	// timesid
	Timesid []int64 `json:"timesid"`
}

// Validate validates this query unbonding o k body
func (o *QueryUnbondingOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *QueryUnbondingOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	for i := 0; i < len(o.Data); i++ {
		if swag.IsZero(o.Data[i]) { // not required
			continue
		}

		if o.Data[i] != nil {
			if err := o.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("queryUnbondingOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *QueryUnbondingOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *QueryUnbondingOKBody) UnmarshalBinary(b []byte) error {
	var res QueryUnbondingOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
