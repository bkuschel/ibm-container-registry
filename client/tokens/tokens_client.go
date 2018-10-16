

package tokens




import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new tokens API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for tokens API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
TokenDelete revokes the specified registry access token
*/
func (a *Client) TokenDelete(params *TokenDeleteParams) (*TokenDeleteOK, *TokenDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTokenDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "tokenDelete",
		Method:             "DELETE",
		PathPattern:        "/api/v1/tokens/{token}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &TokenDeleteReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *TokenDeleteOK:
		return value, nil, nil
	case *TokenDeleteNoContent:
		return nil, value, nil
	}
	return nil, nil, nil

}

/*
TokenDeleteOrg revokes all registry access tokens that match a query
*/
func (a *Client) TokenDeleteOrg(params *TokenDeleteOrgParams) (*TokenDeleteOrgOK, *TokenDeleteOrgNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTokenDeleteOrgParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "tokenDeleteOrg",
		Method:             "DELETE",
		PathPattern:        "/api/v1/tokens",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &TokenDeleteOrgReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *TokenDeleteOrgOK:
		return value, nil, nil
	case *TokenDeleteOrgNoContent:
		return nil, value, nil
	}
	return nil, nil, nil

}

/*
TokenIssue issues a registry access token for the targeted i b m cloud account
*/
func (a *Client) TokenIssue(params *TokenIssueParams) (*TokenIssueOK, *TokenIssueCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTokenIssueParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "tokenIssue",
		Method:             "POST",
		PathPattern:        "/api/v1/tokens",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &TokenIssueReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *TokenIssueOK:
		return value, nil, nil
	case *TokenIssueCreated:
		return nil, value, nil
	}
	return nil, nil, nil

}

/*
TokenListOwner lists all valid registry access tokens for an owner
*/
func (a *Client) TokenListOwner(params *TokenListOwnerParams) (*TokenListOwnerOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTokenListOwnerParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "tokenListOwner",
		Method:             "GET",
		PathPattern:        "/api/v1/tokens",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &TokenListOwnerReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*TokenListOwnerOK), nil

}

/*
TokenRetrieve retrieves the registry token with the given UUID
*/
func (a *Client) TokenRetrieve(params *TokenRetrieveParams) (*TokenRetrieveOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTokenRetrieveParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "tokenRetrieve",
		Method:             "GET",
		PathPattern:        "/api/v1/tokens/{uuid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &TokenRetrieveReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*TokenRetrieveOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
