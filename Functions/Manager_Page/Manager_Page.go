package Manager_Page

import (
	"github.com/WebForEME/Functions"
	"github.com/WebForEME/Functions/CheckService"
	"github.com/WebForEME/sqlOperate"
	"net/http"
)

//Manager的首界面
func ManagerPage(writer http.ResponseWriter, request *http.Request) {
	if(CheckService.CheckManager(writer,request)){
		t := Functions.ParseTemplateFiles("ManagerPage/ManagerLayout","ManagerPage/ManagerBar","ManagerPage/ManagerNotice")
		notices:=sqlOperate.Notices()
		t.Execute(writer, notices)
	}
}
