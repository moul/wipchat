package wiptypes

import (
	"time"

	"github.com/shurcooL/graphql"
)

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
			Product     *struct { // type=Product
				ID      graphql.ID
				Hashtag string
				URL     string
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

type TodosQuery struct {
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

type ProductsQuery struct {
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

type CreatePresignedURLMutation struct {
	CreatePresignedURL struct {
		Fields  graphql.String
		Headers string
		Method  string
		URL     string
	} `graphql:"createPresignedUrl(input:{filename: $filename})"`
}

type TodoInput struct {
	Body        graphql.String    `json:"body"`
	CompletedAt *time.Time        `json:"completed_at,omitempty"`
	Attachments []AttachmentInput `json:"attachments"`
}

type AttachmentInput struct {
	Key      string `json:"key"`
	Filename string `json:"filename"`
	Size     int    `json:"size"`
}

type CreateTodoMutation struct {
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
