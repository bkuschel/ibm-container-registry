

package models




import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ImageInspection image inspection
// model ImageInspection
type ImageInspection struct {

	// The processor architecture used to build this image, and required to run it
	Architecture string `json:"Architecture,omitempty"`

	// The author of the image
	Author string `json:"Author,omitempty"`

	// A plain text description of the image
	Comment string `json:"Comment,omitempty"`

	// Configuration metadata for the image
	Config *Config `json:"Config,omitempty"`

	// The ID of the container which created this image
	Container string `json:"Container,omitempty"`

	// The default configuration for containers started from this image
	ContainerConfig *Config `json:"ContainerConfig,omitempty"`

	// The unix timestamp for the date when the image was created
	Created string `json:"Created,omitempty"`

	// The Docker version used to build this image
	DockerVersion string `json:"DockerVersion,omitempty"`

	// The image ID
	ID string `json:"Id,omitempty"`

	// The operating system family used to build this image, and required to run it
	Os string `json:"Os,omitempty"`

	// The version of the operating system used to build this image
	OsVersion string `json:"OsVersion,omitempty"`

	// The ID of the base image for this image
	Parent string `json:"Parent,omitempty"`

	// Metadata describing the root filesystem for the image
	RootFS *RootFS `json:"RootFS,omitempty"`

	// The size of the image in bytes
	Size int64 `json:"Size,omitempty"`

	// The sum of the size of each layer in the image in bytes
	VirtualSize int64 `json:"VirtualSize,omitempty"`
}

// Validate validates this image inspection
func (m *ImageInspection) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContainerConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRootFS(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ImageInspection) validateConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.Config) { // not required
		return nil
	}

	if m.Config != nil {
		if err := m.Config.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Config")
			}
			return err
		}
	}

	return nil
}

func (m *ImageInspection) validateContainerConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.ContainerConfig) { // not required
		return nil
	}

	if m.ContainerConfig != nil {
		if err := m.ContainerConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ContainerConfig")
			}
			return err
		}
	}

	return nil
}

func (m *ImageInspection) validateRootFS(formats strfmt.Registry) error {

	if swag.IsZero(m.RootFS) { // not required
		return nil
	}

	if m.RootFS != nil {
		if err := m.RootFS.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("RootFS")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ImageInspection) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ImageInspection) UnmarshalBinary(b []byte) error {
	var res ImageInspection
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
