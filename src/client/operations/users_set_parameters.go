// Code generated by go-swagger; DO NOT EDIT.

package operations

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

	"com.azure.iot/iotcentral/iotcgo/models"
)

// NewUsersSetParams creates a new UsersSetParams object
// with the default values initialized.
func NewUsersSetParams() *UsersSetParams {
	var ()
	return &UsersSetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUsersSetParamsWithTimeout creates a new UsersSetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUsersSetParamsWithTimeout(timeout time.Duration) *UsersSetParams {
	var ()
	return &UsersSetParams{

		timeout: timeout,
	}
}

// NewUsersSetParamsWithContext creates a new UsersSetParams object
// with the default values initialized, and the ability to set a context for a request
func NewUsersSetParamsWithContext(ctx context.Context) *UsersSetParams {
	var ()
	return &UsersSetParams{

		Context: ctx,
	}
}

// NewUsersSetParamsWithHTTPClient creates a new UsersSetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUsersSetParamsWithHTTPClient(client *http.Client) *UsersSetParams {
	var ()
	return &UsersSetParams{
		HTTPClient: client,
	}
}

/*UsersSetParams contains all the parameters to send to the API endpoint
for the users set operation typically these are written to a http.Request
*/
type UsersSetParams struct {

	/*Body
	  User body.

	*/
	Body *models.User
	/*UserID
	  Unique ID of the user.

	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the users set params
func (o *UsersSetParams) WithTimeout(timeout time.Duration) *UsersSetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the users set params
func (o *UsersSetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the users set params
func (o *UsersSetParams) WithContext(ctx context.Context) *UsersSetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the users set params
func (o *UsersSetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the users set params
func (o *UsersSetParams) WithHTTPClient(client *http.Client) *UsersSetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the users set params
func (o *UsersSetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the users set params
func (o *UsersSetParams) WithBody(body *models.User) *UsersSetParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the users set params
func (o *UsersSetParams) SetBody(body *models.User) {
	o.Body = body
}

// WithUserID adds the userID to the users set params
func (o *UsersSetParams) WithUserID(userID string) *UsersSetParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the users set params
func (o *UsersSetParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *UsersSetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param user_id
	if err := r.SetPathParam("user_id", o.UserID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
