package api

import (
	"github.com/UnplugCharger/htmx-todo/api/handlers"
	"github.com/UnplugCharger/htmx-todo/config"
	db "github.com/UnplugCharger/htmx-todo/db/sqlc"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
)

type Server struct {
	// dbStore db.Store
	router *chi.Mux
	store  db.Store
	config config.Config
}

// NewServer creates a new server
func NewServer(store db.Store, config config.Config) (*Server, error) {
	server := &Server{
		store:  store,
		config: config,
	}

	server.setupRouter()
	return server, nil
}

// setupRouter sets up the router
func (s *Server) setupRouter() {
	r := chi.NewRouter()
	r.Use(middleware.Logger) // <--<< Logger should come before Recoverer
	r.Use(middleware.Recoverer)

	r.Get("/", handlers.NewGetHomeHandler().ServeHTTP)
	r.Get("/login", handlers.NewGetLoginHandler().ServeHTTP)
	r.Post("/login", handlers.NewPostLoginHandler(s.store).ServeHTTP)

	r.Get("/register", handlers.NewGetRegisterHandler().ServeHTTP)
	r.Post("/register", handlers.NewPostRegisterHandler(s.store).ServeHTTP)

	r.Get("/list-todo", handlers.NewGetListTodoHandler(s.store).ServeHTTP)

	s.router = r
}

// Start starts the server
func (s *Server) Start() error {
	return http.ListenAndServe(s.config.HTTPServerAddress, s.router)
}
