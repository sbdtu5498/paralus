// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1PodDNSConfigOption PodDNSConfigOption defines DNS resolver options of a pod.
//
// swagger:model v1PodDNSConfigOption
type V1PodDNSConfigOption struct {

	// Required.
	Name string `json:"name,omitempty"`

	// +optional
	Value string `json:"value,omitempty"`
}

// Validate validates this v1 pod DNS config option
func (m *V1PodDNSConfigOption) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 pod DNS config option based on context it is used
func (m *V1PodDNSConfigOption) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1PodDNSConfigOption) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1PodDNSConfigOption) UnmarshalBinary(b []byte) error {
	var res V1PodDNSConfigOption
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}