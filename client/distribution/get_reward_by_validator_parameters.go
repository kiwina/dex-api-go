// Code generated by go-swagger; DO NOT EDIT.

package distribution

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

// NewGetRewardByValidatorParams creates a new GetRewardByValidatorParams object
// with the default values initialized.
func NewGetRewardByValidatorParams() *GetRewardByValidatorParams {
	var ()
	return &GetRewardByValidatorParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRewardByValidatorParamsWithTimeout creates a new GetRewardByValidatorParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRewardByValidatorParamsWithTimeout(timeout time.Duration) *GetRewardByValidatorParams {
	var ()
	return &GetRewardByValidatorParams{

		timeout: timeout,
	}
}

// NewGetRewardByValidatorParamsWithContext creates a new GetRewardByValidatorParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRewardByValidatorParamsWithContext(ctx context.Context) *GetRewardByValidatorParams {
	var ()
	return &GetRewardByValidatorParams{

		Context: ctx,
	}
}

// NewGetRewardByValidatorParamsWithHTTPClient creates a new GetRewardByValidatorParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRewardByValidatorParamsWithHTTPClient(client *http.Client) *GetRewardByValidatorParams {
	var ()
	return &GetRewardByValidatorParams{
		HTTPClient: client,
	}
}

/*GetRewardByValidatorParams contains all the parameters to send to the API endpoint
for the get reward by validator operation typically these are written to a http.Request
*/
type GetRewardByValidatorParams struct {

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

// WithTimeout adds the timeout to the get reward by validator params
func (o *GetRewardByValidatorParams) WithTimeout(timeout time.Duration) *GetRewardByValidatorParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get reward by validator params
func (o *GetRewardByValidatorParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get reward by validator params
func (o *GetRewardByValidatorParams) WithContext(ctx context.Context) *GetRewardByValidatorParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get reward by validator params
func (o *GetRewardByValidatorParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get reward by validator params
func (o *GetRewardByValidatorParams) WithHTTPClient(client *http.Client) *GetRewardByValidatorParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get reward by validator params
func (o *GetRewardByValidatorParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDelegatorAddr adds the delegatorAddr to the get reward by validator params
func (o *GetRewardByValidatorParams) WithDelegatorAddr(delegatorAddr string) *GetRewardByValidatorParams {
	o.SetDelegatorAddr(delegatorAddr)
	return o
}

// SetDelegatorAddr adds the delegatorAddr to the get reward by validator params
func (o *GetRewardByValidatorParams) SetDelegatorAddr(delegatorAddr string) {
	o.DelegatorAddr = delegatorAddr
}

// WithValidatorAddr adds the validatorAddr to the get reward by validator params
func (o *GetRewardByValidatorParams) WithValidatorAddr(validatorAddr string) *GetRewardByValidatorParams {
	o.SetValidatorAddr(validatorAddr)
	return o
}

// SetValidatorAddr adds the validatorAddr to the get reward by validator params
func (o *GetRewardByValidatorParams) SetValidatorAddr(validatorAddr string) {
	o.ValidatorAddr = validatorAddr
}

// WriteToRequest writes these params to a swagger request
func (o *GetRewardByValidatorParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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