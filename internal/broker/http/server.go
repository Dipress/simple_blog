package http

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/dipress/blog/internal/ability"
	"github.com/dipress/blog/internal/auth"
	"github.com/dipress/blog/internal/create"
	"github.com/dipress/blog/internal/delete"
	"github.com/dipress/blog/internal/find"
	"github.com/dipress/blog/internal/list"
	"github.com/dipress/blog/internal/reg"
	"github.com/dipress/blog/internal/storage/postgres"
	"github.com/dipress/blog/internal/update"
	"github.com/dipress/blog/internal/validation"
	authEng "github.com/dipress/blog/kit/auth"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

const (
	timeout = 30 * time.Second
)

// NewServer prepares http server.
func NewServer(addr string, db *sql.DB, authenticator *authEng.Authenticator) *http.Server {
	mux := mux.NewRouter()

	repo := postgres.NewRepository(db)
	createService := create.NewService(repo, &validation.Create{})
	findService := find.NewService(repo)
	listService := list.NewService(repo)
	updateService := update.NewService(repo, &validation.Update{}, &ability.PostAbillity{})
	deleteService := delete.NewService(repo, &ability.PostAbillity{})
	registateService := reg.NewService(repo, &validation.Registrate{}, authenticator, time.Hour*24)
	authenticateService := auth.NewService(repo, authenticator, time.Hour*24)

	registrateHandler := RegHandler{
		Registrater: registateService,
	}

	authenticateHandler := AuthHandler{
		Authenticater: authenticateService,
	}

	createHandler := CreateHandler{
		Creater: createService,
	}

	findHandler := FindHandler{
		Finder: findService,
	}

	listHandler := ListHandler{
		Lister: listService,
	}

	updateHandler := UpdateHandler{
		Updater: updateService,
	}

	deleteHandler := DeleteHandler{
		Deleter: deleteService,
	}

	mux.HandleFunc("/signup", httpHandler{
		Handler: &registrateHandler,
	}.ServeHTTP).Methods("POST")

	mux.HandleFunc("/signin", httpHandler{
		Handler: &authenticateHandler,
	}.ServeHTTP).Methods("POST")

	mux.HandleFunc("/posts", AuthMiddleware(httpHandler{
		Handler: &createHandler,
	}, authenticator).ServeHTTP).Methods("POST")

	mux.HandleFunc("/posts/{id}", AuthMiddleware(httpHandler{
		Handler: &updateHandler,
	}, authenticator).ServeHTTP).Methods("PUT")

	mux.HandleFunc("/posts/{id}", AuthMiddleware(httpHandler{
		Handler: &deleteHandler,
	}, authenticator).ServeHTTP).Methods("DELETE")

	mux.HandleFunc("/posts/{id}", httpHandler{
		Handler: &findHandler,
	}.ServeHTTP).Methods("GET")

	mux.HandleFunc("/posts", httpHandler{
		Handler: &listHandler,
	}.ServeHTTP).Methods("GET")

	allowedHeaders := handlers.AllowedHeaders([]string{"Content-Type", "Authorization"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE"})
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})

	s := http.Server{
		Addr:         addr,
		Handler:      handlers.CORS(allowedHeaders, allowedOrigins, allowedMethods)(mux),
		ReadTimeout:  timeout,
		WriteTimeout: timeout,
	}

	return &s
}
