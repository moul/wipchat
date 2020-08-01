package wipchat

import (
	"context"
	"time"

	"github.com/shurcooL/graphql"
)

func (c Client) QueryViewer(ctx context.Context, opts *QueryViewerOptions) (*User, error) {
	if !c.hasKey {
		return nil, ErrTokenRequired
	}
	if opts == nil {
		opts = &QueryViewerOptions{}
		opts.ApplyDefaults()
	}
	var query viewerQuery
	err := c.graphql.Query(ctx, &query, opts.toMap())
	if err != nil {
		return nil, err
	}

	var ret User
	err = typeToType(&query.Viewer, &ret)
	return &ret, err
}

type QueryViewerOptions struct {
	TodosCompleted bool
	TodosLimit     int
	TodosOffset    int
	TodosFilter    string
	TodosOrder     string
	AvatarSize     int
}

func (opts *QueryViewerOptions) ApplyDefaults() {
	opts.TodosCompleted = true
	if opts.TodosLimit == 0 {
		opts.TodosLimit = 20
	}
	if opts.AvatarSize == 0 {
		opts.AvatarSize = 64
	}
}

func (opts *QueryViewerOptions) toMap() map[string]interface{} {
	variables := map[string]interface{}{
		"todosCompleted": graphql.Boolean(opts.TodosCompleted),
		"todosLimit":     graphql.Int(opts.TodosLimit),
		"todosOffset":    graphql.Int(opts.TodosOffset),
		"todosFilter":    graphql.String(opts.TodosFilter),
		"todosOrder":     graphql.String(opts.TodosOrder),
		"avatarSize":     graphql.Int(opts.AvatarSize),
	}
	return variables
}

type viewerQuery struct {
	Viewer struct { // type=User
		ID                  graphql.ID
		URL                 string
		Username            string
		Firstname           string `graphql:"first_name"`
		Lastname            string `graphql:"last_name"`
		AvatarURL           string `graphql:"avatar_url(w: $avatarSize, h: $avatarSize)"`
		CompletedTodosCount int    `graphql:"completed_todos_count"`
		BestStreak          int    `graphql:"best_streak"`
		Streaking           bool
		Todos               []struct { // type=Todo
			ID          graphql.ID
			CreatedAt   *time.Time `graphql:"created_at"`
			CompletedAt *time.Time `graphql:"completed_at"`
			UpdatedAt   *time.Time `graphql:"updated_at"`
			Body        string
			Product     *struct { // type=Product
				ID      graphql.ID
				Hashtag string
				URL     string
			}
		} `graphql:"todos(limit: $todosLimit, completed: $todosCompleted, offset: $todosOffset, filter: $todosFilter, order: $todosOrder)"`
		Products []struct { // type=Product
			ID         graphql.ID
			CreatedAt  *time.Time `graphql:"created_at"`
			Hashtag    string
			Name       string
			Pitch      string
			UpdatedAt  *time.Time `graphql:"updated_at"`
			URL        string
			WebsiteURL string `graphql:"website_url"`
			Makers     []struct {
				ID graphql.ID
			}
			// Todos   []Todo
		}
	}
}
