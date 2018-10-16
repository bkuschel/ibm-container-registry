

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

// NewImageVulnerabilitiesParams creates a new ImageVulnerabilitiesParams object
// with the default values initialized.
func NewImageVulnerabilitiesParams() *ImageVulnerabilitiesParams {
	var ()
	return &ImageVulnerabilitiesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewImageVulnerabilitiesParamsWithTimeout creates a new ImageVulnerabilitiesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewImageVulnerabilitiesParamsWithTimeout(timeout time.Duration) *ImageVulnerabilitiesParams {
	var ()
	return &ImageVulnerabilitiesParams{

		timeout: timeout,
	}
}

// NewImageVulnerabilitiesParamsWithContext creates a new ImageVulnerabilitiesParams object
// with the default values initialized, and the ability to set a context for a request
func NewImageVulnerabilitiesParamsWithContext(ctx context.Context) *ImageVulnerabilitiesParams {
	var ()
	return &ImageVulnerabilitiesParams{

		Context: ctx,
	}
}

// NewImageVulnerabilitiesParamsWithHTTPClient creates a new ImageVulnerabilitiesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewImageVulnerabilitiesParamsWithHTTPClient(client *http.Client) *ImageVulnerabilitiesParams {
	var ()
	return &ImageVulnerabilitiesParams{
		HTTPClient: client,
	}
}

/*ImageVulnerabilitiesParams contains all the parameters to send to the API endpoint
for the image vulnerabilities operation typically these are written to a http.Request
*/
type ImageVulnerabilitiesParams struct {

	/*Account
	  The unique ID for your IBM Cloud account.

	*/
	Account string
	/*Authorization
	  The IBM Cloud IAM access token that you receive when you log in to IBM Cloud. Run `bx iam oauth-tokens` to retrieve your access token, which is prefixed with 'Bearer'.

	*/
	Authorization string
	/*Advisory
	  Specifies to include advisory compliance checks in the report.

	*/
	Advisory *string
	/*All
	  Specifies to include all checks in the report. If not specified or false, only failing checks are returned.

	*/
	All *string
	/*Image
	  The full IBM Cloud registry path to the image that you want to inspect. Run `bx cr images`` or call the `GET /images/json` endpoint to review images that are in the registry.

	*/
	Image string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the image vulnerabilities params
func (o *ImageVulnerabilitiesParams) WithTimeout(timeout time.Duration) *ImageVulnerabilitiesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the image vulnerabilities params
func (o *ImageVulnerabilitiesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the image vulnerabilities params
func (o *ImageVulnerabilitiesParams) WithContext(ctx context.Context) *ImageVulnerabilitiesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the image vulnerabilities params
func (o *ImageVulnerabilitiesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the image vulnerabilities params
func (o *ImageVulnerabilitiesParams) WithHTTPClient(client *http.Client) *ImageVulnerabilitiesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the image vulnerabilities params
func (o *ImageVulnerabilitiesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccount adds the account to the image vulnerabilities params
func (o *ImageVulnerabilitiesParams) WithAccount(account string) *ImageVulnerabilitiesParams {
	o.SetAccount(account)
	return o
}

// SetAccount adds the account to the image vulnerabilities params
func (o *ImageVulnerabilitiesParams) SetAccount(account string) {
	o.Account = account
}

// WithAuthorization adds the authorization to the image vulnerabilities params
func (o *ImageVulnerabilitiesParams) WithAuthorization(authorization string) *ImageVulnerabilitiesParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the image vulnerabilities params
func (o *ImageVulnerabilitiesParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithAdvisory adds the advisory to the image vulnerabilities params
func (o *ImageVulnerabilitiesParams) WithAdvisory(advisory *string) *ImageVulnerabilitiesParams {
	o.SetAdvisory(advisory)
	return o
}

// SetAdvisory adds the advisory to the image vulnerabilities params
func (o *ImageVulnerabilitiesParams) SetAdvisory(advisory *string) {
	o.Advisory = advisory
}

// WithAll adds the all to the image vulnerabilities params
func (o *ImageVulnerabilitiesParams) WithAll(all *string) *ImageVulnerabilitiesParams {
	o.SetAll(all)
	return o
}

// SetAll adds the all to the image vulnerabilities params
func (o *ImageVulnerabilitiesParams) SetAll(all *string) {
	o.All = all
}

// WithImage adds the image to the image vulnerabilities params
func (o *ImageVulnerabilitiesParams) WithImage(image string) *ImageVulnerabilitiesParams {
	o.SetImage(image)
	return o
}

// SetImage adds the image to the image vulnerabilities params
func (o *ImageVulnerabilitiesParams) SetImage(image string) {
	o.Image = image
}

// WriteToRequest writes these params to a request
func (o *ImageVulnerabilitiesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Advisory != nil {

		// query param advisory
		var qrAdvisory string
		if o.Advisory != nil {
			qrAdvisory = *o.Advisory
		}
		qAdvisory := qrAdvisory
		if qAdvisory != "" {
			if err := r.SetQueryParam("advisory", qAdvisory); err != nil {
				return err
			}
		}

	}

	if o.All != nil {

		// query param all
		var qrAll string
		if o.All != nil {
			qrAll = *o.All
		}
		qAll := qrAll
		if qAll != "" {
			if err := r.SetQueryParam("all", qAll); err != nil {
				return err
			}
		}

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
