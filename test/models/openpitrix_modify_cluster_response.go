// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixModifyClusterResponse openpitrix modify cluster response
// swagger:model openpitrixModifyClusterResponse
type OpenpitrixModifyClusterResponse struct {

	// id of cluster modified
	ClusterID string `json:"cluster_id,omitempty"`
}

// Validate validates this openpitrix modify cluster response
func (m *OpenpitrixModifyClusterResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixModifyClusterResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixModifyClusterResponse) UnmarshalBinary(b []byte) error {
	var res OpenpitrixModifyClusterResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}