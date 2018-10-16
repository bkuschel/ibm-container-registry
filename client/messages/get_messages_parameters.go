

package messages




import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetMessagesParams creates a new GetMessagesParams object
// with the default values initialized.
func NewGetMessagesParams() *GetMessagesParams {

	return &GetMessagesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetMessagesParamsWithTimeout creates a new GetMessagesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetMessagesParamsWithTimeout(timeout time.Duration) *GetMessagesParams {

	return &GetMessagesParams{

		timeout: timeout,
	}
}

// NewGetMessagesParamsWithContext creates a new GetMessagesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetMessagesParamsWithContext(ctx context.Context) *GetMessagesParams {

	return &GetMessagesParams{

		Context: ctx,
	}
}

// NewGetMessagesParamsWithHTTPClient creates a new GetMessagesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetMessagesParamsWithHTTPClient(client *http.Client) *GetMessagesParams {

	return &GetMessagesParams{
		HTTPClient: client,
	}
}

/*GetMessagesParams contains all the parameters to send to the API endpoint
for the get messages operation typically these are written to a http.Request
*/
type GetMessagesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get messages params
func (o *GetMessagesParams) WithTimeout(timeout time.Duration) *GetMessagesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get messages params
func (o *GetMessagesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get messages params
func (o *GetMessagesParams) WithContext(ctx context.Context) *GetMessagesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get messages params
func (o *GetMessagesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get messages params
func (o *GetMessagesParams) WithHTTPClient(client *http.Client) *GetMessagesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get messages params
func (o *GetMessagesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a request
func (o *GetMessagesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
