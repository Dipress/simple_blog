package list

import (
	"context"

	"github.com/dipress/blog/internal/post"
	"github.com/pkg/errors"
)

// Repository allows to work with the database.
type Repository interface {
	ListPost(ctx context.Context, pos *post.Posts) error
}

// Service is a use case for posts showing.
type Service struct {
	Repository
}

// NewService factory prepares service for all futher operations.
func NewService(r Repository) *Service {
	s := Service{
		Repository: r,
	}

	return &s
}

// List shows all posts.
func (s *Service) List(ctx context.Context) (*post.Posts, error) {
	var posts post.Posts
	if err := s.Repository.ListPost(ctx, &posts); err != nil {
		return nil, errors.Wrap(err, "list posts")
	}

	return &posts, nil
}
