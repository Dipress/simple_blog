package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/dipress/blog/internal/storage/postgres"
	"github.com/dipress/blog/internal/user"
	"github.com/dipress/blog/kit/auth"
)

func TestCreatePost(t *testing.T) {
	t.Log("with prepared server")
	{
		db, teardown := postgresDB(t)
		defer teardown()

		ctx, cancel := context.WithTimeout(context.Background(), caseTimeout)
		defer cancel()

		repo := postgres.NewRepository(db)

		nu := user.NewUser{
			Username:     "username7",
			Email:        "username7@example.com",
			PasswordHash: "$2y$12$e4.VBLqKAanAZs10dRL65O8.b0kHBC34pcGCN1HdJIchCi9im40Ei",
		}
		var u user.User
		_, _, err := repo.CreateUser(ctx, &nu, &u)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		claims := auth.NewClaims(u.Username, time.Now(), time.Hour)

		token, err := authenticator.GenerateToken(ctx, claims)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		token = "Bearer " + token

		lis, err := net.Listen("tcp", "127.0.0.1:0")
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		s := setupServer(lis.Addr().String(), db, authenticator)
		go s.Serve(lis)
		defer s.Close()

		t.Log("\ttest:0\tshould create a post.")
		{
			postStr := `{"title": "my awesome title", "body": "my awesome body"}`
			req, err := http.NewRequest("POST", fmt.Sprintf("http://%s/posts", s.Addr), strings.NewReader(postStr))
			req.Header.Set("Content-Type", "application/json")
			req.Header.Add("Authorization", token)

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if resp.StatusCode != http.StatusOK {
				t.Errorf("unexpected status code: %d expected: %d", resp.StatusCode, http.StatusOK)
			}
		}
	}
}
