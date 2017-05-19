package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Animation animation
// swagger:model Animation
type Animation struct {

	// file id
	// Required: true
	FileID *string `json:"file_id"`

	// file name
	FileName string `json:"file_name,omitempty"`

	// file size
	FileSize int64 `json:"file_size,omitempty"`

	// mime type
	MimeType string `json:"mime_type,omitempty"`

	// thumb
	Thumb *PhotoSize `json:"thumb,omitempty"`
}

// Validate validates this animation
func (m *Animation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFileID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateThumb(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Animation) validateFileID(formats strfmt.Registry) error {

	if err := validate.Required("file_id", "body", m.FileID); err != nil {
		return err
	}

	return nil
}

func (m *Animation) validateThumb(formats strfmt.Registry) error {

	if swag.IsZero(m.Thumb) { // not required
		return nil
	}

	if m.Thumb != nil {

		if err := m.Thumb.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("thumb")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Animation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Animation) UnmarshalBinary(b []byte) error {
	var res Animation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
