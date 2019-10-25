package converters

import (
	"encoding/json"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
	"strings"
)

type jsonConverter struct {
}

func (jsonConverter) ConvertToYaml(configFile string) ([]byte, error) {
	b, err := ioutil.ReadFile(configFile)
	if err != nil {
		return []byte{}, err
	}
	var data interface{}
	err = json.Unmarshal(b, &data)
	if err != nil {
		return []byte{}, err
	}
	return yaml.Marshal(data)
}

func (jsonConverter) YamlTo(data []byte) ([]byte, error) {
	var dataConvert interface{}
	err := yaml.Unmarshal(data, &dataConvert)
	if err != nil {
		return []byte{}, err
	}
	return json.MarshalIndent(dataConvert, "", "  ")
}

func (jsonConverter) Match(configFile string) bool {
	ext := strings.ToLower(filepath.Ext(configFile))
	return ext == ".json"
}

func (jsonConverter) Type() string {
	return "json"
}
