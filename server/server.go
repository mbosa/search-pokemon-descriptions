package server

import (
	"net/http"

	"github.com/mbosa/search-pokemon-descriptions/db"

	"github.com/jackc/pgx/v5/pgxpool"
)

type server struct {
	router *http.ServeMux
	db     *pgxpool.Pool
	http.Handler
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) Close() {
	s.db.Close()
}

func NewServer() *server {
	server := &server{
		router: http.NewServeMux(),
		db:     db.NewDbPool(),
	}

	server.routes()

	return server
}
