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

// NewSendVenueParams creates a new SendVenueParams object
// with the default values initialized.
func NewSendVenueParams() *SendVenueParams {
	var ()
	return &SendVenueParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSendVenueParamsWithTimeout creates a new SendVenueParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSendVenueParamsWithTimeout(timeout time.Duration) *SendVenueParams {
	var ()
	return &SendVenueParams{

		timeout: timeout,
	}
}

// NewSendVenueParamsWithContext creates a new SendVenueParams object
// with the default values initialized, and the ability to set a context for a request
func NewSendVenueParamsWithContext(ctx context.Context) *SendVenueParams {
	var ()
	return &SendVenueParams{

		Context: ctx,
	}
}

// NewSendVenueParamsWithHTTPClient creates a new SendVenueParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSendVenueParamsWithHTTPClient(client *http.Client) *SendVenueParams {
	var ()
	return &SendVenueParams{
		HTTPClient: client,
	}
}

/*SendVenueParams contains all the parameters to send to the API endpoint
for the send venue operation typically these are written to a http.Request
*/
type SendVenueParams struct {

	/*Body*/
	Body *models.SendVenueBody
	/*Token
	  bot's token to authorize the request

	*/
	Token *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the send venue params
func (o *SendVenueParams) WithTimeout(timeout time.Duration) *SendVenueParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the send venue params
func (o *SendVenueParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the send venue params
func (o *SendVenueParams) WithContext(ctx context.Context) *SendVenueParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the send venue params
func (o *SendVenueParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the send venue params
func (o *SendVenueParams) WithHTTPClient(client *http.Client) *SendVenueParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the send venue params
func (o *SendVenueParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the send venue params
func (o *SendVenueParams) WithBody(body *models.SendVenueBody) *SendVenueParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the send venue params
func (o *SendVenueParams) SetBody(body *models.SendVenueBody) {
	o.Body = body
}

// WithToken adds the token to the send venue params
func (o *SendVenueParams) WithToken(token *string) *SendVenueParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the send venue params
func (o *SendVenueParams) SetToken(token *string) {
	o.Token = token
}

// WriteToRequest writes these params to a swagger request
func (o *SendVenueParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body == nil {
		o.Body = new(models.SendVenueBody)
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
