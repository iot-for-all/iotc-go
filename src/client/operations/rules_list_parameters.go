// Code generated by go-swagger; DO NOT EDIT.

package operations

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

// NewRulesListParams creates a new RulesListParams object
// with the default values initialized.
func NewRulesListParams() *RulesListParams {

	return &RulesListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRulesListParamsWithTimeout creates a new RulesListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRulesListParamsWithTimeout(timeout time.Duration) *RulesListParams {

	return &RulesListParams{

		timeout: timeout,
	}
}

// NewRulesListParamsWithContext creates a new RulesListParams object
// with the default values initialized, and the ability to set a context for a request
func NewRulesListParamsWithContext(ctx context.Context) *RulesListParams {

	return &RulesListParams{

		Context: ctx,
	}
}

// NewRulesListParamsWithHTTPClient creates a new RulesListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRulesListParamsWithHTTPClient(client *http.Client) *RulesListParams {

	return &RulesListParams{
		HTTPClient: client,
	}
}

/*RulesListParams contains all the parameters to send to the API endpoint
for the rules list operation typically these are written to a http.Request
*/
type RulesListParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the rules list params
func (o *RulesListParams) WithTimeout(timeout time.Duration) *RulesListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the rules list params
func (o *RulesListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the rules list params
func (o *RulesListParams) WithContext(ctx context.Context) *RulesListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the rules list params
func (o *RulesListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the rules list params
func (o *RulesListParams) WithHTTPClient(client *http.Client) *RulesListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the rules list params
func (o *RulesListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *RulesListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
