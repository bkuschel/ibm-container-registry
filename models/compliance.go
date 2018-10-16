

package models




import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// Compliance compliance
// model Compliance
type Compliance struct {

	// compliance violations
	ComplianceViolations int32 `json:"compliance_violations,omitempty"`

	// compliant
	Compliant bool `json:"compliant,omitempty"`

	// execution status
	ExecutionStatus string `json:"execution_status,omitempty"`

	// reason
	Reason string `json:"reason,omitempty"`

	// total compliance rules
	TotalComplianceRules int32 `json:"total_compliance_rules,omitempty"`
}

// Validate validates this compliance
func (m *Compliance) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Compliance) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Compliance) UnmarshalBinary(b []byte) error {
	var res Compliance
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
