// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// APIActiveResponseEntry api active response entry
//
// swagger:model apiActiveResponseEntry
type APIActiveResponseEntry struct {

	// entry Id
	EntryID string `json:"entryId,omitempty"`

	// position
	Position int32 `json:"position,omitempty"`
}

// Validate validates this api active response entry
func (m *APIActiveResponseEntry) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *APIActiveResponseEntry) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIActiveResponseEntry) UnmarshalBinary(b []byte) error {
	var res APIActiveResponseEntry
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}