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

// QueryRedelegationReader is a Reader for the QueryRedelegation structure.
type QueryRedelegationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QueryRedelegationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQueryRedelegationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewQueryRedelegationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewQueryRedelegationOK creates a QueryRedelegationOK with default headers values
func NewQueryRedelegationOK() *QueryRedelegationOK {
	return &QueryRedelegationOK{}
}

/*QueryRedelegationOK handles this case with default header values.

OK
*/
type QueryRedelegationOK struct {
	Payload *QueryRedelegationOKBody
}

func (o *QueryRedelegationOK) Error() string {
	return fmt.Sprintf("[GET /expiry/redelegations][%d] queryRedelegationOK  %+v", 200, o.Payload)
}

func (o *QueryRedelegationOK) GetPayload() *QueryRedelegationOKBody {
	return o.Payload
}

func (o *QueryRedelegationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(QueryRedelegationOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryRedelegationInternalServerError creates a QueryRedelegationInternalServerError with default headers values
func NewQueryRedelegationInternalServerError() *QueryRedelegationInternalServerError {
	return &QueryRedelegationInternalServerError{}
}

/*QueryRedelegationInternalServerError handles this case with default header values.

Server internal error
*/
type QueryRedelegationInternalServerError struct {
}

func (o *QueryRedelegationInternalServerError) Error() string {
	return fmt.Sprintf("[GET /expiry/redelegations][%d] queryRedelegationInternalServerError ", 500)
}

func (o *QueryRedelegationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*QueryRedelegationOKBody query redelegation o k body
swagger:model QueryRedelegationOKBody
*/
type QueryRedelegationOKBody struct {

	// data
	Data []*models.Redelegation `json:"data"`

	// timesid
	Timesid []int64 `json:"timesid"`
}

// Validate validates this query redelegation o k body
func (o *QueryRedelegationOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *QueryRedelegationOKBody) validateData(formats strfmt.Registry) error {

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
					return ve.ValidateName("queryRedelegationOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *QueryRedelegationOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *QueryRedelegationOKBody) UnmarshalBinary(b []byte) error {
	var res QueryRedelegationOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
