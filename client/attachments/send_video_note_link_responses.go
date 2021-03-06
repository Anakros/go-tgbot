// Code generated by go-swagger; DO NOT EDIT.

package attachments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/olebedev/go-tgbot/models"
)

// SendVideoNoteLinkReader is a Reader for the SendVideoNoteLink structure.
type SendVideoNoteLinkReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SendVideoNoteLinkReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSendVideoNoteLinkOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewSendVideoNoteLinkBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewSendVideoNoteLinkUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewSendVideoNoteLinkForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewSendVideoNoteLinkNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 420:
		result := NewSendVideoNoteLinkEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewSendVideoNoteLinkInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSendVideoNoteLinkOK creates a SendVideoNoteLinkOK with default headers values
func NewSendVideoNoteLinkOK() *SendVideoNoteLinkOK {
	return &SendVideoNoteLinkOK{}
}

/*SendVideoNoteLinkOK handles this case with default header values.

SendVideoNoteLinkOK send video note link o k
*/
type SendVideoNoteLinkOK struct {
	Payload *models.ResponseMessage
}

func (o *SendVideoNoteLinkOK) Error() string {
	return fmt.Sprintf("[POST /bot{token}/sendVideoNote#link][%d] sendVideoNoteLinkOK  %+v", 200, o.Payload)
}

func (o *SendVideoNoteLinkOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSendVideoNoteLinkBadRequest creates a SendVideoNoteLinkBadRequest with default headers values
func NewSendVideoNoteLinkBadRequest() *SendVideoNoteLinkBadRequest {
	return &SendVideoNoteLinkBadRequest{}
}

/*SendVideoNoteLinkBadRequest handles this case with default header values.

Bad Request
*/
type SendVideoNoteLinkBadRequest struct {
	Payload *models.Error
}

func (o *SendVideoNoteLinkBadRequest) Error() string {
	return fmt.Sprintf("[POST /bot{token}/sendVideoNote#link][%d] sendVideoNoteLinkBadRequest  %+v", 400, o.Payload)
}

func (o *SendVideoNoteLinkBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSendVideoNoteLinkUnauthorized creates a SendVideoNoteLinkUnauthorized with default headers values
func NewSendVideoNoteLinkUnauthorized() *SendVideoNoteLinkUnauthorized {
	return &SendVideoNoteLinkUnauthorized{}
}

/*SendVideoNoteLinkUnauthorized handles this case with default header values.

Unauthorized
*/
type SendVideoNoteLinkUnauthorized struct {
	Payload *models.Error
}

func (o *SendVideoNoteLinkUnauthorized) Error() string {
	return fmt.Sprintf("[POST /bot{token}/sendVideoNote#link][%d] sendVideoNoteLinkUnauthorized  %+v", 401, o.Payload)
}

func (o *SendVideoNoteLinkUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSendVideoNoteLinkForbidden creates a SendVideoNoteLinkForbidden with default headers values
func NewSendVideoNoteLinkForbidden() *SendVideoNoteLinkForbidden {
	return &SendVideoNoteLinkForbidden{}
}

/*SendVideoNoteLinkForbidden handles this case with default header values.

Forbidden
*/
type SendVideoNoteLinkForbidden struct {
	Payload *models.Error
}

func (o *SendVideoNoteLinkForbidden) Error() string {
	return fmt.Sprintf("[POST /bot{token}/sendVideoNote#link][%d] sendVideoNoteLinkForbidden  %+v", 403, o.Payload)
}

func (o *SendVideoNoteLinkForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSendVideoNoteLinkNotFound creates a SendVideoNoteLinkNotFound with default headers values
func NewSendVideoNoteLinkNotFound() *SendVideoNoteLinkNotFound {
	return &SendVideoNoteLinkNotFound{}
}

/*SendVideoNoteLinkNotFound handles this case with default header values.

Not Found
*/
type SendVideoNoteLinkNotFound struct {
	Payload *models.Error
}

func (o *SendVideoNoteLinkNotFound) Error() string {
	return fmt.Sprintf("[POST /bot{token}/sendVideoNote#link][%d] sendVideoNoteLinkNotFound  %+v", 404, o.Payload)
}

func (o *SendVideoNoteLinkNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSendVideoNoteLinkEnhanceYourCalm creates a SendVideoNoteLinkEnhanceYourCalm with default headers values
func NewSendVideoNoteLinkEnhanceYourCalm() *SendVideoNoteLinkEnhanceYourCalm {
	return &SendVideoNoteLinkEnhanceYourCalm{}
}

/*SendVideoNoteLinkEnhanceYourCalm handles this case with default header values.

Flood
*/
type SendVideoNoteLinkEnhanceYourCalm struct {
	Payload *models.Error
}

func (o *SendVideoNoteLinkEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[POST /bot{token}/sendVideoNote#link][%d] sendVideoNoteLinkEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *SendVideoNoteLinkEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSendVideoNoteLinkInternalServerError creates a SendVideoNoteLinkInternalServerError with default headers values
func NewSendVideoNoteLinkInternalServerError() *SendVideoNoteLinkInternalServerError {
	return &SendVideoNoteLinkInternalServerError{}
}

/*SendVideoNoteLinkInternalServerError handles this case with default header values.

Internal
*/
type SendVideoNoteLinkInternalServerError struct {
	Payload *models.Error
}

func (o *SendVideoNoteLinkInternalServerError) Error() string {
	return fmt.Sprintf("[POST /bot{token}/sendVideoNote#link][%d] sendVideoNoteLinkInternalServerError  %+v", 500, o.Payload)
}

func (o *SendVideoNoteLinkInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
