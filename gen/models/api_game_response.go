// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// APIGameResponse api game response
//
// swagger:model apiGameResponse
type APIGameResponse struct {

	// base
	Base *APIQueryResponseBase `json:"base,omitempty"`

	// object
	Object *APIGame `json:"object,omitempty"`
}

// Validate validates this api game response
func (m *APIGameResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBase(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateObject(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APIGameResponse) validateBase(formats strfmt.Registry) error {

	if swag.IsZero(m.Base) { // not required
		return nil
	}

	if m.Base != nil {
		if err := m.Base.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("base")
			}
			return err
		}
	}

	return nil
}

func (m *APIGameResponse) validateObject(formats strfmt.Registry) error {

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
func (m *APIGameResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIGameResponse) UnmarshalBinary(b []byte) error {
	var res APIGameResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
