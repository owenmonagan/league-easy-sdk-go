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

// SeriesFormatServiceDeleteReader is a Reader for the SeriesFormatServiceDelete structure.
type SeriesFormatServiceDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SeriesFormatServiceDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSeriesFormatServiceDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSeriesFormatServiceDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSeriesFormatServiceDeleteOK creates a SeriesFormatServiceDeleteOK with default headers values
func NewSeriesFormatServiceDeleteOK() *SeriesFormatServiceDeleteOK {
	return &SeriesFormatServiceDeleteOK{}
}

/*SeriesFormatServiceDeleteOK handles this case with default header values.

A successful response.
*/
type SeriesFormatServiceDeleteOK struct {
	Payload models.APISeriesFormatDeleteResponse
}

func (o *SeriesFormatServiceDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/format/series/{uuid}][%d] seriesFormatServiceDeleteOK  %+v", 200, o.Payload)
}

func (o *SeriesFormatServiceDeleteOK) GetPayload() models.APISeriesFormatDeleteResponse {
	return o.Payload
}

func (o *SeriesFormatServiceDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSeriesFormatServiceDeleteDefault creates a SeriesFormatServiceDeleteDefault with default headers values
func NewSeriesFormatServiceDeleteDefault(code int) *SeriesFormatServiceDeleteDefault {
	return &SeriesFormatServiceDeleteDefault{
		_statusCode: code,
	}
}

/*SeriesFormatServiceDeleteDefault handles this case with default header values.

An unexpected error response.
*/
type SeriesFormatServiceDeleteDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// Code gets the status code for the series format service delete default response
func (o *SeriesFormatServiceDeleteDefault) Code() int {
	return o._statusCode
}

func (o *SeriesFormatServiceDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/format/series/{uuid}][%d] SeriesFormatService_Delete default  %+v", o._statusCode, o.Payload)
}

func (o *SeriesFormatServiceDeleteDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *SeriesFormatServiceDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}