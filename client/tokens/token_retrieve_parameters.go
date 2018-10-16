

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

// NewTokenRetrieveParams creates a new TokenRetrieveParams object
// with the default values initialized.
func NewTokenRetrieveParams() *TokenRetrieveParams {
	var ()
	return &TokenRetrieveParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTokenRetrieveParamsWithTimeout creates a new TokenRetrieveParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTokenRetrieveParamsWithTimeout(timeout time.Duration) *TokenRetrieveParams {
	var ()
	return &TokenRetrieveParams{

		timeout: timeout,
	}
}

// NewTokenRetrieveParamsWithContext creates a new TokenRetrieveParams object
// with the default values initialized, and the ability to set a context for a request
func NewTokenRetrieveParamsWithContext(ctx context.Context) *TokenRetrieveParams {
	var ()
	return &TokenRetrieveParams{

		Context: ctx,
	}
}

// NewTokenRetrieveParamsWithHTTPClient creates a new TokenRetrieveParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTokenRetrieveParamsWithHTTPClient(client *http.Client) *TokenRetrieveParams {
	var ()
	return &TokenRetrieveParams{
		HTTPClient: client,
	}
}

/*TokenRetrieveParams contains all the parameters to send to the API endpoint
for the token retrieve operation typically these are written to a http.Request
*/
type TokenRetrieveParams struct {

	/*Account
	  The unique ID for your IBM Cloud account.

	*/
	Account string
	/*Authorization
	  The IBM Cloud IAM access token that you receive when you log in to IBM Cloud. Run `bx iam oauth-tokens` to retrieve your access token, which is prefixed with 'Bearer'.

	*/
	Authorization string
	/*UUID
	  The uuid of the token you wish to retrieve

	*/
	UUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the token retrieve params
func (o *TokenRetrieveParams) WithTimeout(timeout time.Duration) *TokenRetrieveParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the token retrieve params
func (o *TokenRetrieveParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the token retrieve params
func (o *TokenRetrieveParams) WithContext(ctx context.Context) *TokenRetrieveParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the token retrieve params
func (o *TokenRetrieveParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the token retrieve params
func (o *TokenRetrieveParams) WithHTTPClient(client *http.Client) *TokenRetrieveParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the token retrieve params
func (o *TokenRetrieveParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccount adds the account to the token retrieve params
func (o *TokenRetrieveParams) WithAccount(account string) *TokenRetrieveParams {
	o.SetAccount(account)
	return o
}

// SetAccount adds the account to the token retrieve params
func (o *TokenRetrieveParams) SetAccount(account string) {
	o.Account = account
}

// WithAuthorization adds the authorization to the token retrieve params
func (o *TokenRetrieveParams) WithAuthorization(authorization string) *TokenRetrieveParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the token retrieve params
func (o *TokenRetrieveParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithUUID adds the uuid to the token retrieve params
func (o *TokenRetrieveParams) WithUUID(uuid string) *TokenRetrieveParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the token retrieve params
func (o *TokenRetrieveParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a request
func (o *TokenRetrieveParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param uuid
	if err := r.SetPathParam("uuid", o.UUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
