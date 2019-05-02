package AboutLab

import (
	"github.com/WebForEME/Functions"
	"github.com/WebForEME/Functions/CheckService"
	"net/http"
)

func Lab(writer http.ResponseWriter, request *http.Request){
	if(CheckService.CheckUser(writer,request)){
		t := Functions.ParseTemplateFiles("aboutLabLay","aboutLab","private.navbar")
		t.Execute(writer,nil)
	}
}
