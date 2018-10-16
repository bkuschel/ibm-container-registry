

package images




import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/bkuschel/ibm-container-registry/models"
)

// ImageInspectReader is a Reader for the ImageInspect structure.
type ImageInspectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ImageInspectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewImageInspectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewImageInspectBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewImageInspectUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewImageInspectForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewImageInspectNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewImageInspectInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewImageInspectServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewImageInspectOK creates a ImageInspectOK with default headers values
func NewImageInspectOK() *ImageInspectOK {
	return &ImageInspectOK{}
}

/*ImageInspectOK handles this case with default header values.

Description was not specified
*/
type ImageInspectOK struct {
	Payload *models.ImageInspection
}

func (o *ImageInspectOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/images/{image}/json][%d] imageInspectOK  %+v", 200, o.Payload)
}

func (o *ImageInspectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ImageInspection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageInspectBadRequest creates a ImageInspectBadRequest with default headers values
func NewImageInspectBadRequest() *ImageInspectBadRequest {
	return &ImageInspectBadRequest{}
}

/*ImageInspectBadRequest handles this case with default header values.

A required header is missing. Correct the request and try again.
*/
type ImageInspectBadRequest struct {
	Payload *models.RequestError
}

func (o *ImageInspectBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v1/images/{image}/json][%d] imageInspectBadRequest  %+v", 400, o.Payload)
}

func (o *ImageInspectBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageInspectUnauthorized creates a ImageInspectUnauthorized with default headers values
func NewImageInspectUnauthorized() *ImageInspectUnauthorized {
	return &ImageInspectUnauthorized{}
}

/*ImageInspectUnauthorized handles this case with default header values.

You are not authorized to view the requested resource, or your IBM Cloud bearer token is invalid. Correct the request and try again.
*/
type ImageInspectUnauthorized struct {
	Payload *models.RequestError
}

func (o *ImageInspectUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/images/{image}/json][%d] imageInspectUnauthorized  %+v", 401, o.Payload)
}

func (o *ImageInspectUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageInspectForbidden creates a ImageInspectForbidden with default headers values
func NewImageInspectForbidden() *ImageInspectForbidden {
	return &ImageInspectForbidden{}
}

/*ImageInspectForbidden handles this case with default header values.

You do not have authorization to access the specified namespace.
*/
type ImageInspectForbidden struct {
	Payload *models.RequestError
}

func (o *ImageInspectForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/images/{image}/json][%d] imageInspectForbidden  %+v", 403, o.Payload)
}

func (o *ImageInspectForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageInspectNotFound creates a ImageInspectNotFound with default headers values
func NewImageInspectNotFound() *ImageInspectNotFound {
	return &ImageInspectNotFound{}
}

/*ImageInspectNotFound handles this case with default header values.

The specified image was not found. Check the spelling of the specified image.
*/
type ImageInspectNotFound struct {
	Payload *models.RequestError
}

func (o *ImageInspectNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/images/{image}/json][%d] imageInspectNotFound  %+v", 404, o.Payload)
}

func (o *ImageInspectNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageInspectInternalServerError creates a ImageInspectInternalServerError with default headers values
func NewImageInspectInternalServerError() *ImageInspectInternalServerError {
	return &ImageInspectInternalServerError{}
}

/*ImageInspectInternalServerError handles this case with default header values.

Internal server error
*/
type ImageInspectInternalServerError struct {
	Payload *models.RequestError
}

func (o *ImageInspectInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1/images/{image}/json][%d] imageInspectInternalServerError  %+v", 500, o.Payload)
}

func (o *ImageInspectInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageInspectServiceUnavailable creates a ImageInspectServiceUnavailable with default headers values
func NewImageInspectServiceUnavailable() *ImageInspectServiceUnavailable {
	return &ImageInspectServiceUnavailable{}
}

/*ImageInspectServiceUnavailable handles this case with default header values.

Unable to authenticate with IBM Cloud. Try again later.
*/
type ImageInspectServiceUnavailable struct {
	Payload *models.RequestError
}

func (o *ImageInspectServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v1/images/{image}/json][%d] imageInspectServiceUnavailable  %+v", 503, o.Payload)
}

func (o *ImageInspectServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
