// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/loxilb-io/loxilb/api/models"
)

// GetConfigEndpointAllHandlerFunc turns a function with the right signature into a get config endpoint all handler
type GetConfigEndpointAllHandlerFunc func(GetConfigEndpointAllParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetConfigEndpointAllHandlerFunc) Handle(params GetConfigEndpointAllParams) middleware.Responder {
	return fn(params)
}

// GetConfigEndpointAllHandler interface for that can handle valid get config endpoint all params
type GetConfigEndpointAllHandler interface {
	Handle(GetConfigEndpointAllParams) middleware.Responder
}

// NewGetConfigEndpointAll creates a new http.Handler for the get config endpoint all operation
func NewGetConfigEndpointAll(ctx *middleware.Context, handler GetConfigEndpointAllHandler) *GetConfigEndpointAll {
	return &GetConfigEndpointAll{Context: ctx, Handler: handler}
}

/*
	GetConfigEndpointAll swagger:route GET /config/endpoint/all getConfigEndpointAll

# Get End-Points State in loxilb

Get End-Points State in loxilb
*/
type GetConfigEndpointAll struct {
	Context *middleware.Context
	Handler GetConfigEndpointAllHandler
}

func (o *GetConfigEndpointAll) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetConfigEndpointAllParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// GetConfigEndpointAllOKBody get config endpoint all o k body
//
// swagger:model GetConfigEndpointAllOKBody
type GetConfigEndpointAllOKBody struct {

	// attr
	Attr []*models.EndPointGetEntry `json:"Attr"`
}

// Validate validates this get config endpoint all o k body
func (o *GetConfigEndpointAllOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAttr(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetConfigEndpointAllOKBody) validateAttr(formats strfmt.Registry) error {
	if swag.IsZero(o.Attr) { // not required
		return nil
	}

	for i := 0; i < len(o.Attr); i++ {
		if swag.IsZero(o.Attr[i]) { // not required
			continue
		}

		if o.Attr[i] != nil {
			if err := o.Attr[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getConfigEndpointAllOK" + "." + "Attr" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getConfigEndpointAllOK" + "." + "Attr" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get config endpoint all o k body based on the context it is used
func (o *GetConfigEndpointAllOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateAttr(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetConfigEndpointAllOKBody) contextValidateAttr(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Attr); i++ {

		if o.Attr[i] != nil {

			if swag.IsZero(o.Attr[i]) { // not required
				return nil
			}

			if err := o.Attr[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getConfigEndpointAllOK" + "." + "Attr" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getConfigEndpointAllOK" + "." + "Attr" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetConfigEndpointAllOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetConfigEndpointAllOKBody) UnmarshalBinary(b []byte) error {
	var res GetConfigEndpointAllOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
