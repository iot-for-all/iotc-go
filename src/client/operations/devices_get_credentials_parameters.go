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

// NewDevicesGetCredentialsParams creates a new DevicesGetCredentialsParams object
// with the default values initialized.
func NewDevicesGetCredentialsParams() *DevicesGetCredentialsParams {
	var ()
	return &DevicesGetCredentialsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDevicesGetCredentialsParamsWithTimeout creates a new DevicesGetCredentialsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDevicesGetCredentialsParamsWithTimeout(timeout time.Duration) *DevicesGetCredentialsParams {
	var ()
	return &DevicesGetCredentialsParams{

		timeout: timeout,
	}
}

// NewDevicesGetCredentialsParamsWithContext creates a new DevicesGetCredentialsParams object
// with the default values initialized, and the ability to set a context for a request
func NewDevicesGetCredentialsParamsWithContext(ctx context.Context) *DevicesGetCredentialsParams {
	var ()
	return &DevicesGetCredentialsParams{

		Context: ctx,
	}
}

// NewDevicesGetCredentialsParamsWithHTTPClient creates a new DevicesGetCredentialsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDevicesGetCredentialsParamsWithHTTPClient(client *http.Client) *DevicesGetCredentialsParams {
	var ()
	return &DevicesGetCredentialsParams{
		HTTPClient: client,
	}
}

/*DevicesGetCredentialsParams contains all the parameters to send to the API endpoint
for the devices get credentials operation typically these are written to a http.Request
*/
type DevicesGetCredentialsParams struct {

	/*DeviceID
	  Unique ID of the device.

	*/
	DeviceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the devices get credentials params
func (o *DevicesGetCredentialsParams) WithTimeout(timeout time.Duration) *DevicesGetCredentialsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the devices get credentials params
func (o *DevicesGetCredentialsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the devices get credentials params
func (o *DevicesGetCredentialsParams) WithContext(ctx context.Context) *DevicesGetCredentialsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the devices get credentials params
func (o *DevicesGetCredentialsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the devices get credentials params
func (o *DevicesGetCredentialsParams) WithHTTPClient(client *http.Client) *DevicesGetCredentialsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the devices get credentials params
func (o *DevicesGetCredentialsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDeviceID adds the deviceID to the devices get credentials params
func (o *DevicesGetCredentialsParams) WithDeviceID(deviceID string) *DevicesGetCredentialsParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the devices get credentials params
func (o *DevicesGetCredentialsParams) SetDeviceID(deviceID string) {
	o.DeviceID = deviceID
}

// WriteToRequest writes these params to a swagger request
func (o *DevicesGetCredentialsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param device_id
	if err := r.SetPathParam("device_id", o.DeviceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
