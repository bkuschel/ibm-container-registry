

package builds




import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewImageBuildParams creates a new ImageBuildParams object
// with the default values initialized.
func NewImageBuildParams() *ImageBuildParams {
	var ()
	return &ImageBuildParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewImageBuildParamsWithTimeout creates a new ImageBuildParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewImageBuildParamsWithTimeout(timeout time.Duration) *ImageBuildParams {
	var ()
	return &ImageBuildParams{

		timeout: timeout,
	}
}

// NewImageBuildParamsWithContext creates a new ImageBuildParams object
// with the default values initialized, and the ability to set a context for a request
func NewImageBuildParamsWithContext(ctx context.Context) *ImageBuildParams {
	var ()
	return &ImageBuildParams{

		Context: ctx,
	}
}

// NewImageBuildParamsWithHTTPClient creates a new ImageBuildParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewImageBuildParamsWithHTTPClient(client *http.Client) *ImageBuildParams {
	var ()
	return &ImageBuildParams{
		HTTPClient: client,
	}
}

/*ImageBuildParams contains all the parameters to send to the API endpoint
for the image build operation typically these are written to a http.Request
*/
type ImageBuildParams struct {

	/*Account
	  The unique ID for your IBM Cloud account.

	*/
	Account string
	/*Authorization
	  The IBM Cloud IAM access token that you receive when you log in to IBM Cloud. Run `bx iam oauth-tokens` to retrieve your access token, which is prefixed with 'Bearer'.

	*/
	Authorization string
	/*RefreshToken
	  The IBM Cloud IAM refresh token that you receive when you log in to IBM Cloud. Your refresh token is stored in your IBM Cloud config file: `~/.bluemix/config.json`

	*/
	RefreshToken string
	/*Buildargs
	  A JSON key-value structure that contains build arguments. The value of the build arguments are available as environment variables when you specify an `ARG` line which matches the key in your Dockerfile.

	*/
	Buildargs *string
	/*F
	  Specify the location of the Dockerfile relative to the build context. If not specified, the default is 'PATH/Dockerfile', where PATH is the root of the build context.

	*/
	F *string
	/*Nocache
	  If set to true, cached image layers from previous builds are not used in this build. Use this option if you expect the result of commands that run in the build to change.

	*/
	Nocache *string
	/*Pull
	  If set to true, the base image is pulled even if an image with a matching tag already exists on the build host. The base image is specified by using the FROM keyword in your Dockerfile. Use this option to update the version of the base image on the build host.

	*/
	Pull *string
	/*Quiet
	  If set to true, build output is suppressed unless an error occurs.

	*/
	Quiet *string
	/*Squash
	  If set to true, the filesystem of the built image is reduced to one layer before it is pushed to the registry. Use this option if the number of layers in your image is close to the maximum for your storage driver.

	*/
	Squash *string
	/*T
	  The full name for the image that you want to build, including the registry URL and namespace.

	*/
	T string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the image build params
func (o *ImageBuildParams) WithTimeout(timeout time.Duration) *ImageBuildParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the image build params
func (o *ImageBuildParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the image build params
func (o *ImageBuildParams) WithContext(ctx context.Context) *ImageBuildParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the image build params
func (o *ImageBuildParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the image build params
func (o *ImageBuildParams) WithHTTPClient(client *http.Client) *ImageBuildParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the image build params
func (o *ImageBuildParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccount adds the account to the image build params
func (o *ImageBuildParams) WithAccount(account string) *ImageBuildParams {
	o.SetAccount(account)
	return o
}

// SetAccount adds the account to the image build params
func (o *ImageBuildParams) SetAccount(account string) {
	o.Account = account
}

// WithAuthorization adds the authorization to the image build params
func (o *ImageBuildParams) WithAuthorization(authorization string) *ImageBuildParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the image build params
func (o *ImageBuildParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithRefreshToken adds the refreshToken to the image build params
func (o *ImageBuildParams) WithRefreshToken(refreshToken string) *ImageBuildParams {
	o.SetRefreshToken(refreshToken)
	return o
}

// SetRefreshToken adds the refreshToken to the image build params
func (o *ImageBuildParams) SetRefreshToken(refreshToken string) {
	o.RefreshToken = refreshToken
}

// WithBuildargs adds the buildargs to the image build params
func (o *ImageBuildParams) WithBuildargs(buildargs *string) *ImageBuildParams {
	o.SetBuildargs(buildargs)
	return o
}

// SetBuildargs adds the buildargs to the image build params
func (o *ImageBuildParams) SetBuildargs(buildargs *string) {
	o.Buildargs = buildargs
}

// WithF adds the f to the image build params
func (o *ImageBuildParams) WithF(f *string) *ImageBuildParams {
	o.SetF(f)
	return o
}

// SetF adds the f to the image build params
func (o *ImageBuildParams) SetF(f *string) {
	o.F = f
}

// WithNocache adds the nocache to the image build params
func (o *ImageBuildParams) WithNocache(nocache *string) *ImageBuildParams {
	o.SetNocache(nocache)
	return o
}

// SetNocache adds the nocache to the image build params
func (o *ImageBuildParams) SetNocache(nocache *string) {
	o.Nocache = nocache
}

// WithPull adds the pull to the image build params
func (o *ImageBuildParams) WithPull(pull *string) *ImageBuildParams {
	o.SetPull(pull)
	return o
}

// SetPull adds the pull to the image build params
func (o *ImageBuildParams) SetPull(pull *string) {
	o.Pull = pull
}

// WithQuiet adds the quiet to the image build params
func (o *ImageBuildParams) WithQuiet(quiet *string) *ImageBuildParams {
	o.SetQuiet(quiet)
	return o
}

// SetQuiet adds the quiet to the image build params
func (o *ImageBuildParams) SetQuiet(quiet *string) {
	o.Quiet = quiet
}

// WithSquash adds the squash to the image build params
func (o *ImageBuildParams) WithSquash(squash *string) *ImageBuildParams {
	o.SetSquash(squash)
	return o
}

// SetSquash adds the squash to the image build params
func (o *ImageBuildParams) SetSquash(squash *string) {
	o.Squash = squash
}

// WithT adds the t to the image build params
func (o *ImageBuildParams) WithT(t string) *ImageBuildParams {
	o.SetT(t)
	return o
}

// SetT adds the t to the image build params
func (o *ImageBuildParams) SetT(t string) {
	o.T = t
}

// WriteToRequest writes these params to a request
func (o *ImageBuildParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// header param RefreshToken
	if err := r.SetHeaderParam("RefreshToken", o.RefreshToken); err != nil {
		return err
	}

	if o.Buildargs != nil {

		// query param buildargs
		var qrBuildargs string
		if o.Buildargs != nil {
			qrBuildargs = *o.Buildargs
		}
		qBuildargs := qrBuildargs
		if qBuildargs != "" {
			if err := r.SetQueryParam("buildargs", qBuildargs); err != nil {
				return err
			}
		}

	}

	if o.F != nil {

		// query param f
		var qrF string
		if o.F != nil {
			qrF = *o.F
		}
		qF := qrF
		if qF != "" {
			if err := r.SetQueryParam("f", qF); err != nil {
				return err
			}
		}

	}

	if o.Nocache != nil {

		// query param nocache
		var qrNocache string
		if o.Nocache != nil {
			qrNocache = *o.Nocache
		}
		qNocache := qrNocache
		if qNocache != "" {
			if err := r.SetQueryParam("nocache", qNocache); err != nil {
				return err
			}
		}

	}

	if o.Pull != nil {

		// query param pull
		var qrPull string
		if o.Pull != nil {
			qrPull = *o.Pull
		}
		qPull := qrPull
		if qPull != "" {
			if err := r.SetQueryParam("pull", qPull); err != nil {
				return err
			}
		}

	}

	if o.Quiet != nil {

		// query param quiet
		var qrQuiet string
		if o.Quiet != nil {
			qrQuiet = *o.Quiet
		}
		qQuiet := qrQuiet
		if qQuiet != "" {
			if err := r.SetQueryParam("quiet", qQuiet); err != nil {
				return err
			}
		}

	}

	if o.Squash != nil {

		// query param squash
		var qrSquash string
		if o.Squash != nil {
			qrSquash = *o.Squash
		}
		qSquash := qrSquash
		if qSquash != "" {
			if err := r.SetQueryParam("squash", qSquash); err != nil {
				return err
			}
		}

	}

	// query param t
	qrT := o.T
	qT := qrT
	if qT != "" {
		if err := r.SetQueryParam("t", qT); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
