// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// APIScore api score
//
// swagger:model apiScore
type APIScore struct {

	// final
	Final bool `json:"final,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// table
	Table []*APIScoreRow `json:"table"`

	// timestamps
	Timestamps *APITimestamps `json:"timestamps,omitempty"`
}

// Validate validates this api score
func (m *APIScore) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTable(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimestamps(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APIScore) validateTable(formats strfmt.Registry) error {

	if swag.IsZero(m.Table) { // not required
		return nil
	}

	for i := 0; i < len(m.Table); i++ {
		if swag.IsZero(m.Table[i]) { // not required
			continue
		}

		if m.Table[i] != nil {
			if err := m.Table[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("table" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *APIScore) validateTimestamps(formats strfmt.Registry) error {

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
func (m *APIScore) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIScore) UnmarshalBinary(b []byte) error {
	var res APIScore
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
