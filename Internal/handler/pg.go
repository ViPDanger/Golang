package handler

import (
	"net/http"
	"strconv"
	"strings"
	"text/template"
	"time"

	"github.com/ViPDanger/Golang/Internal/config"
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
	authors, _ := rep.All_Authors()
	t.TableMaker(w, authors)
}

func PG_AllContent(w http.ResponseWriter, r *http.Request, rep *pg.Repository) {
	w.WriteHeader(http.StatusOK)
	content, _ := rep.All_Content()

	var table structures.Table

	table.Title = ""
	table.First_Line = append(table.First_Line, "ID", "Content", "Author_ID", "Author_Name")

	t.TableMaker(w, content)
}

func PG_FindAuthor(w http.ResponseWriter, r *http.Request, rep *pg.Repository) {
	w.WriteHeader(http.StatusOK)
	id, _ := strconv.Atoi(r.FormValue("id"))
	author, err := rep.Find_Author(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Unable to find Author with this id"))
		return
	}
	t, err := template.ParseFiles("./templates/result.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to parse files"))
		return

	}
	page := structures.Result_Page{
		Title:  "HTTP-PGSQL FIND AUTHOR",
		Result: "ID: " + strconv.Itoa(author.ID) + "; Name: " + author.Name,
	}
	t.Execute(w, &page)
}
func PG_FindContent(w http.ResponseWriter, r *http.Request, rep *pg.Repository) {
	w.WriteHeader(http.StatusOK)
	id, _ := strconv.Atoi(r.FormValue("id"))
	content, err := rep.Find_Content(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Unable to find Content with this id"))
		return
	}
	t, err := template.ParseFiles("./templates/result.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to parse files"))
		return

	}
	page := structures.Result_Page{
		Title:  "HTTP-PGSQL FIND CONTENT",
		Result: "ID: " + strconv.Itoa(content.ID) + "\nName: " + content.Content + "\n Author:" + strconv.Itoa(content.Author.ID) + ": " + content.Author.Name + "\n Date:" + content.Date,
	}
	t.Execute(w, &page)
}

func PG_InsertAuthor(w http.ResponseWriter, r *http.Request, rep *pg.Repository) {
	w.WriteHeader(http.StatusOK)
	id, err := rep.Insert_Author(structures.AuthorDTO{Name: r.FormValue("name")})
	config.Err_log(err)
	t, err := template.ParseFiles("./templates/result.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to parse files"))
		return

	}
	page := structures.Result_Page{
		Title:  "HTTP-PGSQL INSERT AUTHOR",
		Result: "New ID of Author " + r.FormValue("name") + ": " + strconv.Itoa(id),
	}
	t.Execute(w, &page)
}
func PG_InsertContent(w http.ResponseWriter, r *http.Request, rep *pg.Repository) {
	w.WriteHeader(http.StatusOK)
	author_id, _ := strconv.Atoi(r.FormValue("author_id"))
	date := time.Now().Format("02.01.2006  15:04:05")
	id, err := rep.Insert_Content(structures.ContentDTO{Content: r.FormValue("content"), Author_id: author_id, Date: date})
	config.Err_log(err)
	t, err := template.ParseFiles("./templates/result.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to parse files"))
		return
	}
	page := structures.Result_Page{
		Title:  "HTTP-PGSQL INSERT AUTHOR",
		Result: "New ID of Content: " + r.FormValue("content") + ": " + strconv.Itoa(id),
	}
	t.Execute(w, &page)
}

func PG_DeleteAuthor(w http.ResponseWriter, r *http.Request, rep *pg.Repository) {
	w.WriteHeader(http.StatusOK)
	err := rep.Delete_Author(r.FormValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Unable to find Author with this id"))
		return
	}

	t, err := template.ParseFiles("./templates/result.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to parse files"))
		return
	}
	page := structures.Result_Page{
		Title:  "HTTP-PGSQL DELETE AUTHOR",
		Result: "Author ID " + r.FormValue("id") + " deleted succsessfully",
	}
	t.Execute(w, &page)
}
func PG_DeleteContent(w http.ResponseWriter, r *http.Request, rep *pg.Repository) {
	w.WriteHeader(http.StatusOK)
	err := rep.Delete_Content(r.FormValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Unable to find Content with this id"))
		return
	}

	t, err := template.ParseFiles("./templates/result.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to parse files"))
		return
	}
	page := structures.Result_Page{
		Title:  "HTTP-PGSQL DELETE CONTENT",
		Result: "Content ID " + r.FormValue("id") + " deleted succsessfully",
	}
	t.Execute(w, &page)
}
