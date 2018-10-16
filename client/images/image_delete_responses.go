

package images




import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/bkuschel/ibm-container-registry/models"
)

// ImageDeleteReader is a Reader for the ImageDelete structure.
type ImageDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ImageDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewImageDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewImageDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewImageDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewImageDeleteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewImageDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 415:
		result := NewImageDeleteUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewImageDeleteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewImageDeleteServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewImageDeleteOK creates a ImageDeleteOK with default headers values
func NewImageDeleteOK() *ImageDeleteOK {
	return &ImageDeleteOK{}
}

/*ImageDeleteOK handles this case with default header values.

Description was not specified
*/
type ImageDeleteOK struct {
	Payload *models.ImageDeleteResult
}

func (o *ImageDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/images/{image}][%d] imageDeleteOK  %+v", 200, o.Payload)
}

func (o *ImageDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ImageDeleteResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageDeleteBadRequest creates a ImageDeleteBadRequest with default headers values
func NewImageDeleteBadRequest() *ImageDeleteBadRequest {
	return &ImageDeleteBadRequest{}
}

/*ImageDeleteBadRequest handles this case with default header values.

A required header is missing. Correct the request and try again.
*/
type ImageDeleteBadRequest struct {
	Payload *models.RequestError
}

func (o *ImageDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/images/{image}][%d] imageDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *ImageDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageDeleteUnauthorized creates a ImageDeleteUnauthorized with default headers values
func NewImageDeleteUnauthorized() *ImageDeleteUnauthorized {
	return &ImageDeleteUnauthorized{}
}

/*ImageDeleteUnauthorized handles this case with default header values.

You are not authorized to view the requested resource, or your IBM Cloud bearer token is invalid. Correct the request and try again.
*/
type ImageDeleteUnauthorized struct {
	Payload *models.RequestError
}

func (o *ImageDeleteUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/images/{image}][%d] imageDeleteUnauthorized  %+v", 401, o.Payload)
}

func (o *ImageDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageDeleteForbidden creates a ImageDeleteForbidden with default headers values
func NewImageDeleteForbidden() *ImageDeleteForbidden {
	return &ImageDeleteForbidden{}
}

/*ImageDeleteForbidden handles this case with default header values.

You do not have authorization to access the specified namespace.
*/
type ImageDeleteForbidden struct {
	Payload *models.RequestError
}

func (o *ImageDeleteForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/images/{image}][%d] imageDeleteForbidden  %+v", 403, o.Payload)
}

func (o *ImageDeleteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageDeleteNotFound creates a ImageDeleteNotFound with default headers values
func NewImageDeleteNotFound() *ImageDeleteNotFound {
	return &ImageDeleteNotFound{}
}

/*ImageDeleteNotFound handles this case with default header values.

The specified image was not found. Check the spelling of the specified image.
*/
type ImageDeleteNotFound struct {
	Payload *models.RequestError
}

func (o *ImageDeleteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/images/{image}][%d] imageDeleteNotFound  %+v", 404, o.Payload)
}

func (o *ImageDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageDeleteUnsupportedMediaType creates a ImageDeleteUnsupportedMediaType with default headers values
func NewImageDeleteUnsupportedMediaType() *ImageDeleteUnsupportedMediaType {
	return &ImageDeleteUnsupportedMediaType{}
}

/*ImageDeleteUnsupportedMediaType handles this case with default header values.

Inspecting a manifest-list is not supported
*/
type ImageDeleteUnsupportedMediaType struct {
	Payload *models.RequestError
}

func (o *ImageDeleteUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/images/{image}][%d] imageDeleteUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *ImageDeleteUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageDeleteInternalServerError creates a ImageDeleteInternalServerError with default headers values
func NewImageDeleteInternalServerError() *ImageDeleteInternalServerError {
	return &ImageDeleteInternalServerError{}
}

/*ImageDeleteInternalServerError handles this case with default header values.

Internal server error
*/
type ImageDeleteInternalServerError struct {
	Payload *models.RequestError
}

func (o *ImageDeleteInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/images/{image}][%d] imageDeleteInternalServerError  %+v", 500, o.Payload)
}

func (o *ImageDeleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageDeleteServiceUnavailable creates a ImageDeleteServiceUnavailable with default headers values
func NewImageDeleteServiceUnavailable() *ImageDeleteServiceUnavailable {
	return &ImageDeleteServiceUnavailable{}
}

/*ImageDeleteServiceUnavailable handles this case with default header values.

Unable to authenticate with IBM Cloud. Try again later.
*/
type ImageDeleteServiceUnavailable struct {
	Payload *models.RequestError
}

func (o *ImageDeleteServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/images/{image}][%d] imageDeleteServiceUnavailable  %+v", 503, o.Payload)
}

func (o *ImageDeleteServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
