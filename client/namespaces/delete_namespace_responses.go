

package namespaces




import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/bkuschel/ibm-container-registry/models"
)

// DeleteNamespaceReader is a Reader for the DeleteNamespace structure.
type DeleteNamespaceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteNamespaceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteNamespaceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 204:
		result := NewDeleteNamespaceNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteNamespaceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewDeleteNamespaceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewDeleteNamespaceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewDeleteNamespaceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewDeleteNamespaceServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteNamespaceOK creates a DeleteNamespaceOK with default headers values
func NewDeleteNamespaceOK() *DeleteNamespaceOK {
	return &DeleteNamespaceOK{}
}

/*DeleteNamespaceOK handles this case with default header values.

No response was specified
*/
type DeleteNamespaceOK struct {
	Payload string
}

func (o *DeleteNamespaceOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/namespaces/{namespace}][%d] deleteNamespaceOK  %+v", 200, o.Payload)
}

func (o *DeleteNamespaceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteNamespaceNoContent creates a DeleteNamespaceNoContent with default headers values
func NewDeleteNamespaceNoContent() *DeleteNamespaceNoContent {
	return &DeleteNamespaceNoContent{}
}

/*DeleteNamespaceNoContent handles this case with default header values.

Description was not specified
*/
type DeleteNamespaceNoContent struct {
	Payload string
}

func (o *DeleteNamespaceNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/namespaces/{namespace}][%d] deleteNamespaceNoContent  %+v", 204, o.Payload)
}

func (o *DeleteNamespaceNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteNamespaceBadRequest creates a DeleteNamespaceBadRequest with default headers values
func NewDeleteNamespaceBadRequest() *DeleteNamespaceBadRequest {
	return &DeleteNamespaceBadRequest{}
}

/*DeleteNamespaceBadRequest handles this case with default header values.

A required header is missing. Correct the request and try again.
*/
type DeleteNamespaceBadRequest struct {
	Payload *models.RequestError
}

func (o *DeleteNamespaceBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/namespaces/{namespace}][%d] deleteNamespaceBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteNamespaceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteNamespaceUnauthorized creates a DeleteNamespaceUnauthorized with default headers values
func NewDeleteNamespaceUnauthorized() *DeleteNamespaceUnauthorized {
	return &DeleteNamespaceUnauthorized{}
}

/*DeleteNamespaceUnauthorized handles this case with default header values.

Your IBM Cloud bearer token is not valid for the targeted account. Correct the request and try again.
*/
type DeleteNamespaceUnauthorized struct {
	Payload *models.RequestError
}

func (o *DeleteNamespaceUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/namespaces/{namespace}][%d] deleteNamespaceUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteNamespaceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteNamespaceForbidden creates a DeleteNamespaceForbidden with default headers values
func NewDeleteNamespaceForbidden() *DeleteNamespaceForbidden {
	return &DeleteNamespaceForbidden{}
}

/*DeleteNamespaceForbidden handles this case with default header values.

The specified namespace is not owned by the targeted account.
*/
type DeleteNamespaceForbidden struct {
	Payload *models.RequestError
}

func (o *DeleteNamespaceForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/namespaces/{namespace}][%d] deleteNamespaceForbidden  %+v", 403, o.Payload)
}

func (o *DeleteNamespaceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteNamespaceInternalServerError creates a DeleteNamespaceInternalServerError with default headers values
func NewDeleteNamespaceInternalServerError() *DeleteNamespaceInternalServerError {
	return &DeleteNamespaceInternalServerError{}
}

/*DeleteNamespaceInternalServerError handles this case with default header values.

Internal server error
*/
type DeleteNamespaceInternalServerError struct {
	Payload *models.RequestError
}

func (o *DeleteNamespaceInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/namespaces/{namespace}][%d] deleteNamespaceInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteNamespaceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteNamespaceServiceUnavailable creates a DeleteNamespaceServiceUnavailable with default headers values
func NewDeleteNamespaceServiceUnavailable() *DeleteNamespaceServiceUnavailable {
	return &DeleteNamespaceServiceUnavailable{}
}

/*DeleteNamespaceServiceUnavailable handles this case with default header values.

Unable to authenticate with IBM Cloud. Try again later.
*/
type DeleteNamespaceServiceUnavailable struct {
	Payload *models.RequestError
}

func (o *DeleteNamespaceServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/namespaces/{namespace}][%d] deleteNamespaceServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteNamespaceServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
