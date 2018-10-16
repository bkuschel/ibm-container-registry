

package models




import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// QuotaAPIResponse quota API response
// model QuotaAPIResponse
type QuotaAPIResponse struct {

	// limit
	Limit *QuotaResponse `json:"limit,omitempty"`

	// usage
	Usage *QuotaResponse `json:"usage,omitempty"`
}

// Validate validates this quota API response
func (m *QuotaAPIResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLimit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *QuotaAPIResponse) validateLimit(formats strfmt.Registry) error {

	if swag.IsZero(m.Limit) { // not required
		return nil
	}

	if m.Limit != nil {
		if err := m.Limit.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("limit")
			}
			return err
		}
	}

	return nil
}

func (m *QuotaAPIResponse) validateUsage(formats strfmt.Registry) error {

	if swag.IsZero(m.Usage) { // not required
		return nil
	}

	if m.Usage != nil {
		if err := m.Usage.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("usage")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *QuotaAPIResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *QuotaAPIResponse) UnmarshalBinary(b []byte) error {
	var res QuotaAPIResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
