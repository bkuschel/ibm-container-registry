

package models




import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Buildoutput buildoutput
// model buildoutput
type Buildoutput struct {

	// id
	ID string `json:"id,omitempty"`

	// progress detail
	ProgressDetail *Progressdetail `json:"progressDetail,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// stream
	Stream string `json:"stream,omitempty"`
}

// Validate validates this buildoutput
func (m *Buildoutput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProgressDetail(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Buildoutput) validateProgressDetail(formats strfmt.Registry) error {

	if swag.IsZero(m.ProgressDetail) { // not required
		return nil
	}

	if m.ProgressDetail != nil {
		if err := m.ProgressDetail.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("progressDetail")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Buildoutput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Buildoutput) UnmarshalBinary(b []byte) error {
	var res Buildoutput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
