

package builds




import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/bkuschel/ibm-container-registry/models"
)

// ImageBuildReader is a Reader for the ImageBuild structure.
type ImageBuildReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ImageBuildReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewImageBuildOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 500:
		result := NewImageBuildInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewImageBuildServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewImageBuildOK creates a ImageBuildOK with default headers values
func NewImageBuildOK() *ImageBuildOK {
	return &ImageBuildOK{}
}

/*ImageBuildOK handles this case with default header values.

A stream of JSON objects that contain the output from the build host.
*/
type ImageBuildOK struct {
	Payload *models.Buildoutput
}

func (o *ImageBuildOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/builds][%d] imageBuildOK  %+v", 200, o.Payload)
}

func (o *ImageBuildOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Buildoutput)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageBuildInternalServerError creates a ImageBuildInternalServerError with default headers values
func NewImageBuildInternalServerError() *ImageBuildInternalServerError {
	return &ImageBuildInternalServerError{}
}

/*ImageBuildInternalServerError handles this case with default header values.

Internal server error or build failure
*/
type ImageBuildInternalServerError struct {
	Payload *models.Errordetail
}

func (o *ImageBuildInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v1/builds][%d] imageBuildInternalServerError  %+v", 500, o.Payload)
}

func (o *ImageBuildInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Errordetail)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageBuildServiceUnavailable creates a ImageBuildServiceUnavailable with default headers values
func NewImageBuildServiceUnavailable() *ImageBuildServiceUnavailable {
	return &ImageBuildServiceUnavailable{}
}

/*ImageBuildServiceUnavailable handles this case with default header values.

Unable to authenticate with IBM Cloud. Try again later.
*/
type ImageBuildServiceUnavailable struct {
	Payload *models.RequestError
}

func (o *ImageBuildServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v1/builds][%d] imageBuildServiceUnavailable  %+v", 503, o.Payload)
}

func (o *ImageBuildServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
