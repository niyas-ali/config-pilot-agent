package services

import (
	"config-pilot-agent/model"
	"config-pilot-agent/services/json_parser"
	"config-pilot-agent/utils/logger"
)

type PatchManager struct {
	Patches []model.PatchConfiguration
}

func (p *PatchManager) Validate() error {
	for _, item := range p.Patches {
		if err := item.Validate(); err != nil {
			return err
		}
	}
	return nil
}
func (p *PatchManager) LoadConfigurations() {
	if err := json_parser.JsonToModel("patch_configuration.json", &p.Patches); err != nil {
		logger.Fatalln("loading patch configuration failed")
		panic(err)
	}
}
func NewPatchManager() *PatchManager {
	return &PatchManager{}
}
