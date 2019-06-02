package delete

import (
	"context"

	"github.com/dipress/blog/internal/post"
	"github.com/dipress/blog/internal/user"
	"github.com/dipress/blog/kit/auth"
	"github.com/pkg/errors"
)

// go:generate mockgen -source=service.go -package=delete -destination=service.mock.go

// Abillity checks permissions to view posts.
type Abillity interface {
	CanDelete(userID int, post *post.Post) bool
}

// Repository allows to work with the database.
type Repository interface {
	FindByUsername(ctx context.Context, username string, u *user.User) error
	FindPost(ctx context.Context, id int) (*post.Post, error)
	DeletePost(ctx context.Context, id int) error
}

// Service is a use case for post delete.
type Service struct {
	Repository
	Abillity
}

// NewService factory prepares service for all futher operations.
func NewService(r Repository, a Abillity) *Service {
	s := Service{
		Repository: r,
		Abillity:   a,
	}

	return &s
}

// Delete deletes a post.
func (s *Service) Delete(ctx context.Context, id int) error {
	p, err := s.Repository.FindPost(ctx, id)
	if err != nil {
		return errors.Wrap(err, "find post")
	}

	claims, _ := auth.FromContext(ctx)

	var u user.User
	if err := s.Repository.FindByUsername(ctx, claims.Subject, &u); err != nil {
		return errors.Wrap(err, "repository find user")
	}

	ok := s.Abillity.CanDelete(u.ID, p)
	if !ok {
		return post.ErrNotFound
	}

	if err := s.Repository.DeletePost(ctx, p.ID); err != nil {
		return errors.Wrap(err, "delete post")
	}

	return nil
}
