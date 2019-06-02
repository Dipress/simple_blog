package create

import (
	"context"

	"github.com/dipress/blog/internal/post"
	"github.com/dipress/blog/internal/user"
	"github.com/dipress/blog/kit/auth"
	"github.com/pkg/errors"
)

// easyjson service.go
// go:generate mockgen -source=service.go -package=create -destination=service.mock.go

// Validater validates post fields.
type Validater interface {
	Validate(context.Context, *Form) error
}

// Repository allows to work with the database.
type Repository interface {
	FindByUsername(ctx context.Context, username string, u *user.User) error
	CreatePost(ctx context.Context, f *post.NewPost, post *post.Post) error
}

// Form is a post form.
//easyjson:json
type Form struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

// Service is a use case for post validation and creation.
type Service struct {
	Repository
	Validater
}

// NewService factory prepares service for all futher operations.
func NewService(r Repository, v Validater) *Service {
	s := Service{
		Repository: r,
		Validater:  v,
	}

	return &s
}

// Create creates a post.
func (s *Service) Create(ctx context.Context, f *Form) (*post.Post, error) {
	if err := s.Validater.Validate(ctx, f); err != nil {
		return nil, errors.Wrap(err, "validater validate")
	}
	claims, _ := auth.FromContext(ctx)

	var u user.User
	if err := s.Repository.FindByUsername(ctx, claims.Subject, &u); err != nil {
		return nil, errors.Wrap(err, "repository find user")
	}

	np := post.NewPost{
		UserID: u.ID,
		Title:  f.Title,
		Body:   f.Body,
	}

	var p post.Post
	err := s.Repository.CreatePost(ctx, &np, &p)
	if err != nil {
		return nil, errors.Wrap(err, "repository create post")
	}
	return &p, nil
}
