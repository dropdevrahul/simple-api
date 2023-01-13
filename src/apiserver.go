package apiserver

import (
	"log"
	"net/http"

	"github.com/dgraph-io/badger"
	"github.com/go-chi/chi"
)

type Server struct {
	Injectables map[string]Injectable
	Db          Database
	Settings    ServerSettings
}

type ServerSettings struct {
	Port string // of the form :8080
}

type Injectable interface {
}

type Database interface {
}

type DbBadger struct {
	badger *badger.DB
}

func (s *Server) LoadDB() error {
	db, err := badger.Open(badger.DefaultOptions("/tmp/badger"))
	if err != nil {
		log.Fatal(err)
	}

	s.Db = DbBadger{
		badger: db,
	}

	return nil
}

func (s *Server) Inject(key string, injectable Injectable) error {
	s.Injectables[key] = injectable
	return nil
}

func (s *Server) Serve(r *chi.Mux) error {
	return http.ListenAndServe(s.Settings.Port, r)
}
