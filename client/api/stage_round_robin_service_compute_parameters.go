// Code generated by go-swagger; DO NOT EDIT.

package api

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

	"github.com/owenmonagan/toureasy-sdk-go/models"
)

// NewStageRoundRobinServiceComputeParams creates a new StageRoundRobinServiceComputeParams object
// with the default values initialized.
func NewStageRoundRobinServiceComputeParams() *StageRoundRobinServiceComputeParams {
	var ()
	return &StageRoundRobinServiceComputeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewStageRoundRobinServiceComputeParamsWithTimeout creates a new StageRoundRobinServiceComputeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStageRoundRobinServiceComputeParamsWithTimeout(timeout time.Duration) *StageRoundRobinServiceComputeParams {
	var ()
	return &StageRoundRobinServiceComputeParams{

		timeout: timeout,
	}
}

// NewStageRoundRobinServiceComputeParamsWithContext creates a new StageRoundRobinServiceComputeParams object
// with the default values initialized, and the ability to set a context for a request
func NewStageRoundRobinServiceComputeParamsWithContext(ctx context.Context) *StageRoundRobinServiceComputeParams {
	var ()
	return &StageRoundRobinServiceComputeParams{

		Context: ctx,
	}
}

// NewStageRoundRobinServiceComputeParamsWithHTTPClient creates a new StageRoundRobinServiceComputeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStageRoundRobinServiceComputeParamsWithHTTPClient(client *http.Client) *StageRoundRobinServiceComputeParams {
	var ()
	return &StageRoundRobinServiceComputeParams{
		HTTPClient: client,
	}
}

/*StageRoundRobinServiceComputeParams contains all the parameters to send to the API endpoint
for the stage round robin service compute operation typically these are written to a http.Request
*/
type StageRoundRobinServiceComputeParams struct {

	/*Body*/
	Body *models.APIID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the stage round robin service compute params
func (o *StageRoundRobinServiceComputeParams) WithTimeout(timeout time.Duration) *StageRoundRobinServiceComputeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the stage round robin service compute params
func (o *StageRoundRobinServiceComputeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the stage round robin service compute params
func (o *StageRoundRobinServiceComputeParams) WithContext(ctx context.Context) *StageRoundRobinServiceComputeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the stage round robin service compute params
func (o *StageRoundRobinServiceComputeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the stage round robin service compute params
func (o *StageRoundRobinServiceComputeParams) WithHTTPClient(client *http.Client) *StageRoundRobinServiceComputeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the stage round robin service compute params
func (o *StageRoundRobinServiceComputeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the stage round robin service compute params
func (o *StageRoundRobinServiceComputeParams) WithBody(body *models.APIID) *StageRoundRobinServiceComputeParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the stage round robin service compute params
func (o *StageRoundRobinServiceComputeParams) SetBody(body *models.APIID) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *StageRoundRobinServiceComputeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
