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

// NewContinuousDataExportsGetParams creates a new ContinuousDataExportsGetParams object
// with the default values initialized.
func NewContinuousDataExportsGetParams() *ContinuousDataExportsGetParams {
	var ()
	return &ContinuousDataExportsGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewContinuousDataExportsGetParamsWithTimeout creates a new ContinuousDataExportsGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewContinuousDataExportsGetParamsWithTimeout(timeout time.Duration) *ContinuousDataExportsGetParams {
	var ()
	return &ContinuousDataExportsGetParams{

		timeout: timeout,
	}
}

// NewContinuousDataExportsGetParamsWithContext creates a new ContinuousDataExportsGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewContinuousDataExportsGetParamsWithContext(ctx context.Context) *ContinuousDataExportsGetParams {
	var ()
	return &ContinuousDataExportsGetParams{

		Context: ctx,
	}
}

// NewContinuousDataExportsGetParamsWithHTTPClient creates a new ContinuousDataExportsGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewContinuousDataExportsGetParamsWithHTTPClient(client *http.Client) *ContinuousDataExportsGetParams {
	var ()
	return &ContinuousDataExportsGetParams{
		HTTPClient: client,
	}
}

/*ContinuousDataExportsGetParams contains all the parameters to send to the API endpoint
for the continuous data exports get operation typically these are written to a http.Request
*/
type ContinuousDataExportsGetParams struct {

	/*ExportID
	  Unique ID for the continuous data export.

	*/
	ExportID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the continuous data exports get params
func (o *ContinuousDataExportsGetParams) WithTimeout(timeout time.Duration) *ContinuousDataExportsGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the continuous data exports get params
func (o *ContinuousDataExportsGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the continuous data exports get params
func (o *ContinuousDataExportsGetParams) WithContext(ctx context.Context) *ContinuousDataExportsGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the continuous data exports get params
func (o *ContinuousDataExportsGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the continuous data exports get params
func (o *ContinuousDataExportsGetParams) WithHTTPClient(client *http.Client) *ContinuousDataExportsGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the continuous data exports get params
func (o *ContinuousDataExportsGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExportID adds the exportID to the continuous data exports get params
func (o *ContinuousDataExportsGetParams) WithExportID(exportID string) *ContinuousDataExportsGetParams {
	o.SetExportID(exportID)
	return o
}

// SetExportID adds the exportId to the continuous data exports get params
func (o *ContinuousDataExportsGetParams) SetExportID(exportID string) {
	o.ExportID = exportID
}

// WriteToRequest writes these params to a swagger request
func (o *ContinuousDataExportsGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param export_id
	if err := r.SetPathParam("export_id", o.ExportID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}