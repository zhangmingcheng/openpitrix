// Code generated by go-swagger; DO NOT EDIT.

package category_manager

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

// NewModifyCategoryParams creates a new ModifyCategoryParams object
// with the default values initialized.
func NewModifyCategoryParams() *ModifyCategoryParams {
	var ()
	return &ModifyCategoryParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewModifyCategoryParamsWithTimeout creates a new ModifyCategoryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewModifyCategoryParamsWithTimeout(timeout time.Duration) *ModifyCategoryParams {
	var ()
	return &ModifyCategoryParams{

		timeout: timeout,
	}
}

// NewModifyCategoryParamsWithContext creates a new ModifyCategoryParams object
// with the default values initialized, and the ability to set a context for a request
func NewModifyCategoryParamsWithContext(ctx context.Context) *ModifyCategoryParams {
	var ()
	return &ModifyCategoryParams{

		Context: ctx,
	}
}

// NewModifyCategoryParamsWithHTTPClient creates a new ModifyCategoryParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewModifyCategoryParamsWithHTTPClient(client *http.Client) *ModifyCategoryParams {
	var ()
	return &ModifyCategoryParams{
		HTTPClient: client,
	}
}

/*ModifyCategoryParams contains all the parameters to send to the API endpoint
for the modify category operation typically these are written to a http.Request
*/
type ModifyCategoryParams struct {

	/*Body*/
	Body *models.OpenpitrixModifyCategoryRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the modify category params
func (o *ModifyCategoryParams) WithTimeout(timeout time.Duration) *ModifyCategoryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the modify category params
func (o *ModifyCategoryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the modify category params
func (o *ModifyCategoryParams) WithContext(ctx context.Context) *ModifyCategoryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the modify category params
func (o *ModifyCategoryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the modify category params
func (o *ModifyCategoryParams) WithHTTPClient(client *http.Client) *ModifyCategoryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the modify category params
func (o *ModifyCategoryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the modify category params
func (o *ModifyCategoryParams) WithBody(body *models.OpenpitrixModifyCategoryRequest) *ModifyCategoryParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the modify category params
func (o *ModifyCategoryParams) SetBody(body *models.OpenpitrixModifyCategoryRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ModifyCategoryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
