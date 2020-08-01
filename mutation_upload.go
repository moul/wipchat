package wipchat

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"

	"github.com/shurcooL/graphql"
)

func (c Client) uploadAttachment(ctx context.Context, attachment FilePayload) (*attachmentInput, error) {
	if !c.hasKey {
		return nil, ErrTokenRequired
	}
	var mutation createPresignedURLMutation
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
	ret := attachmentInput{
		Key:      fields["key"],
		Filename: attachment.Filename,
		Size:     len(attachment.Bytes),
	}
	return &ret, nil
}

type createPresignedURLMutation struct {
	CreatePresignedURL struct {
		Fields  graphql.String
		Headers string
		Method  string
		URL     string
	} `graphql:"createPresignedUrl(input:{filename: $filename})"`
}

type FilePayload struct {
	Filename string
	Bytes    []byte
}

type attachmentInput struct {
	Key      string `json:"key"`
	Filename string `json:"filename"`
	Size     int    `json:"size"`
}
