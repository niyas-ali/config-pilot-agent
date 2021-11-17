package controller

import (
	"config-pilot-agent/model"
	"context"
	"fmt"

	"github.com/microsoft/azure-devops-go-api/azuredevops"
	azuregit "github.com/microsoft/azure-devops-go-api/azuredevops/git"
)

type AzureDevopsApi struct {
	// (required): Organization name/id
	Organization string
	// (required): Personal Access Token
	Token string
	// (required): Pull Request
	Request model.PullRequest
}

func (az AzureDevopsApi) CreatePr() (string, error) {
	connection := azuredevops.NewPatConnection(fmt.Sprintf("https://dev.azure.com/%s", az.Organization), az.Token)
	ctx := context.Background()
	client, err := azuregit.NewClient(ctx, connection)
	if err != nil {
		return "", err
	}
	pr := azuregit.CreatePullRequestArgs{}
	repoId := az.Request.RepositoryName
	sourceBranch := fmt.Sprintf("refs/heads/%s", az.Request.SourceBranch)
	targetBranch := fmt.Sprintf("refs/heads/%s", az.Request.TargetBranch)
	pr.Project = &az.Request.ProjectName
	pr.RepositoryId = &repoId
	pr.GitPullRequestToCreate = &azuregit.GitPullRequest{
		TargetRefName: &targetBranch,
		SourceRefName: &sourceBranch,
		Description:   &az.Request.Description,
		Title:         &az.Request.Title,
	}
	result, err := client.CreatePullRequest(ctx, pr)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("pull-request id:%d", *result.PullRequestId), nil
}
