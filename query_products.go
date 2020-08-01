package wipchat

import (
	"context"
	"time"

	"github.com/shurcooL/graphql"
)

func (c Client) QueryProducts(ctx context.Context, limit int) ([]Product, error) {
	if !c.hasKey {
		return nil, ErrTokenRequired
	}
	if limit == 0 {
		limit = 20
	}
	var query productsQuery
	err := c.graphql.Query(ctx, &query, map[string]interface{}{
		"limit": graphql.Int(limit),
	})
	if err != nil {
		return nil, err
	}

	var ret []Product
	err = typeToType(&query.Products, &ret)
	return ret, err
}

type productsQuery struct {
	Products []struct {
		ID         graphql.ID `graphql:"id" json:"id,omitempty"`
		CreatedAt  time.Time  `graphql:"created_at" json:"created_at,omitempty"`
		Hashtag    string     `graphql:"hashtag" json:"hashtag,omitempty"`
		Name       string     `graphql:"name" json:"name,omitempty"`
		Pitch      string     `graphql:"pitch" json:"pitch,omitempty"`
		UpdatedAt  time.Time  `graphql:"updated_at" json:"updated_at,omitempty"`
		URL        string     `graphql:"url" json:"url,omitempty"`
		WebsiteURL string     `graphql:"website_url" json:"website_url,omitempty"`
		Makers     []struct {
			ID                  graphql.ID `graphql:"id" json:"id,omitempty"`
			URL                 string     `graphql:"url" json:"url,omitempty"`
			Username            string     `graphql:"username" json:"username,omitempty"`
			Firstname           string     `graphql:"first_name" json:"first_name,omitempty"`
			Lastname            string     `graphql:"last_name" json:"last_name,omitempty"`
			AvatarURL           string     `graphql:"avatar_url" json:"avatar_url,omitempty"`
			CompletedTodosCount int        `graphql:"completed_todos_count" json:"completed_todos_count,omitempty"`
			BestStreak          int        `graphql:"best_streak" json:"best_streak,omitempty"`
			Streaking           bool       `graphql:"streaking" json:"streaking,omitempty"`
			//Todos               []Todo     `graphql:"todos" json:"todos,omitempty"`
			//Products            []Product  `graphql:"products" json:"products,omitempty"`
		} `graphql:"makers" json:"makers,omitempty"`
	} `graphql:"products(limit: $limit)"`
}
