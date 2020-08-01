package wipchat

import (
	"net/http"

	"github.com/shurcooL/graphql"
	"moul.io/roundtripper"
)

type Client struct {
	graphql *graphql.Client
	hasKey  bool
}

func New(apikey string) Client {
	transport := roundtripper.Transport{}
	client := Client{}
	if apikey != "" {
		transport.ExtraHeader = http.Header{"Authorization": []string{"Bearer " + apikey}}
		client.hasKey = true
	}
	httpClient := http.Client{Transport: &transport}
	client.graphql = graphql.NewClient("https://wip.chat/graphql", &httpClient)
	return client
}
