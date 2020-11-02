package wipchat

import (
	"fmt"
	"time"

	"github.com/shurcooL/graphql"
)

type Attachment struct {
	ID          graphql.ID `graphql:"id" json:"id,omitempty"`
	AspectRatio float64    `graphql:"aspect_ratio" json:"aspect_ratio,omitempty"`
	CreatedAt   *time.Time `graphql:"created_at" json:"created_at,omitempty"`
	Filename    string     `graphql:"filename" json:"filename,omitempty"`
	MimeType    string     `graphql:"mime_type" json:"mime_type,omitempty"`
	Size        int        `graphql:"size" json:"size,omitempty"`
	UpdatedAt   *time.Time `graphql:"updated_at" json:"updated_at,omitempty"`
	URL         string     `graphql:"url" json:"url,omitempty"`
}

type Todo struct {
	ID          graphql.ID   `graphql:"id" json:"id,omitempty"`
	Product     *Product     `graphql:"product" json:"product,omitempty"`
	UpdatedAt   *time.Time   `graphql:"updated_at" json:"updated_at,omitempty"`
	User        *User        `graphql:"user" json:"user,omitempty"`
	CreatedAt   *time.Time   `graphql:"created_at" json:"created_at,omitempty"`
	CompletedAt *time.Time   `graphql:"completed_at" json:"completed_at,omitempty"`
	Body        string       `graphql:"body" json:"body,omitempty"`
	Attachments []Attachment `graphql:"attachments" json:"attachments,omitempty"`
}

func (t *Todo) CanonicalURL() string {
	if t.User != nil {
		return fmt.Sprintf("%s/todos/%s", t.User.URL, t.ID)
	}
	return fmt.Sprintf("https://wip.co/@author/todos/%s", t.ID)
}

type Product struct {
	ID         graphql.ID `graphql:"id" json:"id,omitempty"`
	CreatedAt  *time.Time `graphql:"created_at" json:"created_at,omitempty"`
	Hashtag    string     `graphql:"hashtag" json:"hashtag,omitempty"`
	Name       string     `graphql:"name" json:"name,omitempty"`
	Pitch      string     `graphql:"pitch" json:"pitch,omitempty"`
	UpdatedAt  *time.Time `graphql:"updated_at" json:"updated_at,omitempty"`
	URL        string     `graphql:"url" json:"url,omitempty"`
	WebsiteURL string     `graphql:"website_url" json:"website_url,omitempty"`
	Makers     []User     `graphql:"makers" json:"makers,omitempty"`
	Todos      []Todo     `graphql:"todos" json:"todos,omitempty"`
}

type User struct {
	ID                  graphql.ID `graphql:"id" json:"id,omitempty"`
	URL                 string     `graphql:"url" json:"url,omitempty"`
	Username            string     `graphql:"username" json:"username,omitempty"`
	Firstname           string     `graphql:"first_name" json:"first_name,omitempty"`
	Lastname            string     `graphql:"last_name" json:"last_name,omitempty"`
	AvatarURL           string     `graphql:"avatar_url" json:"avatar_url,omitempty"`
	CompletedTodosCount int        `graphql:"completed_todos_count" json:"completed_todos_count,omitempty"`
	BestStreak          int        `graphql:"best_streak" json:"best_streak,omitempty"`
	Streaking           bool       `graphql:"streaking" json:"streaking,omitempty"`
	Todos               []Todo     `graphql:"todos" json:"todos,omitempty"`
	Products            []Product  `graphql:"products" json:"products,omitempty"`
}
