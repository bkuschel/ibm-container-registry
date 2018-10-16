

package plans




import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetPlansParams creates a new GetPlansParams object
// with the default values initialized.
func NewGetPlansParams() *GetPlansParams {
	var ()
	return &GetPlansParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPlansParamsWithTimeout creates a new GetPlansParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPlansParamsWithTimeout(timeout time.Duration) *GetPlansParams {
	var ()
	return &GetPlansParams{

		timeout: timeout,
	}
}

// NewGetPlansParamsWithContext creates a new GetPlansParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPlansParamsWithContext(ctx context.Context) *GetPlansParams {
	var ()
	return &GetPlansParams{

		Context: ctx,
	}
}

// NewGetPlansParamsWithHTTPClient creates a new GetPlansParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPlansParamsWithHTTPClient(client *http.Client) *GetPlansParams {
	var ()
	return &GetPlansParams{
		HTTPClient: client,
	}
}

/*GetPlansParams contains all the parameters to send to the API endpoint
for the get plans operation typically these are written to a http.Request
*/
type GetPlansParams struct {

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

// WithTimeout adds the timeout to the get plans params
func (o *GetPlansParams) WithTimeout(timeout time.Duration) *GetPlansParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get plans params
func (o *GetPlansParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get plans params
func (o *GetPlansParams) WithContext(ctx context.Context) *GetPlansParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get plans params
func (o *GetPlansParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get plans params
func (o *GetPlansParams) WithHTTPClient(client *http.Client) *GetPlansParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get plans params
func (o *GetPlansParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccount adds the account to the get plans params
func (o *GetPlansParams) WithAccount(account string) *GetPlansParams {
	o.SetAccount(account)
	return o
}

// SetAccount adds the account to the get plans params
func (o *GetPlansParams) SetAccount(account string) {
	o.Account = account
}

// WithAuthorization adds the authorization to the get plans params
func (o *GetPlansParams) WithAuthorization(authorization string) *GetPlansParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the get plans params
func (o *GetPlansParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WriteToRequest writes these params to a request
func (o *GetPlansParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
