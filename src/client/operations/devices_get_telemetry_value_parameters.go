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

// NewDevicesGetTelemetryValueParams creates a new DevicesGetTelemetryValueParams object
// with the default values initialized.
func NewDevicesGetTelemetryValueParams() *DevicesGetTelemetryValueParams {
	var ()
	return &DevicesGetTelemetryValueParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDevicesGetTelemetryValueParamsWithTimeout creates a new DevicesGetTelemetryValueParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDevicesGetTelemetryValueParamsWithTimeout(timeout time.Duration) *DevicesGetTelemetryValueParams {
	var ()
	return &DevicesGetTelemetryValueParams{

		timeout: timeout,
	}
}

// NewDevicesGetTelemetryValueParamsWithContext creates a new DevicesGetTelemetryValueParams object
// with the default values initialized, and the ability to set a context for a request
func NewDevicesGetTelemetryValueParamsWithContext(ctx context.Context) *DevicesGetTelemetryValueParams {
	var ()
	return &DevicesGetTelemetryValueParams{

		Context: ctx,
	}
}

// NewDevicesGetTelemetryValueParamsWithHTTPClient creates a new DevicesGetTelemetryValueParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDevicesGetTelemetryValueParamsWithHTTPClient(client *http.Client) *DevicesGetTelemetryValueParams {
	var ()
	return &DevicesGetTelemetryValueParams{
		HTTPClient: client,
	}
}

/*DevicesGetTelemetryValueParams contains all the parameters to send to the API endpoint
for the devices get telemetry value operation typically these are written to a http.Request
*/
type DevicesGetTelemetryValueParams struct {

	/*DeviceID
	  Unique ID of the device.

	*/
	DeviceID string
	/*TelemetryName
	  Name of this device telemetry.

	*/
	TelemetryName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the devices get telemetry value params
func (o *DevicesGetTelemetryValueParams) WithTimeout(timeout time.Duration) *DevicesGetTelemetryValueParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the devices get telemetry value params
func (o *DevicesGetTelemetryValueParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the devices get telemetry value params
func (o *DevicesGetTelemetryValueParams) WithContext(ctx context.Context) *DevicesGetTelemetryValueParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the devices get telemetry value params
func (o *DevicesGetTelemetryValueParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the devices get telemetry value params
func (o *DevicesGetTelemetryValueParams) WithHTTPClient(client *http.Client) *DevicesGetTelemetryValueParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the devices get telemetry value params
func (o *DevicesGetTelemetryValueParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDeviceID adds the deviceID to the devices get telemetry value params
func (o *DevicesGetTelemetryValueParams) WithDeviceID(deviceID string) *DevicesGetTelemetryValueParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the devices get telemetry value params
func (o *DevicesGetTelemetryValueParams) SetDeviceID(deviceID string) {
	o.DeviceID = deviceID
}

// WithTelemetryName adds the telemetryName to the devices get telemetry value params
func (o *DevicesGetTelemetryValueParams) WithTelemetryName(telemetryName string) *DevicesGetTelemetryValueParams {
	o.SetTelemetryName(telemetryName)
	return o
}

// SetTelemetryName adds the telemetryName to the devices get telemetry value params
func (o *DevicesGetTelemetryValueParams) SetTelemetryName(telemetryName string) {
	o.TelemetryName = telemetryName
}

// WriteToRequest writes these params to a swagger request
func (o *DevicesGetTelemetryValueParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param device_id
	if err := r.SetPathParam("device_id", o.DeviceID); err != nil {
		return err
	}

	// path param telemetry_name
	if err := r.SetPathParam("telemetry_name", o.TelemetryName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
