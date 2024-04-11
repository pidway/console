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

// DeleteAccessRuleWithBucketOKCode is the HTTP code returned for type DeleteAccessRuleWithBucketOK
const DeleteAccessRuleWithBucketOKCode int = 200

/*
DeleteAccessRuleWithBucketOK A successful response.

swagger:response deleteAccessRuleWithBucketOK
*/
type DeleteAccessRuleWithBucketOK struct {

	/*
	  In: Body
	*/
	Payload bool `json:"body,omitempty"`
}

// NewDeleteAccessRuleWithBucketOK creates DeleteAccessRuleWithBucketOK with default headers values
func NewDeleteAccessRuleWithBucketOK() *DeleteAccessRuleWithBucketOK {

	return &DeleteAccessRuleWithBucketOK{}
}

// WithPayload adds the payload to the delete access rule with bucket o k response
func (o *DeleteAccessRuleWithBucketOK) WithPayload(payload bool) *DeleteAccessRuleWithBucketOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete access rule with bucket o k response
func (o *DeleteAccessRuleWithBucketOK) SetPayload(payload bool) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteAccessRuleWithBucketOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

/*
DeleteAccessRuleWithBucketDefault Generic error response.

swagger:response deleteAccessRuleWithBucketDefault
*/
type DeleteAccessRuleWithBucketDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.APIError `json:"body,omitempty"`
}

// NewDeleteAccessRuleWithBucketDefault creates DeleteAccessRuleWithBucketDefault with default headers values
func NewDeleteAccessRuleWithBucketDefault(code int) *DeleteAccessRuleWithBucketDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteAccessRuleWithBucketDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete access rule with bucket default response
func (o *DeleteAccessRuleWithBucketDefault) WithStatusCode(code int) *DeleteAccessRuleWithBucketDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete access rule with bucket default response
func (o *DeleteAccessRuleWithBucketDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the delete access rule with bucket default response
func (o *DeleteAccessRuleWithBucketDefault) WithPayload(payload *models.APIError) *DeleteAccessRuleWithBucketDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete access rule with bucket default response
func (o *DeleteAccessRuleWithBucketDefault) SetPayload(payload *models.APIError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteAccessRuleWithBucketDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
