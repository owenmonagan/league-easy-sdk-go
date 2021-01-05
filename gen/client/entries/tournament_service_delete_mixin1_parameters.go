// Code generated by go-swagger; DO NOT EDIT.

package entries

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

// NewTournamentServiceDeleteMixin1Params creates a new TournamentServiceDeleteMixin1Params object
// with the default values initialized.
func NewTournamentServiceDeleteMixin1Params() *TournamentServiceDeleteMixin1Params {
	var ()
	return &TournamentServiceDeleteMixin1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewTournamentServiceDeleteMixin1ParamsWithTimeout creates a new TournamentServiceDeleteMixin1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewTournamentServiceDeleteMixin1ParamsWithTimeout(timeout time.Duration) *TournamentServiceDeleteMixin1Params {
	var ()
	return &TournamentServiceDeleteMixin1Params{

		timeout: timeout,
	}
}

// NewTournamentServiceDeleteMixin1ParamsWithContext creates a new TournamentServiceDeleteMixin1Params object
// with the default values initialized, and the ability to set a context for a request
func NewTournamentServiceDeleteMixin1ParamsWithContext(ctx context.Context) *TournamentServiceDeleteMixin1Params {
	var ()
	return &TournamentServiceDeleteMixin1Params{

		Context: ctx,
	}
}

// NewTournamentServiceDeleteMixin1ParamsWithHTTPClient creates a new TournamentServiceDeleteMixin1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTournamentServiceDeleteMixin1ParamsWithHTTPClient(client *http.Client) *TournamentServiceDeleteMixin1Params {
	var ()
	return &TournamentServiceDeleteMixin1Params{
		HTTPClient: client,
	}
}

/*TournamentServiceDeleteMixin1Params contains all the parameters to send to the API endpoint
for the tournament service delete mixin1 operation typically these are written to a http.Request
*/
type TournamentServiceDeleteMixin1Params struct {

	/*ExternalID*/
	ExternalID *string
	/*UUID*/
	UUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the tournament service delete mixin1 params
func (o *TournamentServiceDeleteMixin1Params) WithTimeout(timeout time.Duration) *TournamentServiceDeleteMixin1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the tournament service delete mixin1 params
func (o *TournamentServiceDeleteMixin1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the tournament service delete mixin1 params
func (o *TournamentServiceDeleteMixin1Params) WithContext(ctx context.Context) *TournamentServiceDeleteMixin1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the tournament service delete mixin1 params
func (o *TournamentServiceDeleteMixin1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the tournament service delete mixin1 params
func (o *TournamentServiceDeleteMixin1Params) WithHTTPClient(client *http.Client) *TournamentServiceDeleteMixin1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the tournament service delete mixin1 params
func (o *TournamentServiceDeleteMixin1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExternalID adds the externalID to the tournament service delete mixin1 params
func (o *TournamentServiceDeleteMixin1Params) WithExternalID(externalID *string) *TournamentServiceDeleteMixin1Params {
	o.SetExternalID(externalID)
	return o
}

// SetExternalID adds the externalId to the tournament service delete mixin1 params
func (o *TournamentServiceDeleteMixin1Params) SetExternalID(externalID *string) {
	o.ExternalID = externalID
}

// WithUUID adds the uuid to the tournament service delete mixin1 params
func (o *TournamentServiceDeleteMixin1Params) WithUUID(uuid string) *TournamentServiceDeleteMixin1Params {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the tournament service delete mixin1 params
func (o *TournamentServiceDeleteMixin1Params) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *TournamentServiceDeleteMixin1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ExternalID != nil {

		// query param externalId
		var qrExternalID string
		if o.ExternalID != nil {
			qrExternalID = *o.ExternalID
		}
		qExternalID := qrExternalID
		if qExternalID != "" {
			if err := r.SetQueryParam("externalId", qExternalID); err != nil {
				return err
			}
		}

	}

	// path param uuid
	if err := r.SetPathParam("uuid", o.UUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
