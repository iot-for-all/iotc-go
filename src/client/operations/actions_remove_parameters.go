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

// NewActionsRemoveParams creates a new ActionsRemoveParams object
// with the default values initialized.
func NewActionsRemoveParams() *ActionsRemoveParams {
	var ()
	return &ActionsRemoveParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewActionsRemoveParamsWithTimeout creates a new ActionsRemoveParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewActionsRemoveParamsWithTimeout(timeout time.Duration) *ActionsRemoveParams {
	var ()
	return &ActionsRemoveParams{

		timeout: timeout,
	}
}

// NewActionsRemoveParamsWithContext creates a new ActionsRemoveParams object
// with the default values initialized, and the ability to set a context for a request
func NewActionsRemoveParamsWithContext(ctx context.Context) *ActionsRemoveParams {
	var ()
	return &ActionsRemoveParams{

		Context: ctx,
	}
}

// NewActionsRemoveParamsWithHTTPClient creates a new ActionsRemoveParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewActionsRemoveParamsWithHTTPClient(client *http.Client) *ActionsRemoveParams {
	var ()
	return &ActionsRemoveParams{
		HTTPClient: client,
	}
}

/*ActionsRemoveParams contains all the parameters to send to the API endpoint
for the actions remove operation typically these are written to a http.Request
*/
type ActionsRemoveParams struct {

	/*ActionID
	  Unique ID of the action.

	*/
	ActionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the actions remove params
func (o *ActionsRemoveParams) WithTimeout(timeout time.Duration) *ActionsRemoveParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the actions remove params
func (o *ActionsRemoveParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the actions remove params
func (o *ActionsRemoveParams) WithContext(ctx context.Context) *ActionsRemoveParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the actions remove params
func (o *ActionsRemoveParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the actions remove params
func (o *ActionsRemoveParams) WithHTTPClient(client *http.Client) *ActionsRemoveParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the actions remove params
func (o *ActionsRemoveParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithActionID adds the actionID to the actions remove params
func (o *ActionsRemoveParams) WithActionID(actionID string) *ActionsRemoveParams {
	o.SetActionID(actionID)
	return o
}

// SetActionID adds the actionId to the actions remove params
func (o *ActionsRemoveParams) SetActionID(actionID string) {
	o.ActionID = actionID
}

// WriteToRequest writes these params to a swagger request
func (o *ActionsRemoveParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param action_id
	if err := r.SetPathParam("action_id", o.ActionID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
