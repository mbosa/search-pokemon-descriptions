package server

import "net/http"

func (s *server) matchExactRoute(path string, handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != path {
			http.NotFound(w, r)
			return
		}
		handler(w, r)
	}
}
