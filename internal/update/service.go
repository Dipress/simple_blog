package update

import (
	"context"

	"github.com/dipress/blog/internal/post"
	"github.com/dipress/blog/internal/user"
	"github.com/dipress/blog/kit/auth"
	"github.com/pkg/errors"
)

// easyjson service.go
// go:generate mockgen -source=service.go -package=update -destination=service.mock.go

// Abillity checks permissions to view posts.
type Abillity interface {
	CanUpdate(userID int, post *post.Post) bool
}

// Validater validates post fields.
type Validater interface {
	Validate(context.Context, *Form) error
}

// Repository allows to work with the database.
type Repository interface {
	FindByUsername(ctx context.Context, username string, u *user.User) error
	FindPost(ctx context.Context, id int) (*post.Post, error)
	UpdatePost(ctx context.Context, id int, p *post.Post) error
}

// Form is a post form.
//easyjson:json
type Form struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

// Service is a use case for post validation and updation.
type Service struct {
	Repository
	Validater
	Abillity
}

// NewService factory prepares service for all futher operations.
func NewService(r Repository, v Validater, a Abillity) *Service {
	s := Service{
		Repository: r,
		Validater:  v,
		Abillity:   a,
	}
	return &s
}

// Update updates a post.
func (s *Service) Update(ctx context.Context, id int, f *Form) (*post.Post, error) {
	if err := s.Validater.Validate(ctx, f); err != nil {
		return nil, errors.Wrap(err, "validater validate")
	}

	p, err := s.Repository.FindPost(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "find post")
	}

	claims, _ := auth.FromContext(ctx)

	var u user.User
	if err := s.Repository.FindByUsername(ctx, claims.Subject, &u); err != nil {
		return nil, errors.Wrap(err, "repository find user")
	}

	ok := s.Abillity.CanUpdate(u.ID, p)
	if !ok {
		return nil, post.ErrNotFound
	}

	p.Title = f.Title
	p.Body = f.Body

	if err := s.Repository.UpdatePost(ctx, id, p); err != nil {
		return nil, errors.Wrap(err, "update post")
	}
	return p, nil
}
