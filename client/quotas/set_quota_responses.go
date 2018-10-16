

package quotas




import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/bkuschel/ibm-container-registry/models"
)

// SetQuotaReader is a Reader for the SetQuota structure.
type SetQuotaReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetQuotaReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSetQuotaOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewSetQuotaBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewSetQuotaForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewSetQuotaInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewSetQuotaServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSetQuotaOK creates a SetQuotaOK with default headers values
func NewSetQuotaOK() *SetQuotaOK {
	return &SetQuotaOK{}
}

/*SetQuotaOK handles this case with default header values.

Description was not specified
*/
type SetQuotaOK struct {
	Payload string
}

func (o *SetQuotaOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/quotas][%d] setQuotaOK  %+v", 200, o.Payload)
}

func (o *SetQuotaOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetQuotaBadRequest creates a SetQuotaBadRequest with default headers values
func NewSetQuotaBadRequest() *SetQuotaBadRequest {
	return &SetQuotaBadRequest{}
}

/*SetQuotaBadRequest handles this case with default header values.

A required header is missing. Correct the request and try again.
*/
type SetQuotaBadRequest struct {
	Payload *models.RequestError
}

func (o *SetQuotaBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/quotas][%d] setQuotaBadRequest  %+v", 400, o.Payload)
}

func (o *SetQuotaBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetQuotaForbidden creates a SetQuotaForbidden with default headers values
func NewSetQuotaForbidden() *SetQuotaForbidden {
	return &SetQuotaForbidden{}
}

/*SetQuotaForbidden handles this case with default header values.

You do not have access to perform the requested operation. You must have either Manager or Billing Manager roles on the targeted account to perform this action.
*/
type SetQuotaForbidden struct {
	Payload *models.RequestError
}

func (o *SetQuotaForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/quotas][%d] setQuotaForbidden  %+v", 403, o.Payload)
}

func (o *SetQuotaForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetQuotaInternalServerError creates a SetQuotaInternalServerError with default headers values
func NewSetQuotaInternalServerError() *SetQuotaInternalServerError {
	return &SetQuotaInternalServerError{}
}

/*SetQuotaInternalServerError handles this case with default header values.

Internal server error
*/
type SetQuotaInternalServerError struct {
	Payload *models.RequestError
}

func (o *SetQuotaInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/quotas][%d] setQuotaInternalServerError  %+v", 500, o.Payload)
}

func (o *SetQuotaInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetQuotaServiceUnavailable creates a SetQuotaServiceUnavailable with default headers values
func NewSetQuotaServiceUnavailable() *SetQuotaServiceUnavailable {
	return &SetQuotaServiceUnavailable{}
}

/*SetQuotaServiceUnavailable handles this case with default header values.

Unable to authenticate with IBM Cloud. Try again later.
*/
type SetQuotaServiceUnavailable struct {
	Payload *models.RequestError
}

func (o *SetQuotaServiceUnavailable) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/quotas][%d] setQuotaServiceUnavailable  %+v", 503, o.Payload)
}

func (o *SetQuotaServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
