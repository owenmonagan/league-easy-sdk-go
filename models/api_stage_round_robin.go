// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// APIStageRoundRobin api stage round robin
//
// swagger:model apiStageRoundRobin
type APIStageRoundRobin struct {

	// id
	ID string `json:"id,omitempty"`

	// max entries per group
	MaxEntriesPerGroup int32 `json:"maxEntriesPerGroup,omitempty"`

	// max qualifiers per group
	MaxQualifiersPerGroup int32 `json:"maxQualifiersPerGroup,omitempty"`

	// min number of repetitions
	MinNumberOfRepetitions int32 `json:"minNumberOfRepetitions,omitempty"`

	// timestamps
	Timestamps *APITimestamps `json:"timestamps,omitempty"`
}

// Validate validates this api stage round robin
func (m *APIStageRoundRobin) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTimestamps(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APIStageRoundRobin) validateTimestamps(formats strfmt.Registry) error {

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
func (m *APIStageRoundRobin) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIStageRoundRobin) UnmarshalBinary(b []byte) error {
	var res APIStageRoundRobin
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}