package detector

import (
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

const configFilePath string = "internal/detector/config.yml"

type BuildType struct {
	Name  string   `yaml:"name"`
	Files []string `yaml:"files"`
}

type Types struct {
	BuildTypes []BuildType `yaml:"types"`
}

func LoadConfig() (Types, error) {
	fp, _ := filepath.Abs(configFilePath)
	bytes, err := ioutil.ReadFile(fp)
	if err != nil {
		return Types{}, err
	}

	var t Types
	err = yaml.Unmarshal(bytes, &t)
	if err != nil {
		return Types{}, err
	}

	return t, nil
}
