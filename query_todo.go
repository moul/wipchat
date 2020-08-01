package wipchat

import (
	"context"
	"time"

	"github.com/shurcooL/graphql"
)

func (c Client) QueryTodos(ctx context.Context) ([]Todo, error) {
	if !c.hasKey {
		return nil, ErrTokenRequired
	}
	var query todosQuery
	err := c.graphql.Query(ctx, &query, nil)
	if err != nil {
		return nil, err
	}

	var ret []Todo
	err = typeToType(&query.Todos, &ret)
	return ret, err
}

type todosQuery struct {
	Todos []struct {
		ID      graphql.ID
		Product struct {
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
		UpdatedAt time.Time `graphql:"updated_at"`
		User      struct {
			ID                  graphql.ID
			URL                 string
			Username            string
			Firstname           string `graphql:"first_name"`
			Lastname            string `graphql:"last_name"`
			AvatarURL           string `graphql:"avatar_url"`
			CompletedTodosCount int    `graphql:"completed_todos_count"`
			BestStreak          int    `graphql:"best_streak"`
			Streaking           bool
			//Products []Product
			//Todos []Todo
		}
		CreatedAt   time.Time `graphql:"created_at"`
		CompletedAt time.Time `graphql:"completed_at"`
		Body        string
		Attachments []struct {
			ID        graphql.ID
			CreatedAt time.Time `graphql:"created_at"`
			Filename  string
			MimeType  string `graphql:"mime_type"`
			Size      int
			UpdatedAt time.Time `graphql:"updated_at"`
			URL       string
			//AspectRatio float64   `graphql:"aspect_ratio"` // buggy
		}
	}
}
