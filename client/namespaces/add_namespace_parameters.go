

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

// NewAddNamespaceParams creates a new AddNamespaceParams object
// with the default values initialized.
func NewAddNamespaceParams() *AddNamespaceParams {
	var ()
	return &AddNamespaceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddNamespaceParamsWithTimeout creates a new AddNamespaceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddNamespaceParamsWithTimeout(timeout time.Duration) *AddNamespaceParams {
	var ()
	return &AddNamespaceParams{

		timeout: timeout,
	}
}

// NewAddNamespaceParamsWithContext creates a new AddNamespaceParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddNamespaceParamsWithContext(ctx context.Context) *AddNamespaceParams {
	var ()
	return &AddNamespaceParams{

		Context: ctx,
	}
}

// NewAddNamespaceParamsWithHTTPClient creates a new AddNamespaceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddNamespaceParamsWithHTTPClient(client *http.Client) *AddNamespaceParams {
	var ()
	return &AddNamespaceParams{
		HTTPClient: client,
	}
}

/*AddNamespaceParams contains all the parameters to send to the API endpoint
for the add namespace operation typically these are written to a http.Request
*/
type AddNamespaceParams struct {

	/*Account
	  The unique ID for your IBM Cloud account.

	*/
	Account string
	/*Authorization
	  The IBM Cloud IAM access token that you receive when you log in to IBM Cloud. Run `bx iam oauth-tokens` to retrieve your access token, which is prefixed with 'Bearer'.

	*/
	Authorization string
	/*Namespace
	  Adds the specified namespace to your IBM Cloud account.

	*/
	Namespace string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add namespace params
func (o *AddNamespaceParams) WithTimeout(timeout time.Duration) *AddNamespaceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add namespace params
func (o *AddNamespaceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add namespace params
func (o *AddNamespaceParams) WithContext(ctx context.Context) *AddNamespaceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add namespace params
func (o *AddNamespaceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add namespace params
func (o *AddNamespaceParams) WithHTTPClient(client *http.Client) *AddNamespaceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add namespace params
func (o *AddNamespaceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccount adds the account to the add namespace params
func (o *AddNamespaceParams) WithAccount(account string) *AddNamespaceParams {
	o.SetAccount(account)
	return o
}

// SetAccount adds the account to the add namespace params
func (o *AddNamespaceParams) SetAccount(account string) {
	o.Account = account
}

// WithAuthorization adds the authorization to the add namespace params
func (o *AddNamespaceParams) WithAuthorization(authorization string) *AddNamespaceParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the add namespace params
func (o *AddNamespaceParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithNamespace adds the namespace to the add namespace params
func (o *AddNamespaceParams) WithNamespace(namespace string) *AddNamespaceParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the add namespace params
func (o *AddNamespaceParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WriteToRequest writes these params to a request
func (o *AddNamespaceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
