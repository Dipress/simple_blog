package http

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/dipress/blog/internal/auth"
	"github.com/dipress/blog/internal/create"
	"github.com/dipress/blog/internal/post"
	"github.com/dipress/blog/internal/reg"
	"github.com/dipress/blog/internal/update"
	"github.com/dipress/blog/internal/validation"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

// Creater abstraction for create service.
type Creater interface {
	Create(ctx context.Context, f *create.Form) (*post.Post, error)
}

// Finder abstraction for find service.
type Finder interface {
	Find(ctx context.Context, id int) (*post.Post, error)
}

// Updater abstraction for update service.
type Updater interface {
	Update(ctx context.Context, id int, f *update.Form) (*post.Post, error)
}

// Deleter abstraction for delete service.
type Deleter interface {
	Delete(ctx context.Context, id int) error
}

// Lister abstraction for list service.
type Lister interface {
	List(ctx context.Context) (*post.Posts, error)
}

// Registrater abstraction for registrate service.
type Registrater interface {
	Registrate(ctx context.Context, f *reg.Form, token *reg.Token) error
}

// Authenticater abstraction for authenticate service.
type Authenticater interface {
	Authenticate(ctx context.Context, email, password string, t *auth.Token) error
}

// Handler allows to handle requests.
type Handler interface {
	Handle(w http.ResponseWriter, r *http.Request) error
}

// RegHandler for registrate request.
type RegHandler struct {
	Registrater
}

// Handle implements Handler interface.
func (h *RegHandler) Handle(w http.ResponseWriter, r *http.Request) error {
	var f reg.Form
	var t reg.Token

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return errors.Wrap(badRequestResponse(w), "read body")
	}

	if err := f.UnmarshalJSON(data); err != nil {
		return errors.Wrap(badRequestResponse(w), "unmarshal json")
	}

	if err := h.Registrater.Registrate(r.Context(), &f, &t); err != nil {
		switch v := errors.Cause(err).(type) {
		case validation.Errors:
			return errors.Wrap(unprocessabeEntityResponse(w, v), "validation response")
		default:
			return errors.Wrap(internalServerErrorResponse(w), "registrate")
		}
	}

	data, err = t.MarshalJSON()
	if err != nil {
		return errors.Wrap(err, "marshal json")
	}

	if _, err := w.Write(data); err != nil {
		return errors.Wrap(err, "write response")
	}

	return nil
}

// AuthHandler for authenticate request.
type AuthHandler struct {
	Authenticater
}

// Handle implements Handler interface.
func (a AuthHandler) Handle(w http.ResponseWriter, r *http.Request) error {
	var f auth.Form
	var t auth.Token

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return errors.Wrap(badRequestResponse(w), "read body")
	}

	if err := f.UnmarshalJSON(data); err != nil {
		return errors.Wrap(badRequestResponse(w), "unmarshal json")
	}

	if err := a.Authenticater.Authenticate(r.Context(), f.Email, f.Password, &t); err != nil {
		switch err := errors.Cause(err); err {
		case auth.ErrNotFound, auth.ErrWrongPassword:
			return errors.Wrap(unauthorizedResponse(w), "find user")
		default:
			return errors.Wrap(internalServerErrorResponse(w), "authenticate")
		}
	}

	data, err = t.MarshalJSON()
	if err != nil {
		return errors.Wrap(err, "marshal json")
	}

	if _, err := w.Write(data); err != nil {
		return errors.Wrap(err, "write response")
	}

	return nil
}

// CreateHandler for create requests.
type CreateHandler struct {
	Creater
}

// Handle implements Handler interface.
func (h *CreateHandler) Handle(w http.ResponseWriter, r *http.Request) error {
	var f create.Form

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return errors.Wrap(badRequestResponse(w), "read body")
	}

	if err := f.UnmarshalJSON(data); err != nil {
		return errors.Wrap(badRequestResponse(w), "unmarshal json")
	}

	post, err := h.Creater.Create(r.Context(), &f)
	if err != nil {
		switch v := errors.Cause(err).(type) {
		case validation.Errors:
			return errors.Wrap(unprocessabeEntityResponse(w, v), "validation response")
		default:
			return errors.Wrap(internalServerErrorResponse(w), "create post")
		}
	}

	data, err = post.MarshalJSON()
	if err != nil {
		return errors.Wrap(err, "marshal json")
	}

	if _, err := w.Write(data); err != nil {
		return errors.Wrap(err, "write response")
	}

	return nil
}

// FindHandler for find requests.
type FindHandler struct {
	Finder
}

// Handle implements Handler interface.
func (f *FindHandler) Handle(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		return errors.Wrapf(badRequestResponse(w), "convert id query param to int: %v", err)
	}

	p, err := f.Finder.Find(r.Context(), id)
	if err != nil {
		switch errors.Cause(err) {
		case post.ErrNotFound:
			return errors.Wrap(notFoundResponse(w), "find")
		default:
			return errors.Wrap(internalServerErrorResponse(w), "find")
		}
	}

	data, err := p.MarshalJSON()
	if err != nil {
		return errors.Wrap(err, "marshal json")
	}

	if _, err := w.Write(data); err != nil {
		return errors.Wrap(err, "write response")
	}

	return nil
}

// UpdateHandler for update requests.
type UpdateHandler struct {
	Updater
}

// Handle implements Handler interface.
func (h *UpdateHandler) Handle(w http.ResponseWriter, r *http.Request) error {
	var f update.Form
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		return errors.Wrapf(badRequestResponse(w), "convert id query param to int: %v", err)
	}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return errors.Wrap(badRequestResponse(w), "read body")
	}

	if err := f.UnmarshalJSON(data); err != nil {
		return errors.Wrap(badRequestResponse(w), "unmarshal json")
	}

	p, err := h.Updater.Update(r.Context(), id, &f)
	if err != nil {
		switch v := errors.Cause(err).(type) {
		case validation.Errors:
			return errors.Wrap(unprocessabeEntityResponse(w, v), "validation response")
		default:
			return errors.Wrap(internalServerErrorResponse(w), "update")
		}
	}

	data, err = p.MarshalJSON()
	if err != nil {
		return errors.Wrap(err, "marshal json")
	}

	if _, err := w.Write(data); err != nil {
		return errors.Wrap(err, "write response")
	}

	return nil
}

// DeleteHandler for delete requests.
type DeleteHandler struct {
	Deleter
}

// Handle implements Handler interface.
func (h DeleteHandler) Handle(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		return errors.Wrapf(badRequestResponse(w), "convert id query param to int: %v", err)
	}
	// TODO:
	// check if permission error from abillity

	if err := h.Deleter.Delete(r.Context(), id); err != nil {
		return errors.Wrap(internalServerErrorResponse(w), "delete")
	}

	return nil
}

// ListHandler for list requests.
type ListHandler struct {
	Lister
}

// Handle implements Handler interface.
func (h ListHandler) Handle(w http.ResponseWriter, r *http.Request) error {
	p, err := h.Lister.List(r.Context())
	if err != nil {
		return errors.Wrap(internalServerErrorResponse(w), "list")
	}

	data, err := p.MarshalJSON()
	if err != nil {
		return errors.Wrap(err, "marshal json")
	}

	if _, err := w.Write(data); err != nil {
		return errors.Wrap(err, "write response")
	}
	return nil
}

// httpHandler allows to implement ServeHTTP for Handler.
type httpHandler struct {
	Handler
}

// ServeHTTP implements http.Handler.
func (h httpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h.Handle(w, r); err != nil {
		log.Printf("serve http: %+v\n", err)
	}
}
