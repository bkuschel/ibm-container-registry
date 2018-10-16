

package quotas




import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new quotas API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for quotas API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetQuota gets quotas for the targeted account
*/
func (a *Client) GetQuota(params *GetQuotaParams) (*GetQuotaOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetQuotaParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getQuota",
		Method:             "GET",
		PathPattern:        "/api/v1/quotas",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetQuotaReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetQuotaOK), nil

}

/*
SetQuota updates quotas for the targeted account
*/
func (a *Client) SetQuota(params *SetQuotaParams) (*SetQuotaOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSetQuotaParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "setQuota",
		Method:             "PATCH",
		PathPattern:        "/api/v1/quotas",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SetQuotaReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SetQuotaOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
