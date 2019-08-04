package User_Plot

import (
	"fmt"
	"github.com/WebForEME/AMethod/SystemTool"
	"github.com/WebForEME/Functions/CheckService"
	"github.com/WebForEME/Functions/DealWrongs"
	"github.com/WebForEME/sqlOperate"
	"net/http"
	"os"
)

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
	//转化为一个文件，再上传然后删除这个文件
	sqlOperate.DownloadData(id[0],name[0])
	path :="UserTempFile"
	if SystemTool.OsSystem() == 1{
		path += "\\"
	}else{
		path += "/"
	}
	path += name[0]+".txt"
	writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", name[0]+".txt"))
	writer.Header().Add("Content-Type", "application/octet-stream")
	http.ServeFile(writer, request, path)
	os.Remove(path)
}
