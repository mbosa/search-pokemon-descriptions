package templates

import (
	"embed"
	"html/template"
)

//go:embed *
var fs embed.FS

var Templates = map[string]*template.Template{
	"search":      template.Must(template.ParseFS(fs, "layout.html", "header.html", "footer.html", "searchContent.html", "searchInput.html")),
	"results":     template.Must(template.ParseFS(fs, "layout.html", "header.html", "footer.html", "resultsContent.html", "searchInput.html")),
	"serverError": template.Must(template.ParseFS(fs, "layout.html", "header.html", "footer.html", "serverErrorContent.html")),
	"notFound":    template.Must(template.ParseFS(fs, "layout.html", "header.html", "footer.html", "notFoundContent.html", "searchInput.html")),
}

type SearchParams struct {
	Title string
}

type Result struct {
	Species_id  string
	Name        string
	Description string
	Url         string
}

type ResulsParams struct {
	Title   string
	Results []Result
}

type ServerErrorParams struct {
	Title string
}

type NotFoundParams struct {
	Title string
}
