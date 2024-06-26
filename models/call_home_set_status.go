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

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CallHomeSetStatus call home set status
//
// swagger:model callHomeSetStatus
type CallHomeSetStatus struct {

	// diag state
	// Required: true
	DiagState *bool `json:"diagState"`

	// logs state
	// Required: true
	LogsState *bool `json:"logsState"`
}

// Validate validates this call home set status
func (m *CallHomeSetStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDiagState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLogsState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CallHomeSetStatus) validateDiagState(formats strfmt.Registry) error {

	if err := validate.Required("diagState", "body", m.DiagState); err != nil {
		return err
	}

	return nil
}

func (m *CallHomeSetStatus) validateLogsState(formats strfmt.Registry) error {

	if err := validate.Required("logsState", "body", m.LogsState); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this call home set status based on context it is used
func (m *CallHomeSetStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CallHomeSetStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CallHomeSetStatus) UnmarshalBinary(b []byte) error {
	var res CallHomeSetStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
