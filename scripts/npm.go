package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

func main() {
	outdir, err := getDistFolderPath()

	if err != nil {
		fmt.Println("[ERROR] [getDistFolderPath] ", err)
		return
	}

	err = createDistFolder(outdir)

	if err != nil {
		fmt.Println("[ERROR] [createDistFolder] ", err)
		return
	}

	pkgString, err := getPackageJsonString()

	if err != nil {
		fmt.Println("[ERROR] [getPackageJsonString] ", err)
		return
	}

	err = writePackageJson(outdir, pkgString)

	if err != nil {
		fmt.Println("[ERROR] [writePackageJson] ", err)
		return
	}

	fmt.Println("package.json created")
}

type Foo struct {
	Version string
	URL     string
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

func getPackageJsonString() (string, error) {
	temp, err := template.ParseFiles("./scripts/templates/package.json")

	if err != nil {
		return "", err
	}

	var buff bytes.Buffer

	temp.Execute(&buff, Foo{Version: "2.0.0",
		URL: `https://github.com/raulfdm/node-versions-cli/releases/download/v{{version}}/myGoPackage_{{version}}_{{platform}}_{{arch}}.tar.gz`,
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
