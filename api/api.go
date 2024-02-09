package api

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"node-versions-cli/data"
)

const nodeVersionURL = "https://nodejs.org/dist/index.json"

func GetNodeVersions() (*[]data.NodeVersion, error) {
	response, err := http.Get(nodeVersionURL)

	if err != nil {
		return nil, err
	}

	if response.StatusCode == http.StatusOK {
		var nodeVersions []data.NodeVersion

		bodyBi, error := io.ReadAll(response.Body)

		if error != nil {
			return nil, error
		}

		json.Unmarshal(bodyBi, &nodeVersions)

		return &nodeVersions, nil
	} else {
		return nil, errors.New("error fetching node versions")
	}

}
