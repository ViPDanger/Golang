// Код был взят с https://stackoverflow.com/questions/19991124/go-template-html-iteration-to-generate-table-from-struct
// Большое спасибо пользователю ANisus.

package templates

import (
	"html/template"
	"log"
	"net/http"
	"reflect"

	"github.com/ViPDanger/Golang/Internal/structures"
)

var templateFuncs = template.FuncMap{"rangeStruct": rangeStructer}

// In the template, we use rangeStruct to turn our struct values
// into a slice we can iterate over
var htmlTemplate = `<!DOCTYPE html>

<html>
	<head>
		<meta charset="utf-8">
		<title>{{.Title}}</title>
	</head>
	<body>
		
        <table border="1" cellpadding="5" cellspacing="5">
			{{range .First_Line}} <td>{{.}}</td>
			{{end}}</tr>
			{{range .Data}}<tr>
			{{range rangeStruct .}} <td>{{.}}</td>
			{{end}}</tr>
			{{end}}
			</table>


			<form action="/" method="post">
				<input type="submit" name="" value= {{.BackButton}}>
			</form>


	</body>
</html>`

func TableMaker(w http.ResponseWriter, Title string, First_Line []string, data any) {

	// We create the template and register out template function
	var err error
	t := template.New("template").Funcs(templateFuncs)
	t, err = t.Parse(htmlTemplate)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to parse files"))
		log.Println(err)
		return

	}
	page := structures.Result_Page{
		Title:      Title,
		First_Line: First_Line,
		BackButton: "Вернуться Назад",
	}
	if reflect.ValueOf(data).Kind() == reflect.Slice {
		page.Data = data
	} else {
		page.Data = append([]any{}, data)
	}
	err = t.Execute(w, page)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("ERR NO EXECUTE FILE"))
		log.Println(err)
		return
	}
}

// RangeStructer takes the first argument, which must be a struct, and
// returns the value of each field in a slice. It will return nil
// if there are no arguments or first argument is not a struct

func rangeStructer(args ...any) []interface{} {
	if len(args) == 0 {
		return nil
	}
	v := reflect.ValueOf(args)
	if v.Kind() != reflect.Struct {
		v = reflect.ValueOf(args[0])
		if v.Kind() != reflect.Struct {
			return nil
		}
	}
	out := structer(v)
	return out
}

func structer(v reflect.Value) []interface{} {
	out := make([]interface{}, 0)
	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).Kind() != reflect.Struct {
			out = append(out, v.Field(i).Interface())
		} else {
			out = append(out, structer(v.Field(i))...)
		}

	}
	return out
}
