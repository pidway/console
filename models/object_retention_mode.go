// Code generated by go-swagger; DO NOT EDIT.

// This file is part of MinIO Console Server
// Copyright (c) 2023 MinIO, Inc.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.
//

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// ObjectRetentionMode object retention mode
//
// swagger:model objectRetentionMode
type ObjectRetentionMode string

func NewObjectRetentionMode(value ObjectRetentionMode) *ObjectRetentionMode {
	return &value
}

// Pointer returns a pointer to a freshly-allocated ObjectRetentionMode.
func (m ObjectRetentionMode) Pointer() *ObjectRetentionMode {
	return &m
}

const (

	// ObjectRetentionModeGovernance captures enum value "governance"
	ObjectRetentionModeGovernance ObjectRetentionMode = "governance"

	// ObjectRetentionModeCompliance captures enum value "compliance"
	ObjectRetentionModeCompliance ObjectRetentionMode = "compliance"
)

// for schema
var objectRetentionModeEnum []interface{}

func init() {
	var res []ObjectRetentionMode
	if err := json.Unmarshal([]byte(`["governance","compliance"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		objectRetentionModeEnum = append(objectRetentionModeEnum, v)
	}
}

func (m ObjectRetentionMode) validateObjectRetentionModeEnum(path, location string, value ObjectRetentionMode) error {
	if err := validate.EnumCase(path, location, value, objectRetentionModeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this object retention mode
func (m ObjectRetentionMode) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateObjectRetentionModeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this object retention mode based on context it is used
func (m ObjectRetentionMode) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
