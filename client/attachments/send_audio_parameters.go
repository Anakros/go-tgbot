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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewSendAudioParams creates a new SendAudioParams object
// with the default values initialized.
func NewSendAudioParams() *SendAudioParams {
	var ()
	return &SendAudioParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSendAudioParamsWithTimeout creates a new SendAudioParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSendAudioParamsWithTimeout(timeout time.Duration) *SendAudioParams {
	var ()
	return &SendAudioParams{

		timeout: timeout,
	}
}

// NewSendAudioParamsWithContext creates a new SendAudioParams object
// with the default values initialized, and the ability to set a context for a request
func NewSendAudioParamsWithContext(ctx context.Context) *SendAudioParams {
	var ()
	return &SendAudioParams{

		Context: ctx,
	}
}

// NewSendAudioParamsWithHTTPClient creates a new SendAudioParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSendAudioParamsWithHTTPClient(client *http.Client) *SendAudioParams {
	var ()
	return &SendAudioParams{
		HTTPClient: client,
	}
}

/*SendAudioParams contains all the parameters to send to the API endpoint
for the send audio operation typically these are written to a http.Request
*/
type SendAudioParams struct {

	/*Audio*/
	Audio runtime.NamedReadCloser
	/*Caption*/
	Caption *string
	/*ChatID*/
	ChatID string
	/*DisableNotification*/
	DisableNotification *bool
	/*Duration*/
	Duration *int64
	/*Performer*/
	Performer *string
	/*ReplyMarkup
	  json string of reply_markup object

	*/
	ReplyMarkup *string
	/*ReplyToMessageID*/
	ReplyToMessageID *int64
	/*Title*/
	Title *string
	/*Token
	  bot's token to authorize the request

	*/
	Token *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the send audio params
func (o *SendAudioParams) WithTimeout(timeout time.Duration) *SendAudioParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the send audio params
func (o *SendAudioParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the send audio params
func (o *SendAudioParams) WithContext(ctx context.Context) *SendAudioParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the send audio params
func (o *SendAudioParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the send audio params
func (o *SendAudioParams) WithHTTPClient(client *http.Client) *SendAudioParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the send audio params
func (o *SendAudioParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAudio adds the audio to the send audio params
func (o *SendAudioParams) WithAudio(audio runtime.NamedReadCloser) *SendAudioParams {
	o.SetAudio(audio)
	return o
}

// SetAudio adds the audio to the send audio params
func (o *SendAudioParams) SetAudio(audio runtime.NamedReadCloser) {
	o.Audio = audio
}

// WithCaption adds the caption to the send audio params
func (o *SendAudioParams) WithCaption(caption *string) *SendAudioParams {
	o.SetCaption(caption)
	return o
}

// SetCaption adds the caption to the send audio params
func (o *SendAudioParams) SetCaption(caption *string) {
	o.Caption = caption
}

// WithChatID adds the chatID to the send audio params
func (o *SendAudioParams) WithChatID(chatID string) *SendAudioParams {
	o.SetChatID(chatID)
	return o
}

// SetChatID adds the chatId to the send audio params
func (o *SendAudioParams) SetChatID(chatID string) {
	o.ChatID = chatID
}

// WithDisableNotification adds the disableNotification to the send audio params
func (o *SendAudioParams) WithDisableNotification(disableNotification *bool) *SendAudioParams {
	o.SetDisableNotification(disableNotification)
	return o
}

// SetDisableNotification adds the disableNotification to the send audio params
func (o *SendAudioParams) SetDisableNotification(disableNotification *bool) {
	o.DisableNotification = disableNotification
}

// WithDuration adds the duration to the send audio params
func (o *SendAudioParams) WithDuration(duration *int64) *SendAudioParams {
	o.SetDuration(duration)
	return o
}

// SetDuration adds the duration to the send audio params
func (o *SendAudioParams) SetDuration(duration *int64) {
	o.Duration = duration
}

// WithPerformer adds the performer to the send audio params
func (o *SendAudioParams) WithPerformer(performer *string) *SendAudioParams {
	o.SetPerformer(performer)
	return o
}

// SetPerformer adds the performer to the send audio params
func (o *SendAudioParams) SetPerformer(performer *string) {
	o.Performer = performer
}

// WithReplyMarkup adds the replyMarkup to the send audio params
func (o *SendAudioParams) WithReplyMarkup(replyMarkup *string) *SendAudioParams {
	o.SetReplyMarkup(replyMarkup)
	return o
}

// SetReplyMarkup adds the replyMarkup to the send audio params
func (o *SendAudioParams) SetReplyMarkup(replyMarkup *string) {
	o.ReplyMarkup = replyMarkup
}

// WithReplyToMessageID adds the replyToMessageID to the send audio params
func (o *SendAudioParams) WithReplyToMessageID(replyToMessageID *int64) *SendAudioParams {
	o.SetReplyToMessageID(replyToMessageID)
	return o
}

// SetReplyToMessageID adds the replyToMessageId to the send audio params
func (o *SendAudioParams) SetReplyToMessageID(replyToMessageID *int64) {
	o.ReplyToMessageID = replyToMessageID
}

// WithTitle adds the title to the send audio params
func (o *SendAudioParams) WithTitle(title *string) *SendAudioParams {
	o.SetTitle(title)
	return o
}

// SetTitle adds the title to the send audio params
func (o *SendAudioParams) SetTitle(title *string) {
	o.Title = title
}

// WithToken adds the token to the send audio params
func (o *SendAudioParams) WithToken(token *string) *SendAudioParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the send audio params
func (o *SendAudioParams) SetToken(token *string) {
	o.Token = token
}

// WriteToRequest writes these params to a swagger request
func (o *SendAudioParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// form file param audio
	if err := r.SetFileParam("audio", o.Audio); err != nil {
		return err
	}

	if o.Caption != nil {

		// form param caption
		var frCaption string
		if o.Caption != nil {
			frCaption = *o.Caption
		}
		fCaption := frCaption
		if fCaption != "" {
			if err := r.SetFormParam("caption", fCaption); err != nil {
				return err
			}
		}

	}

	// form param chat_id
	frChatID := o.ChatID
	fChatID := frChatID
	if fChatID != "" {
		if err := r.SetFormParam("chat_id", fChatID); err != nil {
			return err
		}
	}

	if o.DisableNotification != nil {

		// form param disable_notification
		var frDisableNotification bool
		if o.DisableNotification != nil {
			frDisableNotification = *o.DisableNotification
		}
		fDisableNotification := swag.FormatBool(frDisableNotification)
		if fDisableNotification != "" {
			if err := r.SetFormParam("disable_notification", fDisableNotification); err != nil {
				return err
			}
		}

	}

	if o.Duration != nil {

		// form param duration
		var frDuration int64
		if o.Duration != nil {
			frDuration = *o.Duration
		}
		fDuration := swag.FormatInt64(frDuration)
		if fDuration != "" {
			if err := r.SetFormParam("duration", fDuration); err != nil {
				return err
			}
		}

	}

	if o.Performer != nil {

		// form param performer
		var frPerformer string
		if o.Performer != nil {
			frPerformer = *o.Performer
		}
		fPerformer := frPerformer
		if fPerformer != "" {
			if err := r.SetFormParam("performer", fPerformer); err != nil {
				return err
			}
		}

	}

	if o.ReplyMarkup != nil {

		// form param reply_markup
		var frReplyMarkup string
		if o.ReplyMarkup != nil {
			frReplyMarkup = *o.ReplyMarkup
		}
		fReplyMarkup := frReplyMarkup
		if fReplyMarkup != "" {
			if err := r.SetFormParam("reply_markup", fReplyMarkup); err != nil {
				return err
			}
		}

	}

	if o.ReplyToMessageID != nil {

		// form param reply_to_message_id
		var frReplyToMessageID int64
		if o.ReplyToMessageID != nil {
			frReplyToMessageID = *o.ReplyToMessageID
		}
		fReplyToMessageID := swag.FormatInt64(frReplyToMessageID)
		if fReplyToMessageID != "" {
			if err := r.SetFormParam("reply_to_message_id", fReplyToMessageID); err != nil {
				return err
			}
		}

	}

	if o.Title != nil {

		// form param title
		var frTitle string
		if o.Title != nil {
			frTitle = *o.Title
		}
		fTitle := frTitle
		if fTitle != "" {
			if err := r.SetFormParam("title", fTitle); err != nil {
				return err
			}
		}

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
