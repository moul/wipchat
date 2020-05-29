package wipchat

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
	"time"

	"github.com/shurcooL/graphql"
	"moul.io/roundtripper"
	"moul.io/wipchat/wiptypes"
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

func (c Client) QueryViewer(ctx context.Context) (*wiptypes.ViewerQuery, error) {
	if !c.hasKey {
		return nil, ErrTokenRequired
	}
	var query wiptypes.ViewerQuery
	err := c.graphql.Query(ctx, &query, nil)
	if err != nil {
		return nil, err
	}
	return &query, nil
}

func (c Client) MutateCreateTodo(ctx context.Context, body string, completedAt *time.Time, attachments []Attachment) (*wiptypes.CreateTodoMutation, error) {
	if !c.hasKey {
		return nil, ErrTokenRequired
	}
	attachmentsInput := make([]wiptypes.AttachmentInput, len(attachments))
	for i, attachment := range attachments {
		var mutation wiptypes.CreatePresignedURLMutation
		variables := map[string]interface{}{
			"filename": graphql.String(attachment.Filename),
		}
		err := c.graphql.Mutate(ctx, &mutation, variables)
		if err != nil {
			return nil, err
		}
		if mutation.CreatePresignedURL.Headers != "{}" ||
			mutation.CreatePresignedURL.Method != "post" {
			return nil, fmt.Errorf("unsupported attachment API")
		}
		var fields map[string]string
		err = json.Unmarshal([]byte(mutation.CreatePresignedURL.Fields), &fields)
		if err != nil {
			return nil, err
		}
		var b bytes.Buffer
		w := multipart.NewWriter(&b)
		for key, value := range fields {
			fw, err := w.CreateFormField(key)
			if err != nil {
				return nil, err
			}
			_, err = fw.Write([]byte(value))
			if err != nil {
				return nil, err
			}
		}
		fw, err := w.CreateFormFile("file", attachment.Filename)
		if err != nil {
			return nil, err
		}
		_, err = fw.Write(attachment.Bytes)
		if err != nil {
			return nil, err
		}
		w.Close()
		req, err := http.NewRequest("POST", mutation.CreatePresignedURL.URL, &b)
		if err != nil {
			return nil, err
		}
		req.Header.Set("Content-Type", w.FormDataContentType())
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			return nil, err
		}
		defer res.Body.Close()
		if res.StatusCode != 204 {
			return nil, fmt.Errorf("invalid status code: %d", res.StatusCode)
		}
		attachmentsInput[i] = wiptypes.AttachmentInput{
			Key:      fields["key"],
			Filename: attachment.Filename,
			Size:     len(attachment.Bytes),
		}
	}
	var mutation wiptypes.CreateTodoMutation
	variables := map[string]interface{}{
		"input": wiptypes.TodoInput{
			Body:        graphql.String(body),
			CompletedAt: completedAt,
			Attachments: attachmentsInput,
		},
	}
	err := c.graphql.Mutate(ctx, &mutation, variables)
	if err != nil {
		return nil, err
	}
	return &mutation, nil
}

type Attachment struct {
	Filename string
	Bytes    []byte
}
