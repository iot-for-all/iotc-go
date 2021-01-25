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

// DevicesSetReader is a Reader for the DevicesSet structure.
type DevicesSetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DevicesSetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDevicesSetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDevicesSetOK creates a DevicesSetOK with default headers values
func NewDevicesSetOK() *DevicesSetOK {
	return &DevicesSetOK{}
}

/*DevicesSetOK handles this case with default header values.

Success
*/
type DevicesSetOK struct {
	Payload *models.Device
}

func (o *DevicesSetOK) Error() string {
	return fmt.Sprintf("[PUT /devices/{device_id}][%d] devicesSetOK  %+v", 200, o.Payload)
}

func (o *DevicesSetOK) GetPayload() *models.Device {
	return o.Payload
}

func (o *DevicesSetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Device)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
