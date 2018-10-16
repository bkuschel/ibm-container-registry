

package namespaces




import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/bkuschel/ibm-container-registry/models"
)

// AddNamespaceReader is a Reader for the AddNamespace structure.
type AddNamespaceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddNamespaceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewAddNamespaceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 201:
		result := NewAddNamespaceCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewAddNamespaceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewAddNamespaceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewAddNamespaceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewAddNamespaceConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewAddNamespaceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewAddNamespaceServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAddNamespaceOK creates a AddNamespaceOK with default headers values
func NewAddNamespaceOK() *AddNamespaceOK {
	return &AddNamespaceOK{}
}

/*AddNamespaceOK handles this case with default header values.

The namespace is already assigned to your IBM Cloud account.
*/
type AddNamespaceOK struct {
	Payload *models.NsResponse
}

func (o *AddNamespaceOK) Error() string {
	return fmt.Sprintf("[PUT /api/v1/namespaces/{namespace}][%d] addNamespaceOK  %+v", 200, o.Payload)
}

func (o *AddNamespaceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddNamespaceCreated creates a AddNamespaceCreated with default headers values
func NewAddNamespaceCreated() *AddNamespaceCreated {
	return &AddNamespaceCreated{}
}

/*AddNamespaceCreated handles this case with default header values.

The namespace was successfully added to your IBM Cloud account.
*/
type AddNamespaceCreated struct {
	Payload *models.NsResponse
}

func (o *AddNamespaceCreated) Error() string {
	return fmt.Sprintf("[PUT /api/v1/namespaces/{namespace}][%d] addNamespaceCreated  %+v", 201, o.Payload)
}

func (o *AddNamespaceCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddNamespaceBadRequest creates a AddNamespaceBadRequest with default headers values
func NewAddNamespaceBadRequest() *AddNamespaceBadRequest {
	return &AddNamespaceBadRequest{}
}

/*AddNamespaceBadRequest handles this case with default header values.

CRN0007E: The requested namespace is invalid. Correct the request and try again.
*/
type AddNamespaceBadRequest struct {
	Payload *models.RequestError
}

func (o *AddNamespaceBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v1/namespaces/{namespace}][%d] addNamespaceBadRequest  %+v", 400, o.Payload)
}

func (o *AddNamespaceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddNamespaceUnauthorized creates a AddNamespaceUnauthorized with default headers values
func NewAddNamespaceUnauthorized() *AddNamespaceUnauthorized {
	return &AddNamespaceUnauthorized{}
}

/*AddNamespaceUnauthorized handles this case with default header values.

You are not authorized to view the requested resource, or your IBM Cloud bearer token is invalid. Correct the request and try again.
*/
type AddNamespaceUnauthorized struct {
	Payload *models.RequestError
}

func (o *AddNamespaceUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v1/namespaces/{namespace}][%d] addNamespaceUnauthorized  %+v", 401, o.Payload)
}

func (o *AddNamespaceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddNamespaceForbidden creates a AddNamespaceForbidden with default headers values
func NewAddNamespaceForbidden() *AddNamespaceForbidden {
	return &AddNamespaceForbidden{}
}

/*AddNamespaceForbidden handles this case with default header values.

CRN0008E: The requested namespace is blacklisted. Correct the request and try again.
*/
type AddNamespaceForbidden struct {
	Payload *models.RequestError
}

func (o *AddNamespaceForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v1/namespaces/{namespace}][%d] addNamespaceForbidden  %+v", 403, o.Payload)
}

func (o *AddNamespaceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddNamespaceConflict creates a AddNamespaceConflict with default headers values
func NewAddNamespaceConflict() *AddNamespaceConflict {
	return &AddNamespaceConflict{}
}

/*AddNamespaceConflict handles this case with default header values.

The requested namespace is already in use. Request a different namespace.
*/
type AddNamespaceConflict struct {
	Payload *models.RequestError
}

func (o *AddNamespaceConflict) Error() string {
	return fmt.Sprintf("[PUT /api/v1/namespaces/{namespace}][%d] addNamespaceConflict  %+v", 409, o.Payload)
}

func (o *AddNamespaceConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddNamespaceInternalServerError creates a AddNamespaceInternalServerError with default headers values
func NewAddNamespaceInternalServerError() *AddNamespaceInternalServerError {
	return &AddNamespaceInternalServerError{}
}

/*AddNamespaceInternalServerError handles this case with default header values.

Internal server error
*/
type AddNamespaceInternalServerError struct {
	Payload *models.RequestError
}

func (o *AddNamespaceInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v1/namespaces/{namespace}][%d] addNamespaceInternalServerError  %+v", 500, o.Payload)
}

func (o *AddNamespaceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddNamespaceServiceUnavailable creates a AddNamespaceServiceUnavailable with default headers values
func NewAddNamespaceServiceUnavailable() *AddNamespaceServiceUnavailable {
	return &AddNamespaceServiceUnavailable{}
}

/*AddNamespaceServiceUnavailable handles this case with default header values.

Unable to authenticate with IBM Cloud. Try again later.
*/
type AddNamespaceServiceUnavailable struct {
	Payload *models.RequestError
}

func (o *AddNamespaceServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /api/v1/namespaces/{namespace}][%d] addNamespaceServiceUnavailable  %+v", 503, o.Payload)
}

func (o *AddNamespaceServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
