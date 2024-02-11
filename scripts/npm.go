package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
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

	err = copyFile(
		"README.md",
		outDir,
		"README.md",
	)

	if err != nil {
		fmt.Println("[ERROR] [copy README.md] ", err)
		return
	}

	fmt.Println("README.md copied")

	err = copyFile(
		"./scripts/templates/install-manager.mjs",
		outDir,
		"install-manager.mjs",
	)

	if err != nil {
		fmt.Println("[ERROR] [copy install-manager.mjs] ", err)
		return
	}

	fmt.Println("install-manager.mjs copied")

	err = copyFile(
		"./scripts/templates/bin.mjs",
		outDir,
		"bin.mjs",
	)

	if err != nil {
		fmt.Println("[ERROR] [copy bin.mjs] ", err)
		return
	}

	fmt.Println("bin.mjs copied")

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
		URL:     meta.GetRemoteUrl(),
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

func (r *ReleaserMetaData) GetRemoteUrl() string {
	return fmt.Sprintf("https://github.com/raulfdm/node-versions-cli/releases/download/v{{version}}/%s_{{platform}}_{{arch}}.tar.gz", r.ProjectName)
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

func copyFile(relativeSrc string, outDir string, filename string) error {
	fullPath, err := filepath.Abs("./")

	if err != nil {
		return err
	}

	readmeSrc := filepath.Join(fullPath, relativeSrc)
	readmeDest := filepath.Join(outDir, filename)

	srcFile, err := os.Open(readmeSrc)

	if err != nil {
		return err
	}

	defer srcFile.Close()

	destFile, err := os.Create(readmeDest)

	if err != nil {
		return err
	}

	defer destFile.Close()

	_, err = io.Copy(destFile, srcFile)

	if err != nil {
		return err
	}

	return nil
}
