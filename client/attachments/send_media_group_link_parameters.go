// Code generated by go-swagger; DO NOT EDIT.

package attachments

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

	"github.com/olebedev/go-tgbot/models"
)

// NewSendMediaGroupLinkParams creates a new SendMediaGroupLinkParams object
// with the default values initialized.
func NewSendMediaGroupLinkParams() *SendMediaGroupLinkParams {
	var ()
	return &SendMediaGroupLinkParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSendMediaGroupLinkParamsWithTimeout creates a new SendMediaGroupLinkParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSendMediaGroupLinkParamsWithTimeout(timeout time.Duration) *SendMediaGroupLinkParams {
	var ()
	return &SendMediaGroupLinkParams{

		timeout: timeout,
	}
}

// NewSendMediaGroupLinkParamsWithContext creates a new SendMediaGroupLinkParams object
// with the default values initialized, and the ability to set a context for a request
func NewSendMediaGroupLinkParamsWithContext(ctx context.Context) *SendMediaGroupLinkParams {
	var ()
	return &SendMediaGroupLinkParams{

		Context: ctx,
	}
}

// NewSendMediaGroupLinkParamsWithHTTPClient creates a new SendMediaGroupLinkParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSendMediaGroupLinkParamsWithHTTPClient(client *http.Client) *SendMediaGroupLinkParams {
	var ()
	return &SendMediaGroupLinkParams{
		HTTPClient: client,
	}
}

/*SendMediaGroupLinkParams contains all the parameters to send to the API endpoint
for the send media group link operation typically these are written to a http.Request
*/
type SendMediaGroupLinkParams struct {

	/*Body*/
	Body *models.SendMediaGroupLinkBody
	/*Token
	  bot's token to authorize the request

	*/
	Token *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the send media group link params
func (o *SendMediaGroupLinkParams) WithTimeout(timeout time.Duration) *SendMediaGroupLinkParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the send media group link params
func (o *SendMediaGroupLinkParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the send media group link params
func (o *SendMediaGroupLinkParams) WithContext(ctx context.Context) *SendMediaGroupLinkParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the send media group link params
func (o *SendMediaGroupLinkParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the send media group link params
func (o *SendMediaGroupLinkParams) WithHTTPClient(client *http.Client) *SendMediaGroupLinkParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the send media group link params
func (o *SendMediaGroupLinkParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the send media group link params
func (o *SendMediaGroupLinkParams) WithBody(body *models.SendMediaGroupLinkBody) *SendMediaGroupLinkParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the send media group link params
func (o *SendMediaGroupLinkParams) SetBody(body *models.SendMediaGroupLinkBody) {
	o.Body = body
}

// WithToken adds the token to the send media group link params
func (o *SendMediaGroupLinkParams) WithToken(token *string) *SendMediaGroupLinkParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the send media group link params
func (o *SendMediaGroupLinkParams) SetToken(token *string) {
	o.Token = token
}

// WriteToRequest writes these params to a swagger request
func (o *SendMediaGroupLinkParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body == nil {
		o.Body = new(models.SendMediaGroupLinkBody)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if o.Token != nil {

		// path param token
		if err := r.SetPathParam("token", *o.Token); err != nil {
			return err
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
