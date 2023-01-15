package server

import (
	"net/http"
)

func (s *server) routes() {
	s.router.Handle("/static/", http.StripPrefix("/static/", s.fileServer()))

	s.router.HandleFunc("/", s.matchExactRoute("/", s.handleIndex()))
	s.router.HandleFunc("/search", s.handleSearch())
	s.router.HandleFunc("/searchAll", s.handleSearchAll())
}
