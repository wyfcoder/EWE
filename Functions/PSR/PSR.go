package PSR

import (
	"github.com/WebForEME/Functions"
	"net/http"
)

func PSR(writer http.ResponseWriter, request *http.Request) {

	t := Functions.ParseTemplateFiles("PSR/PSRLayout","PSR/PSRToolBar","PSR/PSR")
	t.Execute(writer, nil)
}

