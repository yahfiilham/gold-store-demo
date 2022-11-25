// Code generated by go-swagger; DO NOT EDIT.

package health_check

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/yahfiilham/gold-store-demo/pkg/models"
)

// GetHealthCheckOKCode is the HTTP code returned for type GetHealthCheckOK
const GetHealthCheckOKCode int = 200

/*GetHealthCheckOK success

swagger:response getHealthCheckOK
*/
type GetHealthCheckOK struct {

	/*
	  In: Body
	*/
	Payload *models.BaseResponse `json:"body,omitempty"`
}

// NewGetHealthCheckOK creates GetHealthCheckOK with default headers values
func NewGetHealthCheckOK() *GetHealthCheckOK {

	return &GetHealthCheckOK{}
}

// WithPayload adds the payload to the get health check o k response
func (o *GetHealthCheckOK) WithPayload(payload *models.BaseResponse) *GetHealthCheckOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get health check o k response
func (o *GetHealthCheckOK) SetPayload(payload *models.BaseResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetHealthCheckOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetHealthCheckDefault error

swagger:response getHealthCheckDefault
*/
type GetHealthCheckDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.BaseResponse `json:"body,omitempty"`
}

// NewGetHealthCheckDefault creates GetHealthCheckDefault with default headers values
func NewGetHealthCheckDefault(code int) *GetHealthCheckDefault {
	if code <= 0 {
		code = 500
	}

	return &GetHealthCheckDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get health check default response
func (o *GetHealthCheckDefault) WithStatusCode(code int) *GetHealthCheckDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get health check default response
func (o *GetHealthCheckDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get health check default response
func (o *GetHealthCheckDefault) WithPayload(payload *models.BaseResponse) *GetHealthCheckDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get health check default response
func (o *GetHealthCheckDefault) SetPayload(payload *models.BaseResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetHealthCheckDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
