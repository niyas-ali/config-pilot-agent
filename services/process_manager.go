package services

import (
	"config-pilot-agent/controller"
	"config-pilot-agent/model"
	"log"
	"os"
)

type ProcessManager struct {
	jobs []GitProcess
}

func (p *ProcessManager) InitializeProcess() {
	token := os.Getenv("TOKEN")
	if token == "" {
		log.Fatal("missing token, add 'TOKEN' in environment varialbe")
	}
	repositoriesConfig := NewRepositoryManager()
	repositoriesConfig.LoadConfigurations()
	patchManagerConfig := NewPatchManager()
	patchManagerConfig.LoadConfigurations()

	for _, repo := range repositoriesConfig.Config.AzureDevops.Repository {
		azController := controller.AzureDevopsApi{
			Organization: repositoriesConfig.Config.Organization,
			Token:        token,
			Request: model.PullRequest{
				RepositoryName: repo.Name,
				ProjectName:    repositoriesConfig.Config.AzureDevops.ProjectName,
				SourceBranch:   repositoriesConfig.Config.CheckoutBranch,
				TargetBranch:   repo.MergeBranch,
				Description:    repositoriesConfig.Config.PrMessage,
				Title:          repositoriesConfig.Config.PrTitle,
			},
		}
		gitProcess := NewGitProcess(patchManagerConfig, repo, azController)
		p.AddProcess(*gitProcess)
	}
	for _, repo := range repositoriesConfig.Config.Github.Repository {
		githubController := controller.GithubApi{
			Organization: repositoriesConfig.Config.Organization,
			Token:        token,
			Request: model.PullRequest{
				RepositoryName: repo.Name,
				SourceBranch:   repositoriesConfig.Config.CheckoutBranch,
				TargetBranch:   repo.MergeBranch,
				Description:    repositoriesConfig.Config.PrMessage,
				Title:          repositoriesConfig.Config.PrTitle,
			},
		}
		gitProcess := NewGitProcess(patchManagerConfig, repo, githubController)
		p.AddProcess(*gitProcess)
	}
}

func (p *ProcessManager) AddProcess(process GitProcess) {
	p.jobs = append(p.jobs, process)
}
func (p *ProcessManager) Run() {
	for _, job := range p.jobs {
		job.Run()
	}
}
