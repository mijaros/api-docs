package image

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/kubeIoT/api-docs/models"
)

/*PostImageCreated Image was created

swagger:response postImageCreated
*/
type PostImageCreated struct {
	/*
	  In: Body
	*/
	Payload *models.Image `json:"body,omitempty"`
}

// NewPostImageCreated creates PostImageCreated with default headers values
func NewPostImageCreated() *PostImageCreated {
	return &PostImageCreated{}
}

// WithPayload adds the payload to the post image created response
func (o *PostImageCreated) WithPayload(payload *models.Image) *PostImageCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post image created response
func (o *PostImageCreated) SetPayload(payload *models.Image) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostImageCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*PostImageForbidden User is unauthorized to create new image

swagger:response postImageForbidden
*/
type PostImageForbidden struct {
	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostImageForbidden creates PostImageForbidden with default headers values
func NewPostImageForbidden() *PostImageForbidden {
	return &PostImageForbidden{}
}

// WithPayload adds the payload to the post image forbidden response
func (o *PostImageForbidden) WithPayload(payload *models.Error) *PostImageForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post image forbidden response
func (o *PostImageForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostImageForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
