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

// NewDevicesListParams creates a new DevicesListParams object
// with the default values initialized.
func NewDevicesListParams() *DevicesListParams {

	return &DevicesListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDevicesListParamsWithTimeout creates a new DevicesListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDevicesListParamsWithTimeout(timeout time.Duration) *DevicesListParams {

	return &DevicesListParams{

		timeout: timeout,
	}
}

// NewDevicesListParamsWithContext creates a new DevicesListParams object
// with the default values initialized, and the ability to set a context for a request
func NewDevicesListParamsWithContext(ctx context.Context) *DevicesListParams {

	return &DevicesListParams{

		Context: ctx,
	}
}

// NewDevicesListParamsWithHTTPClient creates a new DevicesListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDevicesListParamsWithHTTPClient(client *http.Client) *DevicesListParams {

	return &DevicesListParams{
		HTTPClient: client,
	}
}

/*DevicesListParams contains all the parameters to send to the API endpoint
for the devices list operation typically these are written to a http.Request
*/
type DevicesListParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the devices list params
func (o *DevicesListParams) WithTimeout(timeout time.Duration) *DevicesListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the devices list params
func (o *DevicesListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the devices list params
func (o *DevicesListParams) WithContext(ctx context.Context) *DevicesListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the devices list params
func (o *DevicesListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the devices list params
func (o *DevicesListParams) WithHTTPClient(client *http.Client) *DevicesListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the devices list params
func (o *DevicesListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *DevicesListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
