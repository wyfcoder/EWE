package Manager_Feedback

import (
	"github.com/WebForEME/Functions"
	"github.com/WebForEME/Functions/CheckService"
	"github.com/WebForEME/sqlOperate"
	"net/http"
)

//刷新主页
func ManagerFeedback(writer http.ResponseWriter, request *http.Request){
	if(CheckService.CheckManager(writer,request)){
		t := Functions.ParseTemplateFiles("ManagerPage/ManagerLayout","ManagerPage/ManagerBar","ManagerPage/ManagerFeedback")
		files:=sqlOperate.Feedback()
		t.Execute(writer, files)
	}
}



func DeleteFeedback(writer http.ResponseWriter, request *http.Request){
	if(!Functions.DealContent(writer,request)) {
		return
	}
	if(CheckService.CheckManager(writer,request)){
		name := request.Form["Account"][0]
		time := request.Form["Time"][0]
		sqlOperate.DeleteFeedback(name,time)
		http.Redirect(writer,request,"/Manager_Feedback",302)
	}
}

