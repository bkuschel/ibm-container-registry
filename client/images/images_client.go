

package images




import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new images API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for images API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
ImageDelete removes a docker image from the registry
*/
func (a *Client) ImageDelete(params *ImageDeleteParams) (*ImageDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewImageDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ImageDelete",
		Method:             "DELETE",
		PathPattern:        "/api/v1/images/{image}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ImageDeleteReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ImageDeleteOK), nil

}

/*
ImageInspect inspects a docker image in the private registry
*/
func (a *Client) ImageInspect(params *ImageInspectParams) (*ImageInspectOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewImageInspectParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "imageInspect",
		Method:             "GET",
		PathPattern:        "/api/v1/images/{image}/json",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ImageInspectReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ImageInspectOK), nil

}

/*
ImageList lists all images in namespaces in a targeted i b m cloud account
*/
func (a *Client) ImageList(params *ImageListParams) (*ImageListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewImageListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "imageList",
		Method:             "GET",
		PathPattern:        "/api/v1/images",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ImageListReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ImageListOK), nil

}

/*
ImageVulnerabilities views the vulnerability report for a docker image in the private registry
*/
func (a *Client) ImageVulnerabilities(params *ImageVulnerabilitiesParams) (*ImageVulnerabilitiesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewImageVulnerabilitiesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "imageVulnerabilities",
		Method:             "GET",
		PathPattern:        "/api/v1/images/{image}/vulnerabilities",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ImageVulnerabilitiesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ImageVulnerabilitiesOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
