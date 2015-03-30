package view
import (
    "html/template"
    "net/http"
)

var (
    tmpls *template.Template
)

func init(){
    tmpls = template.Must(template.ParseGlob("./template/*.html"))
}

func Exec(w http.ResponseWriter, name string, data interface{}) {
    tmpls.ExecuteTemplate(w, name, data)
}