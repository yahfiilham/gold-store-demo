// Code generated by go-swagger; DO NOT EDIT.

package balance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/yahfiilham/gold-store-demo/pkg/models"
)

// GetBalanceReader is a Reader for the GetBalance structure.
type GetBalanceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBalanceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBalanceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetBalanceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetBalanceOK creates a GetBalanceOK with default headers values
func NewGetBalanceOK() *GetBalanceOK {
	return &GetBalanceOK{}
}

/* GetBalanceOK describes a response with status code 200, with default header values.

success
*/
type GetBalanceOK struct {
	Payload *GetBalanceOKBody
}

// IsSuccess returns true when this get balance o k response has a 2xx status code
func (o *GetBalanceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get balance o k response has a 3xx status code
func (o *GetBalanceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get balance o k response has a 4xx status code
func (o *GetBalanceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get balance o k response has a 5xx status code
func (o *GetBalanceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get balance o k response a status code equal to that given
func (o *GetBalanceOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetBalanceOK) Error() string {
	return fmt.Sprintf("[GET /balance][%d] getBalanceOK  %+v", 200, o.Payload)
}

func (o *GetBalanceOK) String() string {
	return fmt.Sprintf("[GET /balance][%d] getBalanceOK  %+v", 200, o.Payload)
}

func (o *GetBalanceOK) GetPayload() *GetBalanceOKBody {
	return o.Payload
}

func (o *GetBalanceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetBalanceOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBalanceDefault creates a GetBalanceDefault with default headers values
func NewGetBalanceDefault(code int) *GetBalanceDefault {
	return &GetBalanceDefault{
		_statusCode: code,
	}
}

/* GetBalanceDefault describes a response with status code -1, with default header values.

error
*/
type GetBalanceDefault struct {
	_statusCode int

	Payload *models.BaseResponse
}

// Code gets the status code for the get balance default response
func (o *GetBalanceDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this get balance default response has a 2xx status code
func (o *GetBalanceDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get balance default response has a 3xx status code
func (o *GetBalanceDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get balance default response has a 4xx status code
func (o *GetBalanceDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get balance default response has a 5xx status code
func (o *GetBalanceDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get balance default response a status code equal to that given
func (o *GetBalanceDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GetBalanceDefault) Error() string {
	return fmt.Sprintf("[GET /balance][%d] GetBalance default  %+v", o._statusCode, o.Payload)
}

func (o *GetBalanceDefault) String() string {
	return fmt.Sprintf("[GET /balance][%d] GetBalance default  %+v", o._statusCode, o.Payload)
}

func (o *GetBalanceDefault) GetPayload() *models.BaseResponse {
	return o.Payload
}

func (o *GetBalanceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BaseResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetBalanceOKBody get balance o k body
swagger:model GetBalanceOKBody
*/
type GetBalanceOKBody struct {

	// data
	Data *models.Account `json:"data,omitempty"`
}

// Validate validates this get balance o k body
func (o *GetBalanceOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetBalanceOKBody) validateData(formats strfmt.Registry) error {
	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getBalanceOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getBalanceOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get balance o k body based on the context it is used
func (o *GetBalanceOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetBalanceOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {
		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getBalanceOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getBalanceOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetBalanceOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetBalanceOKBody) UnmarshalBinary(b []byte) error {
	var res GetBalanceOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
