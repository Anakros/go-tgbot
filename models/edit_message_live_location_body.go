// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// EditMessageLiveLocationBody edit message live location body
// swagger:model EditMessageLiveLocationBody

type EditMessageLiveLocationBody struct {

	// chat id
	ChatID interface{} `json:"chat_id,omitempty"`

	// inline message id
	InlineMessageID string `json:"inline_message_id,omitempty"`

	// latitude
	// Required: true
	Latitude *float64 `json:"latitude"`

	// longitude
	// Required: true
	Longitude *float64 `json:"longitude"`

	// message id
	MessageID int64 `json:"message_id,omitempty"`

	// reply markup
	ReplyMarkup interface{} `json:"reply_markup,omitempty"`
}

/* polymorph EditMessageLiveLocationBody chat_id false */

/* polymorph EditMessageLiveLocationBody inline_message_id false */

/* polymorph EditMessageLiveLocationBody latitude false */

/* polymorph EditMessageLiveLocationBody longitude false */

/* polymorph EditMessageLiveLocationBody message_id false */

/* polymorph EditMessageLiveLocationBody reply_markup false */

// Validate validates this edit message live location body
func (m *EditMessageLiveLocationBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLatitude(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLongitude(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EditMessageLiveLocationBody) validateLatitude(formats strfmt.Registry) error {

	if err := validate.Required("latitude", "body", m.Latitude); err != nil {
		return err
	}

	return nil
}

func (m *EditMessageLiveLocationBody) validateLongitude(formats strfmt.Registry) error {

	if err := validate.Required("longitude", "body", m.Longitude); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EditMessageLiveLocationBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EditMessageLiveLocationBody) UnmarshalBinary(b []byte) error {
	var res EditMessageLiveLocationBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}