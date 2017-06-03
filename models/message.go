package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Message message
// swagger:model Message
type Message struct {

	// audio
	Audio *Audio `json:"audio,omitempty"`

	// caption
	Caption string `json:"caption,omitempty"`

	// channel chat created
	ChannelChatCreated bool `json:"channel_chat_created,omitempty"`

	// chat
	Chat *Chat `json:"chat,omitempty"`

	// contact
	Contact *Contact `json:"contact,omitempty"`

	// date
	Date int64 `json:"date,omitempty"`

	// delete chat photo
	DeleteChatPhoto bool `json:"delete_chat_photo,omitempty"`

	// document
	Document *Document `json:"document,omitempty"`

	// edit date
	EditDate int64 `json:"edit_date,omitempty"`

	// entities
	Entities []*MessageEntity `json:"entities"`

	// forward date
	ForwardDate int64 `json:"forward_date,omitempty"`

	// forward from
	ForwardFrom *User `json:"forward_from,omitempty"`

	// forward from chat
	ForwardFromChat *Chat `json:"forward_from_chat,omitempty"`

	// forward from message id
	ForwardFromMessageID int64 `json:"forward_from_message_id,omitempty"`

	// from
	From *User `json:"from,omitempty"`

	// game
	Game *Game `json:"game,omitempty"`

	// group chat created
	GroupChatCreated bool `json:"group_chat_created,omitempty"`

	// left chat member
	LeftChatMember *User `json:"left_chat_member,omitempty"`

	// location
	Location *Location `json:"location,omitempty"`

	// message id
	MessageID int64 `json:"message_id,omitempty"`

	// migrate from chat id
	MigrateFromChatID int64 `json:"migrate_from_chat_id,omitempty"`

	// migrate to chat id
	MigrateToChatID int64 `json:"migrate_to_chat_id,omitempty"`

	// new chat members
	NewChatMembers []*User `json:"new_chat_members"`

	// new chat photo
	NewChatPhoto []*PhotoSize `json:"new_chat_photo"`

	// new chat title
	NewChatTitle string `json:"new_chat_title,omitempty"`

	// photo
	Photo []*PhotoSize `json:"photo"`

	// pinned message
	PinnedMessage *Message `json:"pinned_message,omitempty"`

	// reply to message
	ReplyToMessage *Message `json:"reply_to_message,omitempty"`

	// sticker
	Sticker *Sticker `json:"sticker,omitempty"`

	// successful payment
	SuccessfulPayment *SuccessfulPayment `json:"successful_payment,omitempty"`

	// supergroup chat created
	SupergroupChatCreated bool `json:"supergroup_chat_created,omitempty"`

	// text
	Text string `json:"text,omitempty"`

	// venue
	Venue *Venue `json:"venue,omitempty"`

	// video
	Video *Video `json:"video,omitempty"`

	// video note
	VideoNote *VideoNote `json:"video_note,omitempty"`

	// voice
	Voice *Voice `json:"voice,omitempty"`
}

// Validate validates this message
func (m *Message) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAudio(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateChat(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateContact(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDocument(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateEntities(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateForwardFrom(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateForwardFromChat(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateFrom(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateGame(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLeftChatMember(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLocation(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateNewChatMembers(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateNewChatPhoto(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePhoto(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePinnedMessage(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateReplyToMessage(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSticker(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSuccessfulPayment(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateVenue(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateVideo(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateVideoNote(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateVoice(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Message) validateAudio(formats strfmt.Registry) error {

	if swag.IsZero(m.Audio) { // not required
		return nil
	}

	if m.Audio != nil {

		if err := m.Audio.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("audio")
			}
			return err
		}
	}

	return nil
}

func (m *Message) validateChat(formats strfmt.Registry) error {

	if swag.IsZero(m.Chat) { // not required
		return nil
	}

	if m.Chat != nil {

		if err := m.Chat.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("chat")
			}
			return err
		}
	}

	return nil
}

func (m *Message) validateContact(formats strfmt.Registry) error {

	if swag.IsZero(m.Contact) { // not required
		return nil
	}

	if m.Contact != nil {

		if err := m.Contact.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("contact")
			}
			return err
		}
	}

	return nil
}

func (m *Message) validateDocument(formats strfmt.Registry) error {

	if swag.IsZero(m.Document) { // not required
		return nil
	}

	if m.Document != nil {

		if err := m.Document.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("document")
			}
			return err
		}
	}

	return nil
}

func (m *Message) validateEntities(formats strfmt.Registry) error {

	if swag.IsZero(m.Entities) { // not required
		return nil
	}

	for i := 0; i < len(m.Entities); i++ {

		if swag.IsZero(m.Entities[i]) { // not required
			continue
		}

		if m.Entities[i] != nil {

			if err := m.Entities[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("entities" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Message) validateForwardFrom(formats strfmt.Registry) error {

	if swag.IsZero(m.ForwardFrom) { // not required
		return nil
	}

	if m.ForwardFrom != nil {

		if err := m.ForwardFrom.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("forward_from")
			}
			return err
		}
	}

	return nil
}

func (m *Message) validateForwardFromChat(formats strfmt.Registry) error {

	if swag.IsZero(m.ForwardFromChat) { // not required
		return nil
	}

	if m.ForwardFromChat != nil {

		if err := m.ForwardFromChat.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("forward_from_chat")
			}
			return err
		}
	}

	return nil
}

func (m *Message) validateFrom(formats strfmt.Registry) error {

	if swag.IsZero(m.From) { // not required
		return nil
	}

	if m.From != nil {

		if err := m.From.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("from")
			}
			return err
		}
	}

	return nil
}

func (m *Message) validateGame(formats strfmt.Registry) error {

	if swag.IsZero(m.Game) { // not required
		return nil
	}

	if m.Game != nil {

		if err := m.Game.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("game")
			}
			return err
		}
	}

	return nil
}

func (m *Message) validateLeftChatMember(formats strfmt.Registry) error {

	if swag.IsZero(m.LeftChatMember) { // not required
		return nil
	}

	if m.LeftChatMember != nil {

		if err := m.LeftChatMember.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("left_chat_member")
			}
			return err
		}
	}

	return nil
}

func (m *Message) validateLocation(formats strfmt.Registry) error {

	if swag.IsZero(m.Location) { // not required
		return nil
	}

	if m.Location != nil {

		if err := m.Location.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("location")
			}
			return err
		}
	}

	return nil
}

func (m *Message) validateNewChatMembers(formats strfmt.Registry) error {

	if swag.IsZero(m.NewChatMembers) { // not required
		return nil
	}

	for i := 0; i < len(m.NewChatMembers); i++ {

		if swag.IsZero(m.NewChatMembers[i]) { // not required
			continue
		}

		if m.NewChatMembers[i] != nil {

			if err := m.NewChatMembers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("new_chat_members" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Message) validateNewChatPhoto(formats strfmt.Registry) error {

	if swag.IsZero(m.NewChatPhoto) { // not required
		return nil
	}

	for i := 0; i < len(m.NewChatPhoto); i++ {

		if swag.IsZero(m.NewChatPhoto[i]) { // not required
			continue
		}

		if m.NewChatPhoto[i] != nil {

			if err := m.NewChatPhoto[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("new_chat_photo" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Message) validatePhoto(formats strfmt.Registry) error {

	if swag.IsZero(m.Photo) { // not required
		return nil
	}

	for i := 0; i < len(m.Photo); i++ {

		if swag.IsZero(m.Photo[i]) { // not required
			continue
		}

		if m.Photo[i] != nil {

			if err := m.Photo[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("photo" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Message) validatePinnedMessage(formats strfmt.Registry) error {

	if swag.IsZero(m.PinnedMessage) { // not required
		return nil
	}

	if m.PinnedMessage != nil {

		if err := m.PinnedMessage.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pinned_message")
			}
			return err
		}
	}

	return nil
}

func (m *Message) validateReplyToMessage(formats strfmt.Registry) error {

	if swag.IsZero(m.ReplyToMessage) { // not required
		return nil
	}

	if m.ReplyToMessage != nil {

		if err := m.ReplyToMessage.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("reply_to_message")
			}
			return err
		}
	}

	return nil
}

func (m *Message) validateSticker(formats strfmt.Registry) error {

	if swag.IsZero(m.Sticker) { // not required
		return nil
	}

	if m.Sticker != nil {

		if err := m.Sticker.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sticker")
			}
			return err
		}
	}

	return nil
}

func (m *Message) validateSuccessfulPayment(formats strfmt.Registry) error {

	if swag.IsZero(m.SuccessfulPayment) { // not required
		return nil
	}

	if m.SuccessfulPayment != nil {

		if err := m.SuccessfulPayment.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("successful_payment")
			}
			return err
		}
	}

	return nil
}

func (m *Message) validateVenue(formats strfmt.Registry) error {

	if swag.IsZero(m.Venue) { // not required
		return nil
	}

	if m.Venue != nil {

		if err := m.Venue.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("venue")
			}
			return err
		}
	}

	return nil
}

func (m *Message) validateVideo(formats strfmt.Registry) error {

	if swag.IsZero(m.Video) { // not required
		return nil
	}

	if m.Video != nil {

		if err := m.Video.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("video")
			}
			return err
		}
	}

	return nil
}

func (m *Message) validateVideoNote(formats strfmt.Registry) error {

	if swag.IsZero(m.VideoNote) { // not required
		return nil
	}

	if m.VideoNote != nil {

		if err := m.VideoNote.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("video_note")
			}
			return err
		}
	}

	return nil
}

func (m *Message) validateVoice(formats strfmt.Registry) error {

	if swag.IsZero(m.Voice) { // not required
		return nil
	}

	if m.Voice != nil {

		if err := m.Voice.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("voice")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Message) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Message) UnmarshalBinary(b []byte) error {
	var res Message
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
