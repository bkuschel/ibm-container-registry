

package quotas




import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/bkuschel/ibm-container-registry/models"
)

// GetQuotaReader is a Reader for the GetQuota structure.
type GetQuotaReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetQuotaReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetQuotaOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetQuotaBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetQuotaInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewGetQuotaServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetQuotaOK creates a GetQuotaOK with default headers values
func NewGetQuotaOK() *GetQuotaOK {
	return &GetQuotaOK{}
}

/*GetQuotaOK handles this case with default header values.

Description was not specified
*/
type GetQuotaOK struct {
	Payload *models.QuotaAPIResponse
}

func (o *GetQuotaOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/quotas][%d] getQuotaOK  %+v", 200, o.Payload)
}

func (o *GetQuotaOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.QuotaAPIResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQuotaBadRequest creates a GetQuotaBadRequest with default headers values
func NewGetQuotaBadRequest() *GetQuotaBadRequest {
	return &GetQuotaBadRequest{}
}

/*GetQuotaBadRequest handles this case with default header values.

A required header is missing. Correct the request and try again.
*/
type GetQuotaBadRequest struct {
	Payload *models.RequestError
}

func (o *GetQuotaBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v1/quotas][%d] getQuotaBadRequest  %+v", 400, o.Payload)
}

func (o *GetQuotaBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQuotaInternalServerError creates a GetQuotaInternalServerError with default headers values
func NewGetQuotaInternalServerError() *GetQuotaInternalServerError {
	return &GetQuotaInternalServerError{}
}

/*GetQuotaInternalServerError handles this case with default header values.

Internal server error
*/
type GetQuotaInternalServerError struct {
	Payload *models.RequestError
}

func (o *GetQuotaInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1/quotas][%d] getQuotaInternalServerError  %+v", 500, o.Payload)
}

func (o *GetQuotaInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQuotaServiceUnavailable creates a GetQuotaServiceUnavailable with default headers values
func NewGetQuotaServiceUnavailable() *GetQuotaServiceUnavailable {
	return &GetQuotaServiceUnavailable{}
}

/*GetQuotaServiceUnavailable handles this case with default header values.

Unable to authenticate with IBM Cloud. Try again later.
*/
type GetQuotaServiceUnavailable struct {
	Payload *models.RequestError
}

func (o *GetQuotaServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v1/quotas][%d] getQuotaServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetQuotaServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
