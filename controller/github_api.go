package controller

import (
	"bytes"
	"config-pilot-agent/model"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"golang.org/x/oauth2"
)

type GithubApi struct {
	// (required): Organization name/id
	Organization string
	// (required): Personal Access Token
	Token string
	// (required): Pull Request
	Request model.PullRequest
}

func (g GithubApi) CreatePr() string {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/pulls", g.Organization, g.Request.RepositoryName)
	head := fmt.Sprintf("%s:%s", g.Organization, g.Request.SourceBranch)
	postBody, _ := json.Marshal(map[string]string{
		"head":  head,
		"base":  g.Request.TargetBranch,
		"body":  g.Request.Description,
		"title": g.Request.Title,
	})
	responseBody := *bytes.NewBuffer(postBody)
	//sendRequest("6tm32sbwmdqfcj6523qyyihjwebkwhdkimrbtdg6xt2dagaor3pa", url, responseBody)
	return sendRequest(g.Token, url, responseBody)
}

func sendRequest(token string, url string, body bytes.Buffer) string {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{
			AccessToken: token,
		},
	)
	client := oauth2.NewClient(ctx, ts)
	req, err := http.NewRequest("POST", url, &body)
	if err != nil {
		log.Println(err.Error())
		return ""
	}
	req.Header.Set("User-Agent", "config-pilot-agent")
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err.Error())
	}
	defer resp.Body.Close()
	response, _ := ioutil.ReadAll(resp.Body)
	return string(response)
}
