

package namespaces




import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetNamespaceParams creates a new GetNamespaceParams object
// with the default values initialized.
func NewGetNamespaceParams() *GetNamespaceParams {
	var ()
	return &GetNamespaceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetNamespaceParamsWithTimeout creates a new GetNamespaceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetNamespaceParamsWithTimeout(timeout time.Duration) *GetNamespaceParams {
	var ()
	return &GetNamespaceParams{

		timeout: timeout,
	}
}

// NewGetNamespaceParamsWithContext creates a new GetNamespaceParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetNamespaceParamsWithContext(ctx context.Context) *GetNamespaceParams {
	var ()
	return &GetNamespaceParams{

		Context: ctx,
	}
}

// NewGetNamespaceParamsWithHTTPClient creates a new GetNamespaceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetNamespaceParamsWithHTTPClient(client *http.Client) *GetNamespaceParams {
	var ()
	return &GetNamespaceParams{
		HTTPClient: client,
	}
}

/*GetNamespaceParams contains all the parameters to send to the API endpoint
for the get namespace operation typically these are written to a http.Request
*/
type GetNamespaceParams struct {

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

// WithTimeout adds the timeout to the get namespace params
func (o *GetNamespaceParams) WithTimeout(timeout time.Duration) *GetNamespaceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get namespace params
func (o *GetNamespaceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get namespace params
func (o *GetNamespaceParams) WithContext(ctx context.Context) *GetNamespaceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get namespace params
func (o *GetNamespaceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get namespace params
func (o *GetNamespaceParams) WithHTTPClient(client *http.Client) *GetNamespaceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get namespace params
func (o *GetNamespaceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccount adds the account to the get namespace params
func (o *GetNamespaceParams) WithAccount(account string) *GetNamespaceParams {
	o.SetAccount(account)
	return o
}

// SetAccount adds the account to the get namespace params
func (o *GetNamespaceParams) SetAccount(account string) {
	o.Account = account
}

// WithAuthorization adds the authorization to the get namespace params
func (o *GetNamespaceParams) WithAuthorization(authorization string) *GetNamespaceParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the get namespace params
func (o *GetNamespaceParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WriteToRequest writes these params to a request
func (o *GetNamespaceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
