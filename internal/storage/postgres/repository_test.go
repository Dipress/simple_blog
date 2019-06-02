package postgres

import (
	"context"
	"testing"

	"github.com/dipress/blog/internal/post"
	"github.com/dipress/blog/internal/user"
	"github.com/stretchr/testify/assert"
)

func TestCreatePost(t *testing.T) {
	t.Parallel()
	t.Log("with initialized repository")
	{
		db, teardown := postgresDB(t)
		defer teardown()
		r := NewRepository(db)

		t.Log("\ttest:0\tshould insert a new post into the database")
		{
			np := post.NewPost{
				UserID: 1,
				Title:  "Article title",
				Body:   "article body",
			}
			var post post.Post
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			err := r.CreatePost(ctx, &np, &post)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if post.ID == 0 {
				t.Error("expected to parse returned id")
			}
		}
	}
}

func TestFindPost(t *testing.T) {
	t.Log("with initialized repository")
	{
		db, teardown := postgresDB(t)
		defer teardown()
		r := NewRepository(db)

		np := post.NewPost{
			UserID: 2,
			Title:  "Article title",
			Body:   "article body",
		}
		var post post.Post
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		err := r.CreatePost(ctx, &np, &post)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		t.Log("\ttest:0\tshould find the post into the database")
		{
			_, err := r.FindPost(ctx, 1)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		}
	}
}

func TestUpdatePost(t *testing.T) {
	t.Log("with initialized repository")
	{
		db, teardown := postgresDB(t)
		defer teardown()
		r := NewRepository(db)

		np := post.NewPost{
			UserID: 3,
			Title:  "post title",
			Body:   "post body",
		}
		var post post.Post
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		err := r.CreatePost(ctx, &np, &post)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		t.Log("\ttest:0\tshould update the post into the database")
		{
			err := r.UpdatePost(ctx, 1, &post)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		}
	}
}

func TestDeletePost(t *testing.T) {
	t.Log("with initialized repository")
	{
		db, teardown := postgresDB(t)
		defer teardown()
		r := NewRepository(db)

		np := post.NewPost{
			UserID: 4,
			Title:  "post title",
			Body:   "post body",
		}
		var post post.Post
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		err := r.CreatePost(ctx, &np, &post)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		t.Log("\ttest:0\tshould delete the post into the database")
		{
			err := r.DeletePost(ctx, 1)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		}
	}
}

func TestListPost(t *testing.T) {
	t.Log("with initialized repository")
	{
		db, teardown := postgresDB(t)
		defer teardown()
		r := NewRepository(db)

		np := post.NewPost{
			UserID: 5,
			Title:  "post title",
			Body:   "post body",
		}
		var p post.Post
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		err := r.CreatePost(ctx, &np, &p)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		t.Log("\ttest:0\tshould show list of posts into the database")
		{
			var posts post.Posts
			err := r.ListPost(ctx, &posts)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if len(posts.Posts) == 0 {
				t.Error("expected to slice of posts")
			}
		}
	}
}

func TestCreateUser(t *testing.T) {
	t.Parallel()
	t.Log("with initialized repository")
	{
		db, teardown := postgresDB(t)
		defer teardown()
		r := NewRepository(db)

		t.Log("\ttest:0\tshould insert a new user into the database")
		{
			nu := user.NewUser{
				Username:     "username",
				Email:        "username@example.com",
				PasswordHash: "$2y$12$gwoUXq7kCxNcucd.eFxOp.vJYYmo6917fSGuuEowfyNf3E8KySrWC",
			}
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			var user user.User
			_, _, err := r.CreateUser(ctx, &nu, &user)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if user.ID == 0 {
				t.Error("expected to parse returned id")
			}
		}
	}
}

func TestUniqueUsername(t *testing.T) {
	t.Parallel()
	t.Log("with initialized repository")
	{
		db, teardown := postgresDB(t)
		defer teardown()
		r := NewRepository(db)

		nu := user.NewUser{
			Username:     "username1",
			Email:        "username1@example.com",
			PasswordHash: "$2y$12$gwoUXq7kCxNcucd.eFxOp.vJYYmo6917fSGuuEowfyNf3E8KySrWC",
		}

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var user user.User
		_, _, err := r.CreateUser(ctx, &nu, &user)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		t.Log("\ttest:0\tshould return error")
		{
			err := r.UniqueUsername(ctx, "username1")
			assert.Error(t, err, "username already exists")
		}
		t.Log("\ttest:1\tshould return nil")
		{
			err := r.UniqueUsername(ctx, "username2")
			assert.Nil(t, err)
		}
	}
}

func TestUniqueEmail(t *testing.T) {
	t.Parallel()
	t.Log("with initialized repository")
	{
		db, teardown := postgresDB(t)
		defer teardown()
		r := NewRepository(db)

		nu := user.NewUser{
			Username:     "username2",
			Email:        "username2@example.com",
			PasswordHash: "$2y$12$gwoUXq7kCxNcucd.eFxOp.vJYYmo6917fSGuuEowfyNf3E8KySrWC",
		}

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var user user.User
		_, _, err := r.CreateUser(ctx, &nu, &user)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		t.Log("\ttest:0\tshould return error")
		{
			err := r.UniqueEmail(ctx, "username2@example.com")
			assert.Error(t, err, "email already exists")
		}
		t.Log("\ttest:1\tshould return nil")
		{
			err := r.UniqueEmail(ctx, "username3@example.com")
			assert.Nil(t, err)
		}
	}
}

func TestFindByEmail(t *testing.T) {
	t.Parallel()
	t.Log("with initialized repository")
	{
		db, teardown := postgresDB(t)
		defer teardown()
		r := NewRepository(db)

		nu := user.NewUser{
			Username:     "username5",
			Email:        "username5@example.com",
			PasswordHash: "$2y$12$gwoUXq7kCxNcucd.eFxOp.vJYYmo6917fSGuuEowfyNf3E8KySrWC",
		}

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var user user.User
		_, _, err := r.CreateUser(ctx, &nu, &user)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		t.Log("\ttest:0\tshould find user by email")
		{
			err := r.FindByEmail(ctx, user.Email, &user)
			assert.Nil(t, err)
		}
	}
}

func TestFindByUsername(t *testing.T) {
	t.Parallel()
	t.Log("with initialized repository")
	{
		db, teardown := postgresDB(t)
		defer teardown()
		r := NewRepository(db)

		nu := user.NewUser{
			Username:     "username6",
			Email:        "username6@example.com",
			PasswordHash: "$2y$12$gwoUXq7kCxNcucd.eFxOp.vJYYmo6917fSGuuEowfyNf3E8KySrWC",
		}

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var user user.User
		_, _, err := r.CreateUser(ctx, &nu, &user)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		t.Log("\ttest:0\tshould find user by username")
		{
			err := r.FindByUsername(ctx, user.Username, &user)
			assert.Nil(t, err)
		}
	}
}
