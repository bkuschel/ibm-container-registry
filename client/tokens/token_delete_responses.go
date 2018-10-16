

package tokens




import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/bkuschel/ibm-container-registry/models"
)

// TokenDeleteReader is a Reader for the TokenDelete structure.
type TokenDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TokenDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewTokenDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 204:
		result := NewTokenDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewTokenDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewTokenDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewTokenDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewTokenDeleteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewTokenDeleteServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewTokenDeleteOK creates a TokenDeleteOK with default headers values
func NewTokenDeleteOK() *TokenDeleteOK {
	return &TokenDeleteOK{}
}

/*TokenDeleteOK handles this case with default header values.

No response was specified
*/
type TokenDeleteOK struct {
	Payload string
}

func (o *TokenDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/tokens/{token}][%d] tokenDeleteOK  %+v", 200, o.Payload)
}

func (o *TokenDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTokenDeleteNoContent creates a TokenDeleteNoContent with default headers values
func NewTokenDeleteNoContent() *TokenDeleteNoContent {
	return &TokenDeleteNoContent{}
}

/*TokenDeleteNoContent handles this case with default header values.

Description was not specified
*/
type TokenDeleteNoContent struct {
	Payload string
}

func (o *TokenDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/tokens/{token}][%d] tokenDeleteNoContent  %+v", 204, o.Payload)
}

func (o *TokenDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTokenDeleteBadRequest creates a TokenDeleteBadRequest with default headers values
func NewTokenDeleteBadRequest() *TokenDeleteBadRequest {
	return &TokenDeleteBadRequest{}
}

/*TokenDeleteBadRequest handles this case with default header values.

A required header is missing. Correct the request and try again.
*/
type TokenDeleteBadRequest struct {
	Payload *models.RequestError
}

func (o *TokenDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/tokens/{token}][%d] tokenDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *TokenDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTokenDeleteUnauthorized creates a TokenDeleteUnauthorized with default headers values
func NewTokenDeleteUnauthorized() *TokenDeleteUnauthorized {
	return &TokenDeleteUnauthorized{}
}

/*TokenDeleteUnauthorized handles this case with default header values.

You are not authorized to view the requested resource, or your IBM Cloud bearer token is invalid. Correct the request and try again.
*/
type TokenDeleteUnauthorized struct {
	Payload *models.RequestError
}

func (o *TokenDeleteUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/tokens/{token}][%d] tokenDeleteUnauthorized  %+v", 401, o.Payload)
}

func (o *TokenDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTokenDeleteNotFound creates a TokenDeleteNotFound with default headers values
func NewTokenDeleteNotFound() *TokenDeleteNotFound {
	return &TokenDeleteNotFound{}
}

/*TokenDeleteNotFound handles this case with default header values.

The specific token was not found in your account. Correct the request and try again.
*/
type TokenDeleteNotFound struct {
	Payload *models.RequestError
}

func (o *TokenDeleteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/tokens/{token}][%d] tokenDeleteNotFound  %+v", 404, o.Payload)
}

func (o *TokenDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTokenDeleteInternalServerError creates a TokenDeleteInternalServerError with default headers values
func NewTokenDeleteInternalServerError() *TokenDeleteInternalServerError {
	return &TokenDeleteInternalServerError{}
}

/*TokenDeleteInternalServerError handles this case with default header values.

Internal server error
*/
type TokenDeleteInternalServerError struct {
	Payload *models.RequestError
}

func (o *TokenDeleteInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/tokens/{token}][%d] tokenDeleteInternalServerError  %+v", 500, o.Payload)
}

func (o *TokenDeleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTokenDeleteServiceUnavailable creates a TokenDeleteServiceUnavailable with default headers values
func NewTokenDeleteServiceUnavailable() *TokenDeleteServiceUnavailable {
	return &TokenDeleteServiceUnavailable{}
}

/*TokenDeleteServiceUnavailable handles this case with default header values.

Unable to authenticate with IBM Cloud. Try again later.
*/
type TokenDeleteServiceUnavailable struct {
	Payload *models.RequestError
}

func (o *TokenDeleteServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/tokens/{token}][%d] tokenDeleteServiceUnavailable  %+v", 503, o.Payload)
}

func (o *TokenDeleteServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
