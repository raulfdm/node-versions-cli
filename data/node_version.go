package data

type Lts interface{}

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

func (n NodeVersion) IsLts() bool {
	switch n.Lts.(type) {
	case bool:
		return false
	default:
		return true
	}
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
	var ltsVersions []string = []string{}

	for _, version := range n {
		if version.IsLts() {
			ltsVersions = append(ltsVersions, version.Version)
		}
	}

	return ltsVersions
}
