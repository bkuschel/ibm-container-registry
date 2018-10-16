

package models




import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// QuotaResponse quota response
// model QuotaResponse
type QuotaResponse struct {

	// The storage amount in bytes
	Store uint64 `json:"store,omitempty"`

	// The traffic bandwidth amount in bytes
	Traffic uint64 `json:"traffic,omitempty"`
}

// Validate validates this quota response
func (m *QuotaResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *QuotaResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *QuotaResponse) UnmarshalBinary(b []byte) error {
	var res QuotaResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
