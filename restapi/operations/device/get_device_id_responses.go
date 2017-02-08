package device

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/kubeIoT/api-docs/models"
)

/*GetDeviceIDOK Specified device

swagger:response getDeviceIdOK
*/
type GetDeviceIDOK struct {
	/*
	  In: Body
	*/
	Payload *models.Device `json:"body,omitempty"`
}

// NewGetDeviceIDOK creates GetDeviceIDOK with default headers values
func NewGetDeviceIDOK() *GetDeviceIDOK {
	return &GetDeviceIDOK{}
}

// WithPayload adds the payload to the get device Id o k response
func (o *GetDeviceIDOK) WithPayload(payload *models.Device) *GetDeviceIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get device Id o k response
func (o *GetDeviceIDOK) SetPayload(payload *models.Device) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetDeviceIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetDeviceIDNotFound Person was not found

swagger:response getDeviceIdNotFound
*/
type GetDeviceIDNotFound struct {
	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetDeviceIDNotFound creates GetDeviceIDNotFound with default headers values
func NewGetDeviceIDNotFound() *GetDeviceIDNotFound {
	return &GetDeviceIDNotFound{}
}

// WithPayload adds the payload to the get device Id not found response
func (o *GetDeviceIDNotFound) WithPayload(payload *models.Error) *GetDeviceIDNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get device Id not found response
func (o *GetDeviceIDNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetDeviceIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
