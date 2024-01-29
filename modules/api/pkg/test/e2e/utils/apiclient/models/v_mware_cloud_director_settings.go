// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// VMwareCloudDirectorSettings v mware cloud director settings
//
// swagger:model VMwareCloudDirectorSettings
type VMwareCloudDirectorSettings struct {

	// IPAllocationModes are the allowed IP allocation modes for the VMware Cloud Director provider. If not set, all modes are allowed.
	IPAllocationModes []IPAllocationMode `json:"ipAllocationModes"`
}

// Validate validates this v mware cloud director settings
func (m *VMwareCloudDirectorSettings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIPAllocationModes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VMwareCloudDirectorSettings) validateIPAllocationModes(formats strfmt.Registry) error {
	if swag.IsZero(m.IPAllocationModes) { // not required
		return nil
	}

	for i := 0; i < len(m.IPAllocationModes); i++ {

		if err := m.IPAllocationModes[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ipAllocationModes" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ipAllocationModes" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// ContextValidate validate this v mware cloud director settings based on the context it is used
func (m *VMwareCloudDirectorSettings) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateIPAllocationModes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VMwareCloudDirectorSettings) contextValidateIPAllocationModes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.IPAllocationModes); i++ {

		if err := m.IPAllocationModes[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ipAllocationModes" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ipAllocationModes" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *VMwareCloudDirectorSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VMwareCloudDirectorSettings) UnmarshalBinary(b []byte) error {
	var res VMwareCloudDirectorSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
