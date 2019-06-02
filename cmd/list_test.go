package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"testing"

	"github.com/dipress/blog/internal/post"
	"github.com/dipress/blog/internal/storage/postgres"
)

func TestListPost(t *testing.T) {
	t.Log("with prepared server")
	{
		db, teardown := postgresDB(t)
		defer teardown()

		ctx, cancel := context.WithTimeout(context.Background(), caseTimeout)
		defer cancel()

		repo := postgres.NewRepository(db)

		np := post.NewPost{
			UserID: 7,
			Title:  "my title",
			Body:   "my body",
		}
		var p post.Post
		if err := repo.CreatePost(ctx, &np, &p); err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		lis, err := net.Listen("tcp", "127.0.0.1:0")
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		s := setupServer(lis.Addr().String(), db, authenticator)
		go s.Serve(lis)
		defer s.Close()

		t.Log("\ttest:0\tshould show all posts.")
		{
			req, err := http.NewRequest("GET", fmt.Sprintf("http://%s/posts", s.Addr), nil)
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
