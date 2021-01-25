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

// NewDeviceTemplatesSetParams creates a new DeviceTemplatesSetParams object
// with the default values initialized.
func NewDeviceTemplatesSetParams() *DeviceTemplatesSetParams {
	var ()
	return &DeviceTemplatesSetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeviceTemplatesSetParamsWithTimeout creates a new DeviceTemplatesSetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeviceTemplatesSetParamsWithTimeout(timeout time.Duration) *DeviceTemplatesSetParams {
	var ()
	return &DeviceTemplatesSetParams{

		timeout: timeout,
	}
}

// NewDeviceTemplatesSetParamsWithContext creates a new DeviceTemplatesSetParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeviceTemplatesSetParamsWithContext(ctx context.Context) *DeviceTemplatesSetParams {
	var ()
	return &DeviceTemplatesSetParams{

		Context: ctx,
	}
}

// NewDeviceTemplatesSetParamsWithHTTPClient creates a new DeviceTemplatesSetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeviceTemplatesSetParamsWithHTTPClient(client *http.Client) *DeviceTemplatesSetParams {
	var ()
	return &DeviceTemplatesSetParams{
		HTTPClient: client,
	}
}

/*DeviceTemplatesSetParams contains all the parameters to send to the API endpoint
for the device templates set operation typically these are written to a http.Request
*/
type DeviceTemplatesSetParams struct {

	/*Body
	  Device template body.

	*/
	Body *models.DeviceTemplate
	/*DeviceTemplateID
	  Unique ID of the device template.

	*/
	DeviceTemplateID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the device templates set params
func (o *DeviceTemplatesSetParams) WithTimeout(timeout time.Duration) *DeviceTemplatesSetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the device templates set params
func (o *DeviceTemplatesSetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the device templates set params
func (o *DeviceTemplatesSetParams) WithContext(ctx context.Context) *DeviceTemplatesSetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the device templates set params
func (o *DeviceTemplatesSetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the device templates set params
func (o *DeviceTemplatesSetParams) WithHTTPClient(client *http.Client) *DeviceTemplatesSetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the device templates set params
func (o *DeviceTemplatesSetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the device templates set params
func (o *DeviceTemplatesSetParams) WithBody(body *models.DeviceTemplate) *DeviceTemplatesSetParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the device templates set params
func (o *DeviceTemplatesSetParams) SetBody(body *models.DeviceTemplate) {
	o.Body = body
}

// WithDeviceTemplateID adds the deviceTemplateID to the device templates set params
func (o *DeviceTemplatesSetParams) WithDeviceTemplateID(deviceTemplateID string) *DeviceTemplatesSetParams {
	o.SetDeviceTemplateID(deviceTemplateID)
	return o
}

// SetDeviceTemplateID adds the deviceTemplateId to the device templates set params
func (o *DeviceTemplatesSetParams) SetDeviceTemplateID(deviceTemplateID string) {
	o.DeviceTemplateID = deviceTemplateID
}

// WriteToRequest writes these params to a swagger request
func (o *DeviceTemplatesSetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param device_template_id
	if err := r.SetPathParam("device_template_id", o.DeviceTemplateID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
