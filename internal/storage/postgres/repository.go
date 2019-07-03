package postgres

import (
	"context"
	"database/sql"

	"github.com/dipress/blog/internal/auth"
	"github.com/dipress/blog/internal/post"
	"github.com/dipress/blog/internal/reg"
	"github.com/dipress/blog/internal/user"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

const (
	driverName = "postgres"
)

// Repository holds crud actions.
type Repository struct {
	db *sqlx.DB
}

// NewRepository factory prepares repository to work.
func NewRepository(db *sql.DB) *Repository {
	r := Repository{
		db: sqlx.NewDb(db, driverName),
	}

	return &r
}

const createQuery = `INSERT INTO posts (user_id, title, body) VALUES ($1, $2, $3) RETURNING id, user_id, title, body, created_at, updated_at`

// CreatePost inserts a post into a database.
func (r *Repository) CreatePost(ctx context.Context, f *post.NewPost, post *post.Post) error {
	if err := r.db.QueryRowContext(ctx, createQuery, f.UserID, f.Title, f.Body).
		Scan(&post.ID, &post.UserID, &post.Title, &post.Body, &post.CreatedAt, &post.UpdatedAt); err != nil {
		return errors.Wrap(err, "query scan error")
	}

	return nil
}

const findPostQuery = `SELECT id, user_id, title, body, created_at, updated_at FROM posts where id = $1`

// FindPost finds post by id.
func (r *Repository) FindPost(ctx context.Context, id int) (*post.Post, error) {
	var p post.Post
	if err := r.db.QueryRowContext(ctx, findPostQuery, id).
		Scan(&p.ID, &p.UserID, &p.Title, &p.Body, &p.CreatedAt, &p.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, post.ErrNotFound
		}
		return nil, errors.Wrap(err, "query row scan")
	}
	return &p, nil
}

const updatePostQuery = `UPDATE posts SET title=:title, body=:body, updated_at=now() WHERE id=:id`

// UpdatePost updates post by id.
func (r *Repository) UpdatePost(ctx context.Context, id int, p *post.Post) error {
	stmt, err := r.db.PrepareNamed(updatePostQuery)
	if err != nil {
		return errors.Wrap(err, "prepare named")
	}
	defer stmt.Close()

	if _, err := stmt.ExecContext(ctx, map[string]interface{}{
		"id":    id,
		"title": p.Title,
		"body":  p.Body,
	}); err != nil {
		if err == sql.ErrNoRows {
			return post.ErrNotFound
		}
		return errors.Wrap(err, "exec context")
	}

	return nil
}

const deletePostQuery = "DELETE FROM posts WHERE id=:id"

// DeletePost deletes post by id.
func (r *Repository) DeletePost(ctx context.Context, id int) error {
	stmt, err := r.db.PrepareNamed(deletePostQuery)
	if err != nil {
		return errors.Wrap(err, "prepare named")
	}
	defer stmt.Close()

	if _, err := stmt.ExecContext(ctx, map[string]interface{}{
		"id": id,
	}); err != nil {
		if err == sql.ErrNoRows {
			return post.ErrNotFound
		}
		return errors.Wrap(err, "exec context")
	}
	return nil
}

const listPostQuery = `SELECT * FROM posts`

// ListPost shows all posts.
func (r *Repository) ListPost(ctx context.Context, pos *post.Posts) error {
	rows, err := r.db.QueryxContext(ctx, listPostQuery)
	if err != nil {
		return errors.Wrap(err, "query rows")
	}
	defer rows.Close()

	posts := make([]post.Post, 0)

	for rows.Next() {
		var post post.Post
		if err := rows.Scan(&post.ID, &post.UserID, &post.Title, &post.Body, &post.CreatedAt, &post.UpdatedAt); err != nil {
			errors.Wrap(err, "query row scan on loop")
		}
		posts = append(posts, post)
	}
	pos.Posts = posts

	return nil
}

const createUserQuery = `INSERT INTO users (username, email, password_hash) VALUES ($1, $2, $3) RETURNING id, username, email, password_hash, created_at, updated_at`

// CreateUser inserts a new user into the database.
func (r *Repository) CreateUser(ctx context.Context, f *user.NewUser, usr *user.User) (func() error, func() error, error) {
	tx, err := r.db.Beginx()
	if err != nil {
		return nil, nil, errors.Wrap(err, "begin tx")
	}

	if err := tx.QueryRowContext(ctx, createUserQuery, f.Username, f.Email, f.PasswordHash).
		Scan(&usr.ID, &usr.Username, &usr.Email, &usr.PasswordHash, &usr.CreatedAt, &usr.UpdatedAt); err != nil {
		return nil, nil, errors.Wrap(err, "query context scan")
	}
	return tx.Commit, tx.Rollback, nil
}

const uniqueUsernameQuery = `SELECT COUNT(*) FROM users WHERE username = $1`

// UniqueUsername checks that username is unique.
func (r *Repository) UniqueUsername(ctx context.Context, username string) error {
	var c int
	if err := r.db.QueryRowContext(ctx, uniqueUsernameQuery, username).Scan(&c); err != nil {
		return errors.Wrap(err, "scan error")
	}

	if c > 0 {
		return reg.ErrUsernameExists
	}

	return nil
}

const uniqueEmailQuery = `SELECT COUNT(*) FROM users WHERE email = $1`

// UniqueEmail checks that email address is unique.
func (r *Repository) UniqueEmail(ctx context.Context, email string) error {
	var c int
	if err := r.db.QueryRowContext(ctx, uniqueEmailQuery, email).Scan(&c); err != nil {
		return errors.Wrap(err, "scan error")
	}

	if c > 0 {
		return reg.ErrEmailExists
	}

	return nil
}

const emailFindQuery = `SELECT id, username, email, password_hash, created_at, updated_at FROM users WHERE email = $1`

// FindByEmail finds user by email.
func (r *Repository) FindByEmail(ctx context.Context, email string, user *user.User) error {
	if err := r.db.QueryRowContext(ctx, emailFindQuery, email).
		Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash, &user.CreatedAt, &user.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return auth.ErrNotFound
		}
		return errors.Wrap(err, "scan error")
	}

	return nil
}

const usernameFindQuery = `SELECT id, username, email, password_hash, created_at, updated_at FROM users WHERE username = $1`

// FindByUsername finds user by username.
func (r *Repository) FindByUsername(ctx context.Context, username string, u *user.User) error {
	if err := r.db.QueryRowContext(ctx, usernameFindQuery, username).
		Scan(&u.ID, &u.Username, &u.Email, &u.PasswordHash, &u.CreatedAt, &u.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return user.ErrNotFound
		}
		return errors.Wrap(err, "scan error")
	}

	return nil
}
