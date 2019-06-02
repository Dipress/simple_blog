package post

import (
	"errors"
	"time"
)

// easyjson -all model.go

var (
	// ErrNotFound raises when post not found in the database.
	ErrNotFound = errors.New("post not found")
)

// Post contains all post field.
type Post struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// NewPost contains the information which needs to create a new Post.
type NewPost struct {
	UserID int
	Title  string
	Body   string
}

// Posts contains slice of posts.
type Posts struct {
	Posts []Post `json:"posts"`
}
