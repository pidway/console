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

package k_m_s

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/pidway/console/models"
)

// KMSVersionOKCode is the HTTP code returned for type KMSVersionOK
const KMSVersionOKCode int = 200

/*
KMSVersionOK A successful response.

swagger:response kMSVersionOK
*/
type KMSVersionOK struct {

	/*
	  In: Body
	*/
	Payload *models.KmsVersionResponse `json:"body,omitempty"`
}

// NewKMSVersionOK creates KMSVersionOK with default headers values
func NewKMSVersionOK() *KMSVersionOK {

	return &KMSVersionOK{}
}

// WithPayload adds the payload to the k m s version o k response
func (o *KMSVersionOK) WithPayload(payload *models.KmsVersionResponse) *KMSVersionOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the k m s version o k response
func (o *KMSVersionOK) SetPayload(payload *models.KmsVersionResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *KMSVersionOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
KMSVersionDefault Generic error response.

swagger:response kMSVersionDefault
*/
type KMSVersionDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.APIError `json:"body,omitempty"`
}

// NewKMSVersionDefault creates KMSVersionDefault with default headers values
func NewKMSVersionDefault(code int) *KMSVersionDefault {
	if code <= 0 {
		code = 500
	}

	return &KMSVersionDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the k m s version default response
func (o *KMSVersionDefault) WithStatusCode(code int) *KMSVersionDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the k m s version default response
func (o *KMSVersionDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the k m s version default response
func (o *KMSVersionDefault) WithPayload(payload *models.APIError) *KMSVersionDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the k m s version default response
func (o *KMSVersionDefault) SetPayload(payload *models.APIError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *KMSVersionDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
