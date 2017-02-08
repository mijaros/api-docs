package device

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteDeviceIDHandlerFunc turns a function with the right signature into a delete device ID handler
type DeleteDeviceIDHandlerFunc func(DeleteDeviceIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteDeviceIDHandlerFunc) Handle(params DeleteDeviceIDParams) middleware.Responder {
	return fn(params)
}

// DeleteDeviceIDHandler interface for that can handle valid delete device ID params
type DeleteDeviceIDHandler interface {
	Handle(DeleteDeviceIDParams) middleware.Responder
}

// NewDeleteDeviceID creates a new http.Handler for the delete device ID operation
func NewDeleteDeviceID(ctx *middleware.Context, handler DeleteDeviceIDHandler) *DeleteDeviceID {
	return &DeleteDeviceID{Context: ctx, Handler: handler}
}

/*DeleteDeviceID swagger:route DELETE /device/{id} device deleteDeviceId

Delete selected device by ID

*/
type DeleteDeviceID struct {
	Context *middleware.Context
	Handler DeleteDeviceIDHandler
}

func (o *DeleteDeviceID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewDeleteDeviceIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil {
		// bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}