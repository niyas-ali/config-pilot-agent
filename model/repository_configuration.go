package model

type Configuration struct {
	CheckoutBranch string      `json:"checkout_branch"`
	Organization   string      `json:"organization"`
	PrTitle        string      `json:"pr_title"`
	PrMessage      string      `json:"pr_message"`
	AzureDevops    AzureDevops `json:"azure_devops"`
	Github         Github      `json:"github"`
}
type Repository struct {
	Name        string `json:"name"`
	URL         string `json:"url"`
	MergeBranch string `json:"merge_branch"`
}
type AzureDevops struct {
	ProjectName string       `json:"project_name"`
	Repository  []Repository `json:"repository"`
}
type Github struct {
	Repository []Repository `json:"repository"`
}

// type Configuration struct {
// 	CheckoutBranch   string       `json:"checkoutBranch"`
// 	Owner            string       `json:"owner"`
// 	PrRequestTitle   string       `json:"prRequestTitle"`
// 	PrRequestMessage string       `json:"prRequestMessage"`
// 	Repository       []Repository `json:"repository"`
// }
