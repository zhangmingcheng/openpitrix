// Code generated by go-swagger; DO NOT EDIT.

package runtime_manager

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

	"openpitrix.io/openpitrix/test/models"
)

// NewCreateDebugRuntimeParams creates a new CreateDebugRuntimeParams object
// with the default values initialized.
func NewCreateDebugRuntimeParams() *CreateDebugRuntimeParams {
	var ()
	return &CreateDebugRuntimeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateDebugRuntimeParamsWithTimeout creates a new CreateDebugRuntimeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateDebugRuntimeParamsWithTimeout(timeout time.Duration) *CreateDebugRuntimeParams {
	var ()
	return &CreateDebugRuntimeParams{

		timeout: timeout,
	}
}

// NewCreateDebugRuntimeParamsWithContext creates a new CreateDebugRuntimeParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateDebugRuntimeParamsWithContext(ctx context.Context) *CreateDebugRuntimeParams {
	var ()
	return &CreateDebugRuntimeParams{

		Context: ctx,
	}
}

// NewCreateDebugRuntimeParamsWithHTTPClient creates a new CreateDebugRuntimeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateDebugRuntimeParamsWithHTTPClient(client *http.Client) *CreateDebugRuntimeParams {
	var ()
	return &CreateDebugRuntimeParams{
		HTTPClient: client,
	}
}

/*CreateDebugRuntimeParams contains all the parameters to send to the API endpoint
for the create debug runtime operation typically these are written to a http.Request
*/
type CreateDebugRuntimeParams struct {

	/*Body*/
	Body *models.OpenpitrixCreateRuntimeRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create debug runtime params
func (o *CreateDebugRuntimeParams) WithTimeout(timeout time.Duration) *CreateDebugRuntimeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create debug runtime params
func (o *CreateDebugRuntimeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create debug runtime params
func (o *CreateDebugRuntimeParams) WithContext(ctx context.Context) *CreateDebugRuntimeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create debug runtime params
func (o *CreateDebugRuntimeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create debug runtime params
func (o *CreateDebugRuntimeParams) WithHTTPClient(client *http.Client) *CreateDebugRuntimeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create debug runtime params
func (o *CreateDebugRuntimeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create debug runtime params
func (o *CreateDebugRuntimeParams) WithBody(body *models.OpenpitrixCreateRuntimeRequest) *CreateDebugRuntimeParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create debug runtime params
func (o *CreateDebugRuntimeParams) SetBody(body *models.OpenpitrixCreateRuntimeRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateDebugRuntimeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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