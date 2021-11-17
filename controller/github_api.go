package controller

import (
	"bytes"
	"config-pilot-agent/model"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
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

func (g GithubApi) CreatePr() (string, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/pulls", g.Organization, g.Request.RepositoryName)
	head := fmt.Sprintf("%s:%s", g.Organization, g.Request.SourceBranch)
	postBody, _ := json.Marshal(map[string]string{
		"head":  head,
		"base":  g.Request.TargetBranch,
		"body":  g.Request.Description,
		"title": g.Request.Title,
	})
	responseBody := *bytes.NewBuffer(postBody)
	return sendRequest(g.Token, url, responseBody)
}

func sendRequest(token string, url string, body bytes.Buffer) (string, error) {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{
			AccessToken: token,
		},
	)
	client := oauth2.NewClient(ctx, ts)
	req, err := http.NewRequest("POST", url, &body)
	if err != nil {
		return "", err
	}
	req.Header.Set("User-Agent", "config-pilot-agent")
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	response, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return "", nil
	}
	return "", errors.New(string(response))
}
