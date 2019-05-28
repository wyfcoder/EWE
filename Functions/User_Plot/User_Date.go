package User_Plot

import (
	"github.com/WebForEME/Functions"
	"github.com/WebForEME/Functions/CheckService"
	"github.com/WebForEME/sqlOperate"
	"net/http"
)

func Data(writer http.ResponseWriter,request *http.Request){
	if(CheckService.CheckUser(writer,request)){
		t := Functions.ParseTemplateFiles("Tools/layout", "Tools/private.navbar", "Plot/plot")
		datas, _ := sqlOperate.Datas(Functions.GetCookieValue("account", request))
		t.Execute(writer, datas)
	}
}

func DeleteFile(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		Functions.Danger(err, "Cannot parse form.")
		sysW := Functions.SystemWrong()
		Functions.DealWrongCookie(request, writer, sysW.W, sysW.S, sysW.Wa)
		http.Redirect(writer, request, "/deal_wrong", 302)
	}
	if !CheckService.CheckCode(request) {
		t := Functions.ParseTemplateFiles("layout", "close.navbar", "closeWindow")
		t.Execute(writer, "The password has expired.")
		return
	}
	id := request.Form["id"]
	name := request.Form["name"]
	sqlOperate.DeleteData(id[0], name[0])
	http.Redirect(writer, request, "/Data", 302)
}
