// Code generated by go-swagger; DO NOT EDIT.

package market

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
)

// NewGetMarketParamsParams creates a new GetMarketParamsParams object
// with the default values initialized.
func NewGetMarketParamsParams() *GetMarketParamsParams {

	return &GetMarketParamsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetMarketParamsParamsWithTimeout creates a new GetMarketParamsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetMarketParamsParamsWithTimeout(timeout time.Duration) *GetMarketParamsParams {

	return &GetMarketParamsParams{

		timeout: timeout,
	}
}

// NewGetMarketParamsParamsWithContext creates a new GetMarketParamsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetMarketParamsParamsWithContext(ctx context.Context) *GetMarketParamsParams {

	return &GetMarketParamsParams{

		Context: ctx,
	}
}

// NewGetMarketParamsParamsWithHTTPClient creates a new GetMarketParamsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetMarketParamsParamsWithHTTPClient(client *http.Client) *GetMarketParamsParams {

	return &GetMarketParamsParams{
		HTTPClient: client,
	}
}

/*GetMarketParamsParams contains all the parameters to send to the API endpoint
for the get market params operation typically these are written to a http.Request
*/
type GetMarketParamsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get market params params
func (o *GetMarketParamsParams) WithTimeout(timeout time.Duration) *GetMarketParamsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get market params params
func (o *GetMarketParamsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get market params params
func (o *GetMarketParamsParams) WithContext(ctx context.Context) *GetMarketParamsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get market params params
func (o *GetMarketParamsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get market params params
func (o *GetMarketParamsParams) WithHTTPClient(client *http.Client) *GetMarketParamsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get market params params
func (o *GetMarketParamsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetMarketParamsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}