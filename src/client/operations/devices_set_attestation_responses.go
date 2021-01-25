// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"com.azure.iot/iotcentral/iotcgo/models"
)

// DevicesSetAttestationReader is a Reader for the DevicesSetAttestation structure.
type DevicesSetAttestationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DevicesSetAttestationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDevicesSetAttestationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDevicesSetAttestationOK creates a DevicesSetAttestationOK with default headers values
func NewDevicesSetAttestationOK() *DevicesSetAttestationOK {
	return &DevicesSetAttestationOK{}
}

/*DevicesSetAttestationOK handles this case with default header values.

Success
*/
type DevicesSetAttestationOK struct {
	Payload models.Attestation
}

func (o *DevicesSetAttestationOK) Error() string {
	return fmt.Sprintf("[PUT /devices/{device_id}/attestation][%d] devicesSetAttestationOK  %+v", 200, o.Payload)
}

func (o *DevicesSetAttestationOK) GetPayload() models.Attestation {
	return o.Payload
}

func (o *DevicesSetAttestationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalAttestation(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}
