// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/filanov/bm-inventory/models"
)

// ListClustersReader is a Reader for the ListClusters structure.
type ListClustersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListClustersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListClustersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListClustersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListClustersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListClustersOK creates a ListClustersOK with default headers values
func NewListClustersOK() *ListClustersOK {
	return &ListClustersOK{}
}

/*ListClustersOK handles this case with default header values.

Success.
*/
type ListClustersOK struct {
	Payload models.ClusterList
}

func (o *ListClustersOK) Error() string {
	return fmt.Sprintf("[GET /clusters][%d] listClustersOK  %+v", 200, o.Payload)
}

func (o *ListClustersOK) GetPayload() models.ClusterList {
	return o.Payload
}

func (o *ListClustersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListClustersUnauthorized creates a ListClustersUnauthorized with default headers values
func NewListClustersUnauthorized() *ListClustersUnauthorized {
	return &ListClustersUnauthorized{}
}

/*ListClustersUnauthorized handles this case with default header values.

Error.
*/
type ListClustersUnauthorized struct {
	Payload *models.Error
}

func (o *ListClustersUnauthorized) Error() string {
	return fmt.Sprintf("[GET /clusters][%d] listClustersUnauthorized  %+v", 401, o.Payload)
}

func (o *ListClustersUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListClustersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListClustersInternalServerError creates a ListClustersInternalServerError with default headers values
func NewListClustersInternalServerError() *ListClustersInternalServerError {
	return &ListClustersInternalServerError{}
}

/*ListClustersInternalServerError handles this case with default header values.

Error.
*/
type ListClustersInternalServerError struct {
	Payload *models.Error
}

func (o *ListClustersInternalServerError) Error() string {
	return fmt.Sprintf("[GET /clusters][%d] listClustersInternalServerError  %+v", 500, o.Payload)
}

func (o *ListClustersInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListClustersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
