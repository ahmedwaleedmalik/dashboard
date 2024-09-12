// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Kubevirt kubevirt
//
// swagger:model Kubevirt
type Kubevirt struct {

	// If datacenter is set, this preset is only applicable to the
	// configured datacenter.
	Datacenter string `json:"datacenter,omitempty"`

	// Only enabled presets will be available in the KKP dashboard.
	Enabled bool `json:"enabled,omitempty"`

	// IsCustomizable marks a preset as editable on the KKP UI; Customizable presets still have the credentials obscured on the UI, but other fields that are not considered private are displayed during cluster creation. Users can then update those fields, if required.
	// NOTE: This is only supported for OpenStack Cloud Provider in KKP 2.26. Support for other providers will be added later on.
	IsCustomizable bool `json:"isCutomizable,omitempty"`

	// Kubeconfig is the cluster's kubeconfig file, encoded with base64.
	Kubeconfig string `json:"kubeconfig,omitempty"`
}

// Validate validates this kubevirt
func (m *Kubevirt) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this kubevirt based on context it is used
func (m *Kubevirt) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Kubevirt) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Kubevirt) UnmarshalBinary(b []byte) error {
	var res Kubevirt
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
