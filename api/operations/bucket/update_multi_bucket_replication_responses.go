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

package bucket

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/pidway/console/models"
)

// UpdateMultiBucketReplicationCreatedCode is the HTTP code returned for type UpdateMultiBucketReplicationCreated
const UpdateMultiBucketReplicationCreatedCode int = 201

/*
UpdateMultiBucketReplicationCreated A successful response.

swagger:response updateMultiBucketReplicationCreated
*/
type UpdateMultiBucketReplicationCreated struct {
}

// NewUpdateMultiBucketReplicationCreated creates UpdateMultiBucketReplicationCreated with default headers values
func NewUpdateMultiBucketReplicationCreated() *UpdateMultiBucketReplicationCreated {

	return &UpdateMultiBucketReplicationCreated{}
}

// WriteResponse to the client
func (o *UpdateMultiBucketReplicationCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(201)
}

/*
UpdateMultiBucketReplicationDefault Generic error response.

swagger:response updateMultiBucketReplicationDefault
*/
type UpdateMultiBucketReplicationDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.APIError `json:"body,omitempty"`
}

// NewUpdateMultiBucketReplicationDefault creates UpdateMultiBucketReplicationDefault with default headers values
func NewUpdateMultiBucketReplicationDefault(code int) *UpdateMultiBucketReplicationDefault {
	if code <= 0 {
		code = 500
	}

	return &UpdateMultiBucketReplicationDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the update multi bucket replication default response
func (o *UpdateMultiBucketReplicationDefault) WithStatusCode(code int) *UpdateMultiBucketReplicationDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the update multi bucket replication default response
func (o *UpdateMultiBucketReplicationDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the update multi bucket replication default response
func (o *UpdateMultiBucketReplicationDefault) WithPayload(payload *models.APIError) *UpdateMultiBucketReplicationDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update multi bucket replication default response
func (o *UpdateMultiBucketReplicationDefault) SetPayload(payload *models.APIError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateMultiBucketReplicationDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
