// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// APIGroup api group
//
// swagger:model apiGroup
type APIGroup struct {

	// id
	ID string `json:"id,omitempty"`

	// query positions
	QueryPositions []int32 `json:"queryPositions"`

	// repetitions
	Repetitions int32 `json:"repetitions,omitempty"`

	// timestamps
	Timestamps *APITimestamps `json:"timestamps,omitempty"`

	// total entries
	TotalEntries int32 `json:"totalEntries,omitempty"`

	// total qualifiers
	TotalQualifiers int32 `json:"totalQualifiers,omitempty"`
}

// Validate validates this api group
func (m *APIGroup) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTimestamps(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APIGroup) validateTimestamps(formats strfmt.Registry) error {

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
func (m *APIGroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIGroup) UnmarshalBinary(b []byte) error {
	var res APIGroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
