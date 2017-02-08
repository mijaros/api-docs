package capability

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/kubeIoT/api-docs/models"
)

/*GetCapabilityOK List of capabilities

swagger:response getCapabilityOK
*/
type GetCapabilityOK struct {
	/*
	  In: Body
	*/
	Payload []*models.Capability `json:"body,omitempty"`
}

// NewGetCapabilityOK creates GetCapabilityOK with default headers values
func NewGetCapabilityOK() *GetCapabilityOK {
	return &GetCapabilityOK{}
}

// WithPayload adds the payload to the get capability o k response
func (o *GetCapabilityOK) WithPayload(payload []*models.Capability) *GetCapabilityOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get capability o k response
func (o *GetCapabilityOK) SetPayload(payload []*models.Capability) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCapabilityOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make([]*models.Capability, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

/*GetCapabilityForbidden User  is unauthorized

swagger:response getCapabilityForbidden
*/
type GetCapabilityForbidden struct {
	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetCapabilityForbidden creates GetCapabilityForbidden with default headers values
func NewGetCapabilityForbidden() *GetCapabilityForbidden {
	return &GetCapabilityForbidden{}
}

// WithPayload adds the payload to the get capability forbidden response
func (o *GetCapabilityForbidden) WithPayload(payload *models.Error) *GetCapabilityForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get capability forbidden response
func (o *GetCapabilityForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCapabilityForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
