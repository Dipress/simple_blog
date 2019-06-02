package reg

import (
	"context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func Test_Service(t *testing.T) {
	commit := func() error {
		return nil
	}
	rollback := func() error {
		return nil
	}

	tests := []struct {
		name               string
		validateFunc       func(mock *MockValidater)
		repositoryFunc     func(mock *MockRepository)
		tokenGeneratorFunc func(moock *MockTokenGenerator)
		wantErr            bool
	}{
		{
			name: "ok",
			validateFunc: func(m *MockValidater) {
				m.EXPECT().ValidateUser(gomock.Any(), gomock.Any()).Return(nil)
			},
			repositoryFunc: func(m *MockRepository) {
				m.EXPECT().UniqueUsername(gomock.Any(), gomock.Any()).Return(nil)
				m.EXPECT().UniqueEmail(gomock.Any(), gomock.Any()).Return(nil)
				m.EXPECT().CreateUser(gomock.Any(), gomock.Any(), gomock.Any()).Return(commit, rollback, nil)
			},
			tokenGeneratorFunc: func(m *MockTokenGenerator) {
				m.EXPECT().GenerateToken(gomock.Any(), gomock.Any()).Return("", nil)
			},
		},
		{
			name: "validation",
			validateFunc: func(m *MockValidater) {
				m.EXPECT().ValidateUser(gomock.Any(), gomock.Any()).Return(errors.New("mock error"))
			},
			repositoryFunc:     func(m *MockRepository) {},
			tokenGeneratorFunc: func(m *MockTokenGenerator) {},
			wantErr:            true,
		},
		{
			name: "unique username",
			validateFunc: func(m *MockValidater) {
				m.EXPECT().ValidateUser(gomock.Any(), gomock.Any()).Return(nil)
			},
			repositoryFunc: func(m *MockRepository) {
				m.EXPECT().UniqueUsername(gomock.Any(), gomock.Any()).Return(errors.New("mock error"))
			},
			tokenGeneratorFunc: func(m *MockTokenGenerator) {},
			wantErr:            true,
		},
		{
			name: "unique email",
			validateFunc: func(m *MockValidater) {
				m.EXPECT().ValidateUser(gomock.Any(), gomock.Any()).Return(nil)
			},
			repositoryFunc: func(m *MockRepository) {
				m.EXPECT().UniqueUsername(gomock.Any(), gomock.Any()).Return(nil)
				m.EXPECT().UniqueEmail(gomock.Any(), gomock.Any()).Return(errors.New("mock error"))
			},
			tokenGeneratorFunc: func(m *MockTokenGenerator) {},
			wantErr:            true,
		},
		{
			name: "create",
			validateFunc: func(m *MockValidater) {
				m.EXPECT().ValidateUser(gomock.Any(), gomock.Any()).Return(nil)
			},
			repositoryFunc: func(m *MockRepository) {
				m.EXPECT().UniqueUsername(gomock.Any(), gomock.Any()).Return(nil)
				m.EXPECT().UniqueEmail(gomock.Any(), gomock.Any()).Return(nil)
				m.EXPECT().CreateUser(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, rollback, errors.New("mock error"))
			},
			tokenGeneratorFunc: func(m *MockTokenGenerator) {},
			wantErr:            true,
		},
		{
			name: "token",
			validateFunc: func(m *MockValidater) {
				m.EXPECT().ValidateUser(gomock.Any(), gomock.Any()).Return(nil)
			},
			repositoryFunc: func(m *MockRepository) {
				m.EXPECT().UniqueUsername(gomock.Any(), gomock.Any()).Return(nil)
				m.EXPECT().UniqueEmail(gomock.Any(), gomock.Any()).Return(nil)
				m.EXPECT().CreateUser(gomock.Any(), gomock.Any(), gomock.Any()).Return(commit, rollback, nil)
			},
			tokenGeneratorFunc: func(m *MockTokenGenerator) {
				m.EXPECT().GenerateToken(gomock.Any(), gomock.Any()).Return("", errors.New("mock error"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			validator := NewMockValidater(ctrl)
			repo := NewMockRepository(ctrl)
			generator := NewMockTokenGenerator(ctrl)

			tt.validateFunc(validator)
			tt.repositoryFunc(repo)
			tt.tokenGeneratorFunc(generator)

			s := NewService(repo, validator, generator, time.Hour)

			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			form := Form{
				Username: "username",
				Email:    "username@example.com",
				Password: "password123",
			}

			var token Token

			err := s.Registrate(ctx, &form, &token)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.Nil(t, err)
		})

	}
}
