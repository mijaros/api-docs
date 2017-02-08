package application

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetApplicationHandlerFunc turns a function with the right signature into a get application handler
type GetApplicationHandlerFunc func(GetApplicationParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetApplicationHandlerFunc) Handle(params GetApplicationParams) middleware.Responder {
	return fn(params)
}

// GetApplicationHandler interface for that can handle valid get application params
type GetApplicationHandler interface {
	Handle(GetApplicationParams) middleware.Responder
}

// NewGetApplication creates a new http.Handler for the get application operation
func NewGetApplication(ctx *middleware.Context, handler GetApplicationHandler) *GetApplication {
	return &GetApplication{Context: ctx, Handler: handler}
}

/*GetApplication swagger:route GET /application application getApplication

Get all applications

*/
type GetApplication struct {
	Context *middleware.Context
	Handler GetApplicationHandler
}

func (o *GetApplication) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewGetApplicationParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil {
		// bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
