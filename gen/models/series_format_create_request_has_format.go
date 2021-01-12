// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SeriesFormatCreateRequestHasFormat series format create request has format
//
// swagger:model SeriesFormatCreateRequestHasFormat
type SeriesFormatCreateRequestHasFormat struct {

	// collection
	Collection string `json:"collection,omitempty"`

	// uuid
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this series format create request has format
func (m *SeriesFormatCreateRequestHasFormat) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SeriesFormatCreateRequestHasFormat) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SeriesFormatCreateRequestHasFormat) UnmarshalBinary(b []byte) error {
	var res SeriesFormatCreateRequestHasFormat
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
