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

func (m MediaType) Number() int32 {
	switch m {
	case Image:
		return 1
	case Video:
		return 2
	case Audio:
		return 3
	default:
		return 0
	}
}

func FromString(s string) MediaType {
	return MediaType{s}
}

func FromNumber(n int32) MediaType {
	switch n {
	case 1:
		return Image
	case 2:
		return Video
	case 3:
		return Audio
	default:
		return Unknown
	}
}

func (m *MediaType) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.mediaType)
}

func (m *MediaType) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &m.mediaType)
}
