

package tokens




import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/bkuschel/ibm-container-registry/models"
)

// TokenDeleteOrgReader is a Reader for the TokenDeleteOrg structure.
type TokenDeleteOrgReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TokenDeleteOrgReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewTokenDeleteOrgOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 204:
		result := NewTokenDeleteOrgNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewTokenDeleteOrgBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewTokenDeleteOrgUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewTokenDeleteOrgInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewTokenDeleteOrgServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewTokenDeleteOrgOK creates a TokenDeleteOrgOK with default headers values
func NewTokenDeleteOrgOK() *TokenDeleteOrgOK {
	return &TokenDeleteOrgOK{}
}

/*TokenDeleteOrgOK handles this case with default header values.

No response was specified
*/
type TokenDeleteOrgOK struct {
	Payload string
}

func (o *TokenDeleteOrgOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/tokens][%d] tokenDeleteOrgOK  %+v", 200, o.Payload)
}

func (o *TokenDeleteOrgOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTokenDeleteOrgNoContent creates a TokenDeleteOrgNoContent with default headers values
func NewTokenDeleteOrgNoContent() *TokenDeleteOrgNoContent {
	return &TokenDeleteOrgNoContent{}
}

/*TokenDeleteOrgNoContent handles this case with default header values.

Description was not specified
*/
type TokenDeleteOrgNoContent struct {
	Payload string
}

func (o *TokenDeleteOrgNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/tokens][%d] tokenDeleteOrgNoContent  %+v", 204, o.Payload)
}

func (o *TokenDeleteOrgNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTokenDeleteOrgBadRequest creates a TokenDeleteOrgBadRequest with default headers values
func NewTokenDeleteOrgBadRequest() *TokenDeleteOrgBadRequest {
	return &TokenDeleteOrgBadRequest{}
}

/*TokenDeleteOrgBadRequest handles this case with default header values.

A required parameter is missing. Correct the request and try again.
*/
type TokenDeleteOrgBadRequest struct {
	Payload *models.RequestError
}

func (o *TokenDeleteOrgBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/tokens][%d] tokenDeleteOrgBadRequest  %+v", 400, o.Payload)
}

func (o *TokenDeleteOrgBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTokenDeleteOrgUnauthorized creates a TokenDeleteOrgUnauthorized with default headers values
func NewTokenDeleteOrgUnauthorized() *TokenDeleteOrgUnauthorized {
	return &TokenDeleteOrgUnauthorized{}
}

/*TokenDeleteOrgUnauthorized handles this case with default header values.

You are not authorized to view the requested resource, or your IBM Cloud bearer token is invalid. Correct the request and try again.
*/
type TokenDeleteOrgUnauthorized struct {
	Payload *models.RequestError
}

func (o *TokenDeleteOrgUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/tokens][%d] tokenDeleteOrgUnauthorized  %+v", 401, o.Payload)
}

func (o *TokenDeleteOrgUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTokenDeleteOrgInternalServerError creates a TokenDeleteOrgInternalServerError with default headers values
func NewTokenDeleteOrgInternalServerError() *TokenDeleteOrgInternalServerError {
	return &TokenDeleteOrgInternalServerError{}
}

/*TokenDeleteOrgInternalServerError handles this case with default header values.

Internal server error
*/
type TokenDeleteOrgInternalServerError struct {
	Payload *models.RequestError
}

func (o *TokenDeleteOrgInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/tokens][%d] tokenDeleteOrgInternalServerError  %+v", 500, o.Payload)
}

func (o *TokenDeleteOrgInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTokenDeleteOrgServiceUnavailable creates a TokenDeleteOrgServiceUnavailable with default headers values
func NewTokenDeleteOrgServiceUnavailable() *TokenDeleteOrgServiceUnavailable {
	return &TokenDeleteOrgServiceUnavailable{}
}

/*TokenDeleteOrgServiceUnavailable handles this case with default header values.

Unable to authenticate with IBM Cloud. Try again later.
*/
type TokenDeleteOrgServiceUnavailable struct {
	Payload *models.RequestError
}

func (o *TokenDeleteOrgServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/tokens][%d] tokenDeleteOrgServiceUnavailable  %+v", 503, o.Payload)
}

func (o *TokenDeleteOrgServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
