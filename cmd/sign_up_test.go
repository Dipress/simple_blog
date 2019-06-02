package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"
	"testing"
)

func TestSignUp(t *testing.T) {
	t.Log("with prepared server")
	{
		db, teardown := postgresDB(t)
		defer teardown()

		lis, err := net.Listen("tcp", "127.0.0.1:0")
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		s := setupServer(lis.Addr().String(), db, authenticator)
		go s.Serve(lis)
		defer s.Close()

		t.Log("\ttest:0\tshould registrate a new user.")
		{
			regStr := `{"username": "username", "email": "username@example.com", "password": "password123"}`
			req, err := http.NewRequest("POST", fmt.Sprintf("http://%s/signup", s.Addr), strings.NewReader(regStr))
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
