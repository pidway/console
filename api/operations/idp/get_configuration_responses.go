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

	"github.com/go-openapi/runtime"

	"github.com/pidway/console/models"
)

// GetConfigurationOKCode is the HTTP code returned for type GetConfigurationOK
const GetConfigurationOKCode int = 200

/*
GetConfigurationOK A successful response.

swagger:response getConfigurationOK
*/
type GetConfigurationOK struct {

	/*
	  In: Body
	*/
	Payload *models.IdpServerConfiguration `json:"body,omitempty"`
}

// NewGetConfigurationOK creates GetConfigurationOK with default headers values
func NewGetConfigurationOK() *GetConfigurationOK {

	return &GetConfigurationOK{}
}

// WithPayload adds the payload to the get configuration o k response
func (o *GetConfigurationOK) WithPayload(payload *models.IdpServerConfiguration) *GetConfigurationOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get configuration o k response
func (o *GetConfigurationOK) SetPayload(payload *models.IdpServerConfiguration) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetConfigurationOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
GetConfigurationDefault Generic error response.

swagger:response getConfigurationDefault
*/
type GetConfigurationDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.APIError `json:"body,omitempty"`
}

// NewGetConfigurationDefault creates GetConfigurationDefault with default headers values
func NewGetConfigurationDefault(code int) *GetConfigurationDefault {
	if code <= 0 {
		code = 500
	}

	return &GetConfigurationDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get configuration default response
func (o *GetConfigurationDefault) WithStatusCode(code int) *GetConfigurationDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get configuration default response
func (o *GetConfigurationDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get configuration default response
func (o *GetConfigurationDefault) WithPayload(payload *models.APIError) *GetConfigurationDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get configuration default response
func (o *GetConfigurationDefault) SetPayload(payload *models.APIError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetConfigurationDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
