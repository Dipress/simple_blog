package validation

import (
	"context"
	"reflect"
	"testing"

	"github.com/dipress/blog/internal/create"
	"github.com/dipress/blog/internal/reg"
)

func TestCreateValidate(t *testing.T) {
	tests := []struct {
		name    string
		form    create.Form
		wantErr bool
		expect  Errors
	}{
		{
			name: "valid",
			form: create.Form{
				Title: "title",
				Body:  "body",
			},
		},
		{
			name: "missing title",
			form: create.Form{
				Body: "body",
			},
			wantErr: true,
			expect: Errors{
				"title": "cannot be blank",
			},
		},
		{
			name: "blank title",
			form: create.Form{
				Title: "",
				Body:  "body",
			},
			wantErr: true,
			expect: Errors{
				"title": "cannot be blank",
			},
		},
		{
			name: "long title",
			form: create.Form{
				Title: "This is long title, this title is way larger is allowed one, an ti's used for testing.",
				Body:  "body",
			},
			wantErr: true,
			expect: Errors{
				"title": "the length must be between 1 and 50",
			},
		},
		{
			name: "missing body",
			form: create.Form{
				Title: "title",
			},
			wantErr: true,
			expect: Errors{
				"body": "cannot be blank",
			},
		},
		{
			name: "blank body",
			form: create.Form{
				Title: "title",
				Body:  "",
			},
			wantErr: true,
			expect: Errors{
				"body": "cannot be blank",
			},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			var v Create
			err := v.Validate(ctx, &tc.form)

			if tc.wantErr {
				got, ok := err.(Errors)
				if !ok {
					t.Errorf("unknown error: %v", err)
					return
				}

				if !reflect.DeepEqual(tc.expect, got) {
					t.Errorf("expected: %+#v got: %+#v", tc.expect, got)
				}

				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		})
	}
}

func TestCreateUserValidate(t *testing.T) {
	tests := []struct {
		name    string
		form    reg.Form
		wantErr bool
		expect  Errors
	}{
		{
			name: "valid",
			form: reg.Form{
				Username: "shepard",
				Email:    "johndoe@example.com",
				Password: "password123",
			},
		},
		{
			name: "missing username",
			form: reg.Form{
				Email:    "johndoe@example.com",
				Password: "password123",
			},
			wantErr: true,
			expect: Errors{
				"username": "cannot be blank",
			},
		},
		{
			name: "blank username",
			form: reg.Form{
				Username: "",
				Email:    "johndoe@example.com",
				Password: "password123",
			},
			wantErr: true,
			expect: Errors{
				"username": "cannot be blank",
			},
		},
		{
			name: "missing email",
			form: reg.Form{
				Username: "shepard",
				Password: "password123",
			},
			wantErr: true,
			expect: Errors{
				"email": "cannot be blank",
			},
		},
		{
			name: "blank email",
			form: reg.Form{
				Username: "shepard",
				Email:    "",
				Password: "password123",
			},
			wantErr: true,
			expect: Errors{
				"email": "cannot be blank",
			},
		},
		{
			name: "not valid email",
			form: reg.Form{
				Username: "shepard",
				Email:    "johndo",
				Password: "password123",
			},
			wantErr: true,
			expect: Errors{
				"email": "must be a valid email address",
			},
		},
		{
			name: "missing password",
			form: reg.Form{
				Username: "shepard",
				Email:    "johndoe@example.com",
			},
			wantErr: true,
			expect: Errors{
				"password": "cannot be blank",
			},
		},
		{
			name: "blank password",
			form: reg.Form{
				Username: "shepard",
				Email:    "johndoe@example.com",
				Password: "",
			},
			wantErr: true,
			expect: Errors{
				"password": "cannot be blank",
			},
		},
		{
			name: "short password",
			form: reg.Form{
				Username: "shepard",
				Email:    "johndo@example.com",
				Password: "password",
			},
			wantErr: true,
			expect: Errors{
				"password": "the length must be no less than 10",
			},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			var v Registrate
			err := v.ValidateUser(ctx, &tc.form)

			if tc.wantErr {
				got, ok := err.(Errors)
				if !ok {
					t.Errorf("unknown error: %v", err)
					return
				}

				if !reflect.DeepEqual(tc.expect, got) {
					t.Errorf("expected: %+#v got: %+#v", tc.expect, got)
				}

				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		})
	}
}
