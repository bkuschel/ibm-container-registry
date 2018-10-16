

package models




import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ImageDetails image details
// model ImageDetails
type ImageDetails struct {

	// compliance
	Compliance []*ComplianceResult `json:"compliance"`

	// malware
	Malware []*ComplianceResult `json:"malware"`

	// vulnerability
	Vulnerability []*VulnerabilityDetails `json:"vulnerability"`
}

// Validate validates this image details
func (m *ImageDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCompliance(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMalware(formats); err != nil {
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

func (m *ImageDetails) validateCompliance(formats strfmt.Registry) error {

	if swag.IsZero(m.Compliance) { // not required
		return nil
	}

	for i := 0; i < len(m.Compliance); i++ {
		if swag.IsZero(m.Compliance[i]) { // not required
			continue
		}

		if m.Compliance[i] != nil {
			if err := m.Compliance[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("compliance" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ImageDetails) validateMalware(formats strfmt.Registry) error {

	if swag.IsZero(m.Malware) { // not required
		return nil
	}

	for i := 0; i < len(m.Malware); i++ {
		if swag.IsZero(m.Malware[i]) { // not required
			continue
		}

		if m.Malware[i] != nil {
			if err := m.Malware[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("malware" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ImageDetails) validateVulnerability(formats strfmt.Registry) error {

	if swag.IsZero(m.Vulnerability) { // not required
		return nil
	}

	for i := 0; i < len(m.Vulnerability); i++ {
		if swag.IsZero(m.Vulnerability[i]) { // not required
			continue
		}

		if m.Vulnerability[i] != nil {
			if err := m.Vulnerability[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("vulnerability" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ImageDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ImageDetails) UnmarshalBinary(b []byte) error {
	var res ImageDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
