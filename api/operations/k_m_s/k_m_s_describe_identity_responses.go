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

	"github.com/minio/console/models"
)

// KMSDescribeIdentityOKCode is the HTTP code returned for type KMSDescribeIdentityOK
const KMSDescribeIdentityOKCode int = 200

/*
KMSDescribeIdentityOK A successful response.

swagger:response kMSDescribeIdentityOK
*/
type KMSDescribeIdentityOK struct {

	/*
	  In: Body
	*/
	Payload *models.KmsDescribeIdentityResponse `json:"body,omitempty"`
}

// NewKMSDescribeIdentityOK creates KMSDescribeIdentityOK with default headers values
func NewKMSDescribeIdentityOK() *KMSDescribeIdentityOK {

	return &KMSDescribeIdentityOK{}
}

// WithPayload adds the payload to the k m s describe identity o k response
func (o *KMSDescribeIdentityOK) WithPayload(payload *models.KmsDescribeIdentityResponse) *KMSDescribeIdentityOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the k m s describe identity o k response
func (o *KMSDescribeIdentityOK) SetPayload(payload *models.KmsDescribeIdentityResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *KMSDescribeIdentityOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
KMSDescribeIdentityDefault Generic error response.

swagger:response kMSDescribeIdentityDefault
*/
type KMSDescribeIdentityDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.APIError `json:"body,omitempty"`
}

// NewKMSDescribeIdentityDefault creates KMSDescribeIdentityDefault with default headers values
func NewKMSDescribeIdentityDefault(code int) *KMSDescribeIdentityDefault {
	if code <= 0 {
		code = 500
	}

	return &KMSDescribeIdentityDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the k m s describe identity default response
func (o *KMSDescribeIdentityDefault) WithStatusCode(code int) *KMSDescribeIdentityDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the k m s describe identity default response
func (o *KMSDescribeIdentityDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the k m s describe identity default response
func (o *KMSDescribeIdentityDefault) WithPayload(payload *models.APIError) *KMSDescribeIdentityDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the k m s describe identity default response
func (o *KMSDescribeIdentityDefault) SetPayload(payload *models.APIError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *KMSDescribeIdentityDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
