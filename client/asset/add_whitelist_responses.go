// Code generated by go-swagger; DO NOT EDIT.

package asset

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/coinexchain/dex-api-go/models"
)

// AddWhitelistReader is a Reader for the AddWhitelist structure.
type AddWhitelistReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddWhitelistReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddWhitelistOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAddWhitelistBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAddWhitelistInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAddWhitelistOK creates a AddWhitelistOK with default headers values
func NewAddWhitelistOK() *AddWhitelistOK {
	return &AddWhitelistOK{}
}

/*AddWhitelistOK handles this case with default header values.

Add forbid whitelist result
*/
type AddWhitelistOK struct {
	Payload *models.StdTx
}

func (o *AddWhitelistOK) Error() string {
	return fmt.Sprintf("[POST /asset/tokens/{symbol}/forbidden/whitelist][%d] addWhitelistOK  %+v", 200, o.Payload)
}

func (o *AddWhitelistOK) GetPayload() *models.StdTx {
	return o.Payload
}

func (o *AddWhitelistOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StdTx)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddWhitelistBadRequest creates a AddWhitelistBadRequest with default headers values
func NewAddWhitelistBadRequest() *AddWhitelistBadRequest {
	return &AddWhitelistBadRequest{}
}

/*AddWhitelistBadRequest handles this case with default header values.

Invalid Request
*/
type AddWhitelistBadRequest struct {
}

func (o *AddWhitelistBadRequest) Error() string {
	return fmt.Sprintf("[POST /asset/tokens/{symbol}/forbidden/whitelist][%d] addWhitelistBadRequest ", 400)
}

func (o *AddWhitelistBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddWhitelistInternalServerError creates a AddWhitelistInternalServerError with default headers values
func NewAddWhitelistInternalServerError() *AddWhitelistInternalServerError {
	return &AddWhitelistInternalServerError{}
}

/*AddWhitelistInternalServerError handles this case with default header values.

Internal Server Error
*/
type AddWhitelistInternalServerError struct {
}

func (o *AddWhitelistInternalServerError) Error() string {
	return fmt.Sprintf("[POST /asset/tokens/{symbol}/forbidden/whitelist][%d] addWhitelistInternalServerError ", 500)
}

func (o *AddWhitelistInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}