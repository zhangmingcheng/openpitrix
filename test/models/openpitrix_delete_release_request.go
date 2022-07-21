// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixDeleteReleaseRequest openpitrix delete release request
// swagger:model openpitrixDeleteReleaseRequest
type OpenpitrixDeleteReleaseRequest struct {

	// purge
	Purge bool `json:"purge,omitempty"`

	// release name
	ReleaseName string `json:"release_name,omitempty"`

	// runtime id
	RuntimeID string `json:"runtime_id,omitempty"`
}

// Validate validates this openpitrix delete release request
func (m *OpenpitrixDeleteReleaseRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixDeleteReleaseRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixDeleteReleaseRequest) UnmarshalBinary(b []byte) error {
	var res OpenpitrixDeleteReleaseRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}