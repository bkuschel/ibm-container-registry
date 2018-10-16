

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

// NewTokenListOwnerParams creates a new TokenListOwnerParams object
// with the default values initialized.
func NewTokenListOwnerParams() *TokenListOwnerParams {
	var ()
	return &TokenListOwnerParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTokenListOwnerParamsWithTimeout creates a new TokenListOwnerParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTokenListOwnerParamsWithTimeout(timeout time.Duration) *TokenListOwnerParams {
	var ()
	return &TokenListOwnerParams{

		timeout: timeout,
	}
}

// NewTokenListOwnerParamsWithContext creates a new TokenListOwnerParams object
// with the default values initialized, and the ability to set a context for a request
func NewTokenListOwnerParamsWithContext(ctx context.Context) *TokenListOwnerParams {
	var ()
	return &TokenListOwnerParams{

		Context: ctx,
	}
}

// NewTokenListOwnerParamsWithHTTPClient creates a new TokenListOwnerParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTokenListOwnerParamsWithHTTPClient(client *http.Client) *TokenListOwnerParams {
	var ()
	return &TokenListOwnerParams{
		HTTPClient: client,
	}
}

/*TokenListOwnerParams contains all the parameters to send to the API endpoint
for the token list owner operation typically these are written to a http.Request
*/
type TokenListOwnerParams struct {

	/*Account
	  The unique ID for your IBM Cloud account.

	*/
	Account string
	/*Authorization
	  The IBM Cloud IAM access token that you receive when you log in to IBM Cloud. Run `bx iam oauth-tokens` to retrieve your access token, which is prefixed with 'Bearer'.

	*/
	Authorization string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the token list owner params
func (o *TokenListOwnerParams) WithTimeout(timeout time.Duration) *TokenListOwnerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the token list owner params
func (o *TokenListOwnerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the token list owner params
func (o *TokenListOwnerParams) WithContext(ctx context.Context) *TokenListOwnerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the token list owner params
func (o *TokenListOwnerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the token list owner params
func (o *TokenListOwnerParams) WithHTTPClient(client *http.Client) *TokenListOwnerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the token list owner params
func (o *TokenListOwnerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccount adds the account to the token list owner params
func (o *TokenListOwnerParams) WithAccount(account string) *TokenListOwnerParams {
	o.SetAccount(account)
	return o
}

// SetAccount adds the account to the token list owner params
func (o *TokenListOwnerParams) SetAccount(account string) {
	o.Account = account
}

// WithAuthorization adds the authorization to the token list owner params
func (o *TokenListOwnerParams) WithAuthorization(authorization string) *TokenListOwnerParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the token list owner params
func (o *TokenListOwnerParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WriteToRequest writes these params to a request
func (o *TokenListOwnerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
