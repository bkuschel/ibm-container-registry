

package models




import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// SecureConfig secure config
// model SecureConfig
type SecureConfig struct {

	// correct output
	CorrectOutput int32 `json:"correct_output,omitempty"`

	// misconfigured
	Misconfigured int32 `json:"misconfigured,omitempty"`

	// total output docs
	TotalOutputDocs int32 `json:"total_output_docs,omitempty"`
}

// Validate validates this secure config
func (m *SecureConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SecureConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SecureConfig) UnmarshalBinary(b []byte) error {
	var res SecureConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
