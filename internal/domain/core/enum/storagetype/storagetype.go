package storagetype

type StorageType string

const (
	Local StorageType = "local"
)

func (s StorageType) String() string {
	return string(s)
}

func (s StorageType) IsValid() bool {
	return s == Local
}

func FromString(s string) StorageType {
	switch s {
	case "local":
		return Local
	default:
		return StorageType(s)
	}
}
