// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixDeleteGroupsResponse openpitrix delete groups response
// swagger:model openpitrixDeleteGroupsResponse
type OpenpitrixDeleteGroupsResponse struct {

	// ids of group deleted
	GroupID []string `json:"group_id"`
}

// Validate validates this openpitrix delete groups response
func (m *OpenpitrixDeleteGroupsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGroupID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenpitrixDeleteGroupsResponse) validateGroupID(formats strfmt.Registry) error {

	if swag.IsZero(m.GroupID) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixDeleteGroupsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixDeleteGroupsResponse) UnmarshalBinary(b []byte) error {
	var res OpenpitrixDeleteGroupsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}