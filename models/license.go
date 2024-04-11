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

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// License license
//
// swagger:model license
type License struct {

	// account id
	AccountID int64 `json:"account_id,omitempty"`

	// email
	Email string `json:"email,omitempty"`

	// expires at
	ExpiresAt string `json:"expires_at,omitempty"`

	// organization
	Organization string `json:"organization,omitempty"`

	// plan
	Plan string `json:"plan,omitempty"`

	// storage capacity
	StorageCapacity int64 `json:"storage_capacity,omitempty"`
}

// Validate validates this license
func (m *License) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this license based on context it is used
func (m *License) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *License) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *License) UnmarshalBinary(b []byte) error {
	var res License
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
