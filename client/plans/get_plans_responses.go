

package plans




import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/bkuschel/ibm-container-registry/models"
)

// GetPlansReader is a Reader for the GetPlans structure.
type GetPlansReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPlansReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPlansOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetPlansBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetPlansInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewGetPlansServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPlansOK creates a GetPlansOK with default headers values
func NewGetPlansOK() *GetPlansOK {
	return &GetPlansOK{}
}

/*GetPlansOK handles this case with default header values.

Description was not specified
*/
type GetPlansOK struct {
	Payload []*models.PlanResponse
}

func (o *GetPlansOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/plans][%d] getPlansOK  %+v", 200, o.Payload)
}

func (o *GetPlansOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPlansBadRequest creates a GetPlansBadRequest with default headers values
func NewGetPlansBadRequest() *GetPlansBadRequest {
	return &GetPlansBadRequest{}
}

/*GetPlansBadRequest handles this case with default header values.

A required header is missing. Correct the request and try again.
*/
type GetPlansBadRequest struct {
	Payload *models.RequestError
}

func (o *GetPlansBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v1/plans][%d] getPlansBadRequest  %+v", 400, o.Payload)
}

func (o *GetPlansBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPlansInternalServerError creates a GetPlansInternalServerError with default headers values
func NewGetPlansInternalServerError() *GetPlansInternalServerError {
	return &GetPlansInternalServerError{}
}

/*GetPlansInternalServerError handles this case with default header values.

Internal server error
*/
type GetPlansInternalServerError struct {
	Payload *models.RequestError
}

func (o *GetPlansInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1/plans][%d] getPlansInternalServerError  %+v", 500, o.Payload)
}

func (o *GetPlansInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPlansServiceUnavailable creates a GetPlansServiceUnavailable with default headers values
func NewGetPlansServiceUnavailable() *GetPlansServiceUnavailable {
	return &GetPlansServiceUnavailable{}
}

/*GetPlansServiceUnavailable handles this case with default header values.

Unable to authenticate with IBM Cloud. Try again later.
*/
type GetPlansServiceUnavailable struct {
	Payload *models.RequestError
}

func (o *GetPlansServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v1/plans][%d] getPlansServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetPlansServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
