package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// GameHighScore game high score
// swagger:model GameHighScore
type GameHighScore struct {

	// position
	Position int64 `json:"position,omitempty"`

	// score
	Score int64 `json:"score,omitempty"`

	// user
	User *User `json:"user,omitempty"`
}

// Validate validates this game high score
func (m *GameHighScore) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateUser(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GameHighScore) validateUser(formats strfmt.Registry) error {

	if swag.IsZero(m.User) { // not required
		return nil
	}

	if m.User != nil {

		if err := m.User.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("user")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GameHighScore) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GameHighScore) UnmarshalBinary(b []byte) error {
	var res GameHighScore
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
