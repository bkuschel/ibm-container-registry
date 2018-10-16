

package images




import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/bkuschel/ibm-container-registry/models"
)

// ImageListReader is a Reader for the ImageList structure.
type ImageListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ImageListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewImageListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewImageListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewImageListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewImageListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewImageListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewImageListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewImageListServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewImageListOK creates a ImageListOK with default headers values
func NewImageListOK() *ImageListOK {
	return &ImageListOK{}
}

/*ImageListOK handles this case with default header values.

Description was not specified
*/
type ImageListOK struct {
	Payload []*models.RemoteAPIImage
}

func (o *ImageListOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/images][%d] imageListOK  %+v", 200, o.Payload)
}

func (o *ImageListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageListBadRequest creates a ImageListBadRequest with default headers values
func NewImageListBadRequest() *ImageListBadRequest {
	return &ImageListBadRequest{}
}

/*ImageListBadRequest handles this case with default header values.

A required header is missing. Correct the request and try again.
*/
type ImageListBadRequest struct {
	Payload *models.RequestError
}

func (o *ImageListBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v1/images][%d] imageListBadRequest  %+v", 400, o.Payload)
}

func (o *ImageListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageListUnauthorized creates a ImageListUnauthorized with default headers values
func NewImageListUnauthorized() *ImageListUnauthorized {
	return &ImageListUnauthorized{}
}

/*ImageListUnauthorized handles this case with default header values.

You are not authorized to view the requested resource, or your IBM Cloud bearer token is invalid. Correct the request and try again.
*/
type ImageListUnauthorized struct {
	Payload *models.RequestError
}

func (o *ImageListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/images][%d] imageListUnauthorized  %+v", 401, o.Payload)
}

func (o *ImageListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageListForbidden creates a ImageListForbidden with default headers values
func NewImageListForbidden() *ImageListForbidden {
	return &ImageListForbidden{}
}

/*ImageListForbidden handles this case with default header values.

You do not have access to the specified namespace. Choose another namespace.
*/
type ImageListForbidden struct {
	Payload *models.RequestError
}

func (o *ImageListForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/images][%d] imageListForbidden  %+v", 403, o.Payload)
}

func (o *ImageListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageListNotFound creates a ImageListNotFound with default headers values
func NewImageListNotFound() *ImageListNotFound {
	return &ImageListNotFound{}
}

/*ImageListNotFound handles this case with default header values.

No namespaces were found for the targeted account. Add a namespace or target a different account. Alternatively, specify the option to include IBM-provided public images.
*/
type ImageListNotFound struct {
	Payload *models.RequestError
}

func (o *ImageListNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/images][%d] imageListNotFound  %+v", 404, o.Payload)
}

func (o *ImageListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageListInternalServerError creates a ImageListInternalServerError with default headers values
func NewImageListInternalServerError() *ImageListInternalServerError {
	return &ImageListInternalServerError{}
}

/*ImageListInternalServerError handles this case with default header values.

Internal server error
*/
type ImageListInternalServerError struct {
	Payload *models.RequestError
}

func (o *ImageListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1/images][%d] imageListInternalServerError  %+v", 500, o.Payload)
}

func (o *ImageListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageListServiceUnavailable creates a ImageListServiceUnavailable with default headers values
func NewImageListServiceUnavailable() *ImageListServiceUnavailable {
	return &ImageListServiceUnavailable{}
}

/*ImageListServiceUnavailable handles this case with default header values.

Unable to authenticate with IBM Cloud. Try again later.
*/
type ImageListServiceUnavailable struct {
	Payload *models.RequestError
}

func (o *ImageListServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v1/images][%d] imageListServiceUnavailable  %+v", 503, o.Payload)
}

func (o *ImageListServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
