// Code generated by go-swagger; DO NOT EDIT.

package inventory

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/filanov/bm-inventory/models"
)

// ListImagesReader is a Reader for the ListImages structure.
type ListImagesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListImagesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListImagesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewListImagesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListImagesOK creates a ListImagesOK with default headers values
func NewListImagesOK() *ListImagesOK {
	return &ListImagesOK{}
}

/*ListImagesOK handles this case with default header values.

Image list
*/
type ListImagesOK struct {
	Payload models.ImageList
}

func (o *ListImagesOK) Error() string {
	return fmt.Sprintf("[GET /images][%d] listImagesOK  %+v", 200, o.Payload)
}

func (o *ListImagesOK) GetPayload() models.ImageList {
	return o.Payload
}

func (o *ListImagesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListImagesInternalServerError creates a ListImagesInternalServerError with default headers values
func NewListImagesInternalServerError() *ListImagesInternalServerError {
	return &ListImagesInternalServerError{}
}

/*ListImagesInternalServerError handles this case with default header values.

Internal server error
*/
type ListImagesInternalServerError struct {
}

func (o *ListImagesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /images][%d] listImagesInternalServerError ", 500)
}

func (o *ListImagesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
