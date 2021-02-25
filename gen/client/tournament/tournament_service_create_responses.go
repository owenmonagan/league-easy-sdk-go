// Code generated by go-swagger; DO NOT EDIT.

package tournament

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/owenmonagan/toureasy-sdk-go/gen/models"
)

// TournamentServiceCreateReader is a Reader for the TournamentServiceCreate structure.
type TournamentServiceCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TournamentServiceCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTournamentServiceCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewTournamentServiceCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewTournamentServiceCreateOK creates a TournamentServiceCreateOK with default headers values
func NewTournamentServiceCreateOK() *TournamentServiceCreateOK {
	return &TournamentServiceCreateOK{}
}

/*TournamentServiceCreateOK handles this case with default header values.

A successful response.
*/
type TournamentServiceCreateOK struct {
	Payload *models.APITournament
}

func (o *TournamentServiceCreateOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/tournaments][%d] tournamentServiceCreateOK  %+v", 200, o.Payload)
}

func (o *TournamentServiceCreateOK) GetPayload() *models.APITournament {
	return o.Payload
}

func (o *TournamentServiceCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APITournament)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTournamentServiceCreateDefault creates a TournamentServiceCreateDefault with default headers values
func NewTournamentServiceCreateDefault(code int) *TournamentServiceCreateDefault {
	return &TournamentServiceCreateDefault{
		_statusCode: code,
	}
}

/*TournamentServiceCreateDefault handles this case with default header values.

An unexpected error response.
*/
type TournamentServiceCreateDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// Code gets the status code for the tournament service create default response
func (o *TournamentServiceCreateDefault) Code() int {
	return o._statusCode
}

func (o *TournamentServiceCreateDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/tournaments][%d] TournamentService_Create default  %+v", o._statusCode, o.Payload)
}

func (o *TournamentServiceCreateDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *TournamentServiceCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
