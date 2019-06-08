package Manager_Page

import (
	"github.com/WebForEME/Functions"
	"github.com/WebForEME/Functions/CheckService"
	"github.com/WebForEME/sqlOperate"
	"net/http"
)

func ManagerPage(writer http.ResponseWriter, request *http.Request) {
	if(CheckService.CheckManager(writer,request)){
		t := Functions.ParseTemplateFiles("ManagerLayout","ManagerBar","ManagerNotice")
		notices:=sqlOperate.Notices()
		t.Execute(writer, notices)
	}
}
