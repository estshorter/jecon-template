package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"os"

	"github.com/pelletier/go-toml"
)

// Config defines struct for toml file
type Config struct {
	UseBySame          int
	BySameDefinition   string
	BySameDefinitionJP string
	AuthorName         int
	SeiMeiOrder        int
}

func readTomlFile(filename string) (*Config, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	config := Config{}
	err = toml.Unmarshal(content, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

func main() {
	const tomlFilename = "configs.toml"
	const templateFilename = "jecon.bst.tmpl"
	const outputFilename = "jecon.bst"
	config, err := readTomlFile(tomlFilename)
	if err != nil {
		log.Fatal(err)
	}
	tpl := template.Must(template.New(templateFilename).Delims("[[", "]]").ParseFiles(templateFilename))
	file, err := os.Create(outputFilename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	if err := tpl.Execute(file, config); err != nil {
		log.Fatalf("failed to execute template: %v", err)
	}
}
