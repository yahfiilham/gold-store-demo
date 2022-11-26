// Code generated by go-swagger; DO NOT EDIT.

package transaction

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"

	"github.com/go-openapi/swag"
)

// GetMutationURL generates an URL for the get mutation operation
type GetMutationURL struct {
	AccountNo string
	EndDate   int64
	StartDate int64

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetMutationURL) WithBasePath(bp string) *GetMutationURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetMutationURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *GetMutationURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/mutation"

	_basePath := o._basePath
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	accountNoQ := o.AccountNo
	if accountNoQ != "" {
		qs.Set("account_no", accountNoQ)
	}

	endDateQ := swag.FormatInt64(o.EndDate)
	if endDateQ != "" {
		qs.Set("end_date", endDateQ)
	}

	startDateQ := swag.FormatInt64(o.StartDate)
	if startDateQ != "" {
		qs.Set("start_date", startDateQ)
	}

	_result.RawQuery = qs.Encode()

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *GetMutationURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *GetMutationURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *GetMutationURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on GetMutationURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on GetMutationURL")
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
func (o *GetMutationURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
