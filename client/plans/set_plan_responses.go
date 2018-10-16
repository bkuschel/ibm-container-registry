

package plans




import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/bkuschel/ibm-container-registry/models"
)

// SetPlanReader is a Reader for the SetPlan structure.
type SetPlanReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetPlanReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSetPlanOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewSetPlanBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewSetPlanForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewSetPlanInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewSetPlanServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSetPlanOK creates a SetPlanOK with default headers values
func NewSetPlanOK() *SetPlanOK {
	return &SetPlanOK{}
}

/*SetPlanOK handles this case with default header values.

Description was not specified
*/
type SetPlanOK struct {
	Payload string
}

func (o *SetPlanOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/plans][%d] setPlanOK  %+v", 200, o.Payload)
}

func (o *SetPlanOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetPlanBadRequest creates a SetPlanBadRequest with default headers values
func NewSetPlanBadRequest() *SetPlanBadRequest {
	return &SetPlanBadRequest{}
}

/*SetPlanBadRequest handles this case with default header values.

A required header is missing. Correct the request and try again.
*/
type SetPlanBadRequest struct {
	Payload *models.RequestError
}

func (o *SetPlanBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/plans][%d] setPlanBadRequest  %+v", 400, o.Payload)
}

func (o *SetPlanBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetPlanForbidden creates a SetPlanForbidden with default headers values
func NewSetPlanForbidden() *SetPlanForbidden {
	return &SetPlanForbidden{}
}

/*SetPlanForbidden handles this case with default header values.

You do not have access to perform the requested operation. You must have either Manager or Billing Manager roles on the targeted account to perform this action.
*/
type SetPlanForbidden struct {
	Payload *models.RequestError
}

func (o *SetPlanForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/plans][%d] setPlanForbidden  %+v", 403, o.Payload)
}

func (o *SetPlanForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetPlanInternalServerError creates a SetPlanInternalServerError with default headers values
func NewSetPlanInternalServerError() *SetPlanInternalServerError {
	return &SetPlanInternalServerError{}
}

/*SetPlanInternalServerError handles this case with default header values.

Internal server error
*/
type SetPlanInternalServerError struct {
	Payload *models.RequestError
}

func (o *SetPlanInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/plans][%d] setPlanInternalServerError  %+v", 500, o.Payload)
}

func (o *SetPlanInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetPlanServiceUnavailable creates a SetPlanServiceUnavailable with default headers values
func NewSetPlanServiceUnavailable() *SetPlanServiceUnavailable {
	return &SetPlanServiceUnavailable{}
}

/*SetPlanServiceUnavailable handles this case with default header values.

Unable to authenticate with IBM Cloud. Try again later.
*/
type SetPlanServiceUnavailable struct {
	Payload *models.RequestError
}

func (o *SetPlanServiceUnavailable) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/plans][%d] setPlanServiceUnavailable  %+v", 503, o.Payload)
}

func (o *SetPlanServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
