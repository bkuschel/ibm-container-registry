

package models




import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ImageMetadata image metadata
// model ImageMetadata
type ImageMetadata struct {

	// complete
	Complete bool `json:"complete,omitempty"`

	// crawled time
	// Format: date-time
	CrawledTime strfmt.DateTime `json:"crawled_time,omitempty"`

	// namespace
	Namespace string `json:"namespace,omitempty"`

	// os supported
	OsSupported bool `json:"os_supported,omitempty"`
}

// Validate validates this image metadata
func (m *ImageMetadata) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCrawledTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ImageMetadata) validateCrawledTime(formats strfmt.Registry) error {

	if swag.IsZero(m.CrawledTime) { // not required
		return nil
	}

	if err := validate.FormatOf("crawled_time", "body", "date-time", m.CrawledTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ImageMetadata) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ImageMetadata) UnmarshalBinary(b []byte) error {
	var res ImageMetadata
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
