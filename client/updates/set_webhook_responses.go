package updates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/olebedev/go-tgbot/models"
)

// SetWebhookReader is a Reader for the SetWebhook structure.
type SetWebhookReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetWebhookReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSetWebhookOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewSetWebhookBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewSetWebhookUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewSetWebhookForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewSetWebhookNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 420:
		result := NewSetWebhookEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewSetWebhookInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSetWebhookOK creates a SetWebhookOK with default headers values
func NewSetWebhookOK() *SetWebhookOK {
	return &SetWebhookOK{}
}

/*SetWebhookOK handles this case with default header values.

Is OK?
*/
type SetWebhookOK struct {
	Payload *models.ResponseBool
}

func (o *SetWebhookOK) Error() string {
	return fmt.Sprintf("[POST /bot{token}/setWebhook][%d] setWebhookOK  %+v", 200, o.Payload)
}

func (o *SetWebhookOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseBool)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetWebhookBadRequest creates a SetWebhookBadRequest with default headers values
func NewSetWebhookBadRequest() *SetWebhookBadRequest {
	return &SetWebhookBadRequest{}
}

/*SetWebhookBadRequest handles this case with default header values.

Bad Request
*/
type SetWebhookBadRequest struct {
	Payload *models.Error
}

func (o *SetWebhookBadRequest) Error() string {
	return fmt.Sprintf("[POST /bot{token}/setWebhook][%d] setWebhookBadRequest  %+v", 400, o.Payload)
}

func (o *SetWebhookBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetWebhookUnauthorized creates a SetWebhookUnauthorized with default headers values
func NewSetWebhookUnauthorized() *SetWebhookUnauthorized {
	return &SetWebhookUnauthorized{}
}

/*SetWebhookUnauthorized handles this case with default header values.

Unauthorized
*/
type SetWebhookUnauthorized struct {
	Payload *models.Error
}

func (o *SetWebhookUnauthorized) Error() string {
	return fmt.Sprintf("[POST /bot{token}/setWebhook][%d] setWebhookUnauthorized  %+v", 401, o.Payload)
}

func (o *SetWebhookUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetWebhookForbidden creates a SetWebhookForbidden with default headers values
func NewSetWebhookForbidden() *SetWebhookForbidden {
	return &SetWebhookForbidden{}
}

/*SetWebhookForbidden handles this case with default header values.

Forbidden
*/
type SetWebhookForbidden struct {
	Payload *models.Error
}

func (o *SetWebhookForbidden) Error() string {
	return fmt.Sprintf("[POST /bot{token}/setWebhook][%d] setWebhookForbidden  %+v", 403, o.Payload)
}

func (o *SetWebhookForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetWebhookNotFound creates a SetWebhookNotFound with default headers values
func NewSetWebhookNotFound() *SetWebhookNotFound {
	return &SetWebhookNotFound{}
}

/*SetWebhookNotFound handles this case with default header values.

Not Found
*/
type SetWebhookNotFound struct {
	Payload *models.Error
}

func (o *SetWebhookNotFound) Error() string {
	return fmt.Sprintf("[POST /bot{token}/setWebhook][%d] setWebhookNotFound  %+v", 404, o.Payload)
}

func (o *SetWebhookNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetWebhookEnhanceYourCalm creates a SetWebhookEnhanceYourCalm with default headers values
func NewSetWebhookEnhanceYourCalm() *SetWebhookEnhanceYourCalm {
	return &SetWebhookEnhanceYourCalm{}
}

/*SetWebhookEnhanceYourCalm handles this case with default header values.

Flood
*/
type SetWebhookEnhanceYourCalm struct {
	Payload *models.Error
}

func (o *SetWebhookEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[POST /bot{token}/setWebhook][%d] setWebhookEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *SetWebhookEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetWebhookInternalServerError creates a SetWebhookInternalServerError with default headers values
func NewSetWebhookInternalServerError() *SetWebhookInternalServerError {
	return &SetWebhookInternalServerError{}
}

/*SetWebhookInternalServerError handles this case with default header values.

Internal
*/
type SetWebhookInternalServerError struct {
	Payload *models.Error
}

func (o *SetWebhookInternalServerError) Error() string {
	return fmt.Sprintf("[POST /bot{token}/setWebhook][%d] setWebhookInternalServerError  %+v", 500, o.Payload)
}

func (o *SetWebhookInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
