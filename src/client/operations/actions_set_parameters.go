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

// NewActionsSetParams creates a new ActionsSetParams object
// with the default values initialized.
func NewActionsSetParams() *ActionsSetParams {
	var ()
	return &ActionsSetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewActionsSetParamsWithTimeout creates a new ActionsSetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewActionsSetParamsWithTimeout(timeout time.Duration) *ActionsSetParams {
	var ()
	return &ActionsSetParams{

		timeout: timeout,
	}
}

// NewActionsSetParamsWithContext creates a new ActionsSetParams object
// with the default values initialized, and the ability to set a context for a request
func NewActionsSetParamsWithContext(ctx context.Context) *ActionsSetParams {
	var ()
	return &ActionsSetParams{

		Context: ctx,
	}
}

// NewActionsSetParamsWithHTTPClient creates a new ActionsSetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewActionsSetParamsWithHTTPClient(client *http.Client) *ActionsSetParams {
	var ()
	return &ActionsSetParams{
		HTTPClient: client,
	}
}

/*ActionsSetParams contains all the parameters to send to the API endpoint
for the actions set operation typically these are written to a http.Request
*/
type ActionsSetParams struct {

	/*ActionID
	  Unique ID of the action.

	*/
	ActionID string
	/*Body
	  Action body.

	*/
	Body models.Action

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the actions set params
func (o *ActionsSetParams) WithTimeout(timeout time.Duration) *ActionsSetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the actions set params
func (o *ActionsSetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the actions set params
func (o *ActionsSetParams) WithContext(ctx context.Context) *ActionsSetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the actions set params
func (o *ActionsSetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the actions set params
func (o *ActionsSetParams) WithHTTPClient(client *http.Client) *ActionsSetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the actions set params
func (o *ActionsSetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithActionID adds the actionID to the actions set params
func (o *ActionsSetParams) WithActionID(actionID string) *ActionsSetParams {
	o.SetActionID(actionID)
	return o
}

// SetActionID adds the actionId to the actions set params
func (o *ActionsSetParams) SetActionID(actionID string) {
	o.ActionID = actionID
}

// WithBody adds the body to the actions set params
func (o *ActionsSetParams) WithBody(body models.Action) *ActionsSetParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the actions set params
func (o *ActionsSetParams) SetBody(body models.Action) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ActionsSetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param action_id
	if err := r.SetPathParam("action_id", o.ActionID); err != nil {
		return err
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
