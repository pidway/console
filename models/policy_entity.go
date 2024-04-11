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

// PolicyEntity policy entity
//
// swagger:model policyEntity
type PolicyEntity string

func NewPolicyEntity(value PolicyEntity) *PolicyEntity {
	return &value
}

// Pointer returns a pointer to a freshly-allocated PolicyEntity.
func (m PolicyEntity) Pointer() *PolicyEntity {
	return &m
}

const (

	// PolicyEntityUser captures enum value "user"
	PolicyEntityUser PolicyEntity = "user"

	// PolicyEntityGroup captures enum value "group"
	PolicyEntityGroup PolicyEntity = "group"
)

// for schema
var policyEntityEnum []interface{}

func init() {
	var res []PolicyEntity
	if err := json.Unmarshal([]byte(`["user","group"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		policyEntityEnum = append(policyEntityEnum, v)
	}
}

func (m PolicyEntity) validatePolicyEntityEnum(path, location string, value PolicyEntity) error {
	if err := validate.EnumCase(path, location, value, policyEntityEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this policy entity
func (m PolicyEntity) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validatePolicyEntityEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this policy entity based on context it is used
func (m PolicyEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
