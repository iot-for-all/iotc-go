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

// NewDevicesSetParams creates a new DevicesSetParams object
// with the default values initialized.
func NewDevicesSetParams() *DevicesSetParams {
	var ()
	return &DevicesSetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDevicesSetParamsWithTimeout creates a new DevicesSetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDevicesSetParamsWithTimeout(timeout time.Duration) *DevicesSetParams {
	var ()
	return &DevicesSetParams{

		timeout: timeout,
	}
}

// NewDevicesSetParamsWithContext creates a new DevicesSetParams object
// with the default values initialized, and the ability to set a context for a request
func NewDevicesSetParamsWithContext(ctx context.Context) *DevicesSetParams {
	var ()
	return &DevicesSetParams{

		Context: ctx,
	}
}

// NewDevicesSetParamsWithHTTPClient creates a new DevicesSetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDevicesSetParamsWithHTTPClient(client *http.Client) *DevicesSetParams {
	var ()
	return &DevicesSetParams{
		HTTPClient: client,
	}
}

/*DevicesSetParams contains all the parameters to send to the API endpoint
for the devices set operation typically these are written to a http.Request
*/
type DevicesSetParams struct {

	/*Body
	  Device body.

	*/
	Body *models.Device
	/*DeviceID
	  Unique ID of the device.

	*/
	DeviceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the devices set params
func (o *DevicesSetParams) WithTimeout(timeout time.Duration) *DevicesSetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the devices set params
func (o *DevicesSetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the devices set params
func (o *DevicesSetParams) WithContext(ctx context.Context) *DevicesSetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the devices set params
func (o *DevicesSetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the devices set params
func (o *DevicesSetParams) WithHTTPClient(client *http.Client) *DevicesSetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the devices set params
func (o *DevicesSetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the devices set params
func (o *DevicesSetParams) WithBody(body *models.Device) *DevicesSetParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the devices set params
func (o *DevicesSetParams) SetBody(body *models.Device) {
	o.Body = body
}

// WithDeviceID adds the deviceID to the devices set params
func (o *DevicesSetParams) WithDeviceID(deviceID string) *DevicesSetParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the devices set params
func (o *DevicesSetParams) SetDeviceID(deviceID string) {
	o.DeviceID = deviceID
}

// WriteToRequest writes these params to a swagger request
func (o *DevicesSetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
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
