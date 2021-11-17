package model

type Configuration struct {
	CheckoutBranch string      `json:"checkout_branch"`
	PrTitle        string      `json:"pr_title"`
	PrMessage      string      `json:"pr_message"`
	AzureDevops    AzureDevops `json:"azure_devops"`
	Github         Github      `json:"github"`
}
type Repository struct {
	Name        string `json:"name"`
	URL         string `json:"url"`
	MergeBranch string `json:"merge_branch"`
	Reviewer    string `json:"reviewer"`
}
type AzureDevops struct {
	Organization string       `json:"organization"`
	ProjectName  string       `json:"project_name"`
	Repository   []Repository `json:"repository"`
}
type Github struct {
	Organization string       `json:"organization"`
	Repository   []Repository `json:"repository"`
}
