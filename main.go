package main

import (
	"fmt"
	"node-versions-cli/api"
)

func main() {
	nodeVersions, _ := api.GetNodeVersions()

	fmt.Println("Hello, World!", nodeVersions)
}
