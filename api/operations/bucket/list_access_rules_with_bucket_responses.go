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

// ListAccessRulesWithBucketOKCode is the HTTP code returned for type ListAccessRulesWithBucketOK
const ListAccessRulesWithBucketOKCode int = 200

/*
ListAccessRulesWithBucketOK A successful response.

swagger:response listAccessRulesWithBucketOK
*/
type ListAccessRulesWithBucketOK struct {

	/*
	  In: Body
	*/
	Payload *models.ListAccessRulesResponse `json:"body,omitempty"`
}

// NewListAccessRulesWithBucketOK creates ListAccessRulesWithBucketOK with default headers values
func NewListAccessRulesWithBucketOK() *ListAccessRulesWithBucketOK {

	return &ListAccessRulesWithBucketOK{}
}

// WithPayload adds the payload to the list access rules with bucket o k response
func (o *ListAccessRulesWithBucketOK) WithPayload(payload *models.ListAccessRulesResponse) *ListAccessRulesWithBucketOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list access rules with bucket o k response
func (o *ListAccessRulesWithBucketOK) SetPayload(payload *models.ListAccessRulesResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListAccessRulesWithBucketOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
ListAccessRulesWithBucketDefault Generic error response.

swagger:response listAccessRulesWithBucketDefault
*/
type ListAccessRulesWithBucketDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.APIError `json:"body,omitempty"`
}

// NewListAccessRulesWithBucketDefault creates ListAccessRulesWithBucketDefault with default headers values
func NewListAccessRulesWithBucketDefault(code int) *ListAccessRulesWithBucketDefault {
	if code <= 0 {
		code = 500
	}

	return &ListAccessRulesWithBucketDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the list access rules with bucket default response
func (o *ListAccessRulesWithBucketDefault) WithStatusCode(code int) *ListAccessRulesWithBucketDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the list access rules with bucket default response
func (o *ListAccessRulesWithBucketDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the list access rules with bucket default response
func (o *ListAccessRulesWithBucketDefault) WithPayload(payload *models.APIError) *ListAccessRulesWithBucketDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list access rules with bucket default response
func (o *ListAccessRulesWithBucketDefault) SetPayload(payload *models.APIError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListAccessRulesWithBucketDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
