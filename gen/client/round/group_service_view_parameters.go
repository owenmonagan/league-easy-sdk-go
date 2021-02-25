// Code generated by go-swagger; DO NOT EDIT.

package round

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

// NewGroupServiceViewParams creates a new GroupServiceViewParams object
// with the default values initialized.
func NewGroupServiceViewParams() *GroupServiceViewParams {
	var ()
	return &GroupServiceViewParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGroupServiceViewParamsWithTimeout creates a new GroupServiceViewParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGroupServiceViewParamsWithTimeout(timeout time.Duration) *GroupServiceViewParams {
	var ()
	return &GroupServiceViewParams{

		timeout: timeout,
	}
}

// NewGroupServiceViewParamsWithContext creates a new GroupServiceViewParams object
// with the default values initialized, and the ability to set a context for a request
func NewGroupServiceViewParamsWithContext(ctx context.Context) *GroupServiceViewParams {
	var ()
	return &GroupServiceViewParams{

		Context: ctx,
	}
}

// NewGroupServiceViewParamsWithHTTPClient creates a new GroupServiceViewParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGroupServiceViewParamsWithHTTPClient(client *http.Client) *GroupServiceViewParams {
	var ()
	return &GroupServiceViewParams{
		HTTPClient: client,
	}
}

/*GroupServiceViewParams contains all the parameters to send to the API endpoint
for the group service view operation typically these are written to a http.Request
*/
type GroupServiceViewParams struct {

	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the group service view params
func (o *GroupServiceViewParams) WithTimeout(timeout time.Duration) *GroupServiceViewParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the group service view params
func (o *GroupServiceViewParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the group service view params
func (o *GroupServiceViewParams) WithContext(ctx context.Context) *GroupServiceViewParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the group service view params
func (o *GroupServiceViewParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the group service view params
func (o *GroupServiceViewParams) WithHTTPClient(client *http.Client) *GroupServiceViewParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the group service view params
func (o *GroupServiceViewParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the group service view params
func (o *GroupServiceViewParams) WithID(id string) *GroupServiceViewParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the group service view params
func (o *GroupServiceViewParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GroupServiceViewParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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