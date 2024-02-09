package api

import (
	"fmt"
	"net/http"
	"node-versions-cli/data"
)

const nodeVersionURL = "https://nodejs.org/dist/index.json"

func GetNodeVersions() (*[]data.NodeVersion, error) {
	body, err := http.Get(nodeVersionURL)

	if err != nil {
		return nil, err
	}

	var nodeVersions []data.NodeVersion

	fmt.Printf("Body: %v\n", body)
	// err = body.Decode(&nodeVersions)

	return &nodeVersions, nil
}
