

package quotas




import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/bkuschel/ibm-container-registry/models"
)

// NewSetQuotaParams creates a new SetQuotaParams object
// with the default values initialized.
func NewSetQuotaParams() *SetQuotaParams {
	var ()
	return &SetQuotaParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSetQuotaParamsWithTimeout creates a new SetQuotaParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSetQuotaParamsWithTimeout(timeout time.Duration) *SetQuotaParams {
	var ()
	return &SetQuotaParams{

		timeout: timeout,
	}
}

// NewSetQuotaParamsWithContext creates a new SetQuotaParams object
// with the default values initialized, and the ability to set a context for a request
func NewSetQuotaParamsWithContext(ctx context.Context) *SetQuotaParams {
	var ()
	return &SetQuotaParams{

		Context: ctx,
	}
}

// NewSetQuotaParamsWithHTTPClient creates a new SetQuotaParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSetQuotaParamsWithHTTPClient(client *http.Client) *SetQuotaParams {
	var ()
	return &SetQuotaParams{
		HTTPClient: client,
	}
}

/*SetQuotaParams contains all the parameters to send to the API endpoint
for the set quota operation typically these are written to a http.Request
*/
type SetQuotaParams struct {

	/*Account
	  The unique ID for your IBM Cloud account.

	*/
	Account string
	/*Authorization
	  The IBM Cloud IAM access token that you receive when you log in to IBM Cloud. Run `bx iam oauth-tokens` to retrieve your access token, which is prefixed with 'Bearer'.

	*/
	Authorization string
	/*Quota
	  A JSON document containing the requested updates.

	*/
	Quota *models.QuotaResponse

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the set quota params
func (o *SetQuotaParams) WithTimeout(timeout time.Duration) *SetQuotaParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the set quota params
func (o *SetQuotaParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the set quota params
func (o *SetQuotaParams) WithContext(ctx context.Context) *SetQuotaParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the set quota params
func (o *SetQuotaParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the set quota params
func (o *SetQuotaParams) WithHTTPClient(client *http.Client) *SetQuotaParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the set quota params
func (o *SetQuotaParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccount adds the account to the set quota params
func (o *SetQuotaParams) WithAccount(account string) *SetQuotaParams {
	o.SetAccount(account)
	return o
}

// SetAccount adds the account to the set quota params
func (o *SetQuotaParams) SetAccount(account string) {
	o.Account = account
}

// WithAuthorization adds the authorization to the set quota params
func (o *SetQuotaParams) WithAuthorization(authorization string) *SetQuotaParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the set quota params
func (o *SetQuotaParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithQuota adds the quota to the set quota params
func (o *SetQuotaParams) WithQuota(quota *models.QuotaResponse) *SetQuotaParams {
	o.SetQuota(quota)
	return o
}

// SetQuota adds the quota to the set quota params
func (o *SetQuotaParams) SetQuota(quota *models.QuotaResponse) {
	o.Quota = quota
}

// WriteToRequest writes these params to a request
func (o *SetQuotaParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Quota != nil {
		if err := r.SetBodyParam(o.Quota); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
