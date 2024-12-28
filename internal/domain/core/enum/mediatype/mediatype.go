package mediatype

import (
	"encoding/json"
)

var _ json.Unmarshaler = (*MediaType)(nil)
var _ json.Marshaler = (*MediaType)(nil)

type MediaType struct {
	mediaType string
}

var (
	Unknown = MediaType{}
	Image   = MediaType{"image"}
	Video   = MediaType{"video"}
	Audio   = MediaType{"audio"}
)

func (m MediaType) IsValid() bool {
	switch m {
	case Image, Video, Audio:
		return true
	default:
		return false
	}
}

func (m MediaType) String() string {
	return m.mediaType
}

func FromString(s string) MediaType {
	return MediaType{s}
}

func (m *MediaType) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.mediaType)
}

func (m *MediaType) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &m.mediaType)
}
