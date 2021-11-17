package services

import (
	"config-pilot-agent/model"
	"config-pilot-agent/services/json_parser"
	"config-pilot-agent/utils/logger"
)

type RepositoryManager struct {
	Config model.Configuration
}

func (r *RepositoryManager) LoadConfigurations() {
	if err := json_parser.JsonToModel("repository.json", &r.Config); err != nil {
		logger.Fatalln("loading repository configuration failed")
		panic(err)
	}
}

func NewRepositoryManager() *RepositoryManager {
	return &RepositoryManager{}
}
