package handler

import (
	"net/http"
	"strconv"
	"strings"
	"text/template"
	"time"

	pg "github.com/ViPDanger/Golang/Internal/postgres"
	"github.com/ViPDanger/Golang/Internal/structures"
	t "github.com/ViPDanger/Golang/templates"
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
		Title:            "HTTP-Postgress API",
		SearchAllAuthor:  "Search for All Authors",
		SearchAllContent: "Search for All Content",
		SearchAuthorID:   "Search Author by ID",
		SearchContentID:  "Search Content by ID",
		InputAuthor:      "Input new Author",
		InputContent:     "Input new Content",
		DeleteAuthorID:   "Delete Author by ID",
		DeleteContentID:  "Delete Content by ID",
	}
	t.Execute(w, &page)
}

func PG_Response(w http.ResponseWriter, r *http.Request, rep *pg.Repository) {
	w.WriteHeader(http.StatusOK)
	data := r.URL.Path
	data = data[strings.IndexRune(data[1:], '/')+2:]
	switch data {
	case "AllAuthors":
		PG_AllAuthors(w, r, rep)
	case "FindAuthor":
		PG_FindAuthor(w, r, rep)
	case "InsertAuthor":
		PG_InsertAuthor(w, r, rep)
	case "DeleteAuthor":
		PG_DeleteAuthor(w, r, rep)
	case "AllContent":
		PG_AllContent(w, r, rep)
	case "FindContent":
		PG_FindContent(w, r, rep)
	case "InsertContent":
		PG_InsertContent(w, r, rep)
	case "DeleteContent":
		PG_DeleteContent(w, r, rep)
	default:
		w.Write([]byte("Thats... That's not in option right now. Sorry. :'("))
	}

}
func PG_AllAuthors(w http.ResponseWriter, r *http.Request, rep *pg.Repository) {
	w.WriteHeader(http.StatusOK)
	authors, err := rep.All_Authors()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Unable to find Authors"))
		return
	}
	t.TableMaker(w, "AllAuthors", []string{"ID", "Name"}, authors)
}

func PG_AllContent(w http.ResponseWriter, r *http.Request, rep *pg.Repository) {
	w.WriteHeader(http.StatusOK)
	content, err := rep.All_Content()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Unable to find Content"))
		return
	}
	t.TableMaker(w, "AllContent", []string{"ID", "Content", "ID Author", "Author", "Date"}, content)
}

func PG_FindAuthor(w http.ResponseWriter, r *http.Request, rep *pg.Repository) {
	w.WriteHeader(http.StatusOK)
	author, err := rep.Find_Author(r.FormValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Unable to find Author with this id"))
		return
	}
	t.TableMaker(w, "FindAuthor", []string{"ID", "Name"}, author)
}
func PG_FindContent(w http.ResponseWriter, r *http.Request, rep *pg.Repository) {
	w.WriteHeader(http.StatusOK)
	content, err := rep.Find_Content(r.FormValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Unable to find Content with this id"))
		return
	}
	t.TableMaker(w, "FindContent", []string{"ID", "Content", "ID Author", "Author", "Date"}, content)
}

func PG_InsertAuthor(w http.ResponseWriter, r *http.Request, rep *pg.Repository) {
	w.WriteHeader(http.StatusOK)
	id, err := rep.Insert_Author(structures.AuthorDTO{Name: r.FormValue("name")})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Unable to Insert Author"))
		return
	}
	t.TableMaker(w, "InsertAuthor", []string{"Inserted ID", "Inserted Name"}, structures.Author{ID: id, Name: r.FormValue("name")})

}
func PG_InsertContent(w http.ResponseWriter, r *http.Request, rep *pg.Repository) {
	w.WriteHeader(http.StatusOK)

	author_id, _ := strconv.Atoi(r.FormValue("author_id"))
	date := time.Now().Format("02.01.2006  15:04:05")
	id, err := rep.Insert_Content(structures.ContentDTO{Content: r.FormValue("content"), Author_id: author_id, Date: date})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Unable to Insert Content"))
		return
	}
	author, _ := rep.Find_Author(r.FormValue("author_id"))
	t.TableMaker(w, "InsertContent", []string{"Inserted ID", "Inserted Name"}, structures.Content{ID: id, Content: r.FormValue("content"), Author: author, Date: date})

}

func PG_DeleteAuthor(w http.ResponseWriter, r *http.Request, rep *pg.Repository) {
	w.WriteHeader(http.StatusOK)
	author, err := rep.Find_Author(r.FormValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Unable to find Author with this id"))
		return
	}
	err = rep.Delete_Author(r.FormValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Unable to find Author with this id"))
		return
	}
	t.TableMaker(w, "InsertContent", []string{"Deleted ID", "Deleted Name"}, author)

}
func PG_DeleteContent(w http.ResponseWriter, r *http.Request, rep *pg.Repository) {
	w.WriteHeader(http.StatusOK)
	content, err := rep.Find_Content(r.FormValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Unable to find Author with this id"))
		return
	}
	err = rep.Delete_Content(r.FormValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Unable to find Content with this id"))
		return
	}

	t.TableMaker(w, "InsertContent", []string{"Deleted ID", "Deleted Content", "Deleted ID_Author", "Deleted Author Name", "Deleted Date"}, content)
}
