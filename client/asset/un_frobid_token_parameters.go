// Code generated by go-swagger; DO NOT EDIT.

package asset

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

// NewUnFrobidTokenParams creates a new UnFrobidTokenParams object
// with the default values initialized.
func NewUnFrobidTokenParams() *UnFrobidTokenParams {
	var ()
	return &UnFrobidTokenParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUnFrobidTokenParamsWithTimeout creates a new UnFrobidTokenParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUnFrobidTokenParamsWithTimeout(timeout time.Duration) *UnFrobidTokenParams {
	var ()
	return &UnFrobidTokenParams{

		timeout: timeout,
	}
}

// NewUnFrobidTokenParamsWithContext creates a new UnFrobidTokenParams object
// with the default values initialized, and the ability to set a context for a request
func NewUnFrobidTokenParamsWithContext(ctx context.Context) *UnFrobidTokenParams {
	var ()
	return &UnFrobidTokenParams{

		Context: ctx,
	}
}

// NewUnFrobidTokenParamsWithHTTPClient creates a new UnFrobidTokenParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUnFrobidTokenParamsWithHTTPClient(client *http.Client) *UnFrobidTokenParams {
	var ()
	return &UnFrobidTokenParams{
		HTTPClient: client,
	}
}

/*UnFrobidTokenParams contains all the parameters to send to the API endpoint
for the un frobid token operation typically these are written to a http.Request
*/
type UnFrobidTokenParams struct {

	/*BaseReq
	  base req

	*/
	BaseReq UnFrobidTokenBody
	/*Symbol
	  token symbol

	*/
	Symbol string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the un frobid token params
func (o *UnFrobidTokenParams) WithTimeout(timeout time.Duration) *UnFrobidTokenParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the un frobid token params
func (o *UnFrobidTokenParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the un frobid token params
func (o *UnFrobidTokenParams) WithContext(ctx context.Context) *UnFrobidTokenParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the un frobid token params
func (o *UnFrobidTokenParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the un frobid token params
func (o *UnFrobidTokenParams) WithHTTPClient(client *http.Client) *UnFrobidTokenParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the un frobid token params
func (o *UnFrobidTokenParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBaseReq adds the baseReq to the un frobid token params
func (o *UnFrobidTokenParams) WithBaseReq(baseReq UnFrobidTokenBody) *UnFrobidTokenParams {
	o.SetBaseReq(baseReq)
	return o
}

// SetBaseReq adds the baseReq to the un frobid token params
func (o *UnFrobidTokenParams) SetBaseReq(baseReq UnFrobidTokenBody) {
	o.BaseReq = baseReq
}

// WithSymbol adds the symbol to the un frobid token params
func (o *UnFrobidTokenParams) WithSymbol(symbol string) *UnFrobidTokenParams {
	o.SetSymbol(symbol)
	return o
}

// SetSymbol adds the symbol to the un frobid token params
func (o *UnFrobidTokenParams) SetSymbol(symbol string) {
	o.Symbol = symbol
}

// WriteToRequest writes these params to a swagger request
func (o *UnFrobidTokenParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.BaseReq); err != nil {
		return err
	}

	// path param symbol
	if err := r.SetPathParam("symbol", o.Symbol); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
