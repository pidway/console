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
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"

	"github.com/go-openapi/swag"
)

// ListAccessRulesWithBucketURL generates an URL for the list access rules with bucket operation
type ListAccessRulesWithBucketURL struct {
	Bucket string

	Limit  *int32
	Offset *int32

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *ListAccessRulesWithBucketURL) WithBasePath(bp string) *ListAccessRulesWithBucketURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *ListAccessRulesWithBucketURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *ListAccessRulesWithBucketURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/bucket/{bucket}/access-rules"

	bucket := o.Bucket
	if bucket != "" {
		_path = strings.Replace(_path, "{bucket}", bucket, -1)
	} else {
		return nil, errors.New("bucket is required on ListAccessRulesWithBucketURL")
	}

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/api/v1"
	}
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	var limitQ string
	if o.Limit != nil {
		limitQ = swag.FormatInt32(*o.Limit)
	}
	if limitQ != "" {
		qs.Set("limit", limitQ)
	}

	var offsetQ string
	if o.Offset != nil {
		offsetQ = swag.FormatInt32(*o.Offset)
	}
	if offsetQ != "" {
		qs.Set("offset", offsetQ)
	}

	_result.RawQuery = qs.Encode()

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *ListAccessRulesWithBucketURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *ListAccessRulesWithBucketURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *ListAccessRulesWithBucketURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on ListAccessRulesWithBucketURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on ListAccessRulesWithBucketURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *ListAccessRulesWithBucketURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
