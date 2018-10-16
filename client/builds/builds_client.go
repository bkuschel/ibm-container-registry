

package builds




import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new builds API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for builds API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
ImageBuild builds a docker image
*/
func (a *Client) ImageBuild(params *ImageBuildParams) (*ImageBuildOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewImageBuildParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ImageBuild",
		Method:             "POST",
		PathPattern:        "/api/v1/builds",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ImageBuildReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ImageBuildOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
