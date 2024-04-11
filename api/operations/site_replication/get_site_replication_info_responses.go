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

package site_replication

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/minio/console/models"
)

// GetSiteReplicationInfoOKCode is the HTTP code returned for type GetSiteReplicationInfoOK
const GetSiteReplicationInfoOKCode int = 200

/*
GetSiteReplicationInfoOK A successful response.

swagger:response getSiteReplicationInfoOK
*/
type GetSiteReplicationInfoOK struct {

	/*
	  In: Body
	*/
	Payload *models.SiteReplicationInfoResponse `json:"body,omitempty"`
}

// NewGetSiteReplicationInfoOK creates GetSiteReplicationInfoOK with default headers values
func NewGetSiteReplicationInfoOK() *GetSiteReplicationInfoOK {

	return &GetSiteReplicationInfoOK{}
}

// WithPayload adds the payload to the get site replication info o k response
func (o *GetSiteReplicationInfoOK) WithPayload(payload *models.SiteReplicationInfoResponse) *GetSiteReplicationInfoOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get site replication info o k response
func (o *GetSiteReplicationInfoOK) SetPayload(payload *models.SiteReplicationInfoResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSiteReplicationInfoOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
GetSiteReplicationInfoDefault Generic error response.

swagger:response getSiteReplicationInfoDefault
*/
type GetSiteReplicationInfoDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.APIError `json:"body,omitempty"`
}

// NewGetSiteReplicationInfoDefault creates GetSiteReplicationInfoDefault with default headers values
func NewGetSiteReplicationInfoDefault(code int) *GetSiteReplicationInfoDefault {
	if code <= 0 {
		code = 500
	}

	return &GetSiteReplicationInfoDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get site replication info default response
func (o *GetSiteReplicationInfoDefault) WithStatusCode(code int) *GetSiteReplicationInfoDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get site replication info default response
func (o *GetSiteReplicationInfoDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get site replication info default response
func (o *GetSiteReplicationInfoDefault) WithPayload(payload *models.APIError) *GetSiteReplicationInfoDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get site replication info default response
func (o *GetSiteReplicationInfoDefault) SetPayload(payload *models.APIError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSiteReplicationInfoDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
