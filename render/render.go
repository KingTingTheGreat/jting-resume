package render

import (
	"html/template"
	"io"
	"log"
)

var funcMap = template.FuncMap{
	"dec": func(i int) int {
		return i - 1
	},
}

func RenderPage(w io.Writer, data interface{}) {
	tmpl := template.New("").Funcs(funcMap)
	tmpl = template.Must(tmpl.ParseGlob("views/*.html"))

	err := tmpl.ExecuteTemplate(w, "index", data)
	if err != nil {
		log.Fatal("failed to render page", err)
	}
}