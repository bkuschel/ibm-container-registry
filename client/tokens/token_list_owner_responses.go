

package tokens




import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/bkuschel/ibm-container-registry/models"
)

// TokenListOwnerReader is a Reader for the TokenListOwner structure.
type TokenListOwnerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TokenListOwnerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewTokenListOwnerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewTokenListOwnerUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewTokenListOwnerInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewTokenListOwnerServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewTokenListOwnerOK creates a TokenListOwnerOK with default headers values
func NewTokenListOwnerOK() *TokenListOwnerOK {
	return &TokenListOwnerOK{}
}

/*TokenListOwnerOK handles this case with default header values.

Description was not specified
*/
type TokenListOwnerOK struct {
	Payload *models.ListTokensResp
}

func (o *TokenListOwnerOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/tokens][%d] tokenListOwnerOK  %+v", 200, o.Payload)
}

func (o *TokenListOwnerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListTokensResp)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTokenListOwnerUnauthorized creates a TokenListOwnerUnauthorized with default headers values
func NewTokenListOwnerUnauthorized() *TokenListOwnerUnauthorized {
	return &TokenListOwnerUnauthorized{}
}

/*TokenListOwnerUnauthorized handles this case with default header values.

You are not authorized to view the requested resource, or your IBM Cloud bearer token is invalid. Correct the request and try again.
*/
type TokenListOwnerUnauthorized struct {
	Payload *models.RequestError
}

func (o *TokenListOwnerUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/tokens][%d] tokenListOwnerUnauthorized  %+v", 401, o.Payload)
}

func (o *TokenListOwnerUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTokenListOwnerInternalServerError creates a TokenListOwnerInternalServerError with default headers values
func NewTokenListOwnerInternalServerError() *TokenListOwnerInternalServerError {
	return &TokenListOwnerInternalServerError{}
}

/*TokenListOwnerInternalServerError handles this case with default header values.

Internal server error
*/
type TokenListOwnerInternalServerError struct {
	Payload *models.RequestError
}

func (o *TokenListOwnerInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1/tokens][%d] tokenListOwnerInternalServerError  %+v", 500, o.Payload)
}

func (o *TokenListOwnerInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTokenListOwnerServiceUnavailable creates a TokenListOwnerServiceUnavailable with default headers values
func NewTokenListOwnerServiceUnavailable() *TokenListOwnerServiceUnavailable {
	return &TokenListOwnerServiceUnavailable{}
}

/*TokenListOwnerServiceUnavailable handles this case with default header values.

Unable to authenticate with IBM Cloud. Try again later.
*/
type TokenListOwnerServiceUnavailable struct {
	Payload *models.RequestError
}

func (o *TokenListOwnerServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v1/tokens][%d] tokenListOwnerServiceUnavailable  %+v", 503, o.Payload)
}

func (o *TokenListOwnerServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
