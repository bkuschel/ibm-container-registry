

package models




import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// ImageDeleteResult image delete result
// model ImageDeleteResult
type ImageDeleteResult struct {

	// untagged
	Untagged string `json:"Untagged,omitempty"`
}

// Validate validates this image delete result
func (m *ImageDeleteResult) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ImageDeleteResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ImageDeleteResult) UnmarshalBinary(b []byte) error {
	var res ImageDeleteResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
