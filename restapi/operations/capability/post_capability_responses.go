package capability

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/kubeIoT/api-docs/models"
)

/*PostCapabilityCreated capability was created

swagger:response postCapabilityCreated
*/
type PostCapabilityCreated struct {
	/*
	  In: Body
	*/
	Payload *models.Capability `json:"body,omitempty"`
}

// NewPostCapabilityCreated creates PostCapabilityCreated with default headers values
func NewPostCapabilityCreated() *PostCapabilityCreated {
	return &PostCapabilityCreated{}
}

// WithPayload adds the payload to the post capability created response
func (o *PostCapabilityCreated) WithPayload(payload *models.Capability) *PostCapabilityCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post capability created response
func (o *PostCapabilityCreated) SetPayload(payload *models.Capability) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostCapabilityCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*PostCapabilityForbidden User is unauthorized to create new capability

swagger:response postCapabilityForbidden
*/
type PostCapabilityForbidden struct {
	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostCapabilityForbidden creates PostCapabilityForbidden with default headers values
func NewPostCapabilityForbidden() *PostCapabilityForbidden {
	return &PostCapabilityForbidden{}
}

// WithPayload adds the payload to the post capability forbidden response
func (o *PostCapabilityForbidden) WithPayload(payload *models.Error) *PostCapabilityForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post capability forbidden response
func (o *PostCapabilityForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostCapabilityForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
