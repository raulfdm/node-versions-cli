package data

import (
	"encoding/json"
	"fmt"
)

type Lts struct {
	BoolValue   bool
	StringValue string
	IsBool      bool
}

func (lts Lts) UnmarshalJSON(data []byte) error {
	var boolVal bool

	if err := json.Unmarshal(data, &boolVal); err == nil {
		lts.BoolValue = boolVal
		lts.IsBool = true
		return nil
	}

	var strVal string
	if err := json.Unmarshal(data, &strVal); err == nil {
		lts.StringValue = strVal
		lts.IsBool = false
		return nil
	}

	return fmt.Errorf("lts must be a boolean or a string")
}

type NodeVersion struct {
	Version  string   `json:"version"`
	Date     string   `json:"date"`
	Files    []string `json:"files"`
	Npm      string   `json:"npm,omitempty"`
	V8       string   `json:"v8"`
	Uv       string   `json:"uv,omitempty"`
	Zlib     string   `json:"zlib,omitempty"`
	Openssl  string   `json:"openssl,omitempty"`
	Modules  string   `json:"modules,omitempty"`
	Lts      Lts      `json:"lts"`
	Security bool     `json:"security"`
}

type NodeVersions []NodeVersion

func (n NodeVersions) GetAll() []string {
	var allVersions []string
	for _, version := range n {
		allVersions = append(allVersions, version.Version)
	}

	return allVersions
}

func (n NodeVersions) GetLatest() string {
	return n[0].Version
}

func (n NodeVersions) GetCurrentLts() string {

	allLts := n.GetAllLts()

	return allLts[0]
}

func (n NodeVersions) GetAllLts() []string {
	var ltsVersions []string

	for _, version := range n {

		if !version.Lts.IsBool {
			ltsVersions = append(ltsVersions, version.Version)
		}
	}

	return ltsVersions
}
