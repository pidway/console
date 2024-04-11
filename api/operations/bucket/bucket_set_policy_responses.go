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

	"github.com/minio/console/models"
)

// BucketSetPolicyOKCode is the HTTP code returned for type BucketSetPolicyOK
const BucketSetPolicyOKCode int = 200

/*
BucketSetPolicyOK A successful response.

swagger:response bucketSetPolicyOK
*/
type BucketSetPolicyOK struct {

	/*
	  In: Body
	*/
	Payload *models.Bucket `json:"body,omitempty"`
}

// NewBucketSetPolicyOK creates BucketSetPolicyOK with default headers values
func NewBucketSetPolicyOK() *BucketSetPolicyOK {

	return &BucketSetPolicyOK{}
}

// WithPayload adds the payload to the bucket set policy o k response
func (o *BucketSetPolicyOK) WithPayload(payload *models.Bucket) *BucketSetPolicyOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the bucket set policy o k response
func (o *BucketSetPolicyOK) SetPayload(payload *models.Bucket) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *BucketSetPolicyOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
BucketSetPolicyDefault Generic error response.

swagger:response bucketSetPolicyDefault
*/
type BucketSetPolicyDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.APIError `json:"body,omitempty"`
}

// NewBucketSetPolicyDefault creates BucketSetPolicyDefault with default headers values
func NewBucketSetPolicyDefault(code int) *BucketSetPolicyDefault {
	if code <= 0 {
		code = 500
	}

	return &BucketSetPolicyDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the bucket set policy default response
func (o *BucketSetPolicyDefault) WithStatusCode(code int) *BucketSetPolicyDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the bucket set policy default response
func (o *BucketSetPolicyDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the bucket set policy default response
func (o *BucketSetPolicyDefault) WithPayload(payload *models.APIError) *BucketSetPolicyDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the bucket set policy default response
func (o *BucketSetPolicyDefault) SetPayload(payload *models.APIError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *BucketSetPolicyDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
