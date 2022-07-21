// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixValidateUserPasswordRequest openpitrix validate user password request
// swagger:model openpitrixValidateUserPasswordRequest
type OpenpitrixValidateUserPasswordRequest struct {

	// required, user email
	Email string `json:"email,omitempty"`

	// required, user password
	Password string `json:"password,omitempty"`
}

// Validate validates this openpitrix validate user password request
func (m *OpenpitrixValidateUserPasswordRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixValidateUserPasswordRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixValidateUserPasswordRequest) UnmarshalBinary(b []byte) error {
	var res OpenpitrixValidateUserPasswordRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
