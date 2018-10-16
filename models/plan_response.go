

package models




import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// PlanResponse plan response
// model PlanResponse
type PlanResponse struct {

	// plan
	Plan string `json:"plan,omitempty"`
}

// Validate validates this plan response
func (m *PlanResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PlanResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PlanResponse) UnmarshalBinary(b []byte) error {
	var res PlanResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
