

package plans




import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new plans API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for plans API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetPlans gets plans for the targeted account
*/
func (a *Client) GetPlans(params *GetPlansParams) (*GetPlansOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPlansParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getPlans",
		Method:             "GET",
		PathPattern:        "/api/v1/plans",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetPlansReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPlansOK), nil

}

/*
SetPlan updates plans for the targeted account
*/
func (a *Client) SetPlan(params *SetPlanParams) (*SetPlanOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSetPlanParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "setPlan",
		Method:             "PATCH",
		PathPattern:        "/api/v1/plans",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SetPlanReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SetPlanOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
