

package models




import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// ComplianceResult compliance result
// model ComplianceResult
type ComplianceResult struct {

	// compliant
	Compliant bool `json:"compliant,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// policy mandated
	PolicyMandated bool `json:"policy_mandated,omitempty"`

	// reason
	Reason string `json:"reason,omitempty"`
}

// Validate validates this compliance result
func (m *ComplianceResult) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ComplianceResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ComplianceResult) UnmarshalBinary(b []byte) error {
	var res ComplianceResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
