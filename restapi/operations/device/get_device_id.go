package device

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetDeviceIDHandlerFunc turns a function with the right signature into a get device ID handler
type GetDeviceIDHandlerFunc func(GetDeviceIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetDeviceIDHandlerFunc) Handle(params GetDeviceIDParams) middleware.Responder {
	return fn(params)
}

// GetDeviceIDHandler interface for that can handle valid get device ID params
type GetDeviceIDHandler interface {
	Handle(GetDeviceIDParams) middleware.Responder
}

// NewGetDeviceID creates a new http.Handler for the get device ID operation
func NewGetDeviceID(ctx *middleware.Context, handler GetDeviceIDHandler) *GetDeviceID {
	return &GetDeviceID{Context: ctx, Handler: handler}
}

/*GetDeviceID swagger:route GET /device/{id} device getDeviceId

Get device by ID

*/
type GetDeviceID struct {
	Context *middleware.Context
	Handler GetDeviceIDHandler
}

func (o *GetDeviceID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewGetDeviceIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil {
		// bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
