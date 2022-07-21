// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixUser openpitrix user
// swagger:model openpitrixUser
type OpenpitrixUser struct {

	// the time when user create
	CreateTime strfmt.DateTime `json:"create_time,omitempty"`

	// user description
	Description string `json:"description,omitempty"`

	// user email
	Email string `json:"email,omitempty"`

	// user phone number
	PhoneNumber string `json:"phone_number,omitempty"`

	// user status eg.[active|deleted]
	Status string `json:"status,omitempty"`

	// record changed time of status
	StatusTime strfmt.DateTime `json:"status_time,omitempty"`

	// the time when user update
	UpdateTime strfmt.DateTime `json:"update_time,omitempty"`

	// user id, user belong to different group and role, has different permissions
	UserID string `json:"user_id,omitempty"`

	// user name
	Username string `json:"username,omitempty"`
}

// Validate validates this openpitrix user
func (m *OpenpitrixUser) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixUser) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixUser) UnmarshalBinary(b []byte) error {
	var res OpenpitrixUser
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}