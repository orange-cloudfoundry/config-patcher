package main

import (
	"github.com/hashicorp/go-multierror"
	"github.com/orange-cloudfoundry/config-patcher/model"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"path/filepath"
	"github.com/alecthomas/kingpin/v2"
	"github.com/prometheus/common/version"
)

var (
	patchFlag = kingpin.
		Flag("patch", "Set in glob format where to find rules for patching config files").
		Default("/var/vcap/jobs/*/config-patcher/*.yml").
		String()
)

func main() {
	kingpin.Version(version.Print("config-patcher"))
	kingpin.HelpFlag.Short('h')
	kingpin.Parse()

	filesMatch, err := filepath.Glob(*patchFlag)
	if err != nil {
		log.Fatal(err)
	}
	patches := make([]model.Patch, 0)
	for _, file := range filesMatch {
		b, err := ioutil.ReadFile(file)
		if err != nil {
			log.Fatal(err)
		}
		patchTmp := make([]model.Patch, 0)
		err = yaml.Unmarshal(b, &patchTmp)
		if err != nil {
			log.Fatal(err)
		}
		patches = append(patches, patchTmp...)
	}
	var result error
	for _, patch := range patches {
		patcher := NewPatcher(patch)
		err = patcher.Patch()
		if err != nil {
			result = multierror.Append(result, err)
		}
	}
	if result != nil {
		log.Fatal(result)
	}
}
