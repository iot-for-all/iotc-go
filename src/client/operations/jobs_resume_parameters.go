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

// NewJobsResumeParams creates a new JobsResumeParams object
// with the default values initialized.
func NewJobsResumeParams() *JobsResumeParams {
	var ()
	return &JobsResumeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewJobsResumeParamsWithTimeout creates a new JobsResumeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewJobsResumeParamsWithTimeout(timeout time.Duration) *JobsResumeParams {
	var ()
	return &JobsResumeParams{

		timeout: timeout,
	}
}

// NewJobsResumeParamsWithContext creates a new JobsResumeParams object
// with the default values initialized, and the ability to set a context for a request
func NewJobsResumeParamsWithContext(ctx context.Context) *JobsResumeParams {
	var ()
	return &JobsResumeParams{

		Context: ctx,
	}
}

// NewJobsResumeParamsWithHTTPClient creates a new JobsResumeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewJobsResumeParamsWithHTTPClient(client *http.Client) *JobsResumeParams {
	var ()
	return &JobsResumeParams{
		HTTPClient: client,
	}
}

/*JobsResumeParams contains all the parameters to send to the API endpoint
for the jobs resume operation typically these are written to a http.Request
*/
type JobsResumeParams struct {

	/*JobID
	  Unique ID of the job.

	*/
	JobID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the jobs resume params
func (o *JobsResumeParams) WithTimeout(timeout time.Duration) *JobsResumeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the jobs resume params
func (o *JobsResumeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the jobs resume params
func (o *JobsResumeParams) WithContext(ctx context.Context) *JobsResumeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the jobs resume params
func (o *JobsResumeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the jobs resume params
func (o *JobsResumeParams) WithHTTPClient(client *http.Client) *JobsResumeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the jobs resume params
func (o *JobsResumeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithJobID adds the jobID to the jobs resume params
func (o *JobsResumeParams) WithJobID(jobID string) *JobsResumeParams {
	o.SetJobID(jobID)
	return o
}

// SetJobID adds the jobId to the jobs resume params
func (o *JobsResumeParams) SetJobID(jobID string) {
	o.JobID = jobID
}

// WriteToRequest writes these params to a swagger request
func (o *JobsResumeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param job_id
	if err := r.SetPathParam("job_id", o.JobID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
