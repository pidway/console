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

package policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/minio/console/models"
)

// GetUserPolicyOKCode is the HTTP code returned for type GetUserPolicyOK
const GetUserPolicyOKCode int = 200

/*
GetUserPolicyOK A successful response.

swagger:response getUserPolicyOK
*/
type GetUserPolicyOK struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewGetUserPolicyOK creates GetUserPolicyOK with default headers values
func NewGetUserPolicyOK() *GetUserPolicyOK {

	return &GetUserPolicyOK{}
}

// WithPayload adds the payload to the get user policy o k response
func (o *GetUserPolicyOK) WithPayload(payload string) *GetUserPolicyOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get user policy o k response
func (o *GetUserPolicyOK) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUserPolicyOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

/*
GetUserPolicyDefault Generic error response.

swagger:response getUserPolicyDefault
*/
type GetUserPolicyDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.APIError `json:"body,omitempty"`
}

// NewGetUserPolicyDefault creates GetUserPolicyDefault with default headers values
func NewGetUserPolicyDefault(code int) *GetUserPolicyDefault {
	if code <= 0 {
		code = 500
	}

	return &GetUserPolicyDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get user policy default response
func (o *GetUserPolicyDefault) WithStatusCode(code int) *GetUserPolicyDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get user policy default response
func (o *GetUserPolicyDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get user policy default response
func (o *GetUserPolicyDefault) WithPayload(payload *models.APIError) *GetUserPolicyDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get user policy default response
func (o *GetUserPolicyDefault) SetPayload(payload *models.APIError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUserPolicyDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
