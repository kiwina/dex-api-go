// Code generated by go-swagger; DO NOT EDIT.

package bank

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

// NewTransferSupervisedCoinsParams creates a new TransferSupervisedCoinsParams object
// with the default values initialized.
func NewTransferSupervisedCoinsParams() *TransferSupervisedCoinsParams {
	var ()
	return &TransferSupervisedCoinsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTransferSupervisedCoinsParamsWithTimeout creates a new TransferSupervisedCoinsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTransferSupervisedCoinsParamsWithTimeout(timeout time.Duration) *TransferSupervisedCoinsParams {
	var ()
	return &TransferSupervisedCoinsParams{

		timeout: timeout,
	}
}

// NewTransferSupervisedCoinsParamsWithContext creates a new TransferSupervisedCoinsParams object
// with the default values initialized, and the ability to set a context for a request
func NewTransferSupervisedCoinsParamsWithContext(ctx context.Context) *TransferSupervisedCoinsParams {
	var ()
	return &TransferSupervisedCoinsParams{

		Context: ctx,
	}
}

// NewTransferSupervisedCoinsParamsWithHTTPClient creates a new TransferSupervisedCoinsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTransferSupervisedCoinsParamsWithHTTPClient(client *http.Client) *TransferSupervisedCoinsParams {
	var ()
	return &TransferSupervisedCoinsParams{
		HTTPClient: client,
	}
}

/*TransferSupervisedCoinsParams contains all the parameters to send to the API endpoint
for the transfer supervised coins operation typically these are written to a http.Request
*/
type TransferSupervisedCoinsParams struct {

	/*Address
	  Account address in bech32 format

	*/
	Address string
	/*PostTxBody
	  The sender and tx information

	*/
	PostTxBody TransferSupervisedCoinsBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the transfer supervised coins params
func (o *TransferSupervisedCoinsParams) WithTimeout(timeout time.Duration) *TransferSupervisedCoinsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the transfer supervised coins params
func (o *TransferSupervisedCoinsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the transfer supervised coins params
func (o *TransferSupervisedCoinsParams) WithContext(ctx context.Context) *TransferSupervisedCoinsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the transfer supervised coins params
func (o *TransferSupervisedCoinsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the transfer supervised coins params
func (o *TransferSupervisedCoinsParams) WithHTTPClient(client *http.Client) *TransferSupervisedCoinsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the transfer supervised coins params
func (o *TransferSupervisedCoinsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAddress adds the address to the transfer supervised coins params
func (o *TransferSupervisedCoinsParams) WithAddress(address string) *TransferSupervisedCoinsParams {
	o.SetAddress(address)
	return o
}

// SetAddress adds the address to the transfer supervised coins params
func (o *TransferSupervisedCoinsParams) SetAddress(address string) {
	o.Address = address
}

// WithPostTxBody adds the postTxBody to the transfer supervised coins params
func (o *TransferSupervisedCoinsParams) WithPostTxBody(postTxBody TransferSupervisedCoinsBody) *TransferSupervisedCoinsParams {
	o.SetPostTxBody(postTxBody)
	return o
}

// SetPostTxBody adds the postTxBody to the transfer supervised coins params
func (o *TransferSupervisedCoinsParams) SetPostTxBody(postTxBody TransferSupervisedCoinsBody) {
	o.PostTxBody = postTxBody
}

// WriteToRequest writes these params to a swagger request
func (o *TransferSupervisedCoinsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param address
	if err := r.SetPathParam("address", o.Address); err != nil {
		return err
	}

	if err := r.SetBodyParam(o.PostTxBody); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
