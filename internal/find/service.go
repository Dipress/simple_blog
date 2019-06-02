package find

import (
	"context"

	"github.com/dipress/blog/internal/post"
	"github.com/pkg/errors"
)

// Repository allows to work with the database.
type Repository interface {
	FindPost(ctx context.Context, id int) (*post.Post, error)
}

// Service is a use case for post finding.
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

// Find finds post
func (s *Service) Find(ctx context.Context, id int) (*post.Post, error) {
	p, err := s.Repository.FindPost(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "repository find")
	}
	return p, nil
}
