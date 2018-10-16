

package models




import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// RemoteAPIImage remote API image
// model RemoteAPIImage
type RemoteAPIImage struct {

	// configuration issue count
	ConfigurationIssueCount int32 `json:"ConfigurationIssueCount,omitempty"`

	// created
	Created int64 `json:"Created,omitempty"`

	// digest tags
	DigestTags []Tags `json:"DigestTags"`

	// exempt issue count
	ExemptIssueCount int32 `json:"ExemptIssueCount,omitempty"`

	// Id
	ID string `json:"Id,omitempty"`

	// issue count
	IssueCount int32 `json:"IssueCount,omitempty"`

	// labels
	Labels []string `json:"Labels"`

	// parent Id
	ParentID string `json:"ParentId,omitempty"`

	// repo digests
	RepoDigests []string `json:"RepoDigests"`

	// repo tags
	RepoTags []string `json:"RepoTags"`

	// size
	Size int64 `json:"Size,omitempty"`

	// virtual size
	VirtualSize int64 `json:"VirtualSize,omitempty"`

	// vulnerability count
	VulnerabilityCount int32 `json:"VulnerabilityCount,omitempty"`

	// vulnerable
	Vulnerable string `json:"Vulnerable,omitempty"`
}

// Validate validates this remote API image
func (m *RemoteAPIImage) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RemoteAPIImage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RemoteAPIImage) UnmarshalBinary(b []byte) error {
	var res RemoteAPIImage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
