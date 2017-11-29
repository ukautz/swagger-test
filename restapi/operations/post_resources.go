// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PostResourcesHandlerFunc turns a function with the right signature into a post resources handler
type PostResourcesHandlerFunc func(PostResourcesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostResourcesHandlerFunc) Handle(params PostResourcesParams) middleware.Responder {
	return fn(params)
}

// PostResourcesHandler interface for that can handle valid post resources params
type PostResourcesHandler interface {
	Handle(PostResourcesParams) middleware.Responder
}

// NewPostResources creates a new http.Handler for the post resources operation
func NewPostResources(ctx *middleware.Context, handler PostResourcesHandler) *PostResources {
	return &PostResources{Context: ctx, Handler: handler}
}

/*PostResources swagger:route POST /resources postResources

Create new resource

*/
type PostResources struct {
	Context *middleware.Context
	Handler PostResourcesHandler
}

func (o *PostResources) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPostResourcesParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
