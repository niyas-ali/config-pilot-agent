package services

import (
	"config-pilot-job/model"
	"config-pilot-job/services/json_parser"
	"log"
)

type PatchManager struct {
	patches []model.Patch
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
