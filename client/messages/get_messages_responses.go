

package messages




import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/bkuschel/ibm-container-registry/models"
)

// GetMessagesReader is a Reader for the GetMessages structure.
type GetMessagesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMessagesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetMessagesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 204:
		result := NewGetMessagesNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetMessagesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetMessagesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewGetMessagesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetMessagesOK creates a GetMessagesOK with default headers values
func NewGetMessagesOK() *GetMessagesOK {
	return &GetMessagesOK{}
}

/*GetMessagesOK handles this case with default header values.

Description was not specified
*/
type GetMessagesOK struct {
	Payload string
}

func (o *GetMessagesOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/messages][%d] getMessagesOK  %+v", 200, o.Payload)
}

func (o *GetMessagesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMessagesNoContent creates a GetMessagesNoContent with default headers values
func NewGetMessagesNoContent() *GetMessagesNoContent {
	return &GetMessagesNoContent{}
}

/*GetMessagesNoContent handles this case with default header values.

There are no messages to display.
*/
type GetMessagesNoContent struct {
	Payload string
}

func (o *GetMessagesNoContent) Error() string {
	return fmt.Sprintf("[GET /api/v1/messages][%d] getMessagesNoContent  %+v", 204, o.Payload)
}

func (o *GetMessagesNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMessagesBadRequest creates a GetMessagesBadRequest with default headers values
func NewGetMessagesBadRequest() *GetMessagesBadRequest {
	return &GetMessagesBadRequest{}
}

/*GetMessagesBadRequest handles this case with default header values.

A required header is missing. Correct the request and try again.
*/
type GetMessagesBadRequest struct {
	Payload *models.RequestError
}

func (o *GetMessagesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v1/messages][%d] getMessagesBadRequest  %+v", 400, o.Payload)
}

func (o *GetMessagesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMessagesInternalServerError creates a GetMessagesInternalServerError with default headers values
func NewGetMessagesInternalServerError() *GetMessagesInternalServerError {
	return &GetMessagesInternalServerError{}
}

/*GetMessagesInternalServerError handles this case with default header values.

Internal server error
*/
type GetMessagesInternalServerError struct {
	Payload *models.RequestError
}

func (o *GetMessagesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1/messages][%d] getMessagesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetMessagesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMessagesServiceUnavailable creates a GetMessagesServiceUnavailable with default headers values
func NewGetMessagesServiceUnavailable() *GetMessagesServiceUnavailable {
	return &GetMessagesServiceUnavailable{}
}

/*GetMessagesServiceUnavailable handles this case with default header values.

Unable to authenticate with IBM Cloud. Try again later.
*/
type GetMessagesServiceUnavailable struct {
	Payload *models.RequestError
}

func (o *GetMessagesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v1/messages][%d] getMessagesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetMessagesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
