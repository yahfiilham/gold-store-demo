// Code generated by go-swagger; DO NOT EDIT.

package balance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/yahfiilham/gold-store-demo/pkg/models"
)

// GetBalanceOKCode is the HTTP code returned for type GetBalanceOK
const GetBalanceOKCode int = 200

/*GetBalanceOK success

swagger:response getBalanceOK
*/
type GetBalanceOK struct {

	/*
	  In: Body
	*/
	Payload *GetBalanceOKBody `json:"body,omitempty"`
}

// NewGetBalanceOK creates GetBalanceOK with default headers values
func NewGetBalanceOK() *GetBalanceOK {

	return &GetBalanceOK{}
}

// WithPayload adds the payload to the get balance o k response
func (o *GetBalanceOK) WithPayload(payload *GetBalanceOKBody) *GetBalanceOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get balance o k response
func (o *GetBalanceOK) SetPayload(payload *GetBalanceOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetBalanceOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetBalanceDefault error

swagger:response getBalanceDefault
*/
type GetBalanceDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.BaseResponse `json:"body,omitempty"`
}

// NewGetBalanceDefault creates GetBalanceDefault with default headers values
func NewGetBalanceDefault(code int) *GetBalanceDefault {
	if code <= 0 {
		code = 500
	}

	return &GetBalanceDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get balance default response
func (o *GetBalanceDefault) WithStatusCode(code int) *GetBalanceDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get balance default response
func (o *GetBalanceDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get balance default response
func (o *GetBalanceDefault) WithPayload(payload *models.BaseResponse) *GetBalanceDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get balance default response
func (o *GetBalanceDefault) SetPayload(payload *models.BaseResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetBalanceDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
