// Code generated by go-swagger; DO NOT EDIT.

package health_check

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetHealthCheckHandlerFunc turns a function with the right signature into a get health check handler
type GetHealthCheckHandlerFunc func(GetHealthCheckParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetHealthCheckHandlerFunc) Handle(params GetHealthCheckParams) middleware.Responder {
	return fn(params)
}

// GetHealthCheckHandler interface for that can handle valid get health check params
type GetHealthCheckHandler interface {
	Handle(GetHealthCheckParams) middleware.Responder
}

// NewGetHealthCheck creates a new http.Handler for the get health check operation
func NewGetHealthCheck(ctx *middleware.Context, handler GetHealthCheckHandler) *GetHealthCheck {
	return &GetHealthCheck{Context: ctx, Handler: handler}
}

/* GetHealthCheck swagger:route GET /health health check getHealthCheck

check server

*/
type GetHealthCheck struct {
	Context *middleware.Context
	Handler GetHealthCheckHandler
}

func (o *GetHealthCheck) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetHealthCheckParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
