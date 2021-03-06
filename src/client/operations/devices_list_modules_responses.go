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

// DevicesListModulesReader is a Reader for the DevicesListModules structure.
type DevicesListModulesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DevicesListModulesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDevicesListModulesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDevicesListModulesOK creates a DevicesListModulesOK with default headers values
func NewDevicesListModulesOK() *DevicesListModulesOK {
	return &DevicesListModulesOK{}
}

/*DevicesListModulesOK handles this case with default header values.

Success
*/
type DevicesListModulesOK struct {
	Payload *models.Collection
}

func (o *DevicesListModulesOK) Error() string {
	return fmt.Sprintf("[GET /devices/{device_id}/modules][%d] devicesListModulesOK  %+v", 200, o.Payload)
}

func (o *DevicesListModulesOK) GetPayload() *models.Collection {
	return o.Payload
}

func (o *DevicesListModulesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Collection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
