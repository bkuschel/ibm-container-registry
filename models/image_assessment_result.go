

package models




import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ImageAssessmentResult image assessment result
// model ImageAssessmentResult
type ImageAssessmentResult struct {

	// raw verdict results
	RawVerdictResults []*ComplianceResult `json:"raw_verdict_results"`

	// vulnerable
	Vulnerable bool `json:"vulnerable,omitempty"`
}

// Validate validates this image assessment result
func (m *ImageAssessmentResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRawVerdictResults(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ImageAssessmentResult) validateRawVerdictResults(formats strfmt.Registry) error {

	if swag.IsZero(m.RawVerdictResults) { // not required
		return nil
	}

	for i := 0; i < len(m.RawVerdictResults); i++ {
		if swag.IsZero(m.RawVerdictResults[i]) { // not required
			continue
		}

		if m.RawVerdictResults[i] != nil {
			if err := m.RawVerdictResults[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("raw_verdict_results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ImageAssessmentResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ImageAssessmentResult) UnmarshalBinary(b []byte) error {
	var res ImageAssessmentResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
