package Manager_File

import (
	"github.com/WebForEME/Functions"
	"github.com/WebForEME/Functions/CheckService"
	"github.com/WebForEME/ManagerOperator"
	"github.com/WebForEME/sqlOperate"
	"net/http"
)

func ManagerFile(writer http.ResponseWriter, request *http.Request){
	if(CheckService.CheckManager(writer,request)){
		t := Functions.ParseTemplateFiles("ManagerLayout","ManagerBar","ManagerFile")
		files:=sqlOperate.Files()
		t.Execute(writer, files)
	}
}

func UploadFileM(writer http.ResponseWriter, request *http.Request)  {
	if(CheckService.CheckManager(writer,request)){

		err := request.ParseMultipartForm(1024 * 1024 * 13*100) //读取1300M的数据

		if err != nil {
			Functions.Danger(err, "Cannot parse form.")
			sysW := Functions.SystemWrong()
			Functions.DealWrongCookie(request, writer, sysW.W, sysW.S, sysW.Wa)
			http.Redirect(writer, request, "/deal_wrong", 302)
			return
		}

		//文件的验证
		fileHeader := request.MultipartForm.File["upload"][0]

		title :=request.PostFormValue("title")

		tag   :=request.PostFormValue("tag")

		if(!ManagerOperator.CheckUploadTag(tag,fileHeader.Filename)){
			dataW :=Functions.DataWrong()
			Functions.DealWrongCookie(request, writer, dataW.W, dataW.S, "/Manager_File")
			http.Redirect(writer, request, "/deal_wrong", 302)
			return
		}
		ManagerOperator.UploadFile(fileHeader,tag,title)
		http.Redirect(writer,request,"/Manager_File",302)
	}
}
