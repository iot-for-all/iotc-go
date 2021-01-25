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

// NewDevicesGetComponentPropertiesParams creates a new DevicesGetComponentPropertiesParams object
// with the default values initialized.
func NewDevicesGetComponentPropertiesParams() *DevicesGetComponentPropertiesParams {
	var ()
	return &DevicesGetComponentPropertiesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDevicesGetComponentPropertiesParamsWithTimeout creates a new DevicesGetComponentPropertiesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDevicesGetComponentPropertiesParamsWithTimeout(timeout time.Duration) *DevicesGetComponentPropertiesParams {
	var ()
	return &DevicesGetComponentPropertiesParams{

		timeout: timeout,
	}
}

// NewDevicesGetComponentPropertiesParamsWithContext creates a new DevicesGetComponentPropertiesParams object
// with the default values initialized, and the ability to set a context for a request
func NewDevicesGetComponentPropertiesParamsWithContext(ctx context.Context) *DevicesGetComponentPropertiesParams {
	var ()
	return &DevicesGetComponentPropertiesParams{

		Context: ctx,
	}
}

// NewDevicesGetComponentPropertiesParamsWithHTTPClient creates a new DevicesGetComponentPropertiesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDevicesGetComponentPropertiesParamsWithHTTPClient(client *http.Client) *DevicesGetComponentPropertiesParams {
	var ()
	return &DevicesGetComponentPropertiesParams{
		HTTPClient: client,
	}
}

/*DevicesGetComponentPropertiesParams contains all the parameters to send to the API endpoint
for the devices get component properties operation typically these are written to a http.Request
*/
type DevicesGetComponentPropertiesParams struct {

	/*ComponentName
	  Name of the device component.

	*/
	ComponentName string
	/*DeviceID
	  Unique ID of the device.

	*/
	DeviceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the devices get component properties params
func (o *DevicesGetComponentPropertiesParams) WithTimeout(timeout time.Duration) *DevicesGetComponentPropertiesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the devices get component properties params
func (o *DevicesGetComponentPropertiesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the devices get component properties params
func (o *DevicesGetComponentPropertiesParams) WithContext(ctx context.Context) *DevicesGetComponentPropertiesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the devices get component properties params
func (o *DevicesGetComponentPropertiesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the devices get component properties params
func (o *DevicesGetComponentPropertiesParams) WithHTTPClient(client *http.Client) *DevicesGetComponentPropertiesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the devices get component properties params
func (o *DevicesGetComponentPropertiesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithComponentName adds the componentName to the devices get component properties params
func (o *DevicesGetComponentPropertiesParams) WithComponentName(componentName string) *DevicesGetComponentPropertiesParams {
	o.SetComponentName(componentName)
	return o
}

// SetComponentName adds the componentName to the devices get component properties params
func (o *DevicesGetComponentPropertiesParams) SetComponentName(componentName string) {
	o.ComponentName = componentName
}

// WithDeviceID adds the deviceID to the devices get component properties params
func (o *DevicesGetComponentPropertiesParams) WithDeviceID(deviceID string) *DevicesGetComponentPropertiesParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the devices get component properties params
func (o *DevicesGetComponentPropertiesParams) SetDeviceID(deviceID string) {
	o.DeviceID = deviceID
}

// WriteToRequest writes these params to a swagger request
func (o *DevicesGetComponentPropertiesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param component_name
	if err := r.SetPathParam("component_name", o.ComponentName); err != nil {
		return err
	}

	// path param device_id
	if err := r.SetPathParam("device_id", o.DeviceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
