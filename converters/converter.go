package converters

import (
	"fmt"
	"github.com/orange-cloudfoundry/config-patcher/model"
	"strings"
)

type converter interface {
	ConvertToYaml(configFile string) ([]byte, error)
	YamlTo(data []byte) ([]byte, error)
	Match(configFile string) bool
	Type() string
}

var curConverters = []converter{
	&yamlConverter{},
	&jsonConverter{},
	&tomlConverter{},
}

func findConverter(patch model.Patch) converter {
	configType := strings.ToLower(patch.ConfigType)
	for _, c := range curConverters {
		if strings.ToLower(c.Type()) == configType || c.Match(patch.ConfigFile) {
			return c
		}
	}
	return nil
}

func ConvertConfigToYaml(patch model.Patch) ([]byte, error) {
	converter := findConverter(patch)
	if converter == nil {
		return []byte{}, fmt.Errorf("Could not find converter satisfying config file %s", patch.ConfigFile)
	}
	return converter.ConvertToYaml(patch.ConfigFile)
}

func ConfigYamlTo(patch model.Patch, data []byte) ([]byte, error) {
	converter := findConverter(patch)
	if converter == nil {
		return []byte{}, fmt.Errorf("Could not find converter satisfying config file %s", patch.ConfigFile)
	}
	return converter.YamlTo(data)
}
