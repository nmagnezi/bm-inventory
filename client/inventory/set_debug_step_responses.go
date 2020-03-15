// Code generated by go-swagger; DO NOT EDIT.

package inventory

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// SetDebugStepReader is a Reader for the SetDebugStep structure.
type SetDebugStepReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetDebugStepReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSetDebugStepOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewSetDebugStepInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSetDebugStepOK creates a SetDebugStepOK with default headers values
func NewSetDebugStepOK() *SetDebugStepOK {
	return &SetDebugStepOK{}
}

/*SetDebugStepOK handles this case with default header values.

Registered
*/
type SetDebugStepOK struct {
}

func (o *SetDebugStepOK) Error() string {
	return fmt.Sprintf("[POST /debug][%d] setDebugStepOK ", 200)
}

func (o *SetDebugStepOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSetDebugStepInternalServerError creates a SetDebugStepInternalServerError with default headers values
func NewSetDebugStepInternalServerError() *SetDebugStepInternalServerError {
	return &SetDebugStepInternalServerError{}
}

/*SetDebugStepInternalServerError handles this case with default header values.

Internal server error
*/
type SetDebugStepInternalServerError struct {
}

func (o *SetDebugStepInternalServerError) Error() string {
	return fmt.Sprintf("[POST /debug][%d] setDebugStepInternalServerError ", 500)
}

func (o *SetDebugStepInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
