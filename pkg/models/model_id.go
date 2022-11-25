// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ModelID model ID
//
// swagger:model modelID
type ModelID struct {

	// id
	ID int64 `json:"id,omitempty" gorm:"type:int primary key auto_increment"`
}

// Validate validates this model ID
func (m *ModelID) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this model ID based on context it is used
func (m *ModelID) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ModelID) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelID) UnmarshalBinary(b []byte) error {
	var res ModelID
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
