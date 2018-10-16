

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

// NewImageInspectParams creates a new ImageInspectParams object
// with the default values initialized.
func NewImageInspectParams() *ImageInspectParams {
	var ()
	return &ImageInspectParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewImageInspectParamsWithTimeout creates a new ImageInspectParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewImageInspectParamsWithTimeout(timeout time.Duration) *ImageInspectParams {
	var ()
	return &ImageInspectParams{

		timeout: timeout,
	}
}

// NewImageInspectParamsWithContext creates a new ImageInspectParams object
// with the default values initialized, and the ability to set a context for a request
func NewImageInspectParamsWithContext(ctx context.Context) *ImageInspectParams {
	var ()
	return &ImageInspectParams{

		Context: ctx,
	}
}

// NewImageInspectParamsWithHTTPClient creates a new ImageInspectParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewImageInspectParamsWithHTTPClient(client *http.Client) *ImageInspectParams {
	var ()
	return &ImageInspectParams{
		HTTPClient: client,
	}
}

/*ImageInspectParams contains all the parameters to send to the API endpoint
for the image inspect operation typically these are written to a http.Request
*/
type ImageInspectParams struct {

	/*Account
	  The unique ID for your IBM Cloud account.

	*/
	Account string
	/*Authorization
	  The IBM Cloud IAM access token that you receive when you log in to IBM Cloud. Run `bx iam oauth-tokens` to retrieve your access token, which is prefixed with 'Bearer'.

	*/
	Authorization string
	/*Image
	  The full IBM Cloud registry path to the image that you want to inspect. Run `bx cr images` or call the `GET /images/json` endpoint to review images that are in the registry.

	*/
	Image string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the image inspect params
func (o *ImageInspectParams) WithTimeout(timeout time.Duration) *ImageInspectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the image inspect params
func (o *ImageInspectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the image inspect params
func (o *ImageInspectParams) WithContext(ctx context.Context) *ImageInspectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the image inspect params
func (o *ImageInspectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the image inspect params
func (o *ImageInspectParams) WithHTTPClient(client *http.Client) *ImageInspectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the image inspect params
func (o *ImageInspectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccount adds the account to the image inspect params
func (o *ImageInspectParams) WithAccount(account string) *ImageInspectParams {
	o.SetAccount(account)
	return o
}

// SetAccount adds the account to the image inspect params
func (o *ImageInspectParams) SetAccount(account string) {
	o.Account = account
}

// WithAuthorization adds the authorization to the image inspect params
func (o *ImageInspectParams) WithAuthorization(authorization string) *ImageInspectParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the image inspect params
func (o *ImageInspectParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithImage adds the image to the image inspect params
func (o *ImageInspectParams) WithImage(image string) *ImageInspectParams {
	o.SetImage(image)
	return o
}

// SetImage adds the image to the image inspect params
func (o *ImageInspectParams) SetImage(image string) {
	o.Image = image
}

// WriteToRequest writes these params to a request
func (o *ImageInspectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
