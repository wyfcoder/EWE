package User_File

import (
	"github.com/WebForEME/Functions"
	"github.com/WebForEME/Functions/CheckService"
	"github.com/WebForEME/sqlOperate"
	"net/http"
)

func Download(writer http.ResponseWriter, request *http.Request){
	if(CheckService.CheckUser(writer,request)){
		t := Functions.ParseTemplateFiles("layout", "private.navbar", "download")
		files:=sqlOperate.Files()
		t.Execute(writer, files)
	}
}
