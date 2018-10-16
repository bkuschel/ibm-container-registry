

package quotas




import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetQuotaParams creates a new GetQuotaParams object
// with the default values initialized.
func NewGetQuotaParams() *GetQuotaParams {
	var ()
	return &GetQuotaParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetQuotaParamsWithTimeout creates a new GetQuotaParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetQuotaParamsWithTimeout(timeout time.Duration) *GetQuotaParams {
	var ()
	return &GetQuotaParams{

		timeout: timeout,
	}
}

// NewGetQuotaParamsWithContext creates a new GetQuotaParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetQuotaParamsWithContext(ctx context.Context) *GetQuotaParams {
	var ()
	return &GetQuotaParams{

		Context: ctx,
	}
}

// NewGetQuotaParamsWithHTTPClient creates a new GetQuotaParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetQuotaParamsWithHTTPClient(client *http.Client) *GetQuotaParams {
	var ()
	return &GetQuotaParams{
		HTTPClient: client,
	}
}

/*GetQuotaParams contains all the parameters to send to the API endpoint
for the get quota operation typically these are written to a http.Request
*/
type GetQuotaParams struct {

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

// WithTimeout adds the timeout to the get quota params
func (o *GetQuotaParams) WithTimeout(timeout time.Duration) *GetQuotaParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get quota params
func (o *GetQuotaParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get quota params
func (o *GetQuotaParams) WithContext(ctx context.Context) *GetQuotaParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get quota params
func (o *GetQuotaParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get quota params
func (o *GetQuotaParams) WithHTTPClient(client *http.Client) *GetQuotaParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get quota params
func (o *GetQuotaParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccount adds the account to the get quota params
func (o *GetQuotaParams) WithAccount(account string) *GetQuotaParams {
	o.SetAccount(account)
	return o
}

// SetAccount adds the account to the get quota params
func (o *GetQuotaParams) SetAccount(account string) {
	o.Account = account
}

// WithAuthorization adds the authorization to the get quota params
func (o *GetQuotaParams) WithAuthorization(authorization string) *GetQuotaParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the get quota params
func (o *GetQuotaParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WriteToRequest writes these params to a request
func (o *GetQuotaParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
