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


func RayRunHelpScript(writer http.ResponseWriter, request *http.Request){
	t,_:=template.ParseFiles("template/RayRun/DocumentsScript.html")
	t.Execute(writer,nil)
}

func RayRunHelpIRI(writer http.ResponseWriter, request *http.Request){
	t,_:=template.ParseFiles("template/RayRun/DocumentsIRI.html")
	t.Execute(writer,nil)
}

func RayRunHelpPRO(writer http.ResponseWriter, request *http.Request){
	t,_:=template.ParseFiles("template/RayRun/DocumentsPRO.html")
	t.Execute(writer,nil)
}

func RayRunHelpAH0(writer http.ResponseWriter, request *http.Request){
	t,_:=template.ParseFiles("template/RayRun/DocumentsAH0.html")
	t.Execute(writer,nil)
}