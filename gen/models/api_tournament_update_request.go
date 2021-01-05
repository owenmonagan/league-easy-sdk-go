// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// APITournamentUpdateRequest api tournament update request
//
// swagger:model apiTournamentUpdateRequest
type APITournamentUpdateRequest struct {

	// compute
	Compute bool `json:"compute,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// uuid
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this api tournament update request
func (m *APITournamentUpdateRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *APITournamentUpdateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APITournamentUpdateRequest) UnmarshalBinary(b []byte) error {
	var res APITournamentUpdateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
