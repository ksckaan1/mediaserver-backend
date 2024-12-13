package mediatype

type MediaType string

const (
	Image MediaType = "image"
	Video MediaType = "video"
	Audio MediaType = "audio"
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
	return string(m)
}

func FromString(s string) MediaType {
	switch s {
	case "image":
		return Image
	case "video":
		return Video
	case "audio":
		return Audio
	default:
		return MediaType(s)
	}
}
