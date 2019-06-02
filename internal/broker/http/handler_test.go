package http

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/dipress/blog/internal/auth"
	"github.com/dipress/blog/internal/create"
	"github.com/dipress/blog/internal/post"
	"github.com/dipress/blog/internal/reg"
	"github.com/dipress/blog/internal/update"
	"github.com/dipress/blog/internal/validation"
	"github.com/gorilla/mux"
)

func TestCreateHandler(t *testing.T) {
	tests := []struct {
		name        string
		createrFunc func(ctx context.Context, f *create.Form) (*post.Post, error)
		code        int
	}{
		{
			name: "ok",
			createrFunc: func(ctx context.Context, f *create.Form) (*post.Post, error) {
				return &post.Post{}, nil
			},
			code: http.StatusOK,
		},
		{
			name: "validation errors",
			createrFunc: func(ctx context.Context, f *create.Form) (*post.Post, error) {
				return &post.Post{}, make(validation.Errors)
			},
			code: http.StatusUnprocessableEntity,
		},
		{
			name: "internal error",
			createrFunc: func(ctx context.Context, f *create.Form) (*post.Post, error) {
				return &post.Post{}, errors.New("mock error")
			},
			code: http.StatusInternalServerError,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			h := CreateHandler{createrFunc(tc.createrFunc)}
			w := httptest.NewRecorder()
			r := httptest.NewRequest("POST", "http://example.com", strings.NewReader("{}"))

			err := h.Handle(w, r)
			if w.Code != tc.code {
				t.Errorf("unexpected code: %d expected %d error: %v", w.Code, tc.code, err)
			}
		})
	}
}

type createrFunc func(ctx context.Context, f *create.Form) (*post.Post, error)

func (c createrFunc) Create(ctx context.Context, f *create.Form) (*post.Post, error) {
	return c(ctx, f)
}

func TestFindHandler(t *testing.T) {
	tests := []struct {
		name     string
		findFunc func(ctx context.Context, id int) (*post.Post, error)
		code     int
	}{
		{
			name: "ok",
			findFunc: func(ctx context.Context, id int) (*post.Post, error) {
				return &post.Post{}, nil
			},
			code: http.StatusOK,
		},
		{
			name: "internal error",
			findFunc: func(ctx context.Context, id int) (*post.Post, error) {
				return nil, errors.New("mock error")
			},
			code: http.StatusInternalServerError,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			h := FindHandler{findFunc(tc.findFunc)}
			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "http://example.com", strings.NewReader("{}"))
			r = mux.SetURLVars(r, map[string]string{"id": "1"})

			err := h.Handle(w, r)
			if w.Code != tc.code {
				t.Errorf("unexpected code: %d expected %d error: %v", w.Code, tc.code, err)
			}
		})
	}
}

type findFunc func(ctx context.Context, id int) (*post.Post, error)

func (f findFunc) Find(ctx context.Context, id int) (*post.Post, error) {
	return f(ctx, id)
}

func TestUpdateHandler(t *testing.T) {
	tests := []struct {
		name       string
		updateFunc func(ctx context.Context, id int, f *update.Form) (*post.Post, error)
		code       int
	}{
		{
			name: "ok",
			updateFunc: func(ctx context.Context, id int, f *update.Form) (*post.Post, error) {
				return &post.Post{}, nil
			},
			code: http.StatusOK,
		},
		{
			name: "validation errors",
			updateFunc: func(ctx context.Context, id int, f *update.Form) (*post.Post, error) {
				return &post.Post{}, make(validation.Errors)
			},
			code: http.StatusUnprocessableEntity,
		},
		{
			name: "internal error",
			updateFunc: func(ctx context.Context, id int, f *update.Form) (*post.Post, error) {
				return &post.Post{}, errors.New("mock error")
			},
			code: http.StatusInternalServerError,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			h := UpdateHandler{updateFunc(tc.updateFunc)}
			w := httptest.NewRecorder()
			r := httptest.NewRequest("PUT", "http://example.com", strings.NewReader("{}"))
			r = mux.SetURLVars(r, map[string]string{"id": "1"})

			err := h.Handle(w, r)
			if w.Code != tc.code {
				t.Errorf("unexpected code: %d expected %d error: %v", w.Code, tc.code, err)
			}
		})
	}
}

type updateFunc func(ctx context.Context, id int, f *update.Form) (*post.Post, error)

func (u updateFunc) Update(ctx context.Context, id int, f *update.Form) (*post.Post, error) {
	return u(ctx, id, f)
}

func TestDeleteHandler(t *testing.T) {
	tests := []struct {
		name       string
		deleteFunc func(ctx context.Context, id int) error
		code       int
	}{
		{
			name: "ok",
			deleteFunc: func(ctx context.Context, id int) error {
				return nil
			},
			code: http.StatusOK,
		},
		{
			name: "internal error",
			deleteFunc: func(ctx context.Context, id int) error {
				return errors.New("mock error")
			},
			code: http.StatusInternalServerError,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			h := DeleteHandler{deleteFunc(tc.deleteFunc)}
			w := httptest.NewRecorder()
			r := httptest.NewRequest("DELETE", "http://example.com", strings.NewReader("{}"))
			r = mux.SetURLVars(r, map[string]string{"id": "1"})

			err := h.Handle(w, r)
			if w.Code != tc.code {
				t.Errorf("unexpected code: %d expected %d error: %v", w.Code, tc.code, err)
			}
		})
	}
}

type deleteFunc func(ctx context.Context, id int) error

func (d deleteFunc) Delete(ctx context.Context, id int) error {
	return d(ctx, id)
}

func TestListHandler(t *testing.T) {
	tests := []struct {
		name     string
		listFunc func(ctx context.Context) (*post.Posts, error)
		code     int
	}{
		{
			name: "ok",
			listFunc: func(ctx context.Context) (*post.Posts, error) {
				return &post.Posts{}, nil
			},
			code: http.StatusOK,
		},
		{
			name: "internal error",
			listFunc: func(ctx context.Context) (*post.Posts, error) {
				return &post.Posts{}, errors.New("mock error")
			},
			code: http.StatusInternalServerError,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			h := ListHandler{listFunc(tc.listFunc)}
			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "http://example.com", strings.NewReader("{}"))

			err := h.Handle(w, r)
			if w.Code != tc.code {
				t.Errorf("unexpected code: %d expected %d error: %v", w.Code, tc.code, err)
			}
		})
	}
}

type listFunc func(ctx context.Context) (*post.Posts, error)

func (l listFunc) List(ctx context.Context) (*post.Posts, error) {
	return l(ctx)
}

func TestRegHandler(t *testing.T) {
	tests := []struct {
		name    string
		regFunc func(ctx context.Context, f *reg.Form, token *reg.Token) error
		code    int
	}{
		{
			name: "ok",
			regFunc: func(ctx context.Context, f *reg.Form, token *reg.Token) error {
				return nil
			},
			code: http.StatusOK,
		},
		{
			name: "validation errors",
			regFunc: func(ctx context.Context, f *reg.Form, token *reg.Token) error {
				return make(validation.Errors)
			},
			code: http.StatusUnprocessableEntity,
		},
		{
			name: "internal error",
			regFunc: func(ctx context.Context, f *reg.Form, token *reg.Token) error {
				return errors.New("mock error")
			},
			code: http.StatusInternalServerError,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			h := RegHandler{regFunc(tc.regFunc)}
			w := httptest.NewRecorder()
			r := httptest.NewRequest("POST", "http://example.com", strings.NewReader("{}"))

			err := h.Handle(w, r)
			if w.Code != tc.code {
				t.Errorf("unexpected code: %d expected %d error: %v", w.Code, tc.code, err)
			}
		})
	}
}

type regFunc func(ctx context.Context, f *reg.Form, token *reg.Token) error

func (r regFunc) Registrate(ctx context.Context, f *reg.Form, token *reg.Token) error {
	return r(ctx, f, token)
}

func TestAuthHandler(t *testing.T) {
	tests := []struct {
		name     string
		authFunc func(ctx context.Context, email, password string, t *auth.Token) error
		code     int
	}{
		{
			name: "ok",
			authFunc: func(ctx context.Context, email, password string, t *auth.Token) error {
				return nil
			},
			code: http.StatusOK,
		},
		{
			name: "email error",
			authFunc: func(ctx context.Context, email, password string, t *auth.Token) error {
				return auth.ErrNotFound
			},
			code: http.StatusUnauthorized,
		},
		{
			name: "password error",
			authFunc: func(ctx context.Context, email, password string, t *auth.Token) error {
				return auth.ErrWrongPassword
			},
			code: http.StatusUnauthorized,
		},
		{
			name: "internal error",
			authFunc: func(ctx context.Context, email, password string, t *auth.Token) error {
				return errors.New("mock error")
			},
			code: http.StatusInternalServerError,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			h := AuthHandler{authFunc(tc.authFunc)}
			w := httptest.NewRecorder()
			r := httptest.NewRequest("POST", "http://example.com", strings.NewReader("{}"))

			err := h.Handle(w, r)
			if w.Code != tc.code {
				t.Errorf("unexpected code: %d expected %d error: %v", w.Code, tc.code, err)
			}
		})
	}
}

type authFunc func(ctx context.Context, email, password string, t *auth.Token) error

func (a authFunc) Authenticate(ctx context.Context, email, password string, t *auth.Token) error {
	return a(ctx, email, password, t)
}
