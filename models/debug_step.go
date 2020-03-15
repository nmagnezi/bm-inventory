// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DebugStep debug step
//
// swagger:model debug-step
type DebugStep struct {

	// command
	// Required: true
	Command *string `json:"command"`
}

// Validate validates this debug step
func (m *DebugStep) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCommand(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DebugStep) validateCommand(formats strfmt.Registry) error {

	if err := validate.Required("command", "body", m.Command); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DebugStep) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DebugStep) UnmarshalBinary(b []byte) error {
	var res DebugStep
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
