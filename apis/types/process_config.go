// Code generated by go-swagger; DO NOT EDIT.

package types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ProcessConfig ExecProcessConfig holds information about the exec process.
// swagger:model ProcessConfig

type ProcessConfig struct {

	// arguments
	Arguments []string `json:"arguments"`

	// entrypoint
	Entrypoint string `json:"entrypoint,omitempty"`

	// privileged
	Privileged bool `json:"privileged,omitempty"`

	// tty
	Tty bool `json:"tty,omitempty"`

	// user
	User string `json:"user,omitempty"`
}

/* polymorph ProcessConfig arguments false */

/* polymorph ProcessConfig entrypoint false */

/* polymorph ProcessConfig privileged false */

/* polymorph ProcessConfig tty false */

/* polymorph ProcessConfig user false */

// Validate validates this process config
func (m *ProcessConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateArguments(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProcessConfig) validateArguments(formats strfmt.Registry) error {

	if swag.IsZero(m.Arguments) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProcessConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProcessConfig) UnmarshalBinary(b []byte) error {
	var res ProcessConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
