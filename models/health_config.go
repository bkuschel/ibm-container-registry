

package models




import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// HealthConfig health config
// model HealthConfig
type HealthConfig struct {

	// interval
	Interval TimeDuration `json:"Interval,omitempty"`

	// retries
	Retries int32 `json:"Retries,omitempty"`

	// test
	Test []string `json:"Test"`

	// timeout
	Timeout TimeDuration `json:"Timeout,omitempty"`
}

// Validate validates this health config
func (m *HealthConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HealthConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HealthConfig) UnmarshalBinary(b []byte) error {
	var res HealthConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
