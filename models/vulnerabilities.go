

package models




import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// Vulnerabilities vulnerabilities
// model Vulnerabilities
type Vulnerabilities struct {

	// total packages
	TotalPackages int32 `json:"total_packages,omitempty"`

	// total usns for distro
	TotalUsnsForDistro int32 `json:"total_usns_for_distro,omitempty"`

	// vulnerable packages
	VulnerablePackages int32 `json:"vulnerable_packages,omitempty"`

	// vulnerable usns
	VulnerableUsns int32 `json:"vulnerable_usns,omitempty"`
}

// Validate validates this vulnerabilities
func (m *Vulnerabilities) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Vulnerabilities) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Vulnerabilities) UnmarshalBinary(b []byte) error {
	var res Vulnerabilities
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
