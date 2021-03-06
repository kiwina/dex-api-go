// Code generated by go-swagger; DO NOT EDIT.

package comment

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

// NewFollowupCommentParams creates a new FollowupCommentParams object
// with the default values initialized.
func NewFollowupCommentParams() *FollowupCommentParams {
	var ()
	return &FollowupCommentParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFollowupCommentParamsWithTimeout creates a new FollowupCommentParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFollowupCommentParamsWithTimeout(timeout time.Duration) *FollowupCommentParams {
	var ()
	return &FollowupCommentParams{

		timeout: timeout,
	}
}

// NewFollowupCommentParamsWithContext creates a new FollowupCommentParams object
// with the default values initialized, and the ability to set a context for a request
func NewFollowupCommentParamsWithContext(ctx context.Context) *FollowupCommentParams {
	var ()
	return &FollowupCommentParams{

		Context: ctx,
	}
}

// NewFollowupCommentParamsWithHTTPClient creates a new FollowupCommentParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFollowupCommentParamsWithHTTPClient(client *http.Client) *FollowupCommentParams {
	var ()
	return &FollowupCommentParams{
		HTTPClient: client,
	}
}

/*FollowupCommentParams contains all the parameters to send to the API endpoint
for the followup comment operation typically these are written to a http.Request
*/
type FollowupCommentParams struct {

	/*FollowupCommentReq
	  Post a follow-up comment

	*/
	FollowupCommentReq FollowupCommentBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the followup comment params
func (o *FollowupCommentParams) WithTimeout(timeout time.Duration) *FollowupCommentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the followup comment params
func (o *FollowupCommentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the followup comment params
func (o *FollowupCommentParams) WithContext(ctx context.Context) *FollowupCommentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the followup comment params
func (o *FollowupCommentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the followup comment params
func (o *FollowupCommentParams) WithHTTPClient(client *http.Client) *FollowupCommentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the followup comment params
func (o *FollowupCommentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFollowupCommentReq adds the followupCommentReq to the followup comment params
func (o *FollowupCommentParams) WithFollowupCommentReq(followupCommentReq FollowupCommentBody) *FollowupCommentParams {
	o.SetFollowupCommentReq(followupCommentReq)
	return o
}

// SetFollowupCommentReq adds the followupCommentReq to the followup comment params
func (o *FollowupCommentParams) SetFollowupCommentReq(followupCommentReq FollowupCommentBody) {
	o.FollowupCommentReq = followupCommentReq
}

// WriteToRequest writes these params to a swagger request
func (o *FollowupCommentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.FollowupCommentReq); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
