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

// NewGetDelegationsToValidatorParams creates a new GetDelegationsToValidatorParams object
// with the default values initialized.
func NewGetDelegationsToValidatorParams() *GetDelegationsToValidatorParams {
	var ()
	return &GetDelegationsToValidatorParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDelegationsToValidatorParamsWithTimeout creates a new GetDelegationsToValidatorParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDelegationsToValidatorParamsWithTimeout(timeout time.Duration) *GetDelegationsToValidatorParams {
	var ()
	return &GetDelegationsToValidatorParams{

		timeout: timeout,
	}
}

// NewGetDelegationsToValidatorParamsWithContext creates a new GetDelegationsToValidatorParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDelegationsToValidatorParamsWithContext(ctx context.Context) *GetDelegationsToValidatorParams {
	var ()
	return &GetDelegationsToValidatorParams{

		Context: ctx,
	}
}

// NewGetDelegationsToValidatorParamsWithHTTPClient creates a new GetDelegationsToValidatorParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDelegationsToValidatorParamsWithHTTPClient(client *http.Client) *GetDelegationsToValidatorParams {
	var ()
	return &GetDelegationsToValidatorParams{
		HTTPClient: client,
	}
}

/*GetDelegationsToValidatorParams contains all the parameters to send to the API endpoint
for the get delegations to validator operation typically these are written to a http.Request
*/
type GetDelegationsToValidatorParams struct {

	/*DelegatorAddr
	  Bech32 AccAddress of Delegator

	*/
	DelegatorAddr string
	/*ValidatorAddr
	  Bech32 OperatorAddress of validator

	*/
	ValidatorAddr string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get delegations to validator params
func (o *GetDelegationsToValidatorParams) WithTimeout(timeout time.Duration) *GetDelegationsToValidatorParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get delegations to validator params
func (o *GetDelegationsToValidatorParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get delegations to validator params
func (o *GetDelegationsToValidatorParams) WithContext(ctx context.Context) *GetDelegationsToValidatorParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get delegations to validator params
func (o *GetDelegationsToValidatorParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get delegations to validator params
func (o *GetDelegationsToValidatorParams) WithHTTPClient(client *http.Client) *GetDelegationsToValidatorParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get delegations to validator params
func (o *GetDelegationsToValidatorParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDelegatorAddr adds the delegatorAddr to the get delegations to validator params
func (o *GetDelegationsToValidatorParams) WithDelegatorAddr(delegatorAddr string) *GetDelegationsToValidatorParams {
	o.SetDelegatorAddr(delegatorAddr)
	return o
}

// SetDelegatorAddr adds the delegatorAddr to the get delegations to validator params
func (o *GetDelegationsToValidatorParams) SetDelegatorAddr(delegatorAddr string) {
	o.DelegatorAddr = delegatorAddr
}

// WithValidatorAddr adds the validatorAddr to the get delegations to validator params
func (o *GetDelegationsToValidatorParams) WithValidatorAddr(validatorAddr string) *GetDelegationsToValidatorParams {
	o.SetValidatorAddr(validatorAddr)
	return o
}

// SetValidatorAddr adds the validatorAddr to the get delegations to validator params
func (o *GetDelegationsToValidatorParams) SetValidatorAddr(validatorAddr string) {
	o.ValidatorAddr = validatorAddr
}

// WriteToRequest writes these params to a swagger request
func (o *GetDelegationsToValidatorParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param delegatorAddr
	if err := r.SetPathParam("delegatorAddr", o.DelegatorAddr); err != nil {
		return err
	}

	// path param validatorAddr
	if err := r.SetPathParam("validatorAddr", o.ValidatorAddr); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}