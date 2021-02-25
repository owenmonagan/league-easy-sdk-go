// Code generated by go-swagger; DO NOT EDIT.

package api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/owenmonagan/toureasy-sdk-go/models"
)

// StageRoundRobinServiceUpdateReader is a Reader for the StageRoundRobinServiceUpdate structure.
type StageRoundRobinServiceUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StageRoundRobinServiceUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStageRoundRobinServiceUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStageRoundRobinServiceUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStageRoundRobinServiceUpdateOK creates a StageRoundRobinServiceUpdateOK with default headers values
func NewStageRoundRobinServiceUpdateOK() *StageRoundRobinServiceUpdateOK {
	return &StageRoundRobinServiceUpdateOK{}
}

/*StageRoundRobinServiceUpdateOK handles this case with default header values.

A successful response.
*/
type StageRoundRobinServiceUpdateOK struct {
	Payload *models.APIStageRoundRobinResponse
}

func (o *StageRoundRobinServiceUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/stage/rr/{id}][%d] stageRoundRobinServiceUpdateOK  %+v", 200, o.Payload)
}

func (o *StageRoundRobinServiceUpdateOK) GetPayload() *models.APIStageRoundRobinResponse {
	return o.Payload
}

func (o *StageRoundRobinServiceUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIStageRoundRobinResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStageRoundRobinServiceUpdateDefault creates a StageRoundRobinServiceUpdateDefault with default headers values
func NewStageRoundRobinServiceUpdateDefault(code int) *StageRoundRobinServiceUpdateDefault {
	return &StageRoundRobinServiceUpdateDefault{
		_statusCode: code,
	}
}

/*StageRoundRobinServiceUpdateDefault handles this case with default header values.

An unexpected error response.
*/
type StageRoundRobinServiceUpdateDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// Code gets the status code for the stage round robin service update default response
func (o *StageRoundRobinServiceUpdateDefault) Code() int {
	return o._statusCode
}

func (o *StageRoundRobinServiceUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/stage/rr/{id}][%d] StageRoundRobinService_Update default  %+v", o._statusCode, o.Payload)
}

func (o *StageRoundRobinServiceUpdateDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *StageRoundRobinServiceUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}