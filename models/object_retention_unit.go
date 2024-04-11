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

// ObjectRetentionUnit object retention unit
//
// swagger:model objectRetentionUnit
type ObjectRetentionUnit string

func NewObjectRetentionUnit(value ObjectRetentionUnit) *ObjectRetentionUnit {
	return &value
}

// Pointer returns a pointer to a freshly-allocated ObjectRetentionUnit.
func (m ObjectRetentionUnit) Pointer() *ObjectRetentionUnit {
	return &m
}

const (

	// ObjectRetentionUnitDays captures enum value "days"
	ObjectRetentionUnitDays ObjectRetentionUnit = "days"

	// ObjectRetentionUnitYears captures enum value "years"
	ObjectRetentionUnitYears ObjectRetentionUnit = "years"
)

// for schema
var objectRetentionUnitEnum []interface{}

func init() {
	var res []ObjectRetentionUnit
	if err := json.Unmarshal([]byte(`["days","years"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		objectRetentionUnitEnum = append(objectRetentionUnitEnum, v)
	}
}

func (m ObjectRetentionUnit) validateObjectRetentionUnitEnum(path, location string, value ObjectRetentionUnit) error {
	if err := validate.EnumCase(path, location, value, objectRetentionUnitEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this object retention unit
func (m ObjectRetentionUnit) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateObjectRetentionUnitEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this object retention unit based on context it is used
func (m ObjectRetentionUnit) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
