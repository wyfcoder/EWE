package PSR

import (
	"fmt"
	"github.com/WebForEME/AMethod/OSFileTool"
	"github.com/WebForEME/Functions/CheckService"
	"github.com/WebForEME/Functions/DealWrongs"
	PSRDB "github.com/WebForEME/sqlOperate/programDB/PSR"
	"net/http"
)

const FileHead  = "UserTempFile"
const FileBody  = "psrData"


func DownloadFile(writer http.ResponseWriter, request *http.Request){
	err := request.ParseForm()
	if err != nil {
		DealWrongs.DealWrongs(DealWrongs.ErrorCodeForSystem,&writer,request)
	}
	if !CheckService.CheckUCode(request) {
		DealWrongs.DealWrongs(DealWrongs.ErrorCodeForLoginAgain,&writer,request)
	}
	id := request.Form["id"]
	name := request.Form["name"]
	date :=request.Form["date"]

	data :=PSRDB.DownloadFile(id[0],date[0],name[0])

	path := []string{}
	path = append(path,FileHead)
	path = append(path,FileBody)
	path = append(path,id[0]+"_"+date[0]+"_"+name[0])

	_,pathString :=OSFileTool.CreateFile(path, &data)

	writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", name[0]+".txt"))
	writer.Header().Add("Content-Type", "application/octet-stream")
	http.ServeFile(writer, request, pathString)
	OSFileTool.DeleteFile(pathString)
}
