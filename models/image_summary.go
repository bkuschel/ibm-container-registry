

package models




import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ImageSummary image summary
// model ImageSummary
type ImageSummary struct {

	// compliance
	Compliance *Compliance `json:"compliance,omitempty"`

	// malware
	Malware *MalwareCompliance `json:"malware,omitempty"`

	// secureconfig
	Secureconfig *SecureConfig `json:"secureconfig,omitempty"`

	// vulnerability
	Vulnerability *Vulnerabilities `json:"vulnerability,omitempty"`
}

// Validate validates this image summary
func (m *ImageSummary) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCompliance(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMalware(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecureconfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVulnerability(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ImageSummary) validateCompliance(formats strfmt.Registry) error {

	if swag.IsZero(m.Compliance) { // not required
		return nil
	}

	if m.Compliance != nil {
		if err := m.Compliance.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("compliance")
			}
			return err
		}
	}

	return nil
}

func (m *ImageSummary) validateMalware(formats strfmt.Registry) error {

	if swag.IsZero(m.Malware) { // not required
		return nil
	}

	if m.Malware != nil {
		if err := m.Malware.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("malware")
			}
			return err
		}
	}

	return nil
}

func (m *ImageSummary) validateSecureconfig(formats strfmt.Registry) error {

	if swag.IsZero(m.Secureconfig) { // not required
		return nil
	}

	if m.Secureconfig != nil {
		if err := m.Secureconfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("secureconfig")
			}
			return err
		}
	}

	return nil
}

func (m *ImageSummary) validateVulnerability(formats strfmt.Registry) error {

	if swag.IsZero(m.Vulnerability) { // not required
		return nil
	}

	if m.Vulnerability != nil {
		if err := m.Vulnerability.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vulnerability")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ImageSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ImageSummary) UnmarshalBinary(b []byte) error {
	var res ImageSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
