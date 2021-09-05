package main

import (
	"config-pilot-job/services"
)

func main() {
	repositories := services.NewRepositoryManager()
	repositories.LoadConfigurations()
	configManager := services.NewPatchManager()
	configManager.LoadConfigurations()
	processManager := services.ProcessManager{}
	for _, r := range repositories.Config.Repository {
		gitProcess := services.NewGitProcess(repositories.Config, r, configManager)
		processManager.AddProcess(*gitProcess)
	}
	processManager.Run()
}
