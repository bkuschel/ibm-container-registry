

package plans




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

// NewSetPlanParams creates a new SetPlanParams object
// with the default values initialized.
func NewSetPlanParams() *SetPlanParams {
	var ()
	return &SetPlanParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSetPlanParamsWithTimeout creates a new SetPlanParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSetPlanParamsWithTimeout(timeout time.Duration) *SetPlanParams {
	var ()
	return &SetPlanParams{

		timeout: timeout,
	}
}

// NewSetPlanParamsWithContext creates a new SetPlanParams object
// with the default values initialized, and the ability to set a context for a request
func NewSetPlanParamsWithContext(ctx context.Context) *SetPlanParams {
	var ()
	return &SetPlanParams{

		Context: ctx,
	}
}

// NewSetPlanParamsWithHTTPClient creates a new SetPlanParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSetPlanParamsWithHTTPClient(client *http.Client) *SetPlanParams {
	var ()
	return &SetPlanParams{
		HTTPClient: client,
	}
}

/*SetPlanParams contains all the parameters to send to the API endpoint
for the set plan operation typically these are written to a http.Request
*/
type SetPlanParams struct {

	/*Account
	  The unique ID for your IBM Cloud account.

	*/
	Account string
	/*Authorization
	  The IBM Cloud IAM access token that you receive when you log in to IBM Cloud. Run `bx iam oauth-tokens` to retrieve your access token, which is prefixed with 'Bearer'.

	*/
	Authorization string
	/*Plan
	  A JSON document containing the requested updates.

	*/
	Plan *models.PlanResponse

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the set plan params
func (o *SetPlanParams) WithTimeout(timeout time.Duration) *SetPlanParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the set plan params
func (o *SetPlanParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the set plan params
func (o *SetPlanParams) WithContext(ctx context.Context) *SetPlanParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the set plan params
func (o *SetPlanParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the set plan params
func (o *SetPlanParams) WithHTTPClient(client *http.Client) *SetPlanParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the set plan params
func (o *SetPlanParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccount adds the account to the set plan params
func (o *SetPlanParams) WithAccount(account string) *SetPlanParams {
	o.SetAccount(account)
	return o
}

// SetAccount adds the account to the set plan params
func (o *SetPlanParams) SetAccount(account string) {
	o.Account = account
}

// WithAuthorization adds the authorization to the set plan params
func (o *SetPlanParams) WithAuthorization(authorization string) *SetPlanParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the set plan params
func (o *SetPlanParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithPlan adds the plan to the set plan params
func (o *SetPlanParams) WithPlan(plan *models.PlanResponse) *SetPlanParams {
	o.SetPlan(plan)
	return o
}

// SetPlan adds the plan to the set plan params
func (o *SetPlanParams) SetPlan(plan *models.PlanResponse) {
	o.Plan = plan
}

// WriteToRequest writes these params to a request
func (o *SetPlanParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Plan != nil {
		if err := r.SetBodyParam(o.Plan); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
