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

func TestFindPost(t *testing.T) {
	t.Log("with prepared server")
	{
		db, teardown := postgresDB(t)
		defer teardown()

		ctx, cancel := context.WithTimeout(context.Background(), caseTimeout)
		defer cancel()

		repo := postgres.NewRepository(db)

		np := post.NewPost{
			UserID: 10,
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

		t.Log("\ttest:0\tshould find a post.")
		{
			req, err := http.NewRequest("GET", fmt.Sprintf("http://%s/posts/%d", s.Addr, p.ID), nil)
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
