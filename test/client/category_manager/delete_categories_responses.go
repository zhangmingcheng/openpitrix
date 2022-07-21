// Code generated by go-swagger; DO NOT EDIT.

package category_manager

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"openpitrix.io/openpitrix/test/models"
)

// DeleteCategoriesReader is a Reader for the DeleteCategories structure.
type DeleteCategoriesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteCategoriesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteCategoriesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteCategoriesOK creates a DeleteCategoriesOK with default headers values
func NewDeleteCategoriesOK() *DeleteCategoriesOK {
	return &DeleteCategoriesOK{}
}

/*DeleteCategoriesOK handles this case with default header values.

A successful response.
*/
type DeleteCategoriesOK struct {
	Payload *models.OpenpitrixDeleteCategoriesResponse
}

func (o *DeleteCategoriesOK) Error() string {
	return fmt.Sprintf("[DELETE /v1/categories][%d] deleteCategoriesOK  %+v", 200, o.Payload)
}

func (o *DeleteCategoriesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OpenpitrixDeleteCategoriesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}