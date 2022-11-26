// Code generated by go-swagger; DO NOT EDIT.

package buyback

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/yahfiilham/gold-store-demo/pkg/models"
)

// NewSaveBuybackParams creates a new SaveBuybackParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSaveBuybackParams() *SaveBuybackParams {
	return &SaveBuybackParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSaveBuybackParamsWithTimeout creates a new SaveBuybackParams object
// with the ability to set a timeout on a request.
func NewSaveBuybackParamsWithTimeout(timeout time.Duration) *SaveBuybackParams {
	return &SaveBuybackParams{
		timeout: timeout,
	}
}

// NewSaveBuybackParamsWithContext creates a new SaveBuybackParams object
// with the ability to set a context for a request.
func NewSaveBuybackParamsWithContext(ctx context.Context) *SaveBuybackParams {
	return &SaveBuybackParams{
		Context: ctx,
	}
}

// NewSaveBuybackParamsWithHTTPClient creates a new SaveBuybackParams object
// with the ability to set a custom HTTPClient for a request.
func NewSaveBuybackParamsWithHTTPClient(client *http.Client) *SaveBuybackParams {
	return &SaveBuybackParams{
		HTTPClient: client,
	}
}

/* SaveBuybackParams contains all the parameters to send to the API endpoint
   for the save buyback operation.

   Typically these are written to a http.Request.
*/
type SaveBuybackParams struct {

	// Data.
	Data *models.BaseRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the save buyback params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SaveBuybackParams) WithDefaults() *SaveBuybackParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the save buyback params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SaveBuybackParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the save buyback params
func (o *SaveBuybackParams) WithTimeout(timeout time.Duration) *SaveBuybackParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the save buyback params
func (o *SaveBuybackParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the save buyback params
func (o *SaveBuybackParams) WithContext(ctx context.Context) *SaveBuybackParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the save buyback params
func (o *SaveBuybackParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the save buyback params
func (o *SaveBuybackParams) WithHTTPClient(client *http.Client) *SaveBuybackParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the save buyback params
func (o *SaveBuybackParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the save buyback params
func (o *SaveBuybackParams) WithData(data *models.BaseRequest) *SaveBuybackParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the save buyback params
func (o *SaveBuybackParams) SetData(data *models.BaseRequest) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *SaveBuybackParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
