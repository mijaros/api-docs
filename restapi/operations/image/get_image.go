package image

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetImageHandlerFunc turns a function with the right signature into a get image handler
type GetImageHandlerFunc func(GetImageParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetImageHandlerFunc) Handle(params GetImageParams) middleware.Responder {
	return fn(params)
}

// GetImageHandler interface for that can handle valid get image params
type GetImageHandler interface {
	Handle(GetImageParams) middleware.Responder
}

// NewGetImage creates a new http.Handler for the get image operation
func NewGetImage(ctx *middleware.Context, handler GetImageHandler) *GetImage {
	return &GetImage{Context: ctx, Handler: handler}
}

/*GetImage swagger:route GET /image image getImage

Get the images in Registry

*/
type GetImage struct {
	Context *middleware.Context
	Handler GetImageHandler
}

func (o *GetImage) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewGetImageParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil {
		// bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
