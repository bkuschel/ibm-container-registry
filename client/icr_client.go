

package client




import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/bkuschel/ibm-container-registry/client/auth"
	"github.com/bkuschel/ibm-container-registry/client/builds"
	"github.com/bkuschel/ibm-container-registry/client/images"
	"github.com/bkuschel/ibm-container-registry/client/messages"
	"github.com/bkuschel/ibm-container-registry/client/namespaces"
	"github.com/bkuschel/ibm-container-registry/client/plans"
	"github.com/bkuschel/ibm-container-registry/client/quotas"
	"github.com/bkuschel/ibm-container-registry/client/tokens"
)

// Default icr HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "registry.ng.bluemix.net"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"http", "https"}

// NewHTTPClient creates a new icr HTTP client.
func NewHTTPClient(formats strfmt.Registry) *Icr {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new icr HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *Icr {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new icr client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Icr {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(Icr)
	cli.Transport = transport

	cli.Auth = auth.New(transport, formats)

	cli.Builds = builds.New(transport, formats)

	cli.Images = images.New(transport, formats)

	cli.Messages = messages.New(transport, formats)

	cli.Namespaces = namespaces.New(transport, formats)

	cli.Plans = plans.New(transport, formats)

	cli.Quotas = quotas.New(transport, formats)

	cli.Tokens = tokens.New(transport, formats)

	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// Icr is a client for icr
type Icr struct {
	Auth *auth.Client

	Builds *builds.Client

	Images *images.Client

	Messages *messages.Client

	Namespaces *namespaces.Client

	Plans *plans.Client

	Quotas *quotas.Client

	Tokens *tokens.Client

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *Icr) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport

	c.Auth.SetTransport(transport)

	c.Builds.SetTransport(transport)

	c.Images.SetTransport(transport)

	c.Messages.SetTransport(transport)

	c.Namespaces.SetTransport(transport)

	c.Plans.SetTransport(transport)

	c.Quotas.SetTransport(transport)

	c.Tokens.SetTransport(transport)

}
