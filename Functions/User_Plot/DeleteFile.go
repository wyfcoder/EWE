package User_Plot

import (
	"github.com/WebForEME/Functions/CheckService"
	"github.com/WebForEME/Functions/DealWrongs"
	"github.com/WebForEME/sqlOperate"
	"net/http"
)

func DeleteFile(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		DealWrongs.DealWrongs(DealWrongs.ErrorCodeForSystem,&writer,request)
	}
	if !CheckService.CheckUCode(request) {
		DealWrongs.DealWrongs(DealWrongs.ErrorCodeForLoginAgain,&writer,request)
	}
	id := request.Form["id"]
	name := request.Form["name"]
	sqlOperate.DeleteData(id[0], name[0])
	http.Redirect(writer, request, "/PlotCenter", 302)
}
