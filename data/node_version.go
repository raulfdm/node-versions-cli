package data

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
	Lts      bool     `json:"lts"`
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
