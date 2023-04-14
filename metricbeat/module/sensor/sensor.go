package sensor

import (
	"github.com/elastic/beats/v7/metricbeat/mb"
)

func init() {
	// Register the ModuleFactory function for this module.
	if err := mb.Registry.AddModule(ModuleName, NewModule); err != nil {
		panic(err)
	}
}

func NewModule(base mb.BaseModule) (mb.Module, error) {
	// var config Config
	// if err := base.UnpackConfig(&config); err != nil {
	// 	return nil, err
	// }
	return &base, nil
}

// ModuleName is the name of this module.
const ModuleName = "sensor"
