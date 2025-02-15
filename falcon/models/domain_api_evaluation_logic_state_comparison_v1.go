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
	"github.com/go-openapi/validate"
)

// DomainAPIEvaluationLogicStateComparisonV1 domain API evaluation logic state comparison v1
//
// swagger:model domain.APIEvaluationLogicStateComparisonV1
type DomainAPIEvaluationLogicStateComparisonV1 struct {

	// entity comparisons
	// Required: true
	EntityComparisons []*DomainAPIEvaluationLogicEntityComparisonV1 `json:"entity_comparisons"`

	// entity operator
	// Required: true
	EntityOperator *string `json:"entity_operator"`
}

// Validate validates this domain API evaluation logic state comparison v1
func (m *DomainAPIEvaluationLogicStateComparisonV1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEntityComparisons(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntityOperator(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainAPIEvaluationLogicStateComparisonV1) validateEntityComparisons(formats strfmt.Registry) error {

	if err := validate.Required("entity_comparisons", "body", m.EntityComparisons); err != nil {
		return err
	}

	for i := 0; i < len(m.EntityComparisons); i++ {
		if swag.IsZero(m.EntityComparisons[i]) { // not required
			continue
		}

		if m.EntityComparisons[i] != nil {
			if err := m.EntityComparisons[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("entity_comparisons" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("entity_comparisons" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DomainAPIEvaluationLogicStateComparisonV1) validateEntityOperator(formats strfmt.Registry) error {

	if err := validate.Required("entity_operator", "body", m.EntityOperator); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this domain API evaluation logic state comparison v1 based on the context it is used
func (m *DomainAPIEvaluationLogicStateComparisonV1) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEntityComparisons(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainAPIEvaluationLogicStateComparisonV1) contextValidateEntityComparisons(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.EntityComparisons); i++ {

		if m.EntityComparisons[i] != nil {
			if err := m.EntityComparisons[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("entity_comparisons" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("entity_comparisons" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DomainAPIEvaluationLogicStateComparisonV1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainAPIEvaluationLogicStateComparisonV1) UnmarshalBinary(b []byte) error {
	var res DomainAPIEvaluationLogicStateComparisonV1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
