

package images




import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewImageDeleteParams creates a new ImageDeleteParams object
// with the default values initialized.
func NewImageDeleteParams() *ImageDeleteParams {
	var ()
	return &ImageDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewImageDeleteParamsWithTimeout creates a new ImageDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewImageDeleteParamsWithTimeout(timeout time.Duration) *ImageDeleteParams {
	var ()
	return &ImageDeleteParams{

		timeout: timeout,
	}
}

// NewImageDeleteParamsWithContext creates a new ImageDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewImageDeleteParamsWithContext(ctx context.Context) *ImageDeleteParams {
	var ()
	return &ImageDeleteParams{

		Context: ctx,
	}
}

// NewImageDeleteParamsWithHTTPClient creates a new ImageDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewImageDeleteParamsWithHTTPClient(client *http.Client) *ImageDeleteParams {
	var ()
	return &ImageDeleteParams{
		HTTPClient: client,
	}
}

/*ImageDeleteParams contains all the parameters to send to the API endpoint
for the image delete operation typically these are written to a http.Request
*/
type ImageDeleteParams struct {

	/*Account
	  The unique ID for your IBM Cloud account.

	*/
	Account string
	/*Authorization
	  The IBM Cloud IAM access token that you receive when you log in to IBM Cloud. Run `bx iam oauth-tokens` to retrieve your access token, which is prefixed with 'Bearer'.

	*/
	Authorization string
	/*Image
	  The full IBM Cloud registry path to the image that you want to delete, including its tag. If you do not provide a specific tag, the version with the `latest` tag is removed.

	*/
	Image string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the image delete params
func (o *ImageDeleteParams) WithTimeout(timeout time.Duration) *ImageDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the image delete params
func (o *ImageDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the image delete params
func (o *ImageDeleteParams) WithContext(ctx context.Context) *ImageDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the image delete params
func (o *ImageDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the image delete params
func (o *ImageDeleteParams) WithHTTPClient(client *http.Client) *ImageDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the image delete params
func (o *ImageDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccount adds the account to the image delete params
func (o *ImageDeleteParams) WithAccount(account string) *ImageDeleteParams {
	o.SetAccount(account)
	return o
}

// SetAccount adds the account to the image delete params
func (o *ImageDeleteParams) SetAccount(account string) {
	o.Account = account
}

// WithAuthorization adds the authorization to the image delete params
func (o *ImageDeleteParams) WithAuthorization(authorization string) *ImageDeleteParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the image delete params
func (o *ImageDeleteParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithImage adds the image to the image delete params
func (o *ImageDeleteParams) WithImage(image string) *ImageDeleteParams {
	o.SetImage(image)
	return o
}

// SetImage adds the image to the image delete params
func (o *ImageDeleteParams) SetImage(image string) {
	o.Image = image
}

// WriteToRequest writes these params to a request
func (o *ImageDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Account
	if err := r.SetHeaderParam("Account", o.Account); err != nil {
		return err
	}

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	// path param image
	if err := r.SetPathParam("image", o.Image); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
