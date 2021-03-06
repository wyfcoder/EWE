package User_Plot

import (
	"fmt"
	"github.com/WebForEME/AMethod/OSFileTool"
	"github.com/WebForEME/Functions/CheckService"
	"github.com/WebForEME/Functions/DealWrongs"
	"github.com/WebForEME/sqlOperate"
	"net/http"
)

const FileHeader  = "UserTempFile"
const FileBoy  = "plotData"

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
	fileTail := id[0] + "_" +name[0]
	//返回数据即可
	data :=sqlOperate.DownloadData(id[0] , name[0])

	path := []string{}
	path= append(path, FileHeader)
	path= append(path,FileBoy)
	path= append(path,fileTail)

	err,pathString :=OSFileTool.CreateFile(path,&data)

	writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", name[0]+".txt"))
	writer.Header().Add("Content-Type", "application/octet-stream")
	http.ServeFile(writer, request, pathString)

	OSFileTool.DeleteFile(pathString)
}
