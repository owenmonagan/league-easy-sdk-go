// Code generated by go-swagger; DO NOT EDIT.

package entry

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

// NewStageSingleEliminationServiceViewParams creates a new StageSingleEliminationServiceViewParams object
// with the default values initialized.
func NewStageSingleEliminationServiceViewParams() *StageSingleEliminationServiceViewParams {
	var ()
	return &StageSingleEliminationServiceViewParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewStageSingleEliminationServiceViewParamsWithTimeout creates a new StageSingleEliminationServiceViewParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStageSingleEliminationServiceViewParamsWithTimeout(timeout time.Duration) *StageSingleEliminationServiceViewParams {
	var ()
	return &StageSingleEliminationServiceViewParams{

		timeout: timeout,
	}
}

// NewStageSingleEliminationServiceViewParamsWithContext creates a new StageSingleEliminationServiceViewParams object
// with the default values initialized, and the ability to set a context for a request
func NewStageSingleEliminationServiceViewParamsWithContext(ctx context.Context) *StageSingleEliminationServiceViewParams {
	var ()
	return &StageSingleEliminationServiceViewParams{

		Context: ctx,
	}
}

// NewStageSingleEliminationServiceViewParamsWithHTTPClient creates a new StageSingleEliminationServiceViewParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStageSingleEliminationServiceViewParamsWithHTTPClient(client *http.Client) *StageSingleEliminationServiceViewParams {
	var ()
	return &StageSingleEliminationServiceViewParams{
		HTTPClient: client,
	}
}

/*StageSingleEliminationServiceViewParams contains all the parameters to send to the API endpoint
for the stage single elimination service view operation typically these are written to a http.Request
*/
type StageSingleEliminationServiceViewParams struct {

	/*UUID*/
	UUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the stage single elimination service view params
func (o *StageSingleEliminationServiceViewParams) WithTimeout(timeout time.Duration) *StageSingleEliminationServiceViewParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the stage single elimination service view params
func (o *StageSingleEliminationServiceViewParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the stage single elimination service view params
func (o *StageSingleEliminationServiceViewParams) WithContext(ctx context.Context) *StageSingleEliminationServiceViewParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the stage single elimination service view params
func (o *StageSingleEliminationServiceViewParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the stage single elimination service view params
func (o *StageSingleEliminationServiceViewParams) WithHTTPClient(client *http.Client) *StageSingleEliminationServiceViewParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the stage single elimination service view params
func (o *StageSingleEliminationServiceViewParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUUID adds the uuid to the stage single elimination service view params
func (o *StageSingleEliminationServiceViewParams) WithUUID(uuid string) *StageSingleEliminationServiceViewParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the stage single elimination service view params
func (o *StageSingleEliminationServiceViewParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *StageSingleEliminationServiceViewParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
