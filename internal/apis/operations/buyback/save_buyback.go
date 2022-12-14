// Code generated by go-swagger; DO NOT EDIT.

package buyback

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// SaveBuybackHandlerFunc turns a function with the right signature into a save buyback handler
type SaveBuybackHandlerFunc func(SaveBuybackParams) middleware.Responder

// Handle executing the request and returning a response
func (fn SaveBuybackHandlerFunc) Handle(params SaveBuybackParams) middleware.Responder {
	return fn(params)
}

// SaveBuybackHandler interface for that can handle valid save buyback params
type SaveBuybackHandler interface {
	Handle(SaveBuybackParams) middleware.Responder
}

// NewSaveBuyback creates a new http.Handler for the save buyback operation
func NewSaveBuyback(ctx *middleware.Context, handler SaveBuybackHandler) *SaveBuyback {
	return &SaveBuyback{Context: ctx, Handler: handler}
}

/* SaveBuyback swagger:route POST /buyback buyback saveBuyback

buyback

api for buyback

*/
type SaveBuyback struct {
	Context *middleware.Context
	Handler SaveBuybackHandler
}

func (o *SaveBuyback) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewSaveBuybackParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
