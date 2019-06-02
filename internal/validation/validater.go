package validation

import (
	"context"

	"github.com/dipress/blog/internal/create"
	"github.com/dipress/blog/internal/reg"
	"github.com/dipress/blog/internal/update"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

const (
	mismatchMsg   = "mismatch"
	validationMsg = "you have validation errors"
)

// Errors holds validation errors.
type Errors map[string]string

// Error implements error interface.
func (v Errors) Error() string {
	return validationMsg
}

// Create holds create form validations.
type Create struct{}

// Validate validates post form for the create.
func (v *Create) Validate(ctx context.Context, f *create.Form) error {
	ves := make(Errors)

	if err := validation.Validate(f.Title,
		validation.Required,
		validation.Length(1, 50)); err != nil {
		ves["title"] = err.Error()
	}

	if err := validation.Validate(f.Body,
		validation.Required); err != nil {
		ves["body"] = err.Error()
	}

	if len(ves) > 0 {
		return ves
	}

	return nil
}

// Update holds update form validations.
type Update struct{}

// Validate validates post form for the update.
func (v *Update) Validate(ctx context.Context, f *update.Form) error {
	ves := make(Errors)

	if err := validation.Validate(f.Title,
		validation.Required,
		validation.Length(1, 50)); err != nil {
		ves["title"] = err.Error()
	}

	if err := validation.Validate(f.Body,
		validation.Required); err != nil {
		ves["body"] = err.Error()
	}

	if len(ves) > 0 {
		return ves
	}
	return nil
}

// Registrate holds create form validations.
type Registrate struct{}

// ValidateUser validates user form.
func (v *Registrate) ValidateUser(ctx context.Context, f *reg.Form) error {
	ves := make(Errors)

	if err := validation.Validate(f.Username,
		validation.Required,
		validation.Length(0, 50),
	); err != nil {
		ves["username"] = err.Error()
	}

	if err := validation.Validate(f.Email,
		validation.Required,
		is.Email,
	); err != nil {
		ves["email"] = err.Error()
	}

	if err := validation.Validate(f.Password,
		validation.Required,
		validation.Length(10, 0),
	); err != nil {
		ves["password"] = err.Error()
	}

	if len(ves) > 0 {
		return ves
	}

	return nil
}
