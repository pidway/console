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
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/minio/console/models"
)

// GetBucketVersioningHandlerFunc turns a function with the right signature into a get bucket versioning handler
type GetBucketVersioningHandlerFunc func(GetBucketVersioningParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn GetBucketVersioningHandlerFunc) Handle(params GetBucketVersioningParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// GetBucketVersioningHandler interface for that can handle valid get bucket versioning params
type GetBucketVersioningHandler interface {
	Handle(GetBucketVersioningParams, *models.Principal) middleware.Responder
}

// NewGetBucketVersioning creates a new http.Handler for the get bucket versioning operation
func NewGetBucketVersioning(ctx *middleware.Context, handler GetBucketVersioningHandler) *GetBucketVersioning {
	return &GetBucketVersioning{Context: ctx, Handler: handler}
}

/*
	GetBucketVersioning swagger:route GET /buckets/{bucket_name}/versioning Bucket getBucketVersioning

Bucket Versioning
*/
type GetBucketVersioning struct {
	Context *middleware.Context
	Handler GetBucketVersioningHandler
}

func (o *GetBucketVersioning) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetBucketVersioningParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal *models.Principal
	if uprinc != nil {
		principal = uprinc.(*models.Principal) // this is really a models.Principal, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
