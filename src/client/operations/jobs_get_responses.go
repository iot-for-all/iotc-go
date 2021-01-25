// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"com.azure.iot/iotcentral/iotcgo/models"
)

// JobsGetReader is a Reader for the JobsGet structure.
type JobsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *JobsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewJobsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewJobsGetOK creates a JobsGetOK with default headers values
func NewJobsGetOK() *JobsGetOK {
	return &JobsGetOK{}
}

/*JobsGetOK handles this case with default header values.

Success
*/
type JobsGetOK struct {
	Payload *models.Job
}

func (o *JobsGetOK) Error() string {
	return fmt.Sprintf("[GET /jobs/{job_id}][%d] jobsGetOK  %+v", 200, o.Payload)
}

func (o *JobsGetOK) GetPayload() *models.Job {
	return o.Payload
}

func (o *JobsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Job)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
