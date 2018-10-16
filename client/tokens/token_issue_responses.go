

package tokens




import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/bkuschel/ibm-container-registry/models"
)

// TokenIssueReader is a Reader for the TokenIssue structure.
type TokenIssueReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TokenIssueReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewTokenIssueOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 201:
		result := NewTokenIssueCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewTokenIssueBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewTokenIssueUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewTokenIssueInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewTokenIssueServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewTokenIssueOK creates a TokenIssueOK with default headers values
func NewTokenIssueOK() *TokenIssueOK {
	return &TokenIssueOK{}
}

/*TokenIssueOK handles this case with default header values.

No response was specified
*/
type TokenIssueOK struct {
	Payload string
}

func (o *TokenIssueOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/tokens][%d] tokenIssueOK  %+v", 200, o.Payload)
}

func (o *TokenIssueOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTokenIssueCreated creates a TokenIssueCreated with default headers values
func NewTokenIssueCreated() *TokenIssueCreated {
	return &TokenIssueCreated{}
}

/*TokenIssueCreated handles this case with default header values.

Description was not specified
*/
type TokenIssueCreated struct {
	Payload *models.Token
}

func (o *TokenIssueCreated) Error() string {
	return fmt.Sprintf("[POST /api/v1/tokens][%d] tokenIssueCreated  %+v", 201, o.Payload)
}

func (o *TokenIssueCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Token)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTokenIssueBadRequest creates a TokenIssueBadRequest with default headers values
func NewTokenIssueBadRequest() *TokenIssueBadRequest {
	return &TokenIssueBadRequest{}
}

/*TokenIssueBadRequest handles this case with default header values.

A required header is missing. Correct the request and try again.
*/
type TokenIssueBadRequest struct {
	Payload *models.RequestError
}

func (o *TokenIssueBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v1/tokens][%d] tokenIssueBadRequest  %+v", 400, o.Payload)
}

func (o *TokenIssueBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTokenIssueUnauthorized creates a TokenIssueUnauthorized with default headers values
func NewTokenIssueUnauthorized() *TokenIssueUnauthorized {
	return &TokenIssueUnauthorized{}
}

/*TokenIssueUnauthorized handles this case with default header values.

You are not authorized to view the requested resource, or your IBM Cloud bearer token is invalid. Correct the request and try again.
*/
type TokenIssueUnauthorized struct {
	Payload *models.RequestError
}

func (o *TokenIssueUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v1/tokens][%d] tokenIssueUnauthorized  %+v", 401, o.Payload)
}

func (o *TokenIssueUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTokenIssueInternalServerError creates a TokenIssueInternalServerError with default headers values
func NewTokenIssueInternalServerError() *TokenIssueInternalServerError {
	return &TokenIssueInternalServerError{}
}

/*TokenIssueInternalServerError handles this case with default header values.

Internal server error
*/
type TokenIssueInternalServerError struct {
	Payload *models.RequestError
}

func (o *TokenIssueInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v1/tokens][%d] tokenIssueInternalServerError  %+v", 500, o.Payload)
}

func (o *TokenIssueInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTokenIssueServiceUnavailable creates a TokenIssueServiceUnavailable with default headers values
func NewTokenIssueServiceUnavailable() *TokenIssueServiceUnavailable {
	return &TokenIssueServiceUnavailable{}
}

/*TokenIssueServiceUnavailable handles this case with default header values.

Unable to authenticate with IBM Cloud. Try again later.
*/
type TokenIssueServiceUnavailable struct {
	Payload *models.RequestError
}

func (o *TokenIssueServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v1/tokens][%d] tokenIssueServiceUnavailable  %+v", 503, o.Payload)
}

func (o *TokenIssueServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
