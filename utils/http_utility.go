package utils

import (
	"bytes"
	"context"
	"io/ioutil"
	"log"
	"net/http"

	"golang.org/x/oauth2"
)

type Service struct {
	client *http.Client
}

func NewClient() *Service {
	//_token := os.Getenv("token")
	ctx := context.Background()
	token := oauth2.StaticTokenSource(
		&oauth2.Token{
			AccessToken: "",
		},
	)
	client := oauth2.NewClient(ctx, token)
	service := new(Service)
	service.client = client
	return service

}

func (c *Service) SendRequest(url string, body bytes.Buffer) string {
	req, err := http.NewRequest("POST", url, &body)
	req.Header.Set("User-Agent", "config-pilot-job")
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.client.Do(req)
	if err != nil {
		log.Println(err.Error())
	}
	defer resp.Body.Close()
	response, _ := ioutil.ReadAll(resp.Body)
	return string(response)
}
