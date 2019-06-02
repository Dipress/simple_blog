package create

import (
	"context"
	"errors"
	"testing"

	"github.com/dipress/blog/kit/auth"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestServiceCreate(t *testing.T) {
	tests := []struct {
		name           string
		validateFunc   func(mock *MockValidater)
		repositoryFunc func(mock *MockRepository)
		wantErr        bool
	}{
		{
			name: "ok",
			validateFunc: func(m *MockValidater) {
				m.EXPECT().Validate(gomock.Any(), gomock.Any()).Return(nil)
			},
			repositoryFunc: func(m *MockRepository) {
				m.EXPECT().FindByUsername(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
				m.EXPECT().CreatePost(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			},
		},
		{
			name: "validation",
			validateFunc: func(m *MockValidater) {
				m.EXPECT().Validate(gomock.Any(), gomock.Any()).Return(errors.New("mock error"))
			},
			repositoryFunc: func(m *MockRepository) {},
			wantErr:        true,
		},
		{
			name: "find user",
			validateFunc: func(m *MockValidater) {
				m.EXPECT().Validate(gomock.Any(), gomock.Any()).Return(nil)
			},
			repositoryFunc: func(m *MockRepository) {
				m.EXPECT().FindByUsername(gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New("mock error"))
			},
			wantErr: true,
		},
		{
			name: "create post",
			validateFunc: func(m *MockValidater) {
				m.EXPECT().Validate(gomock.Any(), gomock.Any()).Return(nil)
			},
			repositoryFunc: func(m *MockRepository) {
				m.EXPECT().FindByUsername(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
				m.EXPECT().CreatePost(gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New("mock error"))
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

			validator := NewMockValidater(ctrl)
			repo := NewMockRepository(ctrl)
			tc.validateFunc(validator)
			tc.repositoryFunc(repo)

			s := NewService(repo, validator)

			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			claims := auth.Claims{}
			newCtx := auth.ToContext(ctx, &claims)

			form := Form{
				Title: "my awesome titie",
				Body:  "my awesome body",
			}

			_, err := s.Create(newCtx, &form)
			if tc.wantErr {
				assert.Error(t, err)
				return
			}
			assert.Nil(t, err)
		})
	}
}
