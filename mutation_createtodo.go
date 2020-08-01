package wipchat

import (
	"context"
	"time"

	"github.com/shurcooL/graphql"
)

func (c Client) MutateCreateTodo(ctx context.Context, body string, completedAt *time.Time, attachments []FilePayload) (*Todo, error) {
	if !c.hasKey {
		return nil, ErrTokenRequired
	}
	attachmentsInput := make([]attachmentInput, len(attachments))
	for i, attachment := range attachments {
		attachmentInput, err := c.uploadAttachment(ctx, attachment)
		if err != nil {
			return nil, err
		}
		attachmentsInput[i] = *attachmentInput
	}
	var mutation createTodoMutation
	variables := map[string]interface{}{
		"input": TodoInput{
			Body:        graphql.String(body),
			CompletedAt: completedAt,
			Attachments: attachmentsInput,
		},
	}
	err := c.graphql.Mutate(ctx, &mutation, variables)
	if err != nil {
		return nil, err
	}

	var ret Todo
	if err = typeToType(&mutation.CreateTodo, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

type createTodoMutation struct {
	CreateTodo struct {
		ID          graphql.ID
		CreatedAt   time.Time  `graphql:"created_at"`
		CompletedAt *time.Time `graphql:"completed_at" json:"CompletedAt,omitempty"`
		UpdatedAt   time.Time  `graphql:"updated_at"`
		Body        string
		Product     *struct { // type=Product
			ID      graphql.ID
			Hashtag string
			URL     string
		}
		User struct { // type=User
			ID  graphql.ID
			URL string
		}
		Attachments []Attachment
	} `graphql:"createTodo(input: $input)"`
}

type TodoInput struct {
	Body        graphql.String    `json:"body"`
	CompletedAt *time.Time        `json:"completed_at,omitempty"`
	Attachments []attachmentInput `json:"attachments"`
}
