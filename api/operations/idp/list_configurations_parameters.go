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

package idp

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

// NewListConfigurationsParams creates a new ListConfigurationsParams object
//
// There are no default values defined in the spec.
func NewListConfigurationsParams() ListConfigurationsParams {

	return ListConfigurationsParams{}
}

// ListConfigurationsParams contains all the bound params for the list configurations operation
// typically these are obtained from a http.Request
//
// swagger:parameters ListConfigurations
type ListConfigurationsParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*IDP Configuration Type
	  Required: true
	  In: path
	*/
	Type string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewListConfigurationsParams() beforehand.
func (o *ListConfigurationsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rType, rhkType, _ := route.Params.GetOK("type")
	if err := o.bindType(rType, rhkType, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindType binds and validates parameter Type from path.
func (o *ListConfigurationsParams) bindType(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.Type = raw

	return nil
}
