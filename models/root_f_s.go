

package models




import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// RootFS root f s
// model RootFS
type RootFS struct {

	// Descriptor for the base layer in the image
	BaseLayer string `json:"BaseLayer,omitempty"`

	// Descriptors for each layer in the image
	Layers []string `json:"Layers"`

	// The type of filesystem
	Type string `json:"Type,omitempty"`
}

// Validate validates this root f s
func (m *RootFS) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RootFS) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RootFS) UnmarshalBinary(b []byte) error {
	var res RootFS
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
