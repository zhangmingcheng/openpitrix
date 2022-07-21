// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixCategory openpitrix category
// swagger:model openpitrixCategory
type OpenpitrixCategory struct {

	// category id
	CategoryID string `json:"category_id,omitempty"`

	// the time when category create
	CreateTime strfmt.DateTime `json:"create_time,omitempty"`

	// category description
	Description string `json:"description,omitempty"`

	// category icon
	Icon string `json:"icon,omitempty"`

	// the i18n of this category, json format, sample: {"zh_cn": "数据库", "en": "database"}
	Locale string `json:"locale,omitempty"`

	// category name,app belong to a category,eg.[AI|Firewall|cache|...]
	Name string `json:"name,omitempty"`

	// owner
	Owner string `json:"owner,omitempty"`

	// owner path, concat string group_path:user_id
	OwnerPath string `json:"owner_path,omitempty"`

	// the time when category update
	UpdateTime strfmt.DateTime `json:"update_time,omitempty"`
}

// Validate validates this openpitrix category
func (m *OpenpitrixCategory) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixCategory) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixCategory) UnmarshalBinary(b []byte) error {
	var res OpenpitrixCategory
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}