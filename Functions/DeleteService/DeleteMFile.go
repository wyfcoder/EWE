package DeleteService

import (
	"fmt"
	"github.com/WebForEME/Functions"
	"github.com/WebForEME/Functions/CheckService"
	"github.com/WebForEME/sqlOperate"
	"net/http"
	"os"
)

//检查合法的操作性
func DeleteMFile(writer http.ResponseWriter, request *http.Request){
	if(!Functions.DealContent(writer,request)) {
		return
	}
	if(CheckService.CheckManager(writer,request)){
		date :=request.Form["Date"][0]
		title :=request.Form["Name"][0]
		path :=request.Form["Path"][0]
		sqlOperate.DeleteFileM(title,date)
		http.Redirect(writer,request,"/Manager_File",302)
		err:=os.Remove(path)
		fmt.Println(err)
	}
}
