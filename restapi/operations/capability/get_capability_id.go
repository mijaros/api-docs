package capability

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetCapabilityIDHandlerFunc turns a function with the right signature into a get capability ID handler
type GetCapabilityIDHandlerFunc func(GetCapabilityIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetCapabilityIDHandlerFunc) Handle(params GetCapabilityIDParams) middleware.Responder {
	return fn(params)
}

// GetCapabilityIDHandler interface for that can handle valid get capability ID params
type GetCapabilityIDHandler interface {
	Handle(GetCapabilityIDParams) middleware.Responder
}

// NewGetCapabilityID creates a new http.Handler for the get capability ID operation
func NewGetCapabilityID(ctx *middleware.Context, handler GetCapabilityIDHandler) *GetCapabilityID {
	return &GetCapabilityID{Context: ctx, Handler: handler}
}

/*GetCapabilityID swagger:route GET /capability/{id} capability getCapabilityId

Get capability information

*/
type GetCapabilityID struct {
	Context *middleware.Context
	Handler GetCapabilityIDHandler
}

func (o *GetCapabilityID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewGetCapabilityIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil {
		// bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
