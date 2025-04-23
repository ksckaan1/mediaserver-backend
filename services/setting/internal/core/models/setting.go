package models

import "encoding/json"

type Setting struct {
	Key   string          `json:"key"`
	Value json.RawMessage `json:"value"`
}

type SettingList struct {
	List   []*Setting `json:"list"`
	Count  int64      `json:"count"`
	Limit  int64      `json:"limit"`
	Offset int64      `json:"offset"`
}
