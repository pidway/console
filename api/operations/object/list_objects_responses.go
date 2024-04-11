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

package object

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/minio/console/models"
)

// ListObjectsOKCode is the HTTP code returned for type ListObjectsOK
const ListObjectsOKCode int = 200

/*
ListObjectsOK A successful response.

swagger:response listObjectsOK
*/
type ListObjectsOK struct {

	/*
	  In: Body
	*/
	Payload *models.ListObjectsResponse `json:"body,omitempty"`
}

// NewListObjectsOK creates ListObjectsOK with default headers values
func NewListObjectsOK() *ListObjectsOK {

	return &ListObjectsOK{}
}

// WithPayload adds the payload to the list objects o k response
func (o *ListObjectsOK) WithPayload(payload *models.ListObjectsResponse) *ListObjectsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list objects o k response
func (o *ListObjectsOK) SetPayload(payload *models.ListObjectsResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListObjectsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
ListObjectsDefault Generic error response.

swagger:response listObjectsDefault
*/
type ListObjectsDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.APIError `json:"body,omitempty"`
}

// NewListObjectsDefault creates ListObjectsDefault with default headers values
func NewListObjectsDefault(code int) *ListObjectsDefault {
	if code <= 0 {
		code = 500
	}

	return &ListObjectsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the list objects default response
func (o *ListObjectsDefault) WithStatusCode(code int) *ListObjectsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the list objects default response
func (o *ListObjectsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the list objects default response
func (o *ListObjectsDefault) WithPayload(payload *models.APIError) *ListObjectsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list objects default response
func (o *ListObjectsDefault) SetPayload(payload *models.APIError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListObjectsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
