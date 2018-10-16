

package messages




import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new messages API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for messages API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetMessages returns any published system messages
*/
func (a *Client) GetMessages(params *GetMessagesParams) (*GetMessagesOK, *GetMessagesNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetMessagesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getMessages",
		Method:             "GET",
		PathPattern:        "/api/v1/messages",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetMessagesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *GetMessagesOK:
		return value, nil, nil
	case *GetMessagesNoContent:
		return nil, value, nil
	}
	return nil, nil, nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
