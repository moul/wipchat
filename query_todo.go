package wipchat

import (
	"context"
	"time"

	"github.com/shurcooL/graphql"
)

func (c Client) QueryTodos(ctx context.Context, opts *QueryTodosOptions) ([]Todo, error) {
	if !c.hasKey {
		return nil, ErrTokenRequired
	}
	if opts == nil {
		opts = &QueryTodosOptions{}
		opts.ApplyDefaults()
	}

	var query todosQuery
	err := c.graphql.Query(ctx, &query, opts.toMap())
	if err != nil {
		return nil, err
	}

	var ret []Todo
	err = typeToType(&query.Todos, &ret)
	return ret, err
}

type QueryTodosOptions struct {
	TodosCompleted bool
	TodosLimit     int
	TodosFilter    string
}

func (opts *QueryTodosOptions) ApplyDefaults() {
	opts.TodosCompleted = true
	if opts.TodosLimit == 0 {
		opts.TodosLimit = 20
	}
}

func (opts *QueryTodosOptions) toMap() map[string]interface{} {
	variables := map[string]interface{}{
		"todosCompleted": graphql.Boolean(opts.TodosCompleted),
		"todosLimit":     graphql.Int(opts.TodosLimit),
		"todosFilter":    graphql.String(opts.TodosFilter),
	}
	return variables
}

type todosQuery struct {
	Todos []struct {
		ID      graphql.ID
		Product struct {
			ID         graphql.ID
			CreatedAt  *time.Time `graphql:"created_at" json:"created_at,omitempty"`
			Hashtag    string
			Name       string
			Pitch      string
			UpdatedAt  *time.Time `graphql:"updated_at" json:"updated_at,omitempty"`
			URL        string
			WebsiteURL string `graphql:"website_url"`
			Makers     []struct {
				ID graphql.ID
			}
			// Todos   []Todo
		} `json:"product,omitempty"`
		UpdatedAt time.Time `graphql:"updated_at" json:"updated_at,omitempty"`
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
			// Products []Product
			// Todos []Todo
		}
		CreatedAt   *time.Time `graphql:"created_at" json:"created_at,omitempty"`
		CompletedAt *time.Time `graphql:"completed_at" json:"completed_at,omitempty"`
		Body        string
		Attachments []struct {
			ID        graphql.ID
			CreatedAt *time.Time `graphql:"createdAt" json:"created_at,omitempty"`
			Filename  string
			MimeType  string `graphql:"mimeType"`
			Size      int
			UpdatedAt *time.Time `graphql:"updatedAt" json:"updated_at,omitempty"`
			URL       string
			// AspectRatio float64 `graphql:"aspectRatio"` // buggy
		}
	} `graphql:"todos(limit: $todosLimit, completed: $todosCompleted, filter: $todosFilter)"`
}
