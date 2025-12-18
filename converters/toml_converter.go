package converters

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/pelletier/go-toml"
	"gopkg.in/yaml.v2"
)

type tomlConverter struct {
}

func (tomlConverter) ConvertToYaml(configFile string) ([]byte, error) {
	f, err := os.Open(configFile)
	if err != nil {
		return []byte{}, err
	}
	defer func() {
		if closeErr := f.Close(); closeErr != nil && err == nil {
			err = closeErr
		}
	}()
	tree, err := toml.LoadReader(f)
	if err != nil {
		return []byte{}, err
	}
	treeMap := tree.ToMap()
	return yaml.Marshal(treeMap)
}

func (tomlConverter) YamlTo(data []byte) ([]byte, error) {
	var dataConvert map[string]interface{}
	err := yaml.Unmarshal(data, &dataConvert)
	if err != nil {
		return []byte{}, err
	}
	tree, err := toml.TreeFromMap(dataConvert)
	if err != nil {
		return []byte{}, err
	}
	return tree.Marshal()
}

func (tomlConverter) Match(configFile string) bool {
	ext := strings.ToLower(filepath.Ext(configFile))
	return ext == ".toml"
}

func (tomlConverter) Type() string {
	return "toml"
}
