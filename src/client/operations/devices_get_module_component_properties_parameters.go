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

// NewDevicesGetModuleComponentPropertiesParams creates a new DevicesGetModuleComponentPropertiesParams object
// with the default values initialized.
func NewDevicesGetModuleComponentPropertiesParams() *DevicesGetModuleComponentPropertiesParams {
	var ()
	return &DevicesGetModuleComponentPropertiesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDevicesGetModuleComponentPropertiesParamsWithTimeout creates a new DevicesGetModuleComponentPropertiesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDevicesGetModuleComponentPropertiesParamsWithTimeout(timeout time.Duration) *DevicesGetModuleComponentPropertiesParams {
	var ()
	return &DevicesGetModuleComponentPropertiesParams{

		timeout: timeout,
	}
}

// NewDevicesGetModuleComponentPropertiesParamsWithContext creates a new DevicesGetModuleComponentPropertiesParams object
// with the default values initialized, and the ability to set a context for a request
func NewDevicesGetModuleComponentPropertiesParamsWithContext(ctx context.Context) *DevicesGetModuleComponentPropertiesParams {
	var ()
	return &DevicesGetModuleComponentPropertiesParams{

		Context: ctx,
	}
}

// NewDevicesGetModuleComponentPropertiesParamsWithHTTPClient creates a new DevicesGetModuleComponentPropertiesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDevicesGetModuleComponentPropertiesParamsWithHTTPClient(client *http.Client) *DevicesGetModuleComponentPropertiesParams {
	var ()
	return &DevicesGetModuleComponentPropertiesParams{
		HTTPClient: client,
	}
}

/*DevicesGetModuleComponentPropertiesParams contains all the parameters to send to the API endpoint
for the devices get module component properties operation typically these are written to a http.Request
*/
type DevicesGetModuleComponentPropertiesParams struct {

	/*ComponentName
	  Name of the device component.

	*/
	ComponentName string
	/*DeviceID
	  Unique ID of the device.

	*/
	DeviceID string
	/*ModuleName
	  Name of the device module.

	*/
	ModuleName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the devices get module component properties params
func (o *DevicesGetModuleComponentPropertiesParams) WithTimeout(timeout time.Duration) *DevicesGetModuleComponentPropertiesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the devices get module component properties params
func (o *DevicesGetModuleComponentPropertiesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the devices get module component properties params
func (o *DevicesGetModuleComponentPropertiesParams) WithContext(ctx context.Context) *DevicesGetModuleComponentPropertiesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the devices get module component properties params
func (o *DevicesGetModuleComponentPropertiesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the devices get module component properties params
func (o *DevicesGetModuleComponentPropertiesParams) WithHTTPClient(client *http.Client) *DevicesGetModuleComponentPropertiesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the devices get module component properties params
func (o *DevicesGetModuleComponentPropertiesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithComponentName adds the componentName to the devices get module component properties params
func (o *DevicesGetModuleComponentPropertiesParams) WithComponentName(componentName string) *DevicesGetModuleComponentPropertiesParams {
	o.SetComponentName(componentName)
	return o
}

// SetComponentName adds the componentName to the devices get module component properties params
func (o *DevicesGetModuleComponentPropertiesParams) SetComponentName(componentName string) {
	o.ComponentName = componentName
}

// WithDeviceID adds the deviceID to the devices get module component properties params
func (o *DevicesGetModuleComponentPropertiesParams) WithDeviceID(deviceID string) *DevicesGetModuleComponentPropertiesParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the devices get module component properties params
func (o *DevicesGetModuleComponentPropertiesParams) SetDeviceID(deviceID string) {
	o.DeviceID = deviceID
}

// WithModuleName adds the moduleName to the devices get module component properties params
func (o *DevicesGetModuleComponentPropertiesParams) WithModuleName(moduleName string) *DevicesGetModuleComponentPropertiesParams {
	o.SetModuleName(moduleName)
	return o
}

// SetModuleName adds the moduleName to the devices get module component properties params
func (o *DevicesGetModuleComponentPropertiesParams) SetModuleName(moduleName string) {
	o.ModuleName = moduleName
}

// WriteToRequest writes these params to a swagger request
func (o *DevicesGetModuleComponentPropertiesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param module_name
	if err := r.SetPathParam("module_name", o.ModuleName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
