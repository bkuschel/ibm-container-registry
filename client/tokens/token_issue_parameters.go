

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

// NewTokenIssueParams creates a new TokenIssueParams object
// with the default values initialized.
func NewTokenIssueParams() *TokenIssueParams {
	var ()
	return &TokenIssueParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTokenIssueParamsWithTimeout creates a new TokenIssueParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTokenIssueParamsWithTimeout(timeout time.Duration) *TokenIssueParams {
	var ()
	return &TokenIssueParams{

		timeout: timeout,
	}
}

// NewTokenIssueParamsWithContext creates a new TokenIssueParams object
// with the default values initialized, and the ability to set a context for a request
func NewTokenIssueParamsWithContext(ctx context.Context) *TokenIssueParams {
	var ()
	return &TokenIssueParams{

		Context: ctx,
	}
}

// NewTokenIssueParamsWithHTTPClient creates a new TokenIssueParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTokenIssueParamsWithHTTPClient(client *http.Client) *TokenIssueParams {
	var ()
	return &TokenIssueParams{
		HTTPClient: client,
	}
}

/*TokenIssueParams contains all the parameters to send to the API endpoint
for the token issue operation typically these are written to a http.Request
*/
type TokenIssueParams struct {

	/*Account
	  The unique ID for your IBM Cloud account.

	*/
	Account string
	/*Authorization
	  The IBM Cloud IAM access token that you receive when you log in to IBM Cloud. Run `bx iam oauth-tokens` to retrieve your access token, which is prefixed with 'Bearer'.

	*/
	Authorization string
	/*Description
	  Specifies a description for the token so it can be more easily identified. If this option is specified more than once, the last parsed setting is the setting that is used.

	*/
	Description *string
	/*Permanent
	  When specified, the access token does not expire. If this option is specified more than once, the last parsed setting is the setting that is used.

	*/
	Permanent *string
	/*Write
	  When specified, the token provides write access to registry namespaces in your IBM Cloud account. If this option is not specified, or is set to false, the token provides read-only access. If this option is specified more than once, the last parsed setting is the setting that is used.

	*/
	Write *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the token issue params
func (o *TokenIssueParams) WithTimeout(timeout time.Duration) *TokenIssueParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the token issue params
func (o *TokenIssueParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the token issue params
func (o *TokenIssueParams) WithContext(ctx context.Context) *TokenIssueParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the token issue params
func (o *TokenIssueParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the token issue params
func (o *TokenIssueParams) WithHTTPClient(client *http.Client) *TokenIssueParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the token issue params
func (o *TokenIssueParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccount adds the account to the token issue params
func (o *TokenIssueParams) WithAccount(account string) *TokenIssueParams {
	o.SetAccount(account)
	return o
}

// SetAccount adds the account to the token issue params
func (o *TokenIssueParams) SetAccount(account string) {
	o.Account = account
}

// WithAuthorization adds the authorization to the token issue params
func (o *TokenIssueParams) WithAuthorization(authorization string) *TokenIssueParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the token issue params
func (o *TokenIssueParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithDescription adds the description to the token issue params
func (o *TokenIssueParams) WithDescription(description *string) *TokenIssueParams {
	o.SetDescription(description)
	return o
}

// SetDescription adds the description to the token issue params
func (o *TokenIssueParams) SetDescription(description *string) {
	o.Description = description
}

// WithPermanent adds the permanent to the token issue params
func (o *TokenIssueParams) WithPermanent(permanent *string) *TokenIssueParams {
	o.SetPermanent(permanent)
	return o
}

// SetPermanent adds the permanent to the token issue params
func (o *TokenIssueParams) SetPermanent(permanent *string) {
	o.Permanent = permanent
}

// WithWrite adds the write to the token issue params
func (o *TokenIssueParams) WithWrite(write *string) *TokenIssueParams {
	o.SetWrite(write)
	return o
}

// SetWrite adds the write to the token issue params
func (o *TokenIssueParams) SetWrite(write *string) {
	o.Write = write
}

// WriteToRequest writes these params to a request
func (o *TokenIssueParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Description != nil {

		// query param description
		var qrDescription string
		if o.Description != nil {
			qrDescription = *o.Description
		}
		qDescription := qrDescription
		if qDescription != "" {
			if err := r.SetQueryParam("description", qDescription); err != nil {
				return err
			}
		}

	}

	if o.Permanent != nil {

		// query param permanent
		var qrPermanent string
		if o.Permanent != nil {
			qrPermanent = *o.Permanent
		}
		qPermanent := qrPermanent
		if qPermanent != "" {
			if err := r.SetQueryParam("permanent", qPermanent); err != nil {
				return err
			}
		}

	}

	if o.Write != nil {

		// query param write
		var qrWrite string
		if o.Write != nil {
			qrWrite = *o.Write
		}
		qWrite := qrWrite
		if qWrite != "" {
			if err := r.SetQueryParam("write", qWrite); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
