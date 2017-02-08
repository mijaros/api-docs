package device

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteDeviceIDCapabilitiesHandlerFunc turns a function with the right signature into a delete device ID capabilities handler
type DeleteDeviceIDCapabilitiesHandlerFunc func(DeleteDeviceIDCapabilitiesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteDeviceIDCapabilitiesHandlerFunc) Handle(params DeleteDeviceIDCapabilitiesParams) middleware.Responder {
	return fn(params)
}

// DeleteDeviceIDCapabilitiesHandler interface for that can handle valid delete device ID capabilities params
type DeleteDeviceIDCapabilitiesHandler interface {
	Handle(DeleteDeviceIDCapabilitiesParams) middleware.Responder
}

// NewDeleteDeviceIDCapabilities creates a new http.Handler for the delete device ID capabilities operation
func NewDeleteDeviceIDCapabilities(ctx *middleware.Context, handler DeleteDeviceIDCapabilitiesHandler) *DeleteDeviceIDCapabilities {
	return &DeleteDeviceIDCapabilities{Context: ctx, Handler: handler}
}

/*DeleteDeviceIDCapabilities swagger:route DELETE /device/{id}/capabilities device deleteDeviceIdCapabilities

Delete device capability

*/
type DeleteDeviceIDCapabilities struct {
	Context *middleware.Context
	Handler DeleteDeviceIDCapabilitiesHandler
}

func (o *DeleteDeviceIDCapabilities) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewDeleteDeviceIDCapabilitiesParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil {
		// bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
