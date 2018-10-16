

package models




import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// Progressdetail progressdetail
// model progressdetail
type Progressdetail struct {

	// current
	Current string `json:"current,omitempty"`

	// total
	Total string `json:"total,omitempty"`
}

// Validate validates this progressdetail
func (m *Progressdetail) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Progressdetail) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Progressdetail) UnmarshalBinary(b []byte) error {
	var res Progressdetail
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
