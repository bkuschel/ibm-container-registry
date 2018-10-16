

package models




import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// TokenDoc token doc
// model TokenDoc
type TokenDoc struct {

	// id
	ID string `json:"_id,omitempty"`

	// encrypted token
	EncryptedToken string `json:"encrypted_token,omitempty"`

	// expiry
	Expiry int64 `json:"expiry,omitempty"`

	// owner
	Owner string `json:"owner,omitempty"`

	// readonly
	Readonly bool `json:"readonly,omitempty"`

	// revoked
	Revoked bool `json:"revoked,omitempty"`

	// secondary owner
	SecondaryOwner string `json:"secondary_owner,omitempty"`

	// token
	Token string `json:"token,omitempty"`
}

// Validate validates this token doc
func (m *TokenDoc) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TokenDoc) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TokenDoc) UnmarshalBinary(b []byte) error {
	var res TokenDoc
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
