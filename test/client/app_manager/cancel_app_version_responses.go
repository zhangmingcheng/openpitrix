// Code generated by go-swagger; DO NOT EDIT.

package app_manager

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"openpitrix.io/openpitrix/test/models"
)

// CancelAppVersionReader is a Reader for the CancelAppVersion structure.
type CancelAppVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CancelAppVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCancelAppVersionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCancelAppVersionOK creates a CancelAppVersionOK with default headers values
func NewCancelAppVersionOK() *CancelAppVersionOK {
	return &CancelAppVersionOK{}
}

/*CancelAppVersionOK handles this case with default header values.

A successful response.
*/
type CancelAppVersionOK struct {
	Payload *models.OpenpitrixCancelAppVersionResponse
}

func (o *CancelAppVersionOK) Error() string {
	return fmt.Sprintf("[POST /v1/app_version/action/cancel][%d] cancelAppVersionOK  %+v", 200, o.Payload)
}

func (o *CancelAppVersionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OpenpitrixCancelAppVersionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
