package capability

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/kubeIoT/api-docs/models"
)

/*GetCapabilityIDOK capability informations

swagger:response getCapabilityIdOK
*/
type GetCapabilityIDOK struct {
	/*
	  In: Body
	*/
	Payload *models.Capability `json:"body,omitempty"`
}

// NewGetCapabilityIDOK creates GetCapabilityIDOK with default headers values
func NewGetCapabilityIDOK() *GetCapabilityIDOK {
	return &GetCapabilityIDOK{}
}

// WithPayload adds the payload to the get capability Id o k response
func (o *GetCapabilityIDOK) WithPayload(payload *models.Capability) *GetCapabilityIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get capability Id o k response
func (o *GetCapabilityIDOK) SetPayload(payload *models.Capability) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCapabilityIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetCapabilityIDNotFound capability was not found

swagger:response getCapabilityIdNotFound
*/
type GetCapabilityIDNotFound struct {
	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetCapabilityIDNotFound creates GetCapabilityIDNotFound with default headers values
func NewGetCapabilityIDNotFound() *GetCapabilityIDNotFound {
	return &GetCapabilityIDNotFound{}
}

// WithPayload adds the payload to the get capability Id not found response
func (o *GetCapabilityIDNotFound) WithPayload(payload *models.Error) *GetCapabilityIDNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get capability Id not found response
func (o *GetCapabilityIDNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCapabilityIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
