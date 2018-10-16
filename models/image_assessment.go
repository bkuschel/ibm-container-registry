

package models




import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ImageAssessment image assessment
// model ImageAssessment
type ImageAssessment struct {

	// assessment
	Assessment *ImageAssessmentResult `json:"assessment,omitempty"`

	// detail
	Detail *ImageDetails `json:"detail,omitempty"`

	// metadata
	Metadata *ImageMetadata `json:"metadata,omitempty"`

	// summary
	Summary *ImageSummary `json:"summary,omitempty"`
}

// Validate validates this image assessment
func (m *ImageAssessment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAssessment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDetail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSummary(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ImageAssessment) validateAssessment(formats strfmt.Registry) error {

	if swag.IsZero(m.Assessment) { // not required
		return nil
	}

	if m.Assessment != nil {
		if err := m.Assessment.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("assessment")
			}
			return err
		}
	}

	return nil
}

func (m *ImageAssessment) validateDetail(formats strfmt.Registry) error {

	if swag.IsZero(m.Detail) { // not required
		return nil
	}

	if m.Detail != nil {
		if err := m.Detail.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("detail")
			}
			return err
		}
	}

	return nil
}

func (m *ImageAssessment) validateMetadata(formats strfmt.Registry) error {

	if swag.IsZero(m.Metadata) { // not required
		return nil
	}

	if m.Metadata != nil {
		if err := m.Metadata.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metadata")
			}
			return err
		}
	}

	return nil
}

func (m *ImageAssessment) validateSummary(formats strfmt.Registry) error {

	if swag.IsZero(m.Summary) { // not required
		return nil
	}

	if m.Summary != nil {
		if err := m.Summary.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("summary")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ImageAssessment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ImageAssessment) UnmarshalBinary(b []byte) error {
	var res ImageAssessment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
