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

// DeviceTemplatesListDevicesReader is a Reader for the DeviceTemplatesListDevices structure.
type DeviceTemplatesListDevicesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeviceTemplatesListDevicesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeviceTemplatesListDevicesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeviceTemplatesListDevicesOK creates a DeviceTemplatesListDevicesOK with default headers values
func NewDeviceTemplatesListDevicesOK() *DeviceTemplatesListDevicesOK {
	return &DeviceTemplatesListDevicesOK{}
}

/*DeviceTemplatesListDevicesOK handles this case with default header values.

Success
*/
type DeviceTemplatesListDevicesOK struct {
	Payload *models.DeviceCollection
}

func (o *DeviceTemplatesListDevicesOK) Error() string {
	return fmt.Sprintf("[GET /deviceTemplates/{device_template_id}/devices][%d] deviceTemplatesListDevicesOK  %+v", 200, o.Payload)
}

func (o *DeviceTemplatesListDevicesOK) GetPayload() *models.DeviceCollection {
	return o.Payload
}

func (o *DeviceTemplatesListDevicesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceCollection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}