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

// GetBucketRetentionConfigOKCode is the HTTP code returned for type GetBucketRetentionConfigOK
const GetBucketRetentionConfigOKCode int = 200

/*
GetBucketRetentionConfigOK A successful response.

swagger:response getBucketRetentionConfigOK
*/
type GetBucketRetentionConfigOK struct {

	/*
	  In: Body
	*/
	Payload *models.GetBucketRetentionConfig `json:"body,omitempty"`
}

// NewGetBucketRetentionConfigOK creates GetBucketRetentionConfigOK with default headers values
func NewGetBucketRetentionConfigOK() *GetBucketRetentionConfigOK {

	return &GetBucketRetentionConfigOK{}
}

// WithPayload adds the payload to the get bucket retention config o k response
func (o *GetBucketRetentionConfigOK) WithPayload(payload *models.GetBucketRetentionConfig) *GetBucketRetentionConfigOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get bucket retention config o k response
func (o *GetBucketRetentionConfigOK) SetPayload(payload *models.GetBucketRetentionConfig) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetBucketRetentionConfigOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
GetBucketRetentionConfigDefault Generic error response.

swagger:response getBucketRetentionConfigDefault
*/
type GetBucketRetentionConfigDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.APIError `json:"body,omitempty"`
}

// NewGetBucketRetentionConfigDefault creates GetBucketRetentionConfigDefault with default headers values
func NewGetBucketRetentionConfigDefault(code int) *GetBucketRetentionConfigDefault {
	if code <= 0 {
		code = 500
	}

	return &GetBucketRetentionConfigDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get bucket retention config default response
func (o *GetBucketRetentionConfigDefault) WithStatusCode(code int) *GetBucketRetentionConfigDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get bucket retention config default response
func (o *GetBucketRetentionConfigDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get bucket retention config default response
func (o *GetBucketRetentionConfigDefault) WithPayload(payload *models.APIError) *GetBucketRetentionConfigDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get bucket retention config default response
func (o *GetBucketRetentionConfigDefault) SetPayload(payload *models.APIError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetBucketRetentionConfigDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
