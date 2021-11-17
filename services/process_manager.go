package services

import (
	"config-pilot-agent/controller"
	"config-pilot-agent/model"
	"config-pilot-agent/utils/logger"
	"os"
)

type ProcessManager struct {
	jobs []GitProcess
}

func (p *ProcessManager) InitializeProcess() {

	repositoriesConfig := NewRepositoryManager()
	repositoriesConfig.LoadConfigurations()
	patchManagerConfig := NewPatchManager()
	patchManagerConfig.LoadConfigurations()
	if err := repositoriesConfig.Config.Validate(); err != nil {
		logger.Fatalf("validation -> %s", err.Error())
	}
	if err := patchManagerConfig.Validate(); err != nil {
		logger.Fatalf("validation -> %s", err.Error())
	}
	for _, repo := range repositoriesConfig.Config.AzureDevops.Repository {
		token := os.Getenv("AZ_TOKEN")
		if token == "" {
			logger.Fatalln("missing token, add 'AZ_TOKEN' in environment varialbe")
		}
		azController := controller.AzureDevopsApi{
			Organization: repositoriesConfig.Config.AzureDevops.Organization,
			Token:        token,
			Request: model.PullRequest{
				RepositoryName: repo.Name,
				ProjectName:    repositoriesConfig.Config.AzureDevops.ProjectName,
				SourceBranch:   repositoriesConfig.Config.CheckoutBranch,
				TargetBranch:   repo.MergeBranch,
				Description:    repositoriesConfig.Config.PrMessage,
				Title:          repositoriesConfig.Config.PrTitle,
				Reviewer:       repo.Reviewer,
			},
		}
		gitProcess := NewGitProcess(patchManagerConfig, repo, azController)
		p.AddProcess(*gitProcess)
	}
	for _, repo := range repositoriesConfig.Config.Github.Repository {
		token := os.Getenv("GITHUB_TOKEN")
		if token == "" {
			logger.Fatalln("missing token, add 'GITHUB_TOKEN' in environment varialbe")
		}
		githubController := controller.GithubApi{
			Organization: repositoriesConfig.Config.Github.Organization,
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
