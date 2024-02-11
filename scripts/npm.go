package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

func main() {
	outDir, err := getDistFolderPath()

	if err != nil {
		fmt.Println("[ERROR] [getDistFolderPath] ", err)
		return
	}

	meta, err := getReleaserMetaData()

	if err != nil {
		fmt.Println("[ERROR] [getReleaserMetaData] ", err)
		return
	}

	err = createDistFolder(outDir)

	if err != nil {
		fmt.Println("[ERROR] [createDistFolder] ", err)
		return
	}

	pkgString, err := getPackageJsonString(*meta)

	if err != nil {
		fmt.Println("[ERROR] [getPackageJsonString] ", err)
		return
	}

	err = writePackageJson(outDir, pkgString)

	if err != nil {
		fmt.Println("[ERROR] [writePackageJson] ", err)
		return
	}

	fmt.Println("package.json created")
}

func getDistFolderPath() (string, error) {
	fullPath, err := filepath.Abs("./")

	if err != nil {
		return "", err
	}

	distPath := filepath.Join(fullPath, "dist-npm")

	return distPath, nil
}

func createDistFolder(distPath string) error {
	err := os.Mkdir(distPath, 0755)

	if err != nil {
		if !strings.Contains(err.Error(), "file exists") {
			return err
		}
	}

	return nil
}

type PkgJsonTemplate struct {
	Version string
	URL     string
}

func getPackageJsonString(meta ReleaserMetaData) (string, error) {
	temp, err := template.ParseFiles("./scripts/templates/package.json")

	if err != nil {
		return "", err
	}

	var buff bytes.Buffer

	temp.Execute(&buff, PkgJsonTemplate{
		Version: meta.GetVersion(),
		URL:     fmt.Sprintf("https://github.com/raulfdm/node-versions-cli/releases/download/v{{version}}/%s_{{version}}_{{platform}}_{{arch}}.tar.gz", meta.ProjectName),
	})

	result := buff.String()

	return result, nil
}

func writePackageJson(distPath string, pkgString string) error {
	file, err := os.Create(filepath.Join(distPath, "package.json"))

	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write([]byte(pkgString))

	if err != nil {
		return err
	}

	return nil
}

type ReleaserMetaData struct {
	Tag         string `json:"tag"`
	ProjectName string `json:"project_name"`
}

func (r *ReleaserMetaData) GetVersion() string {
	// remove v prefix
	return r.Tag[1:]
}

func getReleaserMetaData() (*ReleaserMetaData, error) {
	fullPath, err := filepath.Abs("./")

	if err != nil {
		return nil, err
	}

	releaserMetaPath := filepath.Join(fullPath, "dist/metadata.json")

	file, err := os.Open(releaserMetaPath)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	var meta ReleaserMetaData

	err = json.NewDecoder(file).Decode(&meta)

	if err != nil {
		return nil, err
	}

	return &meta, nil
}
