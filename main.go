package main

import (
	"config-pilot-agent/services"
)

func main() {
	processManager := services.ProcessManager{}
	processManager.InitializeProcess()
	processManager.Run()
}
