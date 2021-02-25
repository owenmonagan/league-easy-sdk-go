// Code generated by go-swagger; DO NOT EDIT.

package format_score

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

// NewFormatScoreServiceCreateParams creates a new FormatScoreServiceCreateParams object
// with the default values initialized.
func NewFormatScoreServiceCreateParams() *FormatScoreServiceCreateParams {
	var ()
	return &FormatScoreServiceCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFormatScoreServiceCreateParamsWithTimeout creates a new FormatScoreServiceCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFormatScoreServiceCreateParamsWithTimeout(timeout time.Duration) *FormatScoreServiceCreateParams {
	var ()
	return &FormatScoreServiceCreateParams{

		timeout: timeout,
	}
}

// NewFormatScoreServiceCreateParamsWithContext creates a new FormatScoreServiceCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewFormatScoreServiceCreateParamsWithContext(ctx context.Context) *FormatScoreServiceCreateParams {
	var ()
	return &FormatScoreServiceCreateParams{

		Context: ctx,
	}
}

// NewFormatScoreServiceCreateParamsWithHTTPClient creates a new FormatScoreServiceCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFormatScoreServiceCreateParamsWithHTTPClient(client *http.Client) *FormatScoreServiceCreateParams {
	var ()
	return &FormatScoreServiceCreateParams{
		HTTPClient: client,
	}
}

/*FormatScoreServiceCreateParams contains all the parameters to send to the API endpoint
for the format score service create operation typically these are written to a http.Request
*/
type FormatScoreServiceCreateParams struct {

	/*Body*/
	Body *models.APIFormatScoreRequestBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the format score service create params
func (o *FormatScoreServiceCreateParams) WithTimeout(timeout time.Duration) *FormatScoreServiceCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the format score service create params
func (o *FormatScoreServiceCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the format score service create params
func (o *FormatScoreServiceCreateParams) WithContext(ctx context.Context) *FormatScoreServiceCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the format score service create params
func (o *FormatScoreServiceCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the format score service create params
func (o *FormatScoreServiceCreateParams) WithHTTPClient(client *http.Client) *FormatScoreServiceCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the format score service create params
func (o *FormatScoreServiceCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the format score service create params
func (o *FormatScoreServiceCreateParams) WithBody(body *models.APIFormatScoreRequestBody) *FormatScoreServiceCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the format score service create params
func (o *FormatScoreServiceCreateParams) SetBody(body *models.APIFormatScoreRequestBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *FormatScoreServiceCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
