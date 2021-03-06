// Code generated by go-swagger; DO NOT EDIT.

package chats

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/olebedev/go-tgbot/models"
)

// DeleteChatStickerSetReader is a Reader for the DeleteChatStickerSet structure.
type DeleteChatStickerSetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteChatStickerSetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteChatStickerSetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteChatStickerSetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewDeleteChatStickerSetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewDeleteChatStickerSetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteChatStickerSetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 420:
		result := NewDeleteChatStickerSetEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewDeleteChatStickerSetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteChatStickerSetOK creates a DeleteChatStickerSetOK with default headers values
func NewDeleteChatStickerSetOK() *DeleteChatStickerSetOK {
	return &DeleteChatStickerSetOK{}
}

/*DeleteChatStickerSetOK handles this case with default header values.

DeleteChatStickerSetOK delete chat sticker set o k
*/
type DeleteChatStickerSetOK struct {
	Payload *models.ResponseBool
}

func (o *DeleteChatStickerSetOK) Error() string {
	return fmt.Sprintf("[GET /bot{token}/deleteChatStickerSet][%d] deleteChatStickerSetOK  %+v", 200, o.Payload)
}

func (o *DeleteChatStickerSetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseBool)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteChatStickerSetBadRequest creates a DeleteChatStickerSetBadRequest with default headers values
func NewDeleteChatStickerSetBadRequest() *DeleteChatStickerSetBadRequest {
	return &DeleteChatStickerSetBadRequest{}
}

/*DeleteChatStickerSetBadRequest handles this case with default header values.

Bad Request
*/
type DeleteChatStickerSetBadRequest struct {
	Payload *models.Error
}

func (o *DeleteChatStickerSetBadRequest) Error() string {
	return fmt.Sprintf("[GET /bot{token}/deleteChatStickerSet][%d] deleteChatStickerSetBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteChatStickerSetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteChatStickerSetUnauthorized creates a DeleteChatStickerSetUnauthorized with default headers values
func NewDeleteChatStickerSetUnauthorized() *DeleteChatStickerSetUnauthorized {
	return &DeleteChatStickerSetUnauthorized{}
}

/*DeleteChatStickerSetUnauthorized handles this case with default header values.

Unauthorized
*/
type DeleteChatStickerSetUnauthorized struct {
	Payload *models.Error
}

func (o *DeleteChatStickerSetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /bot{token}/deleteChatStickerSet][%d] deleteChatStickerSetUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteChatStickerSetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteChatStickerSetForbidden creates a DeleteChatStickerSetForbidden with default headers values
func NewDeleteChatStickerSetForbidden() *DeleteChatStickerSetForbidden {
	return &DeleteChatStickerSetForbidden{}
}

/*DeleteChatStickerSetForbidden handles this case with default header values.

Forbidden
*/
type DeleteChatStickerSetForbidden struct {
	Payload *models.Error
}

func (o *DeleteChatStickerSetForbidden) Error() string {
	return fmt.Sprintf("[GET /bot{token}/deleteChatStickerSet][%d] deleteChatStickerSetForbidden  %+v", 403, o.Payload)
}

func (o *DeleteChatStickerSetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteChatStickerSetNotFound creates a DeleteChatStickerSetNotFound with default headers values
func NewDeleteChatStickerSetNotFound() *DeleteChatStickerSetNotFound {
	return &DeleteChatStickerSetNotFound{}
}

/*DeleteChatStickerSetNotFound handles this case with default header values.

Not Found
*/
type DeleteChatStickerSetNotFound struct {
	Payload *models.Error
}

func (o *DeleteChatStickerSetNotFound) Error() string {
	return fmt.Sprintf("[GET /bot{token}/deleteChatStickerSet][%d] deleteChatStickerSetNotFound  %+v", 404, o.Payload)
}

func (o *DeleteChatStickerSetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteChatStickerSetEnhanceYourCalm creates a DeleteChatStickerSetEnhanceYourCalm with default headers values
func NewDeleteChatStickerSetEnhanceYourCalm() *DeleteChatStickerSetEnhanceYourCalm {
	return &DeleteChatStickerSetEnhanceYourCalm{}
}

/*DeleteChatStickerSetEnhanceYourCalm handles this case with default header values.

Flood
*/
type DeleteChatStickerSetEnhanceYourCalm struct {
	Payload *models.Error
}

func (o *DeleteChatStickerSetEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /bot{token}/deleteChatStickerSet][%d] deleteChatStickerSetEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *DeleteChatStickerSetEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteChatStickerSetInternalServerError creates a DeleteChatStickerSetInternalServerError with default headers values
func NewDeleteChatStickerSetInternalServerError() *DeleteChatStickerSetInternalServerError {
	return &DeleteChatStickerSetInternalServerError{}
}

/*DeleteChatStickerSetInternalServerError handles this case with default header values.

Internal
*/
type DeleteChatStickerSetInternalServerError struct {
	Payload *models.Error
}

func (o *DeleteChatStickerSetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /bot{token}/deleteChatStickerSet][%d] deleteChatStickerSetInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteChatStickerSetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
