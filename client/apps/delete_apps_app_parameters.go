// Code generated by go-swagger; DO NOT EDIT.

package apps

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

// NewDeleteAppsAppParams creates a new DeleteAppsAppParams object
// with the default values initialized.
func NewDeleteAppsAppParams() *DeleteAppsAppParams {
	var ()
	return &DeleteAppsAppParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAppsAppParamsWithTimeout creates a new DeleteAppsAppParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteAppsAppParamsWithTimeout(timeout time.Duration) *DeleteAppsAppParams {
	var ()
	return &DeleteAppsAppParams{

		timeout: timeout,
	}
}

// NewDeleteAppsAppParamsWithContext creates a new DeleteAppsAppParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteAppsAppParamsWithContext(ctx context.Context) *DeleteAppsAppParams {
	var ()
	return &DeleteAppsAppParams{

		Context: ctx,
	}
}

// NewDeleteAppsAppParamsWithHTTPClient creates a new DeleteAppsAppParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteAppsAppParamsWithHTTPClient(client *http.Client) *DeleteAppsAppParams {
	var ()
	return &DeleteAppsAppParams{
		HTTPClient: client,
	}
}

/*DeleteAppsAppParams contains all the parameters to send to the API endpoint
for the delete apps app operation typically these are written to a http.Request
*/
type DeleteAppsAppParams struct {

	/*App
	  Name of the app.

	*/
	App string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete apps app params
func (o *DeleteAppsAppParams) WithTimeout(timeout time.Duration) *DeleteAppsAppParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete apps app params
func (o *DeleteAppsAppParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete apps app params
func (o *DeleteAppsAppParams) WithContext(ctx context.Context) *DeleteAppsAppParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete apps app params
func (o *DeleteAppsAppParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete apps app params
func (o *DeleteAppsAppParams) WithHTTPClient(client *http.Client) *DeleteAppsAppParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete apps app params
func (o *DeleteAppsAppParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithApp adds the app to the delete apps app params
func (o *DeleteAppsAppParams) WithApp(app string) *DeleteAppsAppParams {
	o.SetApp(app)
	return o
}

// SetApp adds the app to the delete apps app params
func (o *DeleteAppsAppParams) SetApp(app string) {
	o.App = app
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAppsAppParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param app
	if err := r.SetPathParam("app", o.App); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
