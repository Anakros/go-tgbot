package messages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/olebedev/go-tgbot/models"
)

// DeleteMessageReader is a Reader for the DeleteMessage structure.
type DeleteMessageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteMessageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteMessageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteMessageBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewDeleteMessageUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewDeleteMessageForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteMessageNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 420:
		result := NewDeleteMessageEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewDeleteMessageInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteMessageOK creates a DeleteMessageOK with default headers values
func NewDeleteMessageOK() *DeleteMessageOK {
	return &DeleteMessageOK{}
}

/*DeleteMessageOK handles this case with default header values.

DeleteMessageOK delete message o k
*/
type DeleteMessageOK struct {
	Payload bool
}

func (o *DeleteMessageOK) Error() string {
	return fmt.Sprintf("[GET /bot{token}/deleteMessage][%d] deleteMessageOK  %+v", 200, o.Payload)
}

func (o *DeleteMessageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteMessageBadRequest creates a DeleteMessageBadRequest with default headers values
func NewDeleteMessageBadRequest() *DeleteMessageBadRequest {
	return &DeleteMessageBadRequest{}
}

/*DeleteMessageBadRequest handles this case with default header values.

Bad Request
*/
type DeleteMessageBadRequest struct {
	Payload *models.Error
}

func (o *DeleteMessageBadRequest) Error() string {
	return fmt.Sprintf("[GET /bot{token}/deleteMessage][%d] deleteMessageBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteMessageBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteMessageUnauthorized creates a DeleteMessageUnauthorized with default headers values
func NewDeleteMessageUnauthorized() *DeleteMessageUnauthorized {
	return &DeleteMessageUnauthorized{}
}

/*DeleteMessageUnauthorized handles this case with default header values.

Unauthorized
*/
type DeleteMessageUnauthorized struct {
	Payload *models.Error
}

func (o *DeleteMessageUnauthorized) Error() string {
	return fmt.Sprintf("[GET /bot{token}/deleteMessage][%d] deleteMessageUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteMessageUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteMessageForbidden creates a DeleteMessageForbidden with default headers values
func NewDeleteMessageForbidden() *DeleteMessageForbidden {
	return &DeleteMessageForbidden{}
}

/*DeleteMessageForbidden handles this case with default header values.

Forbidden
*/
type DeleteMessageForbidden struct {
	Payload *models.Error
}

func (o *DeleteMessageForbidden) Error() string {
	return fmt.Sprintf("[GET /bot{token}/deleteMessage][%d] deleteMessageForbidden  %+v", 403, o.Payload)
}

func (o *DeleteMessageForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteMessageNotFound creates a DeleteMessageNotFound with default headers values
func NewDeleteMessageNotFound() *DeleteMessageNotFound {
	return &DeleteMessageNotFound{}
}

/*DeleteMessageNotFound handles this case with default header values.

Not Found
*/
type DeleteMessageNotFound struct {
	Payload *models.Error
}

func (o *DeleteMessageNotFound) Error() string {
	return fmt.Sprintf("[GET /bot{token}/deleteMessage][%d] deleteMessageNotFound  %+v", 404, o.Payload)
}

func (o *DeleteMessageNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteMessageEnhanceYourCalm creates a DeleteMessageEnhanceYourCalm with default headers values
func NewDeleteMessageEnhanceYourCalm() *DeleteMessageEnhanceYourCalm {
	return &DeleteMessageEnhanceYourCalm{}
}

/*DeleteMessageEnhanceYourCalm handles this case with default header values.

Flood
*/
type DeleteMessageEnhanceYourCalm struct {
	Payload *models.Error
}

func (o *DeleteMessageEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /bot{token}/deleteMessage][%d] deleteMessageEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *DeleteMessageEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteMessageInternalServerError creates a DeleteMessageInternalServerError with default headers values
func NewDeleteMessageInternalServerError() *DeleteMessageInternalServerError {
	return &DeleteMessageInternalServerError{}
}

/*DeleteMessageInternalServerError handles this case with default header values.

Internal
*/
type DeleteMessageInternalServerError struct {
	Payload *models.Error
}

func (o *DeleteMessageInternalServerError) Error() string {
	return fmt.Sprintf("[GET /bot{token}/deleteMessage][%d] deleteMessageInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteMessageInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
