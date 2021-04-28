// Code generated by go-swagger; DO NOT EDIT.

package read

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/ory/keto-client-go/models"
)

// GetRelationTuplesReader is a Reader for the GetRelationTuples structure.
type GetRelationTuplesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRelationTuplesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRelationTuplesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetRelationTuplesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetRelationTuplesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRelationTuplesOK creates a GetRelationTuplesOK with default headers values
func NewGetRelationTuplesOK() *GetRelationTuplesOK {
	return &GetRelationTuplesOK{}
}

/* GetRelationTuplesOK describes a response with status code 200, with default header values.

getRelationTuplesResponse
*/
type GetRelationTuplesOK struct {
	Payload *models.GetRelationTuplesResponse
}

func (o *GetRelationTuplesOK) Error() string {
	return fmt.Sprintf("[GET /relation-tuples][%d] getRelationTuplesOK  %+v", 200, o.Payload)
}
func (o *GetRelationTuplesOK) GetPayload() *models.GetRelationTuplesResponse {
	return o.Payload
}

func (o *GetRelationTuplesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetRelationTuplesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRelationTuplesNotFound creates a GetRelationTuplesNotFound with default headers values
func NewGetRelationTuplesNotFound() *GetRelationTuplesNotFound {
	return &GetRelationTuplesNotFound{}
}

/* GetRelationTuplesNotFound describes a response with status code 404, with default header values.

The standard error format
*/
type GetRelationTuplesNotFound struct {
	Payload *GetRelationTuplesNotFoundBody
}

func (o *GetRelationTuplesNotFound) Error() string {
	return fmt.Sprintf("[GET /relation-tuples][%d] getRelationTuplesNotFound  %+v", 404, o.Payload)
}
func (o *GetRelationTuplesNotFound) GetPayload() *GetRelationTuplesNotFoundBody {
	return o.Payload
}

func (o *GetRelationTuplesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetRelationTuplesNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRelationTuplesInternalServerError creates a GetRelationTuplesInternalServerError with default headers values
func NewGetRelationTuplesInternalServerError() *GetRelationTuplesInternalServerError {
	return &GetRelationTuplesInternalServerError{}
}

/* GetRelationTuplesInternalServerError describes a response with status code 500, with default header values.

The standard error format
*/
type GetRelationTuplesInternalServerError struct {
	Payload *GetRelationTuplesInternalServerErrorBody
}

func (o *GetRelationTuplesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /relation-tuples][%d] getRelationTuplesInternalServerError  %+v", 500, o.Payload)
}
func (o *GetRelationTuplesInternalServerError) GetPayload() *GetRelationTuplesInternalServerErrorBody {
	return o.Payload
}

func (o *GetRelationTuplesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetRelationTuplesInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetRelationTuplesInternalServerErrorBody get relation tuples internal server error body
swagger:model GetRelationTuplesInternalServerErrorBody
*/
type GetRelationTuplesInternalServerErrorBody struct {

	// code
	Code int64 `json:"code,omitempty"`

	// details
	Details []interface{} `json:"details"`

	// message
	Message string `json:"message,omitempty"`

	// reason
	Reason string `json:"reason,omitempty"`

	// request
	Request string `json:"request,omitempty"`

	// status
	Status string `json:"status,omitempty"`
}

// Validate validates this get relation tuples internal server error body
func (o *GetRelationTuplesInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get relation tuples internal server error body based on context it is used
func (o *GetRelationTuplesInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetRelationTuplesInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetRelationTuplesInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res GetRelationTuplesInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetRelationTuplesNotFoundBody get relation tuples not found body
swagger:model GetRelationTuplesNotFoundBody
*/
type GetRelationTuplesNotFoundBody struct {

	// code
	Code int64 `json:"code,omitempty"`

	// details
	Details []interface{} `json:"details"`

	// message
	Message string `json:"message,omitempty"`

	// reason
	Reason string `json:"reason,omitempty"`

	// request
	Request string `json:"request,omitempty"`

	// status
	Status string `json:"status,omitempty"`
}

// Validate validates this get relation tuples not found body
func (o *GetRelationTuplesNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get relation tuples not found body based on context it is used
func (o *GetRelationTuplesNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetRelationTuplesNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetRelationTuplesNotFoundBody) UnmarshalBinary(b []byte) error {
	var res GetRelationTuplesNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}