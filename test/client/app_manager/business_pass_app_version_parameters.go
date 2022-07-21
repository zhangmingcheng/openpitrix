// Code generated by go-swagger; DO NOT EDIT.

package app_manager

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

// NewBusinessPassAppVersionParams creates a new BusinessPassAppVersionParams object
// with the default values initialized.
func NewBusinessPassAppVersionParams() *BusinessPassAppVersionParams {
	var ()
	return &BusinessPassAppVersionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewBusinessPassAppVersionParamsWithTimeout creates a new BusinessPassAppVersionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewBusinessPassAppVersionParamsWithTimeout(timeout time.Duration) *BusinessPassAppVersionParams {
	var ()
	return &BusinessPassAppVersionParams{

		timeout: timeout,
	}
}

// NewBusinessPassAppVersionParamsWithContext creates a new BusinessPassAppVersionParams object
// with the default values initialized, and the ability to set a context for a request
func NewBusinessPassAppVersionParamsWithContext(ctx context.Context) *BusinessPassAppVersionParams {
	var ()
	return &BusinessPassAppVersionParams{

		Context: ctx,
	}
}

// NewBusinessPassAppVersionParamsWithHTTPClient creates a new BusinessPassAppVersionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewBusinessPassAppVersionParamsWithHTTPClient(client *http.Client) *BusinessPassAppVersionParams {
	var ()
	return &BusinessPassAppVersionParams{
		HTTPClient: client,
	}
}

/*BusinessPassAppVersionParams contains all the parameters to send to the API endpoint
for the business pass app version operation typically these are written to a http.Request
*/
type BusinessPassAppVersionParams struct {

	/*Body*/
	Body *models.OpenpitrixPassAppVersionRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the business pass app version params
func (o *BusinessPassAppVersionParams) WithTimeout(timeout time.Duration) *BusinessPassAppVersionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the business pass app version params
func (o *BusinessPassAppVersionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the business pass app version params
func (o *BusinessPassAppVersionParams) WithContext(ctx context.Context) *BusinessPassAppVersionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the business pass app version params
func (o *BusinessPassAppVersionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the business pass app version params
func (o *BusinessPassAppVersionParams) WithHTTPClient(client *http.Client) *BusinessPassAppVersionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the business pass app version params
func (o *BusinessPassAppVersionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the business pass app version params
func (o *BusinessPassAppVersionParams) WithBody(body *models.OpenpitrixPassAppVersionRequest) *BusinessPassAppVersionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the business pass app version params
func (o *BusinessPassAppVersionParams) SetBody(body *models.OpenpitrixPassAppVersionRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *BusinessPassAppVersionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
