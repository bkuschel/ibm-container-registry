

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

// NewTokenDeleteOrgParams creates a new TokenDeleteOrgParams object
// with the default values initialized.
func NewTokenDeleteOrgParams() *TokenDeleteOrgParams {
	var ()
	return &TokenDeleteOrgParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTokenDeleteOrgParamsWithTimeout creates a new TokenDeleteOrgParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTokenDeleteOrgParamsWithTimeout(timeout time.Duration) *TokenDeleteOrgParams {
	var ()
	return &TokenDeleteOrgParams{

		timeout: timeout,
	}
}

// NewTokenDeleteOrgParamsWithContext creates a new TokenDeleteOrgParams object
// with the default values initialized, and the ability to set a context for a request
func NewTokenDeleteOrgParamsWithContext(ctx context.Context) *TokenDeleteOrgParams {
	var ()
	return &TokenDeleteOrgParams{

		Context: ctx,
	}
}

// NewTokenDeleteOrgParamsWithHTTPClient creates a new TokenDeleteOrgParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTokenDeleteOrgParamsWithHTTPClient(client *http.Client) *TokenDeleteOrgParams {
	var ()
	return &TokenDeleteOrgParams{
		HTTPClient: client,
	}
}

/*TokenDeleteOrgParams contains all the parameters to send to the API endpoint
for the token delete org operation typically these are written to a http.Request
*/
type TokenDeleteOrgParams struct {

	/*Account
	  The unique ID for your IBM Cloud account.

	*/
	Account *string
	/*Authorization
	  The IBM Cloud IAM access token that you receive when you log in to IBM Cloud. Run `bx iam oauth-tokens` to retrieve your access token, which is prefixed with 'Bearer'.

	*/
	Authorization string
	/*SecondaryOwner
	  Specifies a second owner field for the token. If this option is specified more than once, the last parsed setting is the setting that is used.

	*/
	SecondaryOwner string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the token delete org params
func (o *TokenDeleteOrgParams) WithTimeout(timeout time.Duration) *TokenDeleteOrgParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the token delete org params
func (o *TokenDeleteOrgParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the token delete org params
func (o *TokenDeleteOrgParams) WithContext(ctx context.Context) *TokenDeleteOrgParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the token delete org params
func (o *TokenDeleteOrgParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the token delete org params
func (o *TokenDeleteOrgParams) WithHTTPClient(client *http.Client) *TokenDeleteOrgParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the token delete org params
func (o *TokenDeleteOrgParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccount adds the account to the token delete org params
func (o *TokenDeleteOrgParams) WithAccount(account *string) *TokenDeleteOrgParams {
	o.SetAccount(account)
	return o
}

// SetAccount adds the account to the token delete org params
func (o *TokenDeleteOrgParams) SetAccount(account *string) {
	o.Account = account
}

// WithAuthorization adds the authorization to the token delete org params
func (o *TokenDeleteOrgParams) WithAuthorization(authorization string) *TokenDeleteOrgParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the token delete org params
func (o *TokenDeleteOrgParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithSecondaryOwner adds the secondaryOwner to the token delete org params
func (o *TokenDeleteOrgParams) WithSecondaryOwner(secondaryOwner string) *TokenDeleteOrgParams {
	o.SetSecondaryOwner(secondaryOwner)
	return o
}

// SetSecondaryOwner adds the secondaryOwner to the token delete org params
func (o *TokenDeleteOrgParams) SetSecondaryOwner(secondaryOwner string) {
	o.SecondaryOwner = secondaryOwner
}

// WriteToRequest writes these params to a request
func (o *TokenDeleteOrgParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Account != nil {

		// header param Account
		if err := r.SetHeaderParam("Account", *o.Account); err != nil {
			return err
		}

	}

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	// query param secondaryOwner
	qrSecondaryOwner := o.SecondaryOwner
	qSecondaryOwner := qrSecondaryOwner
	if qSecondaryOwner != "" {
		if err := r.SetQueryParam("secondaryOwner", qSecondaryOwner); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
