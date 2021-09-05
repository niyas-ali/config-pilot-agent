package services

import (
	"config-pilot-job/model"
	"config-pilot-job/services/json_parser"
	"log"
)

type RepositoryManager struct {
	Config model.Configuration
}

func (r *RepositoryManager) LoadConfigurations() {
	if err := json_parser.JsonToModel("repository.json", &r.Config); err != nil {
		log.Fatalln("loading repository configuration failed")
		panic(err)
	}
}

func NewRepositoryManager() *RepositoryManager {
	return &RepositoryManager{}
}
