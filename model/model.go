package model

import (
	"github.com/krishicks/yaml-patch"
)

type Patch struct {
	ConfigFile string       `yaml:"config_file"`
	ConfigType string       `yaml:"config_type"`
	Patches    []*Operation `yaml:"patches"`
}

type Operation struct {
	Op    yamlpatch.Op     `yaml:"op,omitempty"`
	Path  yamlpatch.OpPath `yaml:"path,omitempty"`
	From  yamlpatch.OpPath `yaml:"from,omitempty"`
	Value *yamlpatch.Node  `yaml:"value,omitempty"`
	Type  yamlpatch.Op     `yaml:"type,omitempty"`
}

func (c Operation) ToYamlPatch() yamlpatch.Operation {
	return yamlpatch.Operation{
		Op:    c.Op,
		Value: c.Value,
		Path:  c.Path,
		From:  c.From,
	}
}

func (c *Operation) UnmarshalYAML(unmarshal func(interface{}) error) error {
	type plain Operation
	if err := unmarshal((*plain)(c)); err != nil {
		return err
	}
	if c.Op == "" {
		c.Op = c.Type
	}
	return nil
}
