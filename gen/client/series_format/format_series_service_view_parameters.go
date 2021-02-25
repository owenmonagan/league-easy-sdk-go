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
)

// NewFormatSeriesServiceViewParams creates a new FormatSeriesServiceViewParams object
// with the default values initialized.
func NewFormatSeriesServiceViewParams() *FormatSeriesServiceViewParams {
	var ()
	return &FormatSeriesServiceViewParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFormatSeriesServiceViewParamsWithTimeout creates a new FormatSeriesServiceViewParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFormatSeriesServiceViewParamsWithTimeout(timeout time.Duration) *FormatSeriesServiceViewParams {
	var ()
	return &FormatSeriesServiceViewParams{

		timeout: timeout,
	}
}

// NewFormatSeriesServiceViewParamsWithContext creates a new FormatSeriesServiceViewParams object
// with the default values initialized, and the ability to set a context for a request
func NewFormatSeriesServiceViewParamsWithContext(ctx context.Context) *FormatSeriesServiceViewParams {
	var ()
	return &FormatSeriesServiceViewParams{

		Context: ctx,
	}
}

// NewFormatSeriesServiceViewParamsWithHTTPClient creates a new FormatSeriesServiceViewParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFormatSeriesServiceViewParamsWithHTTPClient(client *http.Client) *FormatSeriesServiceViewParams {
	var ()
	return &FormatSeriesServiceViewParams{
		HTTPClient: client,
	}
}

/*FormatSeriesServiceViewParams contains all the parameters to send to the API endpoint
for the format series service view operation typically these are written to a http.Request
*/
type FormatSeriesServiceViewParams struct {

	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the format series service view params
func (o *FormatSeriesServiceViewParams) WithTimeout(timeout time.Duration) *FormatSeriesServiceViewParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the format series service view params
func (o *FormatSeriesServiceViewParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the format series service view params
func (o *FormatSeriesServiceViewParams) WithContext(ctx context.Context) *FormatSeriesServiceViewParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the format series service view params
func (o *FormatSeriesServiceViewParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the format series service view params
func (o *FormatSeriesServiceViewParams) WithHTTPClient(client *http.Client) *FormatSeriesServiceViewParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the format series service view params
func (o *FormatSeriesServiceViewParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the format series service view params
func (o *FormatSeriesServiceViewParams) WithID(id string) *FormatSeriesServiceViewParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the format series service view params
func (o *FormatSeriesServiceViewParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *FormatSeriesServiceViewParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}