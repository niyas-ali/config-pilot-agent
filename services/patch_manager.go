package services

import (
	"config-pilot-agent/model"
	"config-pilot-agent/services/json_parser"
	"log"
)

type PatchManager struct {
	patches []model.PatchConfiguration
}

func (p *PatchManager) LoadConfigurations() {
	if err := json_parser.JsonToModel("patch_configuration.json", &p.patches); err != nil {
		log.Fatalln("loading patch configuration failed")
		panic(err)
	}
}
func NewPatchManager() *PatchManager {
	return &PatchManager{}
}
