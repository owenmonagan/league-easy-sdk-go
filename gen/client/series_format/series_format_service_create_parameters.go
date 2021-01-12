// Code generated by go-swagger; DO NOT EDIT.

package series_format

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/owenmonagan/toureasy-sdk-go/gen/models"
)

// NewSeriesFormatServiceCreateParams creates a new SeriesFormatServiceCreateParams object
// with the default values initialized.
func NewSeriesFormatServiceCreateParams() *SeriesFormatServiceCreateParams {
	var ()
	return &SeriesFormatServiceCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSeriesFormatServiceCreateParamsWithTimeout creates a new SeriesFormatServiceCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSeriesFormatServiceCreateParamsWithTimeout(timeout time.Duration) *SeriesFormatServiceCreateParams {
	var ()
	return &SeriesFormatServiceCreateParams{

		timeout: timeout,
	}
}

// NewSeriesFormatServiceCreateParamsWithContext creates a new SeriesFormatServiceCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewSeriesFormatServiceCreateParamsWithContext(ctx context.Context) *SeriesFormatServiceCreateParams {
	var ()
	return &SeriesFormatServiceCreateParams{

		Context: ctx,
	}
}

// NewSeriesFormatServiceCreateParamsWithHTTPClient creates a new SeriesFormatServiceCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSeriesFormatServiceCreateParamsWithHTTPClient(client *http.Client) *SeriesFormatServiceCreateParams {
	var ()
	return &SeriesFormatServiceCreateParams{
		HTTPClient: client,
	}
}

/*SeriesFormatServiceCreateParams contains all the parameters to send to the API endpoint
for the series format service create operation typically these are written to a http.Request
*/
type SeriesFormatServiceCreateParams struct {

	/*Body*/
	Body *models.APISeriesFormatCreateRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the series format service create params
func (o *SeriesFormatServiceCreateParams) WithTimeout(timeout time.Duration) *SeriesFormatServiceCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the series format service create params
func (o *SeriesFormatServiceCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the series format service create params
func (o *SeriesFormatServiceCreateParams) WithContext(ctx context.Context) *SeriesFormatServiceCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the series format service create params
func (o *SeriesFormatServiceCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the series format service create params
func (o *SeriesFormatServiceCreateParams) WithHTTPClient(client *http.Client) *SeriesFormatServiceCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the series format service create params
func (o *SeriesFormatServiceCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the series format service create params
func (o *SeriesFormatServiceCreateParams) WithBody(body *models.APISeriesFormatCreateRequest) *SeriesFormatServiceCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the series format service create params
func (o *SeriesFormatServiceCreateParams) SetBody(body *models.APISeriesFormatCreateRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *SeriesFormatServiceCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}