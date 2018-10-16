

package models




import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// RequestError request error
// model RequestError
type RequestError struct {

	// code
	Code string `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// rc
	Rc int32 `json:"rc,omitempty"`

	// request id
	RequestID string `json:"request-id,omitempty"`
}

// Validate validates this request error
func (m *RequestError) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RequestError) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RequestError) UnmarshalBinary(b []byte) error {
	var res RequestError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
