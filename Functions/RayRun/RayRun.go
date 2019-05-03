package RayRun

import (
	"github.com/WebForEME/makeTemplate"
	"html/template"
	"net/http"
)

func RayRun(writer http.ResponseWriter, request *http.Request) {
	t := template.New("tmpl.html")
	t, _ = t.Parse(makeTemplate.MakeRayRun())
	t.Execute(writer, nil)
}
