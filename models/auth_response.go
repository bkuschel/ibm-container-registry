

package models




import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// AuthResponse auth response
// model AuthResponse
type AuthResponse struct {

	// Enable role based authorization when authenticating with IBM Cloud IAM
	IamAuthz bool `json:"iam_authz,omitempty"`
}

// Validate validates this auth response
func (m *AuthResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AuthResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AuthResponse) UnmarshalBinary(b []byte) error {
	var res AuthResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
