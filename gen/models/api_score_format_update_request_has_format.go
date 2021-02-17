// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// APIScoreFormatUpdateRequestHasFormat api score format update request has format
//
// swagger:model apiScoreFormatUpdateRequestHasFormat
type APIScoreFormatUpdateRequestHasFormat struct {

	// collection
	Collection string `json:"collection,omitempty"`

	// priority
	Priority int32 `json:"priority,omitempty"`

	// uuid
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this api score format update request has format
func (m *APIScoreFormatUpdateRequestHasFormat) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *APIScoreFormatUpdateRequestHasFormat) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIScoreFormatUpdateRequestHasFormat) UnmarshalBinary(b []byte) error {
	var res APIScoreFormatUpdateRequestHasFormat
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}