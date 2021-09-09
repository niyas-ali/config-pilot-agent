package main

import (
	"config-pilot-job/services"
)

func main() {
	processManager := services.ProcessManager{}
	processManager.InitializeProcess()
	processManager.Run()
}
