

package auth




import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/bkuschel/ibm-container-registry/models"
)

// NewSetAuthParams creates a new SetAuthParams object
// with the default values initialized.
func NewSetAuthParams() *SetAuthParams {
	var ()
	return &SetAuthParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSetAuthParamsWithTimeout creates a new SetAuthParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSetAuthParamsWithTimeout(timeout time.Duration) *SetAuthParams {
	var ()
	return &SetAuthParams{

		timeout: timeout,
	}
}

// NewSetAuthParamsWithContext creates a new SetAuthParams object
// with the default values initialized, and the ability to set a context for a request
func NewSetAuthParamsWithContext(ctx context.Context) *SetAuthParams {
	var ()
	return &SetAuthParams{

		Context: ctx,
	}
}

// NewSetAuthParamsWithHTTPClient creates a new SetAuthParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSetAuthParamsWithHTTPClient(client *http.Client) *SetAuthParams {
	var ()
	return &SetAuthParams{
		HTTPClient: client,
	}
}

/*SetAuthParams contains all the parameters to send to the API endpoint
for the set auth operation typically these are written to a http.Request
*/
type SetAuthParams struct {

	/*Account
	  The unique ID for your IBM Cloud account.

	*/
	Account string
	/*Authorization
	  The IBM Cloud IAM access token that you receive when you log in to IBM Cloud. Run `bx iam oauth-tokens` to retrieve your access token, which is prefixed with 'Bearer'.

	*/
	Authorization string
	/*AuthorizationOptions
	  A JSON document containing the requested updates.

	*/
	AuthorizationOptions *models.AuthResponse

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the set auth params
func (o *SetAuthParams) WithTimeout(timeout time.Duration) *SetAuthParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the set auth params
func (o *SetAuthParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the set auth params
func (o *SetAuthParams) WithContext(ctx context.Context) *SetAuthParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the set auth params
func (o *SetAuthParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the set auth params
func (o *SetAuthParams) WithHTTPClient(client *http.Client) *SetAuthParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the set auth params
func (o *SetAuthParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccount adds the account to the set auth params
func (o *SetAuthParams) WithAccount(account string) *SetAuthParams {
	o.SetAccount(account)
	return o
}

// SetAccount adds the account to the set auth params
func (o *SetAuthParams) SetAccount(account string) {
	o.Account = account
}

// WithAuthorization adds the authorization to the set auth params
func (o *SetAuthParams) WithAuthorization(authorization string) *SetAuthParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the set auth params
func (o *SetAuthParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithAuthorizationOptions adds the authorizationOptions to the set auth params
func (o *SetAuthParams) WithAuthorizationOptions(authorizationOptions *models.AuthResponse) *SetAuthParams {
	o.SetAuthorizationOptions(authorizationOptions)
	return o
}

// SetAuthorizationOptions adds the authorizationOptions to the set auth params
func (o *SetAuthParams) SetAuthorizationOptions(authorizationOptions *models.AuthResponse) {
	o.AuthorizationOptions = authorizationOptions
}

// WriteToRequest writes these params to a request
func (o *SetAuthParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.AuthorizationOptions != nil {
		if err := r.SetBodyParam(o.AuthorizationOptions); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
