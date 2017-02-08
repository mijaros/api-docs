package application

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

/*DeleteApplicationIDNoContent Application was Deleted

swagger:response deleteApplicationIdNoContent
*/
type DeleteApplicationIDNoContent struct {
}

// NewDeleteApplicationIDNoContent creates DeleteApplicationIDNoContent with default headers values
func NewDeleteApplicationIDNoContent() *DeleteApplicationIDNoContent {
	return &DeleteApplicationIDNoContent{}
}

// WriteResponse to the client
func (o *DeleteApplicationIDNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(204)
}

/*DeleteApplicationIDNotFound Application was not found

swagger:response deleteApplicationIdNotFound
*/
type DeleteApplicationIDNotFound struct {
}

// NewDeleteApplicationIDNotFound creates DeleteApplicationIDNotFound with default headers values
func NewDeleteApplicationIDNotFound() *DeleteApplicationIDNotFound {
	return &DeleteApplicationIDNotFound{}
}

// WriteResponse to the client
func (o *DeleteApplicationIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
}