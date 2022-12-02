// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/loxilb-io/loxilb/api/models"
)

// GetConfigTunnelVxlanAllOKCode is the HTTP code returned for type GetConfigTunnelVxlanAllOK
const GetConfigTunnelVxlanAllOKCode int = 200

/*
GetConfigTunnelVxlanAllOK OK

swagger:response getConfigTunnelVxlanAllOK
*/
type GetConfigTunnelVxlanAllOK struct {

	/*
	  In: Body
	*/
	Payload *GetConfigTunnelVxlanAllOKBody `json:"body,omitempty"`
}

// NewGetConfigTunnelVxlanAllOK creates GetConfigTunnelVxlanAllOK with default headers values
func NewGetConfigTunnelVxlanAllOK() *GetConfigTunnelVxlanAllOK {

	return &GetConfigTunnelVxlanAllOK{}
}

// WithPayload adds the payload to the get config tunnel vxlan all o k response
func (o *GetConfigTunnelVxlanAllOK) WithPayload(payload *GetConfigTunnelVxlanAllOKBody) *GetConfigTunnelVxlanAllOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get config tunnel vxlan all o k response
func (o *GetConfigTunnelVxlanAllOK) SetPayload(payload *GetConfigTunnelVxlanAllOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetConfigTunnelVxlanAllOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetConfigTunnelVxlanAllUnauthorizedCode is the HTTP code returned for type GetConfigTunnelVxlanAllUnauthorized
const GetConfigTunnelVxlanAllUnauthorizedCode int = 401

/*
GetConfigTunnelVxlanAllUnauthorized Invalid authentication credentials

swagger:response getConfigTunnelVxlanAllUnauthorized
*/
type GetConfigTunnelVxlanAllUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetConfigTunnelVxlanAllUnauthorized creates GetConfigTunnelVxlanAllUnauthorized with default headers values
func NewGetConfigTunnelVxlanAllUnauthorized() *GetConfigTunnelVxlanAllUnauthorized {

	return &GetConfigTunnelVxlanAllUnauthorized{}
}

// WithPayload adds the payload to the get config tunnel vxlan all unauthorized response
func (o *GetConfigTunnelVxlanAllUnauthorized) WithPayload(payload *models.Error) *GetConfigTunnelVxlanAllUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get config tunnel vxlan all unauthorized response
func (o *GetConfigTunnelVxlanAllUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetConfigTunnelVxlanAllUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetConfigTunnelVxlanAllInternalServerErrorCode is the HTTP code returned for type GetConfigTunnelVxlanAllInternalServerError
const GetConfigTunnelVxlanAllInternalServerErrorCode int = 500

/*
GetConfigTunnelVxlanAllInternalServerError Internal service error

swagger:response getConfigTunnelVxlanAllInternalServerError
*/
type GetConfigTunnelVxlanAllInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetConfigTunnelVxlanAllInternalServerError creates GetConfigTunnelVxlanAllInternalServerError with default headers values
func NewGetConfigTunnelVxlanAllInternalServerError() *GetConfigTunnelVxlanAllInternalServerError {

	return &GetConfigTunnelVxlanAllInternalServerError{}
}

// WithPayload adds the payload to the get config tunnel vxlan all internal server error response
func (o *GetConfigTunnelVxlanAllInternalServerError) WithPayload(payload *models.Error) *GetConfigTunnelVxlanAllInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get config tunnel vxlan all internal server error response
func (o *GetConfigTunnelVxlanAllInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetConfigTunnelVxlanAllInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetConfigTunnelVxlanAllServiceUnavailableCode is the HTTP code returned for type GetConfigTunnelVxlanAllServiceUnavailable
const GetConfigTunnelVxlanAllServiceUnavailableCode int = 503

/*
GetConfigTunnelVxlanAllServiceUnavailable Maintanence mode

swagger:response getConfigTunnelVxlanAllServiceUnavailable
*/
type GetConfigTunnelVxlanAllServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetConfigTunnelVxlanAllServiceUnavailable creates GetConfigTunnelVxlanAllServiceUnavailable with default headers values
func NewGetConfigTunnelVxlanAllServiceUnavailable() *GetConfigTunnelVxlanAllServiceUnavailable {

	return &GetConfigTunnelVxlanAllServiceUnavailable{}
}

// WithPayload adds the payload to the get config tunnel vxlan all service unavailable response
func (o *GetConfigTunnelVxlanAllServiceUnavailable) WithPayload(payload *models.Error) *GetConfigTunnelVxlanAllServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get config tunnel vxlan all service unavailable response
func (o *GetConfigTunnelVxlanAllServiceUnavailable) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetConfigTunnelVxlanAllServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
