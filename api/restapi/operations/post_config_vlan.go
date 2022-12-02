// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PostConfigVlanHandlerFunc turns a function with the right signature into a post config vlan handler
type PostConfigVlanHandlerFunc func(PostConfigVlanParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostConfigVlanHandlerFunc) Handle(params PostConfigVlanParams) middleware.Responder {
	return fn(params)
}

// PostConfigVlanHandler interface for that can handle valid post config vlan params
type PostConfigVlanHandler interface {
	Handle(PostConfigVlanParams) middleware.Responder
}

// NewPostConfigVlan creates a new http.Handler for the post config vlan operation
func NewPostConfigVlan(ctx *middleware.Context, handler PostConfigVlanHandler) *PostConfigVlan {
	return &PostConfigVlan{Context: ctx, Handler: handler}
}

/*
	PostConfigVlan swagger:route POST /config/vlan postConfigVlan

# Create vlan interface in the device

Create vlan interface in the device
*/
type PostConfigVlan struct {
	Context *middleware.Context
	Handler PostConfigVlanHandler
}

func (o *PostConfigVlan) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPostConfigVlanParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
