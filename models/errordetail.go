

package models




import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// Errordetail errordetail
// model errordetail
type Errordetail struct {

	// code
	Code string `json:"code,omitempty"`

	// error
	Error string `json:"error,omitempty"`

	// error detail
	ErrorDetail string `json:"errorDetail,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// rc
	Rc int32 `json:"rc,omitempty"`

	// request id
	RequestID string `json:"request-id,omitempty"`
}

// Validate validates this errordetail
func (m *Errordetail) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Errordetail) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Errordetail) UnmarshalBinary(b []byte) error {
	var res Errordetail
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
