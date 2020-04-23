// Code generated by go-swagger; DO NOT EDIT.

package staking

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

// NewGetDelegationsParams creates a new GetDelegationsParams object
// with the default values initialized.
func NewGetDelegationsParams() *GetDelegationsParams {
	var ()
	return &GetDelegationsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDelegationsParamsWithTimeout creates a new GetDelegationsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDelegationsParamsWithTimeout(timeout time.Duration) *GetDelegationsParams {
	var ()
	return &GetDelegationsParams{

		timeout: timeout,
	}
}

// NewGetDelegationsParamsWithContext creates a new GetDelegationsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDelegationsParamsWithContext(ctx context.Context) *GetDelegationsParams {
	var ()
	return &GetDelegationsParams{

		Context: ctx,
	}
}

// NewGetDelegationsParamsWithHTTPClient creates a new GetDelegationsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDelegationsParamsWithHTTPClient(client *http.Client) *GetDelegationsParams {
	var ()
	return &GetDelegationsParams{
		HTTPClient: client,
	}
}

/*GetDelegationsParams contains all the parameters to send to the API endpoint
for the get delegations operation typically these are written to a http.Request
*/
type GetDelegationsParams struct {

	/*DelegatorAddr
	  Bech32 AccAddress of Delegator

	*/
	DelegatorAddr string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get delegations params
func (o *GetDelegationsParams) WithTimeout(timeout time.Duration) *GetDelegationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get delegations params
func (o *GetDelegationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get delegations params
func (o *GetDelegationsParams) WithContext(ctx context.Context) *GetDelegationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get delegations params
func (o *GetDelegationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get delegations params
func (o *GetDelegationsParams) WithHTTPClient(client *http.Client) *GetDelegationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get delegations params
func (o *GetDelegationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDelegatorAddr adds the delegatorAddr to the get delegations params
func (o *GetDelegationsParams) WithDelegatorAddr(delegatorAddr string) *GetDelegationsParams {
	o.SetDelegatorAddr(delegatorAddr)
	return o
}

// SetDelegatorAddr adds the delegatorAddr to the get delegations params
func (o *GetDelegationsParams) SetDelegatorAddr(delegatorAddr string) {
	o.DelegatorAddr = delegatorAddr
}

// WriteToRequest writes these params to a swagger request
func (o *GetDelegationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param delegatorAddr
	if err := r.SetPathParam("delegatorAddr", o.DelegatorAddr); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}