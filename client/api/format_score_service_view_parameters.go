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
)

// NewFormatScoreServiceViewParams creates a new FormatScoreServiceViewParams object
// with the default values initialized.
func NewFormatScoreServiceViewParams() *FormatScoreServiceViewParams {
	var ()
	return &FormatScoreServiceViewParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFormatScoreServiceViewParamsWithTimeout creates a new FormatScoreServiceViewParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFormatScoreServiceViewParamsWithTimeout(timeout time.Duration) *FormatScoreServiceViewParams {
	var ()
	return &FormatScoreServiceViewParams{

		timeout: timeout,
	}
}

// NewFormatScoreServiceViewParamsWithContext creates a new FormatScoreServiceViewParams object
// with the default values initialized, and the ability to set a context for a request
func NewFormatScoreServiceViewParamsWithContext(ctx context.Context) *FormatScoreServiceViewParams {
	var ()
	return &FormatScoreServiceViewParams{

		Context: ctx,
	}
}

// NewFormatScoreServiceViewParamsWithHTTPClient creates a new FormatScoreServiceViewParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFormatScoreServiceViewParamsWithHTTPClient(client *http.Client) *FormatScoreServiceViewParams {
	var ()
	return &FormatScoreServiceViewParams{
		HTTPClient: client,
	}
}

/*FormatScoreServiceViewParams contains all the parameters to send to the API endpoint
for the format score service view operation typically these are written to a http.Request
*/
type FormatScoreServiceViewParams struct {

	/*ID*/
	ID string
	/*TournamentID*/
	TournamentID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the format score service view params
func (o *FormatScoreServiceViewParams) WithTimeout(timeout time.Duration) *FormatScoreServiceViewParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the format score service view params
func (o *FormatScoreServiceViewParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the format score service view params
func (o *FormatScoreServiceViewParams) WithContext(ctx context.Context) *FormatScoreServiceViewParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the format score service view params
func (o *FormatScoreServiceViewParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the format score service view params
func (o *FormatScoreServiceViewParams) WithHTTPClient(client *http.Client) *FormatScoreServiceViewParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the format score service view params
func (o *FormatScoreServiceViewParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the format score service view params
func (o *FormatScoreServiceViewParams) WithID(id string) *FormatScoreServiceViewParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the format score service view params
func (o *FormatScoreServiceViewParams) SetID(id string) {
	o.ID = id
}

// WithTournamentID adds the tournamentID to the format score service view params
func (o *FormatScoreServiceViewParams) WithTournamentID(tournamentID *string) *FormatScoreServiceViewParams {
	o.SetTournamentID(tournamentID)
	return o
}

// SetTournamentID adds the tournamentId to the format score service view params
func (o *FormatScoreServiceViewParams) SetTournamentID(tournamentID *string) {
	o.TournamentID = tournamentID
}

// WriteToRequest writes these params to a swagger request
func (o *FormatScoreServiceViewParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if o.TournamentID != nil {

		// query param tournamentId
		var qrTournamentID string
		if o.TournamentID != nil {
			qrTournamentID = *o.TournamentID
		}
		qTournamentID := qrTournamentID
		if qTournamentID != "" {
			if err := r.SetQueryParam("tournamentId", qTournamentID); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
