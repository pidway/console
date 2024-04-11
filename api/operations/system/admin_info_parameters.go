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

package system

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewAdminInfoParams creates a new AdminInfoParams object
// with the default values initialized.
func NewAdminInfoParams() AdminInfoParams {

	var (
		// initialize parameters with default values

		defaultOnlyDefault = bool(false)
	)

	return AdminInfoParams{
		DefaultOnly: &defaultOnlyDefault,
	}
}

// AdminInfoParams contains all the bound params for the admin info operation
// typically these are obtained from a http.Request
//
// swagger:parameters AdminInfo
type AdminInfoParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  In: query
	  Default: false
	*/
	DefaultOnly *bool
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewAdminInfoParams() beforehand.
func (o *AdminInfoParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qDefaultOnly, qhkDefaultOnly, _ := qs.GetOK("defaultOnly")
	if err := o.bindDefaultOnly(qDefaultOnly, qhkDefaultOnly, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindDefaultOnly binds and validates parameter DefaultOnly from query.
func (o *AdminInfoParams) bindDefaultOnly(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewAdminInfoParams()
		return nil
	}

	value, err := swag.ConvertBool(raw)
	if err != nil {
		return errors.InvalidType("defaultOnly", "query", "bool", raw)
	}
	o.DefaultOnly = &value

	return nil
}
