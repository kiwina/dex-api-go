// Code generated by go-swagger; DO NOT EDIT.

package misc

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// QueryLastBlockReader is a Reader for the QueryLastBlock structure.
type QueryLastBlockReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QueryLastBlockReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQueryLastBlockOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewQueryLastBlockInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewQueryLastBlockOK creates a QueryLastBlockOK with default headers values
func NewQueryLastBlockOK() *QueryLastBlockOK {
	return &QueryLastBlockOK{}
}

/*QueryLastBlockOK handles this case with default header values.

OK
*/
type QueryLastBlockOK struct {
	Payload *QueryLastBlockOKBody
}

func (o *QueryLastBlockOK) Error() string {
	return fmt.Sprintf("[GET /misc/height][%d] queryLastBlockOK  %+v", 200, o.Payload)
}

func (o *QueryLastBlockOK) GetPayload() *QueryLastBlockOKBody {
	return o.Payload
}

func (o *QueryLastBlockOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(QueryLastBlockOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryLastBlockInternalServerError creates a QueryLastBlockInternalServerError with default headers values
func NewQueryLastBlockInternalServerError() *QueryLastBlockInternalServerError {
	return &QueryLastBlockInternalServerError{}
}

/*QueryLastBlockInternalServerError handles this case with default header values.

Server internal error
*/
type QueryLastBlockInternalServerError struct {
}

func (o *QueryLastBlockInternalServerError) Error() string {
	return fmt.Sprintf("[GET /misc/height][%d] queryLastBlockInternalServerError ", 500)
}

func (o *QueryLastBlockInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*QueryLastBlockOKBody query last block o k body
swagger:model QueryLastBlockOKBody
*/
type QueryLastBlockOKBody struct {

	// Block height
	Height int64 `json:"height,omitempty"`

	// Block timestamp
	Timestamp int64 `json:"timestamp,omitempty"`
}

// Validate validates this query last block o k body
func (o *QueryLastBlockOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *QueryLastBlockOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *QueryLastBlockOKBody) UnmarshalBinary(b []byte) error {
	var res QueryLastBlockOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
