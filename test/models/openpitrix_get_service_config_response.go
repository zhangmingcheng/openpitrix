// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixGetServiceConfigResponse openpitrix get service config response
// swagger:model openpitrixGetServiceConfigResponse
type OpenpitrixGetServiceConfigResponse struct {

	// basic config
	BasicConfig *OpenpitrixBasicConfig `json:"basic_config,omitempty"`

	// notification config
	NotificationConfig *OpenpitrixNotificationConfig `json:"notification_config,omitempty"`

	// runtime config
	RuntimeConfig *OpenpitrixRuntimeConfig `json:"runtime_config,omitempty"`
}

// Validate validates this openpitrix get service config response
func (m *OpenpitrixGetServiceConfigResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBasicConfig(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateNotificationConfig(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRuntimeConfig(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenpitrixGetServiceConfigResponse) validateBasicConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.BasicConfig) { // not required
		return nil
	}

	if m.BasicConfig != nil {

		if err := m.BasicConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("basic_config")
			}
			return err
		}
	}

	return nil
}

func (m *OpenpitrixGetServiceConfigResponse) validateNotificationConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.NotificationConfig) { // not required
		return nil
	}

	if m.NotificationConfig != nil {

		if err := m.NotificationConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("notification_config")
			}
			return err
		}
	}

	return nil
}

func (m *OpenpitrixGetServiceConfigResponse) validateRuntimeConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.RuntimeConfig) { // not required
		return nil
	}

	if m.RuntimeConfig != nil {

		if err := m.RuntimeConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("runtime_config")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixGetServiceConfigResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixGetServiceConfigResponse) UnmarshalBinary(b []byte) error {
	var res OpenpitrixGetServiceConfigResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
