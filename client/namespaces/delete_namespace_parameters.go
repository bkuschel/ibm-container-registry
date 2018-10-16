

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

// NewDeleteNamespaceParams creates a new DeleteNamespaceParams object
// with the default values initialized.
func NewDeleteNamespaceParams() *DeleteNamespaceParams {
	var ()
	return &DeleteNamespaceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteNamespaceParamsWithTimeout creates a new DeleteNamespaceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteNamespaceParamsWithTimeout(timeout time.Duration) *DeleteNamespaceParams {
	var ()
	return &DeleteNamespaceParams{

		timeout: timeout,
	}
}

// NewDeleteNamespaceParamsWithContext creates a new DeleteNamespaceParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteNamespaceParamsWithContext(ctx context.Context) *DeleteNamespaceParams {
	var ()
	return &DeleteNamespaceParams{

		Context: ctx,
	}
}

// NewDeleteNamespaceParamsWithHTTPClient creates a new DeleteNamespaceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteNamespaceParamsWithHTTPClient(client *http.Client) *DeleteNamespaceParams {
	var ()
	return &DeleteNamespaceParams{
		HTTPClient: client,
	}
}

/*DeleteNamespaceParams contains all the parameters to send to the API endpoint
for the delete namespace operation typically these are written to a http.Request
*/
type DeleteNamespaceParams struct {

	/*Account
	  The unique ID for your IBM Cloud account.

	*/
	Account string
	/*Authorization
	  The IBM Cloud IAM access token that you receive when you log in to IBM Cloud. Run `bx iam oauth-tokens` to retrieve your access token, which is prefixed with 'Bearer'.

	*/
	Authorization string
	/*Namespace
	  The namespace that you want to delete.

	*/
	Namespace string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete namespace params
func (o *DeleteNamespaceParams) WithTimeout(timeout time.Duration) *DeleteNamespaceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete namespace params
func (o *DeleteNamespaceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete namespace params
func (o *DeleteNamespaceParams) WithContext(ctx context.Context) *DeleteNamespaceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete namespace params
func (o *DeleteNamespaceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete namespace params
func (o *DeleteNamespaceParams) WithHTTPClient(client *http.Client) *DeleteNamespaceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete namespace params
func (o *DeleteNamespaceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccount adds the account to the delete namespace params
func (o *DeleteNamespaceParams) WithAccount(account string) *DeleteNamespaceParams {
	o.SetAccount(account)
	return o
}

// SetAccount adds the account to the delete namespace params
func (o *DeleteNamespaceParams) SetAccount(account string) {
	o.Account = account
}

// WithAuthorization adds the authorization to the delete namespace params
func (o *DeleteNamespaceParams) WithAuthorization(authorization string) *DeleteNamespaceParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the delete namespace params
func (o *DeleteNamespaceParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithNamespace adds the namespace to the delete namespace params
func (o *DeleteNamespaceParams) WithNamespace(namespace string) *DeleteNamespaceParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the delete namespace params
func (o *DeleteNamespaceParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WriteToRequest writes these params to a request
func (o *DeleteNamespaceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
