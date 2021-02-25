// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// APIGame api game
//
// swagger:model apiGame
type APIGame struct {

	// id
	ID string `json:"id,omitempty"`

	// position
	Position int32 `json:"position,omitempty"`

	// timestamps
	Timestamps *APITimestamps `json:"timestamps,omitempty"`
}

// Validate validates this api game
func (m *APIGame) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTimestamps(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APIGame) validateTimestamps(formats strfmt.Registry) error {

	if swag.IsZero(m.Timestamps) { // not required
		return nil
	}

	if m.Timestamps != nil {
		if err := m.Timestamps.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("timestamps")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *APIGame) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIGame) UnmarshalBinary(b []byte) error {
	var res APIGame
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
