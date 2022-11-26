// Code generated by go-swagger; DO NOT EDIT.

package transaction

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

	"github.com/yahfiilham/gold-store-demo/pkg/models"
)

// GetMutationHandlerFunc turns a function with the right signature into a get mutation handler
type GetMutationHandlerFunc func(GetMutationParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetMutationHandlerFunc) Handle(params GetMutationParams) middleware.Responder {
	return fn(params)
}

// GetMutationHandler interface for that can handle valid get mutation params
type GetMutationHandler interface {
	Handle(GetMutationParams) middleware.Responder
}

// NewGetMutation creates a new http.Handler for the get mutation operation
func NewGetMutation(ctx *middleware.Context, handler GetMutationHandler) *GetMutation {
	return &GetMutation{Context: ctx, Handler: handler}
}

/* GetMutation swagger:route GET /mutation transaction getMutation

get data transaction

api for get data transaction

*/
type GetMutation struct {
	Context *middleware.Context
	Handler GetMutationHandler
}

func (o *GetMutation) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetMutationParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// GetMutationOKBody get mutation o k body
//
// swagger:model GetMutationOKBody
type GetMutationOKBody struct {

	// data
	Data []*models.Transaction `json:"data"`
}

// Validate validates this get mutation o k body
func (o *GetMutationOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetMutationOKBody) validateData(formats strfmt.Registry) error {
	if swag.IsZero(o.Data) { // not required
		return nil
	}

	for i := 0; i < len(o.Data); i++ {
		if swag.IsZero(o.Data[i]) { // not required
			continue
		}

		if o.Data[i] != nil {
			if err := o.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getMutationOK" + "." + "data" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getMutationOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get mutation o k body based on the context it is used
func (o *GetMutationOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetMutationOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Data); i++ {

		if o.Data[i] != nil {
			if err := o.Data[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getMutationOK" + "." + "data" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getMutationOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetMutationOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetMutationOKBody) UnmarshalBinary(b []byte) error {
	var res GetMutationOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
