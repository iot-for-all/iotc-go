// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// JobsResumeReader is a Reader for the JobsResume structure.
type JobsResumeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *JobsResumeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewJobsResumeNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewJobsResumeNoContent creates a JobsResumeNoContent with default headers values
func NewJobsResumeNoContent() *JobsResumeNoContent {
	return &JobsResumeNoContent{}
}

/*JobsResumeNoContent handles this case with default header values.

Success
*/
type JobsResumeNoContent struct {
}

func (o *JobsResumeNoContent) Error() string {
	return fmt.Sprintf("[POST /jobs/{job_id}/resume][%d] jobsResumeNoContent ", 204)
}

func (o *JobsResumeNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
