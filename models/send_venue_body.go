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

// SendVenueBody send venue body
// swagger:model SendVenueBody
type SendVenueBody struct {

	// address
	// Required: true
	Address *string `json:"address"`

	// chat id
	// Required: true
	ChatID interface{} `json:"chat_id"`

	// disable notification
	DisableNotification bool `json:"disable_notification,omitempty"`

	// foursquare id
	FoursquareID string `json:"foursquare_id,omitempty"`

	// latitude
	// Required: true
	Latitude *float64 `json:"latitude"`

	// longitude
	// Required: true
	Longitude *float64 `json:"longitude"`

	// reply markup
	ReplyMarkup interface{} `json:"reply_markup,omitempty"`

	// reply to message id
	ReplyToMessageID int64 `json:"reply_to_message_id,omitempty"`

	// title
	// Required: true
	Title *string `json:"title"`
}

// Validate validates this send venue body
func (m *SendVenueBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddress(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateChatID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLatitude(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLongitude(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTitle(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SendVenueBody) validateAddress(formats strfmt.Registry) error {

	if err := validate.Required("address", "body", m.Address); err != nil {
		return err
	}

	return nil
}

func (m *SendVenueBody) validateChatID(formats strfmt.Registry) error {

	return nil
}

func (m *SendVenueBody) validateLatitude(formats strfmt.Registry) error {

	if err := validate.Required("latitude", "body", m.Latitude); err != nil {
		return err
	}

	return nil
}

func (m *SendVenueBody) validateLongitude(formats strfmt.Registry) error {

	if err := validate.Required("longitude", "body", m.Longitude); err != nil {
		return err
	}

	return nil
}

func (m *SendVenueBody) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("title", "body", m.Title); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SendVenueBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SendVenueBody) UnmarshalBinary(b []byte) error {
	var res SendVenueBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
