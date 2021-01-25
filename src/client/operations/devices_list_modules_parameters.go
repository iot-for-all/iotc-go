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

// NewDevicesListModulesParams creates a new DevicesListModulesParams object
// with the default values initialized.
func NewDevicesListModulesParams() *DevicesListModulesParams {
	var ()
	return &DevicesListModulesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDevicesListModulesParamsWithTimeout creates a new DevicesListModulesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDevicesListModulesParamsWithTimeout(timeout time.Duration) *DevicesListModulesParams {
	var ()
	return &DevicesListModulesParams{

		timeout: timeout,
	}
}

// NewDevicesListModulesParamsWithContext creates a new DevicesListModulesParams object
// with the default values initialized, and the ability to set a context for a request
func NewDevicesListModulesParamsWithContext(ctx context.Context) *DevicesListModulesParams {
	var ()
	return &DevicesListModulesParams{

		Context: ctx,
	}
}

// NewDevicesListModulesParamsWithHTTPClient creates a new DevicesListModulesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDevicesListModulesParamsWithHTTPClient(client *http.Client) *DevicesListModulesParams {
	var ()
	return &DevicesListModulesParams{
		HTTPClient: client,
	}
}

/*DevicesListModulesParams contains all the parameters to send to the API endpoint
for the devices list modules operation typically these are written to a http.Request
*/
type DevicesListModulesParams struct {

	/*DeviceID
	  Unique ID of the device.

	*/
	DeviceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the devices list modules params
func (o *DevicesListModulesParams) WithTimeout(timeout time.Duration) *DevicesListModulesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the devices list modules params
func (o *DevicesListModulesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the devices list modules params
func (o *DevicesListModulesParams) WithContext(ctx context.Context) *DevicesListModulesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the devices list modules params
func (o *DevicesListModulesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the devices list modules params
func (o *DevicesListModulesParams) WithHTTPClient(client *http.Client) *DevicesListModulesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the devices list modules params
func (o *DevicesListModulesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDeviceID adds the deviceID to the devices list modules params
func (o *DevicesListModulesParams) WithDeviceID(deviceID string) *DevicesListModulesParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the devices list modules params
func (o *DevicesListModulesParams) SetDeviceID(deviceID string) {
	o.DeviceID = deviceID
}

// WriteToRequest writes these params to a swagger request
func (o *DevicesListModulesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
