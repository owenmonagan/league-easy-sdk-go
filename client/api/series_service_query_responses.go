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

// SeriesServiceQueryReader is a Reader for the SeriesServiceQuery structure.
type SeriesServiceQueryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SeriesServiceQueryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSeriesServiceQueryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSeriesServiceQueryDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSeriesServiceQueryOK creates a SeriesServiceQueryOK with default headers values
func NewSeriesServiceQueryOK() *SeriesServiceQueryOK {
	return &SeriesServiceQueryOK{}
}

/*SeriesServiceQueryOK handles this case with default header values.

A successful response.
*/
type SeriesServiceQueryOK struct {
	Payload *models.APISeriesResponseList
}

func (o *SeriesServiceQueryOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/series/query][%d] seriesServiceQueryOK  %+v", 200, o.Payload)
}

func (o *SeriesServiceQueryOK) GetPayload() *models.APISeriesResponseList {
	return o.Payload
}

func (o *SeriesServiceQueryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APISeriesResponseList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSeriesServiceQueryDefault creates a SeriesServiceQueryDefault with default headers values
func NewSeriesServiceQueryDefault(code int) *SeriesServiceQueryDefault {
	return &SeriesServiceQueryDefault{
		_statusCode: code,
	}
}

/*SeriesServiceQueryDefault handles this case with default header values.

An unexpected error response.
*/
type SeriesServiceQueryDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// Code gets the status code for the series service query default response
func (o *SeriesServiceQueryDefault) Code() int {
	return o._statusCode
}

func (o *SeriesServiceQueryDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/series/query][%d] SeriesService_Query default  %+v", o._statusCode, o.Payload)
}

func (o *SeriesServiceQueryDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *SeriesServiceQueryDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}