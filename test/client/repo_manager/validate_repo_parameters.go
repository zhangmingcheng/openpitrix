// Code generated by go-swagger; DO NOT EDIT.

package repo_manager

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewValidateRepoParams creates a new ValidateRepoParams object
// with the default values initialized.
func NewValidateRepoParams() *ValidateRepoParams {
	var ()
	return &ValidateRepoParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewValidateRepoParamsWithTimeout creates a new ValidateRepoParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewValidateRepoParamsWithTimeout(timeout time.Duration) *ValidateRepoParams {
	var ()
	return &ValidateRepoParams{

		timeout: timeout,
	}
}

// NewValidateRepoParamsWithContext creates a new ValidateRepoParams object
// with the default values initialized, and the ability to set a context for a request
func NewValidateRepoParamsWithContext(ctx context.Context) *ValidateRepoParams {
	var ()
	return &ValidateRepoParams{

		Context: ctx,
	}
}

// NewValidateRepoParamsWithHTTPClient creates a new ValidateRepoParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewValidateRepoParamsWithHTTPClient(client *http.Client) *ValidateRepoParams {
	var ()
	return &ValidateRepoParams{
		HTTPClient: client,
	}
}

/*ValidateRepoParams contains all the parameters to send to the API endpoint
for the validate repo operation typically these are written to a http.Request
*/
type ValidateRepoParams struct {

	/*Credential
	  required, credential of visiting the repository.

	*/
	Credential *string
	/*Type
	  required, type of repository.

	*/
	Type *string
	/*URL
	  required, url of visiting the repository.

	*/
	URL *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the validate repo params
func (o *ValidateRepoParams) WithTimeout(timeout time.Duration) *ValidateRepoParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the validate repo params
func (o *ValidateRepoParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the validate repo params
func (o *ValidateRepoParams) WithContext(ctx context.Context) *ValidateRepoParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the validate repo params
func (o *ValidateRepoParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the validate repo params
func (o *ValidateRepoParams) WithHTTPClient(client *http.Client) *ValidateRepoParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the validate repo params
func (o *ValidateRepoParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCredential adds the credential to the validate repo params
func (o *ValidateRepoParams) WithCredential(credential *string) *ValidateRepoParams {
	o.SetCredential(credential)
	return o
}

// SetCredential adds the credential to the validate repo params
func (o *ValidateRepoParams) SetCredential(credential *string) {
	o.Credential = credential
}

// WithType adds the typeVar to the validate repo params
func (o *ValidateRepoParams) WithType(typeVar *string) *ValidateRepoParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the validate repo params
func (o *ValidateRepoParams) SetType(typeVar *string) {
	o.Type = typeVar
}

// WithURL adds the url to the validate repo params
func (o *ValidateRepoParams) WithURL(url *string) *ValidateRepoParams {
	o.SetURL(url)
	return o
}

// SetURL adds the url to the validate repo params
func (o *ValidateRepoParams) SetURL(url *string) {
	o.URL = url
}

// WriteToRequest writes these params to a swagger request
func (o *ValidateRepoParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Credential != nil {

		// query param credential
		var qrCredential string
		if o.Credential != nil {
			qrCredential = *o.Credential
		}
		qCredential := qrCredential
		if qCredential != "" {
			if err := r.SetQueryParam("credential", qCredential); err != nil {
				return err
			}
		}

	}

	if o.Type != nil {

		// query param type
		var qrType string
		if o.Type != nil {
			qrType = *o.Type
		}
		qType := qrType
		if qType != "" {
			if err := r.SetQueryParam("type", qType); err != nil {
				return err
			}
		}

	}

	if o.URL != nil {

		// query param url
		var qrURL string
		if o.URL != nil {
			qrURL = *o.URL
		}
		qURL := qrURL
		if qURL != "" {
			if err := r.SetQueryParam("url", qURL); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}