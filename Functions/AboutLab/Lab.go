package AboutLab

import (
	"github.com/WebForEME/Functions"
	"github.com/WebForEME/Functions/CheckService"
	"net/http"
)

func Lab(writer http.ResponseWriter, request *http.Request){
	if(CheckService.CheckUser(writer,request)){
		t := Functions.ParseTemplateFiles("Lab/aboutLabLay","Lab/aboutLab","Tools/private.navbar")
		t.Execute(writer,nil)
	}
}
