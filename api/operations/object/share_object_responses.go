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

// ShareObjectOKCode is the HTTP code returned for type ShareObjectOK
const ShareObjectOKCode int = 200

/*
ShareObjectOK A successful response.

swagger:response shareObjectOK
*/
type ShareObjectOK struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewShareObjectOK creates ShareObjectOK with default headers values
func NewShareObjectOK() *ShareObjectOK {

	return &ShareObjectOK{}
}

// WithPayload adds the payload to the share object o k response
func (o *ShareObjectOK) WithPayload(payload string) *ShareObjectOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the share object o k response
func (o *ShareObjectOK) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ShareObjectOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

/*
ShareObjectDefault Generic error response.

swagger:response shareObjectDefault
*/
type ShareObjectDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.APIError `json:"body,omitempty"`
}

// NewShareObjectDefault creates ShareObjectDefault with default headers values
func NewShareObjectDefault(code int) *ShareObjectDefault {
	if code <= 0 {
		code = 500
	}

	return &ShareObjectDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the share object default response
func (o *ShareObjectDefault) WithStatusCode(code int) *ShareObjectDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the share object default response
func (o *ShareObjectDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the share object default response
func (o *ShareObjectDefault) WithPayload(payload *models.APIError) *ShareObjectDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the share object default response
func (o *ShareObjectDefault) SetPayload(payload *models.APIError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ShareObjectDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
