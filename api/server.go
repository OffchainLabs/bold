// Package api defines an API server for BOLD, allowing retrieval of information
// from both the assertion chain and the challenge manager contracts in order to
// understand ongoing challenges.
package api

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var (
	ErrNoConfig                 = errors.New("no config provided")
	ErrNoEdgesProvider          = errors.New("no edges provider")
	ErrNoAssertionsProvider     = errors.New("no assertions provider")
	ErrAlreadyRegisteredMethods = errors.New("already registered methods")
)

type Config struct {
	Address            string
	EdgesProvider      EdgesProvider
	AssertionsProvider AssertionsProvider
	DBConfig           *DBConfig
}

type Server struct {
	srv *http.Server

	edges      EdgesProvider
	assertions AssertionsProvider
	database   *Database

	router *mux.Router

	registered bool
}

func NewServer(cfg *Config) (*Server, error) {
	if cfg == nil {
		return nil, ErrNoConfig
	}

	if cfg.Address == "" {
		cfg.Address = ":8080"
	}

	if cfg.EdgesProvider == nil {
		return nil, ErrNoEdgesProvider
	}
	if cfg.AssertionsProvider == nil {
		return nil, ErrNoAssertionsProvider
	}

	r := mux.NewRouter()

	s := &Server{
		srv: &http.Server{
			Handler:           r,
			Addr:              cfg.Address,
			WriteTimeout:      15 * time.Second,
			ReadTimeout:       15 * time.Second,
			ReadHeaderTimeout: 15 * time.Second,
		},
		edges:      cfg.EdgesProvider,
		assertions: cfg.AssertionsProvider,
		router:     r,
	}
	if cfg.DBConfig != nil && cfg.DBConfig.Enable {
		database, err := NewDatabase(cfg.DBConfig, cfg.EdgesProvider)
		if err != nil {
			return nil, err
		}
		s.database = database
	}

	if err := s.registerMethods(); err != nil {
		return nil, err
	}

	return s, nil
}

func (s *Server) Start(ctx context.Context) error {
	go s.database.Start(ctx)
	return s.srv.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	s.database.close()
	return s.srv.Shutdown(ctx)
}

func (s *Server) registerMethods() error {
	if s.registered {
		return ErrAlreadyRegisteredMethods
	}

	s.router.HandleFunc("/healthz", healthzHandler).Methods("GET")

	// Assertions
	s.router.HandleFunc("/assertions", s.listAssertionsHandler).Methods("GET")
	s.router.HandleFunc("/assertions/{id}", s.getAssertionHandler).Methods("GET")

	// Edges
	s.router.HandleFunc("/honest-edges", s.listHonestEdgesHandler).Methods("GET")
	s.router.HandleFunc("/edges", s.listEdgesHandler).Methods("GET")
	s.router.HandleFunc("/edges/{id}", s.getEdgeHandler).Methods("GET")
	s.router.HandleFunc("/honest-confirmable-edges", s.listHonestConfirmableEdgesHandler).Methods("GET")
	s.router.HandleFunc("/evil-confirmed-edges", s.listEvilConfirmedEdgesHandler).Methods("GET")

	// Stakes
	s.router.HandleFunc("/mini-stakes", s.listMiniStakesHandler).Methods("GET")

	// Database query
	if s.database != nil {
		s.router.HandleFunc("/query-database/{query}", s.queryDatabaseHandler).Methods("GET")
		s.router.HandleFunc("/update-database", s.updateDatabaseHandler).Methods("PUT")
	}
	s.registered = true
	return nil
}

func writeJSONResponse(w http.ResponseWriter, code int, data any) error {
	body, err := json.Marshal(data)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, err = w.Write(body)
	return err
}

func writeError(w http.ResponseWriter, code int, err error) {
	w.WriteHeader(code)
	if _, err2 := w.Write([]byte(err.Error())); err != nil {
		log.Error("could not write response body", "err", err2, "status", code)
	}
}
