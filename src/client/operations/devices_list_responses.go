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

// DevicesListReader is a Reader for the DevicesList structure.
type DevicesListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DevicesListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDevicesListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDevicesListOK creates a DevicesListOK with default headers values
func NewDevicesListOK() *DevicesListOK {
	return &DevicesListOK{}
}

/*DevicesListOK handles this case with default header values.

Success
*/
type DevicesListOK struct {
	Payload *models.DeviceCollection
}

func (o *DevicesListOK) Error() string {
	return fmt.Sprintf("[GET /devices][%d] devicesListOK  %+v", 200, o.Payload)
}

func (o *DevicesListOK) GetPayload() *models.DeviceCollection {
	return o.Payload
}

func (o *DevicesListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceCollection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}