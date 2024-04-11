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

package group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/pidway/console/models"
)

// UpdateGroupOKCode is the HTTP code returned for type UpdateGroupOK
const UpdateGroupOKCode int = 200

/*
UpdateGroupOK A successful response.

swagger:response updateGroupOK
*/
type UpdateGroupOK struct {

	/*
	  In: Body
	*/
	Payload *models.Group `json:"body,omitempty"`
}

// NewUpdateGroupOK creates UpdateGroupOK with default headers values
func NewUpdateGroupOK() *UpdateGroupOK {

	return &UpdateGroupOK{}
}

// WithPayload adds the payload to the update group o k response
func (o *UpdateGroupOK) WithPayload(payload *models.Group) *UpdateGroupOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update group o k response
func (o *UpdateGroupOK) SetPayload(payload *models.Group) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateGroupOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
UpdateGroupDefault Generic error response.

swagger:response updateGroupDefault
*/
type UpdateGroupDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.APIError `json:"body,omitempty"`
}

// NewUpdateGroupDefault creates UpdateGroupDefault with default headers values
func NewUpdateGroupDefault(code int) *UpdateGroupDefault {
	if code <= 0 {
		code = 500
	}

	return &UpdateGroupDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the update group default response
func (o *UpdateGroupDefault) WithStatusCode(code int) *UpdateGroupDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the update group default response
func (o *UpdateGroupDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the update group default response
func (o *UpdateGroupDefault) WithPayload(payload *models.APIError) *UpdateGroupDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update group default response
func (o *UpdateGroupDefault) SetPayload(payload *models.APIError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateGroupDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
