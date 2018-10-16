

package auth




import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/bkuschel/ibm-container-registry/models"
)

// SetAuthReader is a Reader for the SetAuth structure.
type SetAuthReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetAuthReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSetAuthOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 204:
		result := NewSetAuthNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewSetAuthBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewSetAuthForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewSetAuthInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewSetAuthServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSetAuthOK creates a SetAuthOK with default headers values
func NewSetAuthOK() *SetAuthOK {
	return &SetAuthOK{}
}

/*SetAuthOK handles this case with default header values.

No response was specified
*/
type SetAuthOK struct {
	Payload string
}

func (o *SetAuthOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/auth][%d] setAuthOK  %+v", 200, o.Payload)
}

func (o *SetAuthOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetAuthNoContent creates a SetAuthNoContent with default headers values
func NewSetAuthNoContent() *SetAuthNoContent {
	return &SetAuthNoContent{}
}

/*SetAuthNoContent handles this case with default header values.

Description was not specified
*/
type SetAuthNoContent struct {
	Payload string
}

func (o *SetAuthNoContent) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/auth][%d] setAuthNoContent  %+v", 204, o.Payload)
}

func (o *SetAuthNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetAuthBadRequest creates a SetAuthBadRequest with default headers values
func NewSetAuthBadRequest() *SetAuthBadRequest {
	return &SetAuthBadRequest{}
}

/*SetAuthBadRequest handles this case with default header values.

A required header is missing. Correct the request and try again.
*/
type SetAuthBadRequest struct {
	Payload *models.RequestError
}

func (o *SetAuthBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/auth][%d] setAuthBadRequest  %+v", 400, o.Payload)
}

func (o *SetAuthBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetAuthForbidden creates a SetAuthForbidden with default headers values
func NewSetAuthForbidden() *SetAuthForbidden {
	return &SetAuthForbidden{}
}

/*SetAuthForbidden handles this case with default header values.

You do not have access to perform the requested operation. You must have either Manager or Billing Manager roles on the targeted account to perform this action.
*/
type SetAuthForbidden struct {
	Payload *models.RequestError
}

func (o *SetAuthForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/auth][%d] setAuthForbidden  %+v", 403, o.Payload)
}

func (o *SetAuthForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetAuthInternalServerError creates a SetAuthInternalServerError with default headers values
func NewSetAuthInternalServerError() *SetAuthInternalServerError {
	return &SetAuthInternalServerError{}
}

/*SetAuthInternalServerError handles this case with default header values.

Internal server error
*/
type SetAuthInternalServerError struct {
	Payload *models.RequestError
}

func (o *SetAuthInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/auth][%d] setAuthInternalServerError  %+v", 500, o.Payload)
}

func (o *SetAuthInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetAuthServiceUnavailable creates a SetAuthServiceUnavailable with default headers values
func NewSetAuthServiceUnavailable() *SetAuthServiceUnavailable {
	return &SetAuthServiceUnavailable{}
}

/*SetAuthServiceUnavailable handles this case with default header values.

Unable to authenticate with IBM Cloud. Try again later.
*/
type SetAuthServiceUnavailable struct {
	Payload *models.RequestError
}

func (o *SetAuthServiceUnavailable) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/auth][%d] setAuthServiceUnavailable  %+v", 503, o.Payload)
}

func (o *SetAuthServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
