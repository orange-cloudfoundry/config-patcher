package converters

import (
	"io/ioutil"
	"path/filepath"
	"strings"
)

type yamlConverter struct {
}

func (yamlConverter) ConvertToYaml(configFile string) ([]byte, error) {
	return ioutil.ReadFile(configFile)
}

func (yamlConverter) YamlTo(data []byte) ([]byte, error) {
	return data, nil
}

func (yamlConverter) Match(configFile string) bool {
	ext := strings.ToLower(filepath.Ext(configFile))
	return ext == ".yml" || ext == ".yaml"
}

func (yamlConverter) Type() string {
	return "yaml"
}
