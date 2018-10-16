

package namespaces




import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/bkuschel/ibm-container-registry/models"
)

// GetNamespaceReader is a Reader for the GetNamespace structure.
type GetNamespaceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNamespaceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetNamespaceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetNamespaceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetNamespaceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetNamespaceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewGetNamespaceServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetNamespaceOK creates a GetNamespaceOK with default headers values
func NewGetNamespaceOK() *GetNamespaceOK {
	return &GetNamespaceOK{}
}

/*GetNamespaceOK handles this case with default header values.

Description was not specified
*/
type GetNamespaceOK struct {
	Payload []string
}

func (o *GetNamespaceOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/namespaces][%d] getNamespaceOK  %+v", 200, o.Payload)
}

func (o *GetNamespaceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNamespaceBadRequest creates a GetNamespaceBadRequest with default headers values
func NewGetNamespaceBadRequest() *GetNamespaceBadRequest {
	return &GetNamespaceBadRequest{}
}

/*GetNamespaceBadRequest handles this case with default header values.

A required header is missing. Correct the request and try again.
*/
type GetNamespaceBadRequest struct {
	Payload *models.RequestError
}

func (o *GetNamespaceBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v1/namespaces][%d] getNamespaceBadRequest  %+v", 400, o.Payload)
}

func (o *GetNamespaceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNamespaceUnauthorized creates a GetNamespaceUnauthorized with default headers values
func NewGetNamespaceUnauthorized() *GetNamespaceUnauthorized {
	return &GetNamespaceUnauthorized{}
}

/*GetNamespaceUnauthorized handles this case with default header values.

You are not authorized to view the requested resource, or your IBM Cloud bearer token is invalid. Correct the request and try again.
*/
type GetNamespaceUnauthorized struct {
	Payload *models.RequestError
}

func (o *GetNamespaceUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/namespaces][%d] getNamespaceUnauthorized  %+v", 401, o.Payload)
}

func (o *GetNamespaceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNamespaceInternalServerError creates a GetNamespaceInternalServerError with default headers values
func NewGetNamespaceInternalServerError() *GetNamespaceInternalServerError {
	return &GetNamespaceInternalServerError{}
}

/*GetNamespaceInternalServerError handles this case with default header values.

Internal server error
*/
type GetNamespaceInternalServerError struct {
	Payload *models.RequestError
}

func (o *GetNamespaceInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1/namespaces][%d] getNamespaceInternalServerError  %+v", 500, o.Payload)
}

func (o *GetNamespaceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNamespaceServiceUnavailable creates a GetNamespaceServiceUnavailable with default headers values
func NewGetNamespaceServiceUnavailable() *GetNamespaceServiceUnavailable {
	return &GetNamespaceServiceUnavailable{}
}

/*GetNamespaceServiceUnavailable handles this case with default header values.

Unable to authenticate with IBM Cloud. Try again later.
*/
type GetNamespaceServiceUnavailable struct {
	Payload *models.RequestError
}

func (o *GetNamespaceServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v1/namespaces][%d] getNamespaceServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetNamespaceServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
