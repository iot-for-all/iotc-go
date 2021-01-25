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

// DevicesUpdateComponentPropertiesReader is a Reader for the DevicesUpdateComponentProperties structure.
type DevicesUpdateComponentPropertiesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DevicesUpdateComponentPropertiesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewDevicesUpdateComponentPropertiesAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDevicesUpdateComponentPropertiesAccepted creates a DevicesUpdateComponentPropertiesAccepted with default headers values
func NewDevicesUpdateComponentPropertiesAccepted() *DevicesUpdateComponentPropertiesAccepted {
	return &DevicesUpdateComponentPropertiesAccepted{}
}

/*DevicesUpdateComponentPropertiesAccepted handles this case with default header values.

Success
*/
type DevicesUpdateComponentPropertiesAccepted struct {
	Payload models.DeviceProperties
}

func (o *DevicesUpdateComponentPropertiesAccepted) Error() string {
	return fmt.Sprintf("[PUT /devices/{device_id}/components/{component_name}/properties][%d] devicesUpdateComponentPropertiesAccepted  %+v", 202, o.Payload)
}

func (o *DevicesUpdateComponentPropertiesAccepted) GetPayload() models.DeviceProperties {
	return o.Payload
}

func (o *DevicesUpdateComponentPropertiesAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
