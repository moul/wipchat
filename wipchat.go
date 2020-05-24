package wipchat

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/shurcooL/graphql"
	"moul.io/roundtripper"
)

type Client struct {
	graphql *graphql.Client
}

func New(apikey string) Client {
	transport := roundtripper.Transport{}
	if apikey != "" {
		transport.ExtraHeader = http.Header{"Authorization": []string{"Bearer " + apikey}}
	}
	httpClient := http.Client{Transport: &transport}
	gqlClient := graphql.NewClient("https://wip.chat/graphql", &httpClient)
	return Client{graphql: gqlClient}
}

func (c Client) QueryViewer(ctx context.Context) (*ViewerQuery, error) {
	var query ViewerQuery
	err := c.graphql.Query(ctx, &query, nil)
	if err != nil {
		return nil, err
	}
	return &query, nil
}

type ViewerQuery struct {
	Viewer struct { // type=User
		ID                  graphql.ID
		URL                 string
		Username            string
		Firstname           string `graphql:"first_name"`
		Lastname            string `graphql:"last_name"`
		AvatarURL           string `graphql:"avatar_url"`
		CompletedTodosCount int    `graphql:"completed_todos_count"`
		BestStreak          int    `graphql:"best_streak"`
		Streaking           bool
		Todos               []struct { // type=Todo
			ID          graphql.ID
			CreatedAt   time.Time `graphql:"created_at"`
			CompletedAt time.Time `graphql:"completed_at"`
			UpdatedAt   time.Time `graphql:"updated_at"`
			Body        string
			Product     struct { // type=Product
				ID      graphql.ID
				Hashtag string
				URL     string
			}
			User struct { // type=User
				ID  graphql.ID
				URL string
			}
		} `graphql:"todos(limit:5)"`
		Products []struct { // type=Product
			ID         graphql.ID
			CreatedAt  time.Time `graphql:"created_at"`
			Hashtag    string
			Name       string
			Pitch      string
			UpdatedAt  time.Time `graphql:"updated_at"`
			URL        string
			WebsiteURL string `graphql:"website_url"`
			// Makers  []User
			// Todos   []Todo
		}
	}
}

func (c Client) MutateCreateTodo(ctx context.Context, body string, completedAt *time.Time, attachments []io.Reader) (*CreateTodoMutation, error) {
	if len(attachments) > 0 {
		return nil, fmt.Errorf("attachments are not implemented yet")
	}
	var mutation CreateTodoMutation
	variables := map[string]interface{}{
		"input": TodoInput{
			Body:        graphql.String(body),
			CompletedAt: completedAt,
			// Attachments:
		},
	}
	err := c.graphql.Mutate(ctx, &mutation, variables)
	if err != nil {
		return nil, err
	}
	return &mutation, nil
}

type TodoInput struct {
	Body        graphql.String `json:"body"`
	CompletedAt *time.Time     `json:"completed_at"`
	// Attachments:
}

type CreateTodoMutation struct {
	CreateTodo struct {
		ID          graphql.ID
		CreatedAt   time.Time  `graphql:"created_at"`
		CompletedAt *time.Time `graphql:"completed_at" json:"CompletedAt,omitempty"`
		UpdatedAt   time.Time  `graphql:"updated_at"`
		Body        string
		Product     struct { // type=Product
			ID      graphql.ID
			Hashtag string
			URL     string
		}
		User struct { // type=User
			ID  graphql.ID
			URL string
		}
	} `graphql:"createTodo(input: $input)"`
}
