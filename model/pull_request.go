package model

type PullRequest struct {
	RepositoryName string
	ProjectName    string
	SourceBranch   string
	TargetBranch   string
	Description    string
	Title          string
}
