// Code generated by go-swagger; DO NOT EDIT.

package series_format

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/owenmonagan/toureasy-sdk-go/gen/models"
)

// FormatSeriesServiceDeleteReader is a Reader for the FormatSeriesServiceDelete structure.
type FormatSeriesServiceDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FormatSeriesServiceDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFormatSeriesServiceDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFormatSeriesServiceDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFormatSeriesServiceDeleteOK creates a FormatSeriesServiceDeleteOK with default headers values
func NewFormatSeriesServiceDeleteOK() *FormatSeriesServiceDeleteOK {
	return &FormatSeriesServiceDeleteOK{}
}

/*FormatSeriesServiceDeleteOK handles this case with default header values.

A successful response.
*/
type FormatSeriesServiceDeleteOK struct {
	Payload models.APIEmpty
}

func (o *FormatSeriesServiceDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/format/series/{id}][%d] formatSeriesServiceDeleteOK  %+v", 200, o.Payload)
}

func (o *FormatSeriesServiceDeleteOK) GetPayload() models.APIEmpty {
	return o.Payload
}

func (o *FormatSeriesServiceDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFormatSeriesServiceDeleteDefault creates a FormatSeriesServiceDeleteDefault with default headers values
func NewFormatSeriesServiceDeleteDefault(code int) *FormatSeriesServiceDeleteDefault {
	return &FormatSeriesServiceDeleteDefault{
		_statusCode: code,
	}
}

/*FormatSeriesServiceDeleteDefault handles this case with default header values.

An unexpected error response.
*/
type FormatSeriesServiceDeleteDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// Code gets the status code for the format series service delete default response
func (o *FormatSeriesServiceDeleteDefault) Code() int {
	return o._statusCode
}

func (o *FormatSeriesServiceDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/format/series/{id}][%d] FormatSeriesService_Delete default  %+v", o._statusCode, o.Payload)
}

func (o *FormatSeriesServiceDeleteDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *FormatSeriesServiceDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
