

package auth




import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new auth API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for auth API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetAuth gets authorization options for the targeted account
*/
func (a *Client) GetAuth(params *GetAuthParams) (*GetAuthOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAuthParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getAuth",
		Method:             "GET",
		PathPattern:        "/api/v1/auth",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetAuthReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetAuthOK), nil

}

/*
SetAuth updates authorization options for the targeted account
*/
func (a *Client) SetAuth(params *SetAuthParams) (*SetAuthOK, *SetAuthNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSetAuthParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "setAuth",
		Method:             "PATCH",
		PathPattern:        "/api/v1/auth",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SetAuthReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *SetAuthOK:
		return value, nil, nil
	case *SetAuthNoContent:
		return nil, value, nil
	}
	return nil, nil, nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
