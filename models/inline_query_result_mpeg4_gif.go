package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// InlineQueryResultMpeg4Gif inline query result mpeg4 gif
// swagger:model InlineQueryResultMpeg4Gif
type InlineQueryResultMpeg4Gif struct {

	// caption
	Caption string `json:"caption,omitempty"`

	// id
	// Required: true
	ID *string `json:"id"`

	// input message content
	InputMessageContent interface{} `json:"input_message_content,omitempty"`

	// mpeg4 duration
	Mpeg4Duration int64 `json:"mpeg4_duration,omitempty"`

	// mpeg4 height
	Mpeg4Height int64 `json:"mpeg4_height,omitempty"`

	// mpeg4 url
	// Required: true
	Mpeg4URL *string `json:"mpeg4_url"`

	// mpeg4 width
	Mpeg4Width int64 `json:"mpeg4_width,omitempty"`

	// reply markup
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`

	// thumb url
	// Required: true
	ThumbURL *string `json:"thumb_url"`

	// title
	Title string `json:"title,omitempty"`

	// type
	// Required: true
	Type InlineType `json:"type"`
}

// Validate validates this inline query result mpeg4 gif
func (m *InlineQueryResultMpeg4Gif) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMpeg4URL(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateReplyMarkup(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateThumbURL(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InlineQueryResultMpeg4Gif) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *InlineQueryResultMpeg4Gif) validateMpeg4URL(formats strfmt.Registry) error {

	if err := validate.Required("mpeg4_url", "body", m.Mpeg4URL); err != nil {
		return err
	}

	return nil
}

func (m *InlineQueryResultMpeg4Gif) validateReplyMarkup(formats strfmt.Registry) error {

	if swag.IsZero(m.ReplyMarkup) { // not required
		return nil
	}

	if m.ReplyMarkup != nil {

		if err := m.ReplyMarkup.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("reply_markup")
			}
			return err
		}
	}

	return nil
}

func (m *InlineQueryResultMpeg4Gif) validateThumbURL(formats strfmt.Registry) error {

	if err := validate.Required("thumb_url", "body", m.ThumbURL); err != nil {
		return err
	}

	return nil
}

func (m *InlineQueryResultMpeg4Gif) validateType(formats strfmt.Registry) error {

	if err := m.Type.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("type")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InlineQueryResultMpeg4Gif) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InlineQueryResultMpeg4Gif) UnmarshalBinary(b []byte) error {
	var res InlineQueryResultMpeg4Gif
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
