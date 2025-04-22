package usertype

import (
	"encoding/json"
)

var _ json.Unmarshaler = (*UserType)(nil)
var _ json.Marshaler = (*UserType)(nil)

type UserType struct {
	userType string
}

var (
	Unknown = UserType{}
	Admin   = UserType{"admin"}
	Viewer  = UserType{"viewer"}
)

func (m UserType) IsValid() bool {
	switch m {
	case Admin, Viewer:
		return true
	default:
		return false
	}
}

func (m UserType) String() string {
	return m.userType
}

func FromString(s string) UserType {
	return UserType{s}
}

func (m *UserType) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.userType)
}

func (m *UserType) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &m.userType)
}
