package delete

import (
	"context"
	"errors"
	"testing"

	post "github.com/dipress/blog/internal/post"
	"github.com/dipress/blog/kit/auth"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestServiceDelete(t *testing.T) {
	tests := []struct {
		name           string
		repositoryFunc func(mock *MockRepository)
		abilityFunc    func(mock *MockAbillity)
		wantErr        bool
	}{
		{
			name: "ok",
			repositoryFunc: func(m *MockRepository) {
				m.EXPECT().FindPost(gomock.Any(), gomock.Any()).Return(&post.Post{}, nil)
				m.EXPECT().FindByUsername(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
				m.EXPECT().DeletePost(gomock.Any(), gomock.Any()).Return(nil)
			},
			abilityFunc: func(m *MockAbillity) {
				m.EXPECT().CanDelete(gomock.Any(), gomock.Any()).Return(true)
			},
		},
		{
			name: "find post error",
			repositoryFunc: func(m *MockRepository) {
				m.EXPECT().FindPost(gomock.Any(), gomock.Any()).Return(&post.Post{}, errors.New("mock error"))
			},
			abilityFunc: func(m *MockAbillity) {},
			wantErr:     true,
		},
		{
			name: "find user error",
			repositoryFunc: func(m *MockRepository) {
				m.EXPECT().FindPost(gomock.Any(), gomock.Any()).Return(&post.Post{}, nil)
				m.EXPECT().FindByUsername(gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New("mock error"))
			},
			abilityFunc: func(m *MockAbillity) {},
			wantErr:     true,
		},
		{
			name: "ability error",
			repositoryFunc: func(m *MockRepository) {
				m.EXPECT().FindPost(gomock.Any(), gomock.Any()).Return(&post.Post{}, nil)
				m.EXPECT().FindByUsername(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			},
			abilityFunc: func(m *MockAbillity) {
				m.EXPECT().CanDelete(gomock.Any(), gomock.Any()).Return(false)
			},
			wantErr: true,
		},
		{
			name: "delete error",
			repositoryFunc: func(m *MockRepository) {
				m.EXPECT().FindPost(gomock.Any(), gomock.Any()).Return(&post.Post{}, nil)
				m.EXPECT().FindByUsername(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
				m.EXPECT().DeletePost(gomock.Any(), gomock.Any()).Return(errors.New("mock error"))
			},
			abilityFunc: func(m *MockAbillity) {
				m.EXPECT().CanDelete(gomock.Any(), gomock.Any()).Return(true)
			},
			wantErr: true,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := NewMockRepository(ctrl)
			ability := NewMockAbillity(ctrl)
			tc.repositoryFunc(repo)
			tc.abilityFunc(ability)

			s := NewService(repo, ability)

			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			claims := auth.Claims{}
			newCtx := auth.ToContext(ctx, &claims)

			err := s.Delete(newCtx, 1)

			if tc.wantErr {
				assert.Error(t, err)
				return
			}
			assert.Nil(t, err)
		})
	}
}
