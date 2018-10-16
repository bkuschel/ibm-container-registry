

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

// NewImageListParams creates a new ImageListParams object
// with the default values initialized.
func NewImageListParams() *ImageListParams {
	var ()
	return &ImageListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewImageListParamsWithTimeout creates a new ImageListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewImageListParamsWithTimeout(timeout time.Duration) *ImageListParams {
	var ()
	return &ImageListParams{

		timeout: timeout,
	}
}

// NewImageListParamsWithContext creates a new ImageListParams object
// with the default values initialized, and the ability to set a context for a request
func NewImageListParamsWithContext(ctx context.Context) *ImageListParams {
	var ()
	return &ImageListParams{

		Context: ctx,
	}
}

// NewImageListParamsWithHTTPClient creates a new ImageListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewImageListParamsWithHTTPClient(client *http.Client) *ImageListParams {
	var ()
	return &ImageListParams{
		HTTPClient: client,
	}
}

/*ImageListParams contains all the parameters to send to the API endpoint
for the image list operation typically these are written to a http.Request
*/
type ImageListParams struct {

	/*Account
	  The unique ID for your IBM Cloud account.

	*/
	Account string
	/*Authorization
	  The IBM Cloud IAM access token that you receive when you log in to IBM Cloud. Run `bx iam oauth-tokens` to retrieve your access token, which is prefixed with 'Bearer'.

	*/
	Authorization string
	/*IncludeIBM
	  Includes IBM-provided public images in the list of images. If this option is not specified, private images are listed only. If this option is specified more than once, the last parsed setting is the setting that is used.

	*/
	IncludeIBM *string
	/*IncludePrivate
	  Includes private images in the list of images. If this option is not specified, private images are listed. If this option is specified more than once, the last parsed setting is the setting that is used.

	*/
	IncludePrivate *string
	/*Namespace
	  Lists images that are stored in the specified namespace only. Query multiple namespaces by specifying this option for each namespace. If this option is not specified, images from all namespaces in the specified IBM Cloud account are listed.

	*/
	Namespace *string
	/*Repository
	  Lists images that are stored in the specified repository, under your namespaces. Query multiple repositories by specifying this option for each repository. If this option is not specified, images from all repos are listed.

	*/
	Repository *string
	/*Vulnerabilities
	  Displays Vulnerability Advisor status for the listed images. If this option is specified more than once, the last parsed setting is the setting that is used.

	*/
	Vulnerabilities *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the image list params
func (o *ImageListParams) WithTimeout(timeout time.Duration) *ImageListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the image list params
func (o *ImageListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the image list params
func (o *ImageListParams) WithContext(ctx context.Context) *ImageListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the image list params
func (o *ImageListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the image list params
func (o *ImageListParams) WithHTTPClient(client *http.Client) *ImageListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the image list params
func (o *ImageListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccount adds the account to the image list params
func (o *ImageListParams) WithAccount(account string) *ImageListParams {
	o.SetAccount(account)
	return o
}

// SetAccount adds the account to the image list params
func (o *ImageListParams) SetAccount(account string) {
	o.Account = account
}

// WithAuthorization adds the authorization to the image list params
func (o *ImageListParams) WithAuthorization(authorization string) *ImageListParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the image list params
func (o *ImageListParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithIncludeIBM adds the includeIBM to the image list params
func (o *ImageListParams) WithIncludeIBM(includeIBM *string) *ImageListParams {
	o.SetIncludeIBM(includeIBM)
	return o
}

// SetIncludeIBM adds the includeIBM to the image list params
func (o *ImageListParams) SetIncludeIBM(includeIBM *string) {
	o.IncludeIBM = includeIBM
}

// WithIncludePrivate adds the includePrivate to the image list params
func (o *ImageListParams) WithIncludePrivate(includePrivate *string) *ImageListParams {
	o.SetIncludePrivate(includePrivate)
	return o
}

// SetIncludePrivate adds the includePrivate to the image list params
func (o *ImageListParams) SetIncludePrivate(includePrivate *string) {
	o.IncludePrivate = includePrivate
}

// WithNamespace adds the namespace to the image list params
func (o *ImageListParams) WithNamespace(namespace *string) *ImageListParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the image list params
func (o *ImageListParams) SetNamespace(namespace *string) {
	o.Namespace = namespace
}

// WithRepository adds the repository to the image list params
func (o *ImageListParams) WithRepository(repository *string) *ImageListParams {
	o.SetRepository(repository)
	return o
}

// SetRepository adds the repository to the image list params
func (o *ImageListParams) SetRepository(repository *string) {
	o.Repository = repository
}

// WithVulnerabilities adds the vulnerabilities to the image list params
func (o *ImageListParams) WithVulnerabilities(vulnerabilities *string) *ImageListParams {
	o.SetVulnerabilities(vulnerabilities)
	return o
}

// SetVulnerabilities adds the vulnerabilities to the image list params
func (o *ImageListParams) SetVulnerabilities(vulnerabilities *string) {
	o.Vulnerabilities = vulnerabilities
}

// WriteToRequest writes these params to a request
func (o *ImageListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.IncludeIBM != nil {

		// query param includeIBM
		var qrIncludeIBM string
		if o.IncludeIBM != nil {
			qrIncludeIBM = *o.IncludeIBM
		}
		qIncludeIBM := qrIncludeIBM
		if qIncludeIBM != "" {
			if err := r.SetQueryParam("includeIBM", qIncludeIBM); err != nil {
				return err
			}
		}

	}

	if o.IncludePrivate != nil {

		// query param includePrivate
		var qrIncludePrivate string
		if o.IncludePrivate != nil {
			qrIncludePrivate = *o.IncludePrivate
		}
		qIncludePrivate := qrIncludePrivate
		if qIncludePrivate != "" {
			if err := r.SetQueryParam("includePrivate", qIncludePrivate); err != nil {
				return err
			}
		}

	}

	if o.Namespace != nil {

		// query param namespace
		var qrNamespace string
		if o.Namespace != nil {
			qrNamespace = *o.Namespace
		}
		qNamespace := qrNamespace
		if qNamespace != "" {
			if err := r.SetQueryParam("namespace", qNamespace); err != nil {
				return err
			}
		}

	}

	if o.Repository != nil {

		// query param repository
		var qrRepository string
		if o.Repository != nil {
			qrRepository = *o.Repository
		}
		qRepository := qrRepository
		if qRepository != "" {
			if err := r.SetQueryParam("repository", qRepository); err != nil {
				return err
			}
		}

	}

	if o.Vulnerabilities != nil {

		// query param vulnerabilities
		var qrVulnerabilities string
		if o.Vulnerabilities != nil {
			qrVulnerabilities = *o.Vulnerabilities
		}
		qVulnerabilities := qrVulnerabilities
		if qVulnerabilities != "" {
			if err := r.SetQueryParam("vulnerabilities", qVulnerabilities); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
