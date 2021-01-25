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

	"com.azure.iot/iotcentral/iotcgo/models"
)

// NewAPITokensSetParams creates a new APITokensSetParams object
// with the default values initialized.
func NewAPITokensSetParams() *APITokensSetParams {
	var ()
	return &APITokensSetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAPITokensSetParamsWithTimeout creates a new APITokensSetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAPITokensSetParamsWithTimeout(timeout time.Duration) *APITokensSetParams {
	var ()
	return &APITokensSetParams{

		timeout: timeout,
	}
}

// NewAPITokensSetParamsWithContext creates a new APITokensSetParams object
// with the default values initialized, and the ability to set a context for a request
func NewAPITokensSetParamsWithContext(ctx context.Context) *APITokensSetParams {
	var ()
	return &APITokensSetParams{

		Context: ctx,
	}
}

// NewAPITokensSetParamsWithHTTPClient creates a new APITokensSetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAPITokensSetParamsWithHTTPClient(client *http.Client) *APITokensSetParams {
	var ()
	return &APITokensSetParams{
		HTTPClient: client,
	}
}

/*APITokensSetParams contains all the parameters to send to the API endpoint
for the Api tokens set operation typically these are written to a http.Request
*/
type APITokensSetParams struct {

	/*Body
	  API token body.

	*/
	Body *models.APIToken
	/*TokenID
	  Unique ID for the API token.

	*/
	TokenID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the Api tokens set params
func (o *APITokensSetParams) WithTimeout(timeout time.Duration) *APITokensSetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the Api tokens set params
func (o *APITokensSetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the Api tokens set params
func (o *APITokensSetParams) WithContext(ctx context.Context) *APITokensSetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the Api tokens set params
func (o *APITokensSetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the Api tokens set params
func (o *APITokensSetParams) WithHTTPClient(client *http.Client) *APITokensSetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the Api tokens set params
func (o *APITokensSetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the Api tokens set params
func (o *APITokensSetParams) WithBody(body *models.APIToken) *APITokensSetParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the Api tokens set params
func (o *APITokensSetParams) SetBody(body *models.APIToken) {
	o.Body = body
}

// WithTokenID adds the tokenID to the Api tokens set params
func (o *APITokensSetParams) WithTokenID(tokenID string) *APITokensSetParams {
	o.SetTokenID(tokenID)
	return o
}

// SetTokenID adds the tokenId to the Api tokens set params
func (o *APITokensSetParams) SetTokenID(tokenID string) {
	o.TokenID = tokenID
}

// WriteToRequest writes these params to a swagger request
func (o *APITokensSetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param token_id
	if err := r.SetPathParam("token_id", o.TokenID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
