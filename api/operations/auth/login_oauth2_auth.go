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

package auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// LoginOauth2AuthHandlerFunc turns a function with the right signature into a login oauth2 auth handler
type LoginOauth2AuthHandlerFunc func(LoginOauth2AuthParams) middleware.Responder

// Handle executing the request and returning a response
func (fn LoginOauth2AuthHandlerFunc) Handle(params LoginOauth2AuthParams) middleware.Responder {
	return fn(params)
}

// LoginOauth2AuthHandler interface for that can handle valid login oauth2 auth params
type LoginOauth2AuthHandler interface {
	Handle(LoginOauth2AuthParams) middleware.Responder
}

// NewLoginOauth2Auth creates a new http.Handler for the login oauth2 auth operation
func NewLoginOauth2Auth(ctx *middleware.Context, handler LoginOauth2AuthHandler) *LoginOauth2Auth {
	return &LoginOauth2Auth{Context: ctx, Handler: handler}
}

/*
	LoginOauth2Auth swagger:route POST /login/oauth2/auth Auth loginOauth2Auth

Identity Provider oauth2 callback endpoint.
*/
type LoginOauth2Auth struct {
	Context *middleware.Context
	Handler LoginOauth2AuthHandler
}

func (o *LoginOauth2Auth) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewLoginOauth2AuthParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
