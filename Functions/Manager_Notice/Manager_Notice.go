package Manager_Notice

import (
	"github.com/WebForEME/Functions"
	"github.com/WebForEME/Functions/CheckService"
	"github.com/WebForEME/sqlOperate"
	"net/http"
)

func AddNotice(writer http.ResponseWriter, request *http.Request){
	if(!Functions.DealContent(writer,request)) {
		return
	}
	if(CheckService.CheckManager(writer,request)){
		name := request.PostFormValue("name")
		content := request.PostFormValue("content")
		sqlOperate.AddNotice(name,content)
		http.Redirect(writer,request,"/Manager_Notice",302)
	}
}


func DeleteNotice(writer http.ResponseWriter, request *http.Request){
	if(!Functions.DealContent(writer,request)) {
		return
	}
	if(CheckService.CheckManager(writer,request)){
		name := request.Form["Name"][0]
		date := request.Form["Date"][0]
		sqlOperate.DeleteNotice(name,date)
		http.Redirect(writer,request,"/Manager_Notice",302)
	}
}
