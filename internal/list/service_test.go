package list

import (
	"context"
	"errors"
	"testing"

	"github.com/dipress/blog/internal/post"
	"github.com/stretchr/testify/assert"
)

func TestServiceList(t *testing.T) {
	tests := []struct {
		name           string
		repositoryFunc func(ctx context.Context, pos *post.Posts) error
		wantErr        bool
	}{
		{
			name: "ok",
			repositoryFunc: func(ctx context.Context, pos *post.Posts) error {
				return nil
			},
		},
		{
			name: "repository error",
			repositoryFunc: func(ctx context.Context, pos *post.Posts) error {
				return errors.New("mock error")
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

			_, err := s.List(ctx)

			if tc.wantErr {
				assert.Error(t, err)
				return
			}
			assert.Nil(t, err)
		})
	}
}

type repositoryFunc func(ctx context.Context, pos *post.Posts) error

func (r repositoryFunc) ListPost(ctx context.Context, pos *post.Posts) error {
	return r(ctx, pos)
}
