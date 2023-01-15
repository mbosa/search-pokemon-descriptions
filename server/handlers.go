package server

import (
	"context"
	"net/http"
	"unicode"

	"github.com/mbosa/search-pokemon-descriptions/db/queries"
	"github.com/mbosa/search-pokemon-descriptions/templates"
)

func (s *server) fileServer() http.Handler {
	return http.FileServer(http.Dir("./static"))
}

func (s *server) handleIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := templates.SearchParams{
			Title: "Layout",
		}
		templates.Templates["search"].Execute(w, params)
	}
}

func (s *server) handleSearch() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		query := r.FormValue("q")
		if query == "" {
			params := templates.SearchParams{
				Title: "Layout",
			}
			templates.Templates["search"].Execute(w, params)
			return
		}

		rows, err := s.db.Query(context.Background(), queries.Search, query)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)

			params := templates.ServerErrorParams{
				Title: "Error",
			}

			templates.Templates["serverError"].Execute(w, params)
			return
		}
		var results []templates.Result
		for rows.Next() {
			var (
				species_id  string
				name        string
				description string
				url         string
			)
			rows.Scan(&species_id, &name, &description, &url)

			results = append(results, templates.Result{
				Species_id:  species_id,
				Name:        capitalize(name),
				Description: description,
				Url:         url,
			})
		}

		if len(results) == 0 {
			w.WriteHeader(http.StatusNotFound)

			params := templates.NotFoundParams{
				Title: "Test",
			}

			templates.Templates["notFound"].Execute(w, params)

			return
		}

		params := templates.ResulsParams{
			Title:   "Test",
			Results: results,
		}

		templates.Templates["results"].Execute(w, params)
	}
}

func (s *server) handleSearchAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		rows, err := s.db.Query(context.Background(), queries.SearchAll)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)

			params := templates.ServerErrorParams{
				Title: "Error",
			}

			templates.Templates["serverError"].Execute(w, params)
			return
		}

		var results []templates.Result
		for rows.Next() {
			var (
				species_id  string
				name        string
				description string
				url         string
			)
			rows.Scan(&species_id, &name, &description, &url)

			results = append(results, templates.Result{
				Species_id:  species_id,
				Name:        capitalize(name),
				Description: description,
				Url:         url,
			})
		}

		params := templates.ResulsParams{
			Title:   "Test",
			Results: results,
		}

		templates.Templates["results"].Execute(w, params)
	}
}

func capitalize(s string) string {
	r := []rune(s)
	r[0] = unicode.ToUpper(r[0])

	return string(r)
}
