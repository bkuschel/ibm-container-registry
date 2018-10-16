

package auth




import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetAuthParams creates a new GetAuthParams object
// with the default values initialized.
func NewGetAuthParams() *GetAuthParams {
	var ()
	return &GetAuthParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAuthParamsWithTimeout creates a new GetAuthParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAuthParamsWithTimeout(timeout time.Duration) *GetAuthParams {
	var ()
	return &GetAuthParams{

		timeout: timeout,
	}
}

// NewGetAuthParamsWithContext creates a new GetAuthParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAuthParamsWithContext(ctx context.Context) *GetAuthParams {
	var ()
	return &GetAuthParams{

		Context: ctx,
	}
}

// NewGetAuthParamsWithHTTPClient creates a new GetAuthParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAuthParamsWithHTTPClient(client *http.Client) *GetAuthParams {
	var ()
	return &GetAuthParams{
		HTTPClient: client,
	}
}

/*GetAuthParams contains all the parameters to send to the API endpoint
for the get auth operation typically these are written to a http.Request
*/
type GetAuthParams struct {

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

// WithTimeout adds the timeout to the get auth params
func (o *GetAuthParams) WithTimeout(timeout time.Duration) *GetAuthParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get auth params
func (o *GetAuthParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get auth params
func (o *GetAuthParams) WithContext(ctx context.Context) *GetAuthParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get auth params
func (o *GetAuthParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get auth params
func (o *GetAuthParams) WithHTTPClient(client *http.Client) *GetAuthParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get auth params
func (o *GetAuthParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccount adds the account to the get auth params
func (o *GetAuthParams) WithAccount(account string) *GetAuthParams {
	o.SetAccount(account)
	return o
}

// SetAccount adds the account to the get auth params
func (o *GetAuthParams) SetAccount(account string) {
	o.Account = account
}

// WithAuthorization adds the authorization to the get auth params
func (o *GetAuthParams) WithAuthorization(authorization string) *GetAuthParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the get auth params
func (o *GetAuthParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WriteToRequest writes these params to a request
func (o *GetAuthParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
