package games

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewSetGameScoreParams creates a new SetGameScoreParams object
// with the default values initialized.
func NewSetGameScoreParams() *SetGameScoreParams {
	var ()
	return &SetGameScoreParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSetGameScoreParamsWithTimeout creates a new SetGameScoreParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSetGameScoreParamsWithTimeout(timeout time.Duration) *SetGameScoreParams {
	var ()
	return &SetGameScoreParams{

		timeout: timeout,
	}
}

// NewSetGameScoreParamsWithContext creates a new SetGameScoreParams object
// with the default values initialized, and the ability to set a context for a request
func NewSetGameScoreParamsWithContext(ctx context.Context) *SetGameScoreParams {
	var ()
	return &SetGameScoreParams{

		Context: ctx,
	}
}

// NewSetGameScoreParamsWithHTTPClient creates a new SetGameScoreParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSetGameScoreParamsWithHTTPClient(client *http.Client) *SetGameScoreParams {
	var ()
	return &SetGameScoreParams{
		HTTPClient: client,
	}
}

/*SetGameScoreParams contains all the parameters to send to the API endpoint
for the set game score operation typically these are written to a http.Request
*/
type SetGameScoreParams struct {

	/*ChatID*/
	ChatID *int64
	/*DisableEditMessage*/
	DisableEditMessage *bool
	/*Force*/
	Force *bool
	/*InlineMessageID*/
	InlineMessageID *string
	/*MessageID*/
	MessageID *int64
	/*Score*/
	Score int64
	/*Token
	  bot's token to authorize the request

	*/
	Token *string
	/*UserID*/
	UserID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the set game score params
func (o *SetGameScoreParams) WithTimeout(timeout time.Duration) *SetGameScoreParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the set game score params
func (o *SetGameScoreParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the set game score params
func (o *SetGameScoreParams) WithContext(ctx context.Context) *SetGameScoreParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the set game score params
func (o *SetGameScoreParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the set game score params
func (o *SetGameScoreParams) WithHTTPClient(client *http.Client) *SetGameScoreParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the set game score params
func (o *SetGameScoreParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithChatID adds the chatID to the set game score params
func (o *SetGameScoreParams) WithChatID(chatID *int64) *SetGameScoreParams {
	o.SetChatID(chatID)
	return o
}

// SetChatID adds the chatId to the set game score params
func (o *SetGameScoreParams) SetChatID(chatID *int64) {
	o.ChatID = chatID
}

// WithDisableEditMessage adds the disableEditMessage to the set game score params
func (o *SetGameScoreParams) WithDisableEditMessage(disableEditMessage *bool) *SetGameScoreParams {
	o.SetDisableEditMessage(disableEditMessage)
	return o
}

// SetDisableEditMessage adds the disableEditMessage to the set game score params
func (o *SetGameScoreParams) SetDisableEditMessage(disableEditMessage *bool) {
	o.DisableEditMessage = disableEditMessage
}

// WithForce adds the force to the set game score params
func (o *SetGameScoreParams) WithForce(force *bool) *SetGameScoreParams {
	o.SetForce(force)
	return o
}

// SetForce adds the force to the set game score params
func (o *SetGameScoreParams) SetForce(force *bool) {
	o.Force = force
}

// WithInlineMessageID adds the inlineMessageID to the set game score params
func (o *SetGameScoreParams) WithInlineMessageID(inlineMessageID *string) *SetGameScoreParams {
	o.SetInlineMessageID(inlineMessageID)
	return o
}

// SetInlineMessageID adds the inlineMessageId to the set game score params
func (o *SetGameScoreParams) SetInlineMessageID(inlineMessageID *string) {
	o.InlineMessageID = inlineMessageID
}

// WithMessageID adds the messageID to the set game score params
func (o *SetGameScoreParams) WithMessageID(messageID *int64) *SetGameScoreParams {
	o.SetMessageID(messageID)
	return o
}

// SetMessageID adds the messageId to the set game score params
func (o *SetGameScoreParams) SetMessageID(messageID *int64) {
	o.MessageID = messageID
}

// WithScore adds the score to the set game score params
func (o *SetGameScoreParams) WithScore(score int64) *SetGameScoreParams {
	o.SetScore(score)
	return o
}

// SetScore adds the score to the set game score params
func (o *SetGameScoreParams) SetScore(score int64) {
	o.Score = score
}

// WithToken adds the token to the set game score params
func (o *SetGameScoreParams) WithToken(token *string) *SetGameScoreParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the set game score params
func (o *SetGameScoreParams) SetToken(token *string) {
	o.Token = token
}

// WithUserID adds the userID to the set game score params
func (o *SetGameScoreParams) WithUserID(userID int64) *SetGameScoreParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the set game score params
func (o *SetGameScoreParams) SetUserID(userID int64) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *SetGameScoreParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ChatID != nil {

		// query param chat_id
		var qrChatID int64
		if o.ChatID != nil {
			qrChatID = *o.ChatID
		}
		qChatID := swag.FormatInt64(qrChatID)
		if qChatID != "" {
			if err := r.SetQueryParam("chat_id", qChatID); err != nil {
				return err
			}
		}

	}

	if o.DisableEditMessage != nil {

		// query param disable_edit_message
		var qrDisableEditMessage bool
		if o.DisableEditMessage != nil {
			qrDisableEditMessage = *o.DisableEditMessage
		}
		qDisableEditMessage := swag.FormatBool(qrDisableEditMessage)
		if qDisableEditMessage != "" {
			if err := r.SetQueryParam("disable_edit_message", qDisableEditMessage); err != nil {
				return err
			}
		}

	}

	if o.Force != nil {

		// query param force
		var qrForce bool
		if o.Force != nil {
			qrForce = *o.Force
		}
		qForce := swag.FormatBool(qrForce)
		if qForce != "" {
			if err := r.SetQueryParam("force", qForce); err != nil {
				return err
			}
		}

	}

	if o.InlineMessageID != nil {

		// query param inline_message_id
		var qrInlineMessageID string
		if o.InlineMessageID != nil {
			qrInlineMessageID = *o.InlineMessageID
		}
		qInlineMessageID := qrInlineMessageID
		if qInlineMessageID != "" {
			if err := r.SetQueryParam("inline_message_id", qInlineMessageID); err != nil {
				return err
			}
		}

	}

	if o.MessageID != nil {

		// query param message_id
		var qrMessageID int64
		if o.MessageID != nil {
			qrMessageID = *o.MessageID
		}
		qMessageID := swag.FormatInt64(qrMessageID)
		if qMessageID != "" {
			if err := r.SetQueryParam("message_id", qMessageID); err != nil {
				return err
			}
		}

	}

	// query param score
	qrScore := o.Score
	qScore := swag.FormatInt64(qrScore)
	if qScore != "" {
		if err := r.SetQueryParam("score", qScore); err != nil {
			return err
		}
	}

	if o.Token != nil {

		// path param token
		if err := r.SetPathParam("token", *o.Token); err != nil {
			return err
		}

	}

	// query param user_id
	qrUserID := o.UserID
	qUserID := swag.FormatInt64(qrUserID)
	if qUserID != "" {
		if err := r.SetQueryParam("user_id", qUserID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
