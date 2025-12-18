package main

import (
	"fmt"
	"os"

	yamlpatch "github.com/krishicks/yaml-patch"
	"github.com/orange-cloudfoundry/config-patcher/converters"
	"github.com/orange-cloudfoundry/config-patcher/model"
)

type Patcher struct {
	patch model.Patch
}

func NewPatcher(patch model.Patch) *Patcher {
	return &Patcher{
		patch: patch,
	}
}

func (c Patcher) Patch() error {
	configFilePerm, err := c.configFilePerm()
	if err != nil {
		return err
	}
	config, err := converters.ConvertConfigToYaml(c.patch)
	if err != nil {
		return err
	}

	opPatch := make(yamlpatch.Patch, len(c.patch.Patches))
	for i, p := range c.patch.Patches {
		opPatch[i] = p.ToYamlPatch()
	}
	bs, err := opPatch.Apply(config)
	if err != nil {
		return fmt.Errorf("applying patch failed: %s", err)
	}
	dataReconvert, err := converters.ConfigYamlTo(c.patch, bs)
	if err != nil {
		return fmt.Errorf("reconvert to config format failed: %s", err)
	}
	return os.WriteFile(c.patch.ConfigFile, dataReconvert, configFilePerm)
}

func (c Patcher) configFilePerm() (os.FileMode, error) {
	f, err := os.Open(c.patch.ConfigFile)
	if err != nil {
		return 0, err
	}
	defer func() {
		if closeErr := f.Close(); closeErr != nil && err == nil {
			err = closeErr
		}
	}()
	stat, err := f.Stat()
	if err != nil {
		return 0, err
	}
	return stat.Mode(), nil
}
