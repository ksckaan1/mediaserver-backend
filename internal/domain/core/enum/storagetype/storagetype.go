package storagetype

import "encoding/json"

var _ json.Unmarshaler = (*StorageType)(nil)
var _ json.Marshaler = (*StorageType)(nil)

type StorageType struct {
	storageType string
}

var (
	Unknown = StorageType{}
	Local   = StorageType{"local"}
)

func (s StorageType) String() string {
	return s.storageType
}

func (s StorageType) IsValid() bool {
	return s == Local
}

func FromString(s string) StorageType {
	return StorageType{s}
}

func (m *StorageType) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.storageType)
}

func (m *StorageType) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &m.storageType)
}
