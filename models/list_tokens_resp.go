

package models




import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ListTokensResp list tokens resp
// model ListTokensResp
type ListTokensResp struct {

	// tokens
	Tokens []*TokenDoc `json:"tokens"`
}

// Validate validates this list tokens resp
func (m *ListTokensResp) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTokens(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListTokensResp) validateTokens(formats strfmt.Registry) error {

	if swag.IsZero(m.Tokens) { // not required
		return nil
	}

	for i := 0; i < len(m.Tokens); i++ {
		if swag.IsZero(m.Tokens[i]) { // not required
			continue
		}

		if m.Tokens[i] != nil {
			if err := m.Tokens[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tokens" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ListTokensResp) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListTokensResp) UnmarshalBinary(b []byte) error {
	var res ListTokensResp
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
