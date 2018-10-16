

package models




import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Config config
// model Config
type Config struct {

	// True if command is already escaped (Windows specific)
	ArgsEscaped bool `json:"ArgsEscaped,omitempty"`

	// If true, standard error is attached
	AttachStderr bool `json:"AttachStderr,omitempty"`

	// If true, standard input is attached, which makes possible user interaction
	AttachStdin bool `json:"AttachStdin,omitempty"`

	// If true, standard output is attached
	AttachStdout bool `json:"AttachStdout,omitempty"`

	// Command that is run when starting the container
	Cmd []string `json:"Cmd"`

	// The FQDN for the container
	Domainname string `json:"Domainname,omitempty"`

	// Entrypoint to run when starting the container
	Entrypoint []string `json:"Entrypoint"`

	// List of environment variables to set in the container
	Env []string `json:"Env"`

	// A list of exposed ports in a format [123:{},456:{}]
	ExposedPorts []Set `json:"ExposedPorts"`

	// Describes how to check that the container is healthy
	Healthcheck *HealthConfig `json:"Healthcheck,omitempty"`

	// The host name of the container
	Hostname string `json:"Hostname,omitempty"`

	// Name of the image as it was passed by the operator (eg. could be symbolic)
	Image string `json:"Image,omitempty"`

	// List of labels set to this container
	Labels []string `json:"Labels"`

	// The MAC Address of the container
	MacAddress string `json:"MacAddress,omitempty"`

	// If true, containers are not given network access
	NetworkDisabled bool `json:"NetworkDisabled,omitempty"`

	// ONBUILD metadata that were defined on the image Dockerfile https://docs.docker.com/engine/reference/builder/#onbuild
	OnBuild []string `json:"OnBuild"`

	// Open stdin
	OpenStdin bool `json:"OpenStdin,omitempty"`

	// Shell for shell-form of RUN, CMD, ENTRYPOINT
	Shell []string `json:"Shell"`

	// If true, close stdin after the 1 attached client disconnects.
	StdinOnce bool `json:"StdinOnce,omitempty"`

	// Signal to stop a container
	StopSignal string `json:"StopSignal,omitempty"`

	// Timeout (in seconds) to stop a container
	StopTimeout int32 `json:"StopTimeout,omitempty"`

	// Attach standard streams to a tty, including stdin if it is not closed.
	Tty bool `json:"Tty,omitempty"`

	// The user that will run the command(s) inside the container
	User string `json:"User,omitempty"`

	// List of volumes (mounts) used for the container
	Volumes []Set `json:"Volumes"`

	// Current working directory (PWD) in the command will be launched
	WorkingDir string `json:"WorkingDir,omitempty"`
}

// Validate validates this config
func (m *Config) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHealthcheck(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Config) validateHealthcheck(formats strfmt.Registry) error {

	if swag.IsZero(m.Healthcheck) { // not required
		return nil
	}

	if m.Healthcheck != nil {
		if err := m.Healthcheck.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Healthcheck")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Config) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Config) UnmarshalBinary(b []byte) error {
	var res Config
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
