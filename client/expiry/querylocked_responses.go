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

// QuerylockedReader is a Reader for the Querylocked structure.
type QuerylockedReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QuerylockedReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQuerylockedOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewQuerylockedInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewQuerylockedOK creates a QuerylockedOK with default headers values
func NewQuerylockedOK() *QuerylockedOK {
	return &QuerylockedOK{}
}

/*QuerylockedOK handles this case with default header values.

OK
*/
type QuerylockedOK struct {
	Payload *QuerylockedOKBody
}

func (o *QuerylockedOK) Error() string {
	return fmt.Sprintf("[GET /expiry/lockeds][%d] querylockedOK  %+v", 200, o.Payload)
}

func (o *QuerylockedOK) GetPayload() *QuerylockedOKBody {
	return o.Payload
}

func (o *QuerylockedOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(QuerylockedOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuerylockedInternalServerError creates a QuerylockedInternalServerError with default headers values
func NewQuerylockedInternalServerError() *QuerylockedInternalServerError {
	return &QuerylockedInternalServerError{}
}

/*QuerylockedInternalServerError handles this case with default header values.

Server internal error
*/
type QuerylockedInternalServerError struct {
}

func (o *QuerylockedInternalServerError) Error() string {
	return fmt.Sprintf("[GET /expiry/lockeds][%d] querylockedInternalServerError ", 500)
}

func (o *QuerylockedInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*QuerylockedOKBody querylocked o k body
swagger:model QuerylockedOKBody
*/
type QuerylockedOKBody struct {

	// data
	Data []*models.Locked `json:"data"`

	// timesid
	Timesid []int64 `json:"timesid"`
}

// Validate validates this querylocked o k body
func (o *QuerylockedOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *QuerylockedOKBody) validateData(formats strfmt.Registry) error {

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
					return ve.ValidateName("querylockedOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *QuerylockedOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *QuerylockedOKBody) UnmarshalBinary(b []byte) error {
	var res QuerylockedOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
