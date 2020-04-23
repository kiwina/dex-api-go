// Code generated by go-swagger; DO NOT EDIT.

package governance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetVotingParametersReader is a Reader for the GetVotingParameters structure.
type GetVotingParametersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVotingParametersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetVotingParametersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetVotingParametersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetVotingParametersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetVotingParametersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetVotingParametersOK creates a GetVotingParametersOK with default headers values
func NewGetVotingParametersOK() *GetVotingParametersOK {
	return &GetVotingParametersOK{}
}

/*GetVotingParametersOK handles this case with default header values.

OK
*/
type GetVotingParametersOK struct {
	Payload *GetVotingParametersOKBody
}

func (o *GetVotingParametersOK) Error() string {
	return fmt.Sprintf("[GET /gov/parameters/voting][%d] getVotingParametersOK  %+v", 200, o.Payload)
}

func (o *GetVotingParametersOK) GetPayload() *GetVotingParametersOKBody {
	return o.Payload
}

func (o *GetVotingParametersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetVotingParametersOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVotingParametersBadRequest creates a GetVotingParametersBadRequest with default headers values
func NewGetVotingParametersBadRequest() *GetVotingParametersBadRequest {
	return &GetVotingParametersBadRequest{}
}

/*GetVotingParametersBadRequest handles this case with default header values.

<other_path> is not a valid query request path
*/
type GetVotingParametersBadRequest struct {
}

func (o *GetVotingParametersBadRequest) Error() string {
	return fmt.Sprintf("[GET /gov/parameters/voting][%d] getVotingParametersBadRequest ", 400)
}

func (o *GetVotingParametersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetVotingParametersNotFound creates a GetVotingParametersNotFound with default headers values
func NewGetVotingParametersNotFound() *GetVotingParametersNotFound {
	return &GetVotingParametersNotFound{}
}

/*GetVotingParametersNotFound handles this case with default header values.

Found no voting parameters
*/
type GetVotingParametersNotFound struct {
}

func (o *GetVotingParametersNotFound) Error() string {
	return fmt.Sprintf("[GET /gov/parameters/voting][%d] getVotingParametersNotFound ", 404)
}

func (o *GetVotingParametersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetVotingParametersInternalServerError creates a GetVotingParametersInternalServerError with default headers values
func NewGetVotingParametersInternalServerError() *GetVotingParametersInternalServerError {
	return &GetVotingParametersInternalServerError{}
}

/*GetVotingParametersInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetVotingParametersInternalServerError struct {
}

func (o *GetVotingParametersInternalServerError) Error() string {
	return fmt.Sprintf("[GET /gov/parameters/voting][%d] getVotingParametersInternalServerError ", 500)
}

func (o *GetVotingParametersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*GetVotingParametersOKBody get voting parameters o k body
swagger:model GetVotingParametersOKBody
*/
type GetVotingParametersOKBody struct {

	// height
	Height string `json:"height,omitempty"`

	// result
	Result *GetVotingParametersOKBodyResult `json:"result,omitempty"`
}

// Validate validates this get voting parameters o k body
func (o *GetVotingParametersOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetVotingParametersOKBody) validateResult(formats strfmt.Registry) error {

	if swag.IsZero(o.Result) { // not required
		return nil
	}

	if o.Result != nil {
		if err := o.Result.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getVotingParametersOK" + "." + "result")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetVotingParametersOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetVotingParametersOKBody) UnmarshalBinary(b []byte) error {
	var res GetVotingParametersOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetVotingParametersOKBodyResult get voting parameters o k body result
swagger:model GetVotingParametersOKBodyResult
*/
type GetVotingParametersOKBodyResult struct {

	// voting period
	VotingPeriod string `json:"voting_period,omitempty"`
}

// Validate validates this get voting parameters o k body result
func (o *GetVotingParametersOKBodyResult) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetVotingParametersOKBodyResult) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetVotingParametersOKBodyResult) UnmarshalBinary(b []byte) error {
	var res GetVotingParametersOKBodyResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}