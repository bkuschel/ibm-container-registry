

package tokens




import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/bkuschel/ibm-container-registry/models"
)

// TokenRetrieveReader is a Reader for the TokenRetrieve structure.
type TokenRetrieveReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TokenRetrieveReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewTokenRetrieveOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewTokenRetrieveBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewTokenRetrieveUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewTokenRetrieveNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewTokenRetrieveInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewTokenRetrieveServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewTokenRetrieveOK creates a TokenRetrieveOK with default headers values
func NewTokenRetrieveOK() *TokenRetrieveOK {
	return &TokenRetrieveOK{}
}

/*TokenRetrieveOK handles this case with default header values.

Description was not specified
*/
type TokenRetrieveOK struct {
	Payload []*models.Token
}

func (o *TokenRetrieveOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/tokens/{uuid}][%d] tokenRetrieveOK  %+v", 200, o.Payload)
}

func (o *TokenRetrieveOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTokenRetrieveBadRequest creates a TokenRetrieveBadRequest with default headers values
func NewTokenRetrieveBadRequest() *TokenRetrieveBadRequest {
	return &TokenRetrieveBadRequest{}
}

/*TokenRetrieveBadRequest handles this case with default header values.

A required header is missing. Correct the request and try again.
*/
type TokenRetrieveBadRequest struct {
	Payload *models.RequestError
}

func (o *TokenRetrieveBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v1/tokens/{uuid}][%d] tokenRetrieveBadRequest  %+v", 400, o.Payload)
}

func (o *TokenRetrieveBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTokenRetrieveUnauthorized creates a TokenRetrieveUnauthorized with default headers values
func NewTokenRetrieveUnauthorized() *TokenRetrieveUnauthorized {
	return &TokenRetrieveUnauthorized{}
}

/*TokenRetrieveUnauthorized handles this case with default header values.

You are not authorized to view the requested resource, or your IBM Cloud bearer token is invalid. Correct the request and try again.
*/
type TokenRetrieveUnauthorized struct {
	Payload *models.RequestError
}

func (o *TokenRetrieveUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/tokens/{uuid}][%d] tokenRetrieveUnauthorized  %+v", 401, o.Payload)
}

func (o *TokenRetrieveUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTokenRetrieveNotFound creates a TokenRetrieveNotFound with default headers values
func NewTokenRetrieveNotFound() *TokenRetrieveNotFound {
	return &TokenRetrieveNotFound{}
}

/*TokenRetrieveNotFound handles this case with default header values.

No token with the specific ID was found in your account. Correct the request and try again.
*/
type TokenRetrieveNotFound struct {
	Payload *models.RequestError
}

func (o *TokenRetrieveNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/tokens/{uuid}][%d] tokenRetrieveNotFound  %+v", 404, o.Payload)
}

func (o *TokenRetrieveNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTokenRetrieveInternalServerError creates a TokenRetrieveInternalServerError with default headers values
func NewTokenRetrieveInternalServerError() *TokenRetrieveInternalServerError {
	return &TokenRetrieveInternalServerError{}
}

/*TokenRetrieveInternalServerError handles this case with default header values.

Internal server error
*/
type TokenRetrieveInternalServerError struct {
	Payload *models.RequestError
}

func (o *TokenRetrieveInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1/tokens/{uuid}][%d] tokenRetrieveInternalServerError  %+v", 500, o.Payload)
}

func (o *TokenRetrieveInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTokenRetrieveServiceUnavailable creates a TokenRetrieveServiceUnavailable with default headers values
func NewTokenRetrieveServiceUnavailable() *TokenRetrieveServiceUnavailable {
	return &TokenRetrieveServiceUnavailable{}
}

/*TokenRetrieveServiceUnavailable handles this case with default header values.

Unable to authenticate with IBM Cloud. Try again later.
*/
type TokenRetrieveServiceUnavailable struct {
	Payload *models.RequestError
}

func (o *TokenRetrieveServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v1/tokens/{uuid}][%d] tokenRetrieveServiceUnavailable  %+v", 503, o.Payload)
}

func (o *TokenRetrieveServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
