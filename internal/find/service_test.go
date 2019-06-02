package find

import (
	"context"
	"errors"
	"testing"

	"github.com/dipress/blog/internal/post"
	"github.com/stretchr/testify/assert"
)

func Test_Service(t *testing.T) {
	tests := []struct {
		name           string
		repositoryFunc func(ctx context.Context, id int) (*post.Post, error)
		wantErr        bool
	}{
		{
			name: "ok",
			repositoryFunc: func(ctx context.Context, id int) (*post.Post, error) {
				return &post.Post{}, nil
			},
		},
		{
			name: "repository error",
			repositoryFunc: func(ctx context.Context, id int) (*post.Post, error) {
				return nil, errors.New("mock error")
			},
			wantErr: true,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			s := NewService(repositoryFunc(tc.repositoryFunc))
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			_, err := s.FindPost(ctx, 1)
			if tc.wantErr {
				assert.Error(t, err)
				return
			}

			assert.Nil(t, err)
		})
	}
}

type repositoryFunc func(ctx context.Context, id int) (*post.Post, error)

func (r repositoryFunc) FindPost(ctx context.Context, id int) (*post.Post, error) {
	return r(ctx, id)
}
