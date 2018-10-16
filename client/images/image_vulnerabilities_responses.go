

package images




import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/bkuschel/ibm-container-registry/models"
)

// ImageVulnerabilitiesReader is a Reader for the ImageVulnerabilities structure.
type ImageVulnerabilitiesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ImageVulnerabilitiesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewImageVulnerabilitiesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewImageVulnerabilitiesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewImageVulnerabilitiesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewImageVulnerabilitiesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewImageVulnerabilitiesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewImageVulnerabilitiesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewImageVulnerabilitiesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewImageVulnerabilitiesOK creates a ImageVulnerabilitiesOK with default headers values
func NewImageVulnerabilitiesOK() *ImageVulnerabilitiesOK {
	return &ImageVulnerabilitiesOK{}
}

/*ImageVulnerabilitiesOK handles this case with default header values.

Description was not specified
*/
type ImageVulnerabilitiesOK struct {
	Payload *models.ImageAssessment
}

func (o *ImageVulnerabilitiesOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/images/{image}/vulnerabilities][%d] imageVulnerabilitiesOK  %+v", 200, o.Payload)
}

func (o *ImageVulnerabilitiesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ImageAssessment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageVulnerabilitiesBadRequest creates a ImageVulnerabilitiesBadRequest with default headers values
func NewImageVulnerabilitiesBadRequest() *ImageVulnerabilitiesBadRequest {
	return &ImageVulnerabilitiesBadRequest{}
}

/*ImageVulnerabilitiesBadRequest handles this case with default header values.

A required header is missing. Correct the request and try again.
*/
type ImageVulnerabilitiesBadRequest struct {
	Payload *models.RequestError
}

func (o *ImageVulnerabilitiesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v1/images/{image}/vulnerabilities][%d] imageVulnerabilitiesBadRequest  %+v", 400, o.Payload)
}

func (o *ImageVulnerabilitiesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageVulnerabilitiesUnauthorized creates a ImageVulnerabilitiesUnauthorized with default headers values
func NewImageVulnerabilitiesUnauthorized() *ImageVulnerabilitiesUnauthorized {
	return &ImageVulnerabilitiesUnauthorized{}
}

/*ImageVulnerabilitiesUnauthorized handles this case with default header values.

You are not authorized to view the requested resource, or your IBM Cloud bearer token is invalid. Correct the request and try again.
*/
type ImageVulnerabilitiesUnauthorized struct {
	Payload *models.RequestError
}

func (o *ImageVulnerabilitiesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/images/{image}/vulnerabilities][%d] imageVulnerabilitiesUnauthorized  %+v", 401, o.Payload)
}

func (o *ImageVulnerabilitiesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageVulnerabilitiesForbidden creates a ImageVulnerabilitiesForbidden with default headers values
func NewImageVulnerabilitiesForbidden() *ImageVulnerabilitiesForbidden {
	return &ImageVulnerabilitiesForbidden{}
}

/*ImageVulnerabilitiesForbidden handles this case with default header values.

You do not have authorization to access the specified namespace.
*/
type ImageVulnerabilitiesForbidden struct {
	Payload *models.RequestError
}

func (o *ImageVulnerabilitiesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/images/{image}/vulnerabilities][%d] imageVulnerabilitiesForbidden  %+v", 403, o.Payload)
}

func (o *ImageVulnerabilitiesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageVulnerabilitiesNotFound creates a ImageVulnerabilitiesNotFound with default headers values
func NewImageVulnerabilitiesNotFound() *ImageVulnerabilitiesNotFound {
	return &ImageVulnerabilitiesNotFound{}
}

/*ImageVulnerabilitiesNotFound handles this case with default header values.

The specified image was not found. Check the spelling of the specified image.
*/
type ImageVulnerabilitiesNotFound struct {
	Payload *models.RequestError
}

func (o *ImageVulnerabilitiesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/images/{image}/vulnerabilities][%d] imageVulnerabilitiesNotFound  %+v", 404, o.Payload)
}

func (o *ImageVulnerabilitiesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageVulnerabilitiesInternalServerError creates a ImageVulnerabilitiesInternalServerError with default headers values
func NewImageVulnerabilitiesInternalServerError() *ImageVulnerabilitiesInternalServerError {
	return &ImageVulnerabilitiesInternalServerError{}
}

/*ImageVulnerabilitiesInternalServerError handles this case with default header values.

Internal server error
*/
type ImageVulnerabilitiesInternalServerError struct {
	Payload *models.RequestError
}

func (o *ImageVulnerabilitiesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1/images/{image}/vulnerabilities][%d] imageVulnerabilitiesInternalServerError  %+v", 500, o.Payload)
}

func (o *ImageVulnerabilitiesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageVulnerabilitiesServiceUnavailable creates a ImageVulnerabilitiesServiceUnavailable with default headers values
func NewImageVulnerabilitiesServiceUnavailable() *ImageVulnerabilitiesServiceUnavailable {
	return &ImageVulnerabilitiesServiceUnavailable{}
}

/*ImageVulnerabilitiesServiceUnavailable handles this case with default header values.

Unable to authenticate with IBM Cloud. Try again later.
*/
type ImageVulnerabilitiesServiceUnavailable struct {
	Payload *models.RequestError
}

func (o *ImageVulnerabilitiesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v1/images/{image}/vulnerabilities][%d] imageVulnerabilitiesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *ImageVulnerabilitiesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
