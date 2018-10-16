

package tokens




import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewTokenDeleteParams creates a new TokenDeleteParams object
// with the default values initialized.
func NewTokenDeleteParams() *TokenDeleteParams {
	var ()
	return &TokenDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTokenDeleteParamsWithTimeout creates a new TokenDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTokenDeleteParamsWithTimeout(timeout time.Duration) *TokenDeleteParams {
	var ()
	return &TokenDeleteParams{

		timeout: timeout,
	}
}

// NewTokenDeleteParamsWithContext creates a new TokenDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewTokenDeleteParamsWithContext(ctx context.Context) *TokenDeleteParams {
	var ()
	return &TokenDeleteParams{

		Context: ctx,
	}
}

// NewTokenDeleteParamsWithHTTPClient creates a new TokenDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTokenDeleteParamsWithHTTPClient(client *http.Client) *TokenDeleteParams {
	var ()
	return &TokenDeleteParams{
		HTTPClient: client,
	}
}

/*TokenDeleteParams contains all the parameters to send to the API endpoint
for the token delete operation typically these are written to a http.Request
*/
type TokenDeleteParams struct {

	/*Account
	  The unique ID for your IBM Cloud account.

	*/
	Account string
	/*Authorization
	  The IBM Cloud IAM access token that you receive when you log in to IBM Cloud. Run `bx iam oauth-tokens` to retrieve your access token, which is prefixed with 'Bearer'.

	*/
	Authorization string
	/*Token
	  The token that you want to delete.

	*/
	Token string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the token delete params
func (o *TokenDeleteParams) WithTimeout(timeout time.Duration) *TokenDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the token delete params
func (o *TokenDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the token delete params
func (o *TokenDeleteParams) WithContext(ctx context.Context) *TokenDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the token delete params
func (o *TokenDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the token delete params
func (o *TokenDeleteParams) WithHTTPClient(client *http.Client) *TokenDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the token delete params
func (o *TokenDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccount adds the account to the token delete params
func (o *TokenDeleteParams) WithAccount(account string) *TokenDeleteParams {
	o.SetAccount(account)
	return o
}

// SetAccount adds the account to the token delete params
func (o *TokenDeleteParams) SetAccount(account string) {
	o.Account = account
}

// WithAuthorization adds the authorization to the token delete params
func (o *TokenDeleteParams) WithAuthorization(authorization string) *TokenDeleteParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the token delete params
func (o *TokenDeleteParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithToken adds the token to the token delete params
func (o *TokenDeleteParams) WithToken(token string) *TokenDeleteParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the token delete params
func (o *TokenDeleteParams) SetToken(token string) {
	o.Token = token
}

// WriteToRequest writes these params to a request
func (o *TokenDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param token
	if err := r.SetPathParam("token", o.Token); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
