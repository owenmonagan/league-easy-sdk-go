// Code generated by go-swagger; DO NOT EDIT.

package round_robin

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

// NewStageRoundRobinServiceViewParams creates a new StageRoundRobinServiceViewParams object
// with the default values initialized.
func NewStageRoundRobinServiceViewParams() *StageRoundRobinServiceViewParams {
	var ()
	return &StageRoundRobinServiceViewParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewStageRoundRobinServiceViewParamsWithTimeout creates a new StageRoundRobinServiceViewParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStageRoundRobinServiceViewParamsWithTimeout(timeout time.Duration) *StageRoundRobinServiceViewParams {
	var ()
	return &StageRoundRobinServiceViewParams{

		timeout: timeout,
	}
}

// NewStageRoundRobinServiceViewParamsWithContext creates a new StageRoundRobinServiceViewParams object
// with the default values initialized, and the ability to set a context for a request
func NewStageRoundRobinServiceViewParamsWithContext(ctx context.Context) *StageRoundRobinServiceViewParams {
	var ()
	return &StageRoundRobinServiceViewParams{

		Context: ctx,
	}
}

// NewStageRoundRobinServiceViewParamsWithHTTPClient creates a new StageRoundRobinServiceViewParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStageRoundRobinServiceViewParamsWithHTTPClient(client *http.Client) *StageRoundRobinServiceViewParams {
	var ()
	return &StageRoundRobinServiceViewParams{
		HTTPClient: client,
	}
}

/*StageRoundRobinServiceViewParams contains all the parameters to send to the API endpoint
for the stage round robin service view operation typically these are written to a http.Request
*/
type StageRoundRobinServiceViewParams struct {

	/*UUID*/
	UUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the stage round robin service view params
func (o *StageRoundRobinServiceViewParams) WithTimeout(timeout time.Duration) *StageRoundRobinServiceViewParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the stage round robin service view params
func (o *StageRoundRobinServiceViewParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the stage round robin service view params
func (o *StageRoundRobinServiceViewParams) WithContext(ctx context.Context) *StageRoundRobinServiceViewParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the stage round robin service view params
func (o *StageRoundRobinServiceViewParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the stage round robin service view params
func (o *StageRoundRobinServiceViewParams) WithHTTPClient(client *http.Client) *StageRoundRobinServiceViewParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the stage round robin service view params
func (o *StageRoundRobinServiceViewParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUUID adds the uuid to the stage round robin service view params
func (o *StageRoundRobinServiceViewParams) WithUUID(uuid string) *StageRoundRobinServiceViewParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the stage round robin service view params
func (o *StageRoundRobinServiceViewParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *StageRoundRobinServiceViewParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param uuid
	if err := r.SetPathParam("uuid", o.UUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
