// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// APIStageSingleElimination api stage single elimination
//
// swagger:model apiStageSingleElimination
type APIStageSingleElimination struct {

	// created at
	CreatedAt string `json:"createdAt,omitempty"`

	// max entries
	MaxEntries int32 `json:"maxEntries,omitempty"`

	// max entries enabled
	MaxEntriesEnabled bool `json:"maxEntriesEnabled,omitempty"`

	// max qualifying round position
	MaxQualifyingRoundPosition int32 `json:"maxQualifyingRoundPosition,omitempty"`

	// updated at
	UpdatedAt string `json:"updatedAt,omitempty"`

	// uuid
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this api stage single elimination
func (m *APIStageSingleElimination) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *APIStageSingleElimination) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIStageSingleElimination) UnmarshalBinary(b []byte) error {
	var res APIStageSingleElimination
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
