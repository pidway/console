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

// KMSListKeysOKCode is the HTTP code returned for type KMSListKeysOK
const KMSListKeysOKCode int = 200

/*
KMSListKeysOK A successful response.

swagger:response kMSListKeysOK
*/
type KMSListKeysOK struct {

	/*
	  In: Body
	*/
	Payload *models.KmsListKeysResponse `json:"body,omitempty"`
}

// NewKMSListKeysOK creates KMSListKeysOK with default headers values
func NewKMSListKeysOK() *KMSListKeysOK {

	return &KMSListKeysOK{}
}

// WithPayload adds the payload to the k m s list keys o k response
func (o *KMSListKeysOK) WithPayload(payload *models.KmsListKeysResponse) *KMSListKeysOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the k m s list keys o k response
func (o *KMSListKeysOK) SetPayload(payload *models.KmsListKeysResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *KMSListKeysOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
KMSListKeysDefault Generic error response.

swagger:response kMSListKeysDefault
*/
type KMSListKeysDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.APIError `json:"body,omitempty"`
}

// NewKMSListKeysDefault creates KMSListKeysDefault with default headers values
func NewKMSListKeysDefault(code int) *KMSListKeysDefault {
	if code <= 0 {
		code = 500
	}

	return &KMSListKeysDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the k m s list keys default response
func (o *KMSListKeysDefault) WithStatusCode(code int) *KMSListKeysDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the k m s list keys default response
func (o *KMSListKeysDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the k m s list keys default response
func (o *KMSListKeysDefault) WithPayload(payload *models.APIError) *KMSListKeysDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the k m s list keys default response
func (o *KMSListKeysDefault) SetPayload(payload *models.APIError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *KMSListKeysDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
