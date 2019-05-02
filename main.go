package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	yaml "gopkg.in/yaml.v2"
)

type BundleFile struct {
	Description string   `yaml:"description"`
	Version     string   `yaml:"version"`
	Tenets      []string `yaml:"tenets"`
	Tags        []string `yaml:"tags"`
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal(errors.New("supply path to lingo_bundle.yaml"))
	}
	wd := os.Args[1]

	dirs, err := ioutil.ReadDir(wd)
	if err != nil {
		log.Fatal(err)
	}

	yamlData, err := ioutil.ReadFile(wd + "/lingo_bundle.yaml")
	if err != nil {
		log.Fatal(err)
	}

	var file BundleFile
	if err := yaml.Unmarshal(yamlData, &file); err != nil {
		log.Fatal(err)
	}

	var tenets []string
	for _, dir := range dirs {
		if !dir.IsDir() {
			continue
		}
		tenets = append(tenets, dir.Name())
	}

	file.Tenets = tenets

	newData, err := yaml.Marshal(file)
	if err != nil {
		log.Fatal(err)
	}
	path := filepath.Join(wd, "lingo_bundle.yaml")
	err = ioutil.WriteFile(path, newData, 0664)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Wrote lingo_bundle.yaml in", wd)
}
