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

// APIQueryResponseBase api query response base
//
// swagger:model apiQueryResponseBase
type APIQueryResponseBase struct {

	// entry list
	EntryList []*APIPosition `json:"entryList"`

	// score list
	ScoreList []*APIPosition `json:"scoreList"`

	// status
	Status *APIStatus `json:"status,omitempty"`
}

// Validate validates this api query response base
func (m *APIQueryResponseBase) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEntryList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScoreList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APIQueryResponseBase) validateEntryList(formats strfmt.Registry) error {

	if swag.IsZero(m.EntryList) { // not required
		return nil
	}

	for i := 0; i < len(m.EntryList); i++ {
		if swag.IsZero(m.EntryList[i]) { // not required
			continue
		}

		if m.EntryList[i] != nil {
			if err := m.EntryList[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("entryList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *APIQueryResponseBase) validateScoreList(formats strfmt.Registry) error {

	if swag.IsZero(m.ScoreList) { // not required
		return nil
	}

	for i := 0; i < len(m.ScoreList); i++ {
		if swag.IsZero(m.ScoreList[i]) { // not required
			continue
		}

		if m.ScoreList[i] != nil {
			if err := m.ScoreList[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("scoreList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *APIQueryResponseBase) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if m.Status != nil {
		if err := m.Status.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *APIQueryResponseBase) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIQueryResponseBase) UnmarshalBinary(b []byte) error {
	var res APIQueryResponseBase
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
