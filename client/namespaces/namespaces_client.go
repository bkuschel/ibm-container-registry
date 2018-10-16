

package namespaces




import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new namespaces API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for namespaces API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
AddNamespace adds a namespace to the targeted i b m cloud account
*/
func (a *Client) AddNamespace(params *AddNamespaceParams) (*AddNamespaceOK, *AddNamespaceCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddNamespaceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "addNamespace",
		Method:             "PUT",
		PathPattern:        "/api/v1/namespaces/{namespace}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AddNamespaceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *AddNamespaceOK:
		return value, nil, nil
	case *AddNamespaceCreated:
		return nil, value, nil
	}
	return nil, nil, nil

}

/*
DeleteNamespace deletes the i b m cloud container registry namespace from the targeted i b m cloud account and removes all images that were in that namespace
*/
func (a *Client) DeleteNamespace(params *DeleteNamespaceParams) (*DeleteNamespaceOK, *DeleteNamespaceNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteNamespaceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteNamespace",
		Method:             "DELETE",
		PathPattern:        "/api/v1/namespaces/{namespace}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteNamespaceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *DeleteNamespaceOK:
		return value, nil, nil
	case *DeleteNamespaceNoContent:
		return nil, value, nil
	}
	return nil, nil, nil

}

/*
GetNamespace lists authorized namespaces in the targeted i b m cloud account
*/
func (a *Client) GetNamespace(params *GetNamespaceParams) (*GetNamespaceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNamespaceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getNamespace",
		Method:             "GET",
		PathPattern:        "/api/v1/namespaces",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetNamespaceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetNamespaceOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
