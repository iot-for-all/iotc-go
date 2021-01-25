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

// NewDevicesGetComponentCommandHistoryParams creates a new DevicesGetComponentCommandHistoryParams object
// with the default values initialized.
func NewDevicesGetComponentCommandHistoryParams() *DevicesGetComponentCommandHistoryParams {
	var ()
	return &DevicesGetComponentCommandHistoryParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDevicesGetComponentCommandHistoryParamsWithTimeout creates a new DevicesGetComponentCommandHistoryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDevicesGetComponentCommandHistoryParamsWithTimeout(timeout time.Duration) *DevicesGetComponentCommandHistoryParams {
	var ()
	return &DevicesGetComponentCommandHistoryParams{

		timeout: timeout,
	}
}

// NewDevicesGetComponentCommandHistoryParamsWithContext creates a new DevicesGetComponentCommandHistoryParams object
// with the default values initialized, and the ability to set a context for a request
func NewDevicesGetComponentCommandHistoryParamsWithContext(ctx context.Context) *DevicesGetComponentCommandHistoryParams {
	var ()
	return &DevicesGetComponentCommandHistoryParams{

		Context: ctx,
	}
}

// NewDevicesGetComponentCommandHistoryParamsWithHTTPClient creates a new DevicesGetComponentCommandHistoryParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDevicesGetComponentCommandHistoryParamsWithHTTPClient(client *http.Client) *DevicesGetComponentCommandHistoryParams {
	var ()
	return &DevicesGetComponentCommandHistoryParams{
		HTTPClient: client,
	}
}

/*DevicesGetComponentCommandHistoryParams contains all the parameters to send to the API endpoint
for the devices get component command history operation typically these are written to a http.Request
*/
type DevicesGetComponentCommandHistoryParams struct {

	/*CommandName
	  Name of this device command.

	*/
	CommandName string
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

// WithTimeout adds the timeout to the devices get component command history params
func (o *DevicesGetComponentCommandHistoryParams) WithTimeout(timeout time.Duration) *DevicesGetComponentCommandHistoryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the devices get component command history params
func (o *DevicesGetComponentCommandHistoryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the devices get component command history params
func (o *DevicesGetComponentCommandHistoryParams) WithContext(ctx context.Context) *DevicesGetComponentCommandHistoryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the devices get component command history params
func (o *DevicesGetComponentCommandHistoryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the devices get component command history params
func (o *DevicesGetComponentCommandHistoryParams) WithHTTPClient(client *http.Client) *DevicesGetComponentCommandHistoryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the devices get component command history params
func (o *DevicesGetComponentCommandHistoryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCommandName adds the commandName to the devices get component command history params
func (o *DevicesGetComponentCommandHistoryParams) WithCommandName(commandName string) *DevicesGetComponentCommandHistoryParams {
	o.SetCommandName(commandName)
	return o
}

// SetCommandName adds the commandName to the devices get component command history params
func (o *DevicesGetComponentCommandHistoryParams) SetCommandName(commandName string) {
	o.CommandName = commandName
}

// WithComponentName adds the componentName to the devices get component command history params
func (o *DevicesGetComponentCommandHistoryParams) WithComponentName(componentName string) *DevicesGetComponentCommandHistoryParams {
	o.SetComponentName(componentName)
	return o
}

// SetComponentName adds the componentName to the devices get component command history params
func (o *DevicesGetComponentCommandHistoryParams) SetComponentName(componentName string) {
	o.ComponentName = componentName
}

// WithDeviceID adds the deviceID to the devices get component command history params
func (o *DevicesGetComponentCommandHistoryParams) WithDeviceID(deviceID string) *DevicesGetComponentCommandHistoryParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the devices get component command history params
func (o *DevicesGetComponentCommandHistoryParams) SetDeviceID(deviceID string) {
	o.DeviceID = deviceID
}

// WriteToRequest writes these params to a swagger request
func (o *DevicesGetComponentCommandHistoryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param command_name
	if err := r.SetPathParam("command_name", o.CommandName); err != nil {
		return err
	}

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
