

package auth




import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/bkuschel/ibm-container-registry/models"
)

// GetAuthReader is a Reader for the GetAuth structure.
type GetAuthReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAuthReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAuthOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetAuthBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetAuthInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewGetAuthServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAuthOK creates a GetAuthOK with default headers values
func NewGetAuthOK() *GetAuthOK {
	return &GetAuthOK{}
}

/*GetAuthOK handles this case with default header values.

Description was not specified
*/
type GetAuthOK struct {
	Payload *models.AuthResponse
}

func (o *GetAuthOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/auth][%d] getAuthOK  %+v", 200, o.Payload)
}

func (o *GetAuthOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AuthResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthBadRequest creates a GetAuthBadRequest with default headers values
func NewGetAuthBadRequest() *GetAuthBadRequest {
	return &GetAuthBadRequest{}
}

/*GetAuthBadRequest handles this case with default header values.

A required header is missing. Correct the request and try again.
*/
type GetAuthBadRequest struct {
	Payload *models.RequestError
}

func (o *GetAuthBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v1/auth][%d] getAuthBadRequest  %+v", 400, o.Payload)
}

func (o *GetAuthBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthInternalServerError creates a GetAuthInternalServerError with default headers values
func NewGetAuthInternalServerError() *GetAuthInternalServerError {
	return &GetAuthInternalServerError{}
}

/*GetAuthInternalServerError handles this case with default header values.

Internal server error
*/
type GetAuthInternalServerError struct {
	Payload *models.RequestError
}

func (o *GetAuthInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1/auth][%d] getAuthInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAuthInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthServiceUnavailable creates a GetAuthServiceUnavailable with default headers values
func NewGetAuthServiceUnavailable() *GetAuthServiceUnavailable {
	return &GetAuthServiceUnavailable{}
}

/*GetAuthServiceUnavailable handles this case with default header values.

Unable to authenticate with IBM Cloud. Try again later.
*/
type GetAuthServiceUnavailable struct {
	Payload *models.RequestError
}

func (o *GetAuthServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v1/auth][%d] getAuthServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetAuthServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
