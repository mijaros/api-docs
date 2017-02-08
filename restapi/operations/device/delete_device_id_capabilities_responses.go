package device

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/kubeIoT/api-docs/models"
)

/*DeleteDeviceIDCapabilitiesNoContent Device capability was Deleted

swagger:response deleteDeviceIdCapabilitiesNoContent
*/
type DeleteDeviceIDCapabilitiesNoContent struct {
}

// NewDeleteDeviceIDCapabilitiesNoContent creates DeleteDeviceIDCapabilitiesNoContent with default headers values
func NewDeleteDeviceIDCapabilitiesNoContent() *DeleteDeviceIDCapabilitiesNoContent {
	return &DeleteDeviceIDCapabilitiesNoContent{}
}

// WriteResponse to the client
func (o *DeleteDeviceIDCapabilitiesNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(204)
}

/*DeleteDeviceIDCapabilitiesNotFound Device capability was not found

swagger:response deleteDeviceIdCapabilitiesNotFound
*/
type DeleteDeviceIDCapabilitiesNotFound struct {
	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteDeviceIDCapabilitiesNotFound creates DeleteDeviceIDCapabilitiesNotFound with default headers values
func NewDeleteDeviceIDCapabilitiesNotFound() *DeleteDeviceIDCapabilitiesNotFound {
	return &DeleteDeviceIDCapabilitiesNotFound{}
}

// WithPayload adds the payload to the delete device Id capabilities not found response
func (o *DeleteDeviceIDCapabilitiesNotFound) WithPayload(payload *models.Error) *DeleteDeviceIDCapabilitiesNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete device Id capabilities not found response
func (o *DeleteDeviceIDCapabilitiesNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteDeviceIDCapabilitiesNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
