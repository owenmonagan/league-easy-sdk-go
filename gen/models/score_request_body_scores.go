// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ScoreRequestBodyScores score request body scores
//
// swagger:model ScoreRequestBodyScores
type ScoreRequestBodyScores struct {

	// game Id
	GameID string `json:"gameId,omitempty"`

	// object
	Object *APIScore `json:"object,omitempty"`
}

// Validate validates this score request body scores
func (m *ScoreRequestBodyScores) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateObject(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ScoreRequestBodyScores) validateObject(formats strfmt.Registry) error {

	if swag.IsZero(m.Object) { // not required
		return nil
	}

	if m.Object != nil {
		if err := m.Object.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("object")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ScoreRequestBodyScores) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ScoreRequestBodyScores) UnmarshalBinary(b []byte) error {
	var res ScoreRequestBodyScores
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
