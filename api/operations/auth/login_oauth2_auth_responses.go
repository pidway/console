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
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/pidway/console/models"
)

// LoginOauth2AuthNoContentCode is the HTTP code returned for type LoginOauth2AuthNoContent
const LoginOauth2AuthNoContentCode int = 204

/*
LoginOauth2AuthNoContent A successful login.

swagger:response loginOauth2AuthNoContent
*/
type LoginOauth2AuthNoContent struct {
}

// NewLoginOauth2AuthNoContent creates LoginOauth2AuthNoContent with default headers values
func NewLoginOauth2AuthNoContent() *LoginOauth2AuthNoContent {

	return &LoginOauth2AuthNoContent{}
}

// WriteResponse to the client
func (o *LoginOauth2AuthNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

/*
LoginOauth2AuthDefault Generic error response.

swagger:response loginOauth2AuthDefault
*/
type LoginOauth2AuthDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.APIError `json:"body,omitempty"`
}

// NewLoginOauth2AuthDefault creates LoginOauth2AuthDefault with default headers values
func NewLoginOauth2AuthDefault(code int) *LoginOauth2AuthDefault {
	if code <= 0 {
		code = 500
	}

	return &LoginOauth2AuthDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the login oauth2 auth default response
func (o *LoginOauth2AuthDefault) WithStatusCode(code int) *LoginOauth2AuthDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the login oauth2 auth default response
func (o *LoginOauth2AuthDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the login oauth2 auth default response
func (o *LoginOauth2AuthDefault) WithPayload(payload *models.APIError) *LoginOauth2AuthDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the login oauth2 auth default response
func (o *LoginOauth2AuthDefault) SetPayload(payload *models.APIError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *LoginOauth2AuthDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
