// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RequestsSensorUpdateBuildV1 requests sensor update build v1
//
// swagger:model requests.SensorUpdateBuildV1
type RequestsSensorUpdateBuildV1 struct {

	// build
	// Required: true
	Build *string `json:"build"`

	// platform
	// Required: true
	Platform *string `json:"platform"`
}

// Validate validates this requests sensor update build v1
func (m *RequestsSensorUpdateBuildV1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBuild(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlatform(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RequestsSensorUpdateBuildV1) validateBuild(formats strfmt.Registry) error {

	if err := validate.Required("build", "body", m.Build); err != nil {
		return err
	}

	return nil
}

func (m *RequestsSensorUpdateBuildV1) validatePlatform(formats strfmt.Registry) error {

	if err := validate.Required("platform", "body", m.Platform); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this requests sensor update build v1 based on context it is used
func (m *RequestsSensorUpdateBuildV1) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RequestsSensorUpdateBuildV1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RequestsSensorUpdateBuildV1) UnmarshalBinary(b []byte) error {
	var res RequestsSensorUpdateBuildV1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
