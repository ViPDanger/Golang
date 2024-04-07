package Handler

import (
	"net/http"
	"text/template"

	pg "github.com/ViPDanger/Golang/Internal/postgres"
	"github.com/ViPDanger/Golang/Internal/structures"
)

func Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html")

	t, err := template.ParseFiles("./templates/get_request.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to parse files"))
		return

	}
	page := structures.Get_Page{
		Title:            "HTTP - Postgress API",
		SearchAllAuthor:  "Search for All Authors",
		SearchAllContent: "Search for All Content",
		SearchAutrorID:   "Search Author by ID",
		SearchContentID:  "Search Content by ID",
		InputAuthor:      "Input new Author",
		InputContent:     "Input new Content",
		DeleteAuthorID:   "Delete Author by ID",
		DeleteContentID:  "Delete Content by ID",
	}
	t.Execute(w, &page)
	w.WriteHeader(http.StatusOK)
}

func PG_Response(w http.ResponseWriter, r *http.Request, rep *pg.Repository) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("RESPOOONSE!!!"))
}
