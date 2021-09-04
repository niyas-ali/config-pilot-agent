package main

import (
	"config-pilot-job/services"
	"fmt"
)

func main() {
	repositories := services.NewRepositoryManager()
	repositories.LoadConfigurations()
	configManager := services.NewPatchManager()
	configManager.LoadConfigurations()
	processManager := services.ProcessManager{}
	for _, r := range repositories.Repositories {
		gitProcess := services.NewGitProcess(r.Name, r.Url, configManager)
		processManager.AddProcess(*gitProcess)
		fmt.Println(r)
	}
	processManager.Run()
}
