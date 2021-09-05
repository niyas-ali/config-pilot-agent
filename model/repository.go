package model

type Repository struct {
	Name   string `json:"name"`
	URL    string `json:"url"`
	Branch string `json:"branch"`
}
type Configuration struct {
	CheckoutBranch   string       `json:"checkoutBranch"`
	Owner            string       `json:"owner"`
	PrRequestTitle   string       `json:"prRequestTitle"`
	PrRequestMessage string       `json:"prRequestMessage"`
	Repository       []Repository `json:"repository"`
}
