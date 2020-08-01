package wipchat

import (
	"context"
	"time"

	"github.com/shurcooL/graphql"
)

func (c Client) QueryProducts(ctx context.Context) ([]Product, error) {
	if !c.hasKey {
		return nil, ErrTokenRequired
	}
	var query productsQuery
	err := c.graphql.Query(ctx, &query, nil)
	if err != nil {
		return nil, err
	}

	var ret []Product
	err = typeToType(&query.Products, &ret)
	return ret, err
}

type productsQuery struct {
	Products []struct {
		ID         graphql.ID `graphql:"id"`
		CreatedAt  time.Time  `graphql:"created_at"`
		Hashtag    string     `graphql:"hashtag"`
		Name       string     `graphql:"name"`
		Pitch      string     `graphql:"pitch"`
		UpdatedAt  time.Time  `graphql:"updated_at"`
		URL        string     `graphql:"url"`
		WebsiteURL string     `graphql:"website_url"`
		Makers     []struct {
			ID                  graphql.ID `graphql:"id"`
			URL                 string     `graphql:"url"`
			Username            string     `graphql:"username"`
			Firstname           string     `graphql:"first_name"`
			Lastname            string     `graphql:"last_name"`
			AvatarURL           string     `graphql:"avatar_url"`
			CompletedTodosCount int        `graphql:"completed_todos_count"`
			BestStreak          int        `graphql:"best_streak"`
			Streaking           bool       `graphql:"streaking"`
			//Todos               []Todo     `graphql:"todos" json:"todos,omitempty"`
			//Products            []Product  `graphql:"products" json:"products,omitempty"`
		} `graphql:"makers" json:"makers,omitempty"`
	}
}
