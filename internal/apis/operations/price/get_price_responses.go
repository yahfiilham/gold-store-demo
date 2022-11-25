// Code generated by go-swagger; DO NOT EDIT.

package price

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/yahfiilham/gold-store-demo/pkg/models"
)

// GetPriceOKCode is the HTTP code returned for type GetPriceOK
const GetPriceOKCode int = 200

/*GetPriceOK success

swagger:response getPriceOK
*/
type GetPriceOK struct {

	/*
	  In: Body
	*/
	Payload *GetPriceOKBody `json:"body,omitempty"`
}

// NewGetPriceOK creates GetPriceOK with default headers values
func NewGetPriceOK() *GetPriceOK {

	return &GetPriceOK{}
}

// WithPayload adds the payload to the get price o k response
func (o *GetPriceOK) WithPayload(payload *GetPriceOKBody) *GetPriceOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get price o k response
func (o *GetPriceOK) SetPayload(payload *GetPriceOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPriceOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetPriceDefault error

swagger:response getPriceDefault
*/
type GetPriceDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.BaseResponse `json:"body,omitempty"`
}

// NewGetPriceDefault creates GetPriceDefault with default headers values
func NewGetPriceDefault(code int) *GetPriceDefault {
	if code <= 0 {
		code = 500
	}

	return &GetPriceDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get price default response
func (o *GetPriceDefault) WithStatusCode(code int) *GetPriceDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get price default response
func (o *GetPriceDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get price default response
func (o *GetPriceDefault) WithPayload(payload *models.BaseResponse) *GetPriceDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get price default response
func (o *GetPriceDefault) SetPayload(payload *models.BaseResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPriceDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
