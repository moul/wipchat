package wiptypes

import (
	"time"

	"github.com/shurcooL/graphql"
)

type Attachment struct {
	ID          graphql.ID `graphql:"id"`
	AspectRatio float64    `graphql:"aspect_ratio"`
	CreatedAt   time.Time  `graphql:"created_at"`
	Filename    string     `graphql:"filename"`
	MimeType    string     `graphql:"mime_type"`
	Size        int        `graphql:"size"`
	UpdatedAt   time.Time  `graphql:"updated_at"`
	URL         string     `graphql:"url"`
}

type Todo struct {
	ID          graphql.ID   `graphql:"id"`
	Product     *Product     `graphql:"product" json:"product,omitempty"`
	UpdatedAt   time.Time    `graphql:"updated_at"`
	User        *User        `graphql:"user" json:"user,omitempty"`
	CreatedAt   time.Time    `graphql:"created_at"`
	CompletedAt time.Time    `graphql:"completed_at"`
	Body        string       `graphql:"body"`
	Attachments []Attachment `graphql:"attachments" json:"attachments,omitempty"`
}

type Product struct {
	ID         graphql.ID `graphql:"id"`
	CreatedAt  time.Time  `graphql:"created_at"`
	Hashtag    string     `graphql:"hashtag"`
	Name       string     `graphql:"name"`
	Pitch      string     `graphql:"pitch"`
	UpdatedAt  time.Time  `graphql:"updated_at"`
	URL        string     `graphql:"url"`
	WebsiteURL string     `graphql:"website_url"`
	Makers     []User     `graphql:"makers" json:"makers,omitempty"`
	Todos      []Todo     `graphql:"todos" json:"todos,omitempty"`
}

type User struct {
	ID                  graphql.ID `graphql:"id"`
	URL                 string     `graphql:"url"`
	Username            string     `graphql:"username"`
	Firstname           string     `graphql:"first_name"`
	Lastname            string     `graphql:"last_name"`
	AvatarURL           string     `graphql:"avatar_url"`
	CompletedTodosCount int        `graphql:"completed_todos_count"`
	BestStreak          int        `graphql:"best_streak"`
	Streaking           bool       `graphql:"streaking"`
	Todos               []Todo     `graphql:"todos" json:"todos,omitempty"`
	Products            []Product  `graphql:"products" json:"products,omitempty"`
}
