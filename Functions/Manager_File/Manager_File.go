package Manager_File

import (
	"github.com/WebForEME/Functions"
	"github.com/WebForEME/Functions/CheckService"
	"github.com/WebForEME/Functions/DealWrongs"
	"github.com/WebForEME/ManagerOperator"
	"github.com/WebForEME/sqlOperate"
	"net/http"
)

const ErrorCodeForSystem    = -1 //系统错误码
const ErrorCodeForTag       =  5 // 文件后缀与标签不符
const MaxMemory = 1024*1024*1024*100 //100G的文件

func ManagerFile(writer http.ResponseWriter, request *http.Request){
	if(CheckService.CheckManager(writer,request)){
		t := Functions.ParseTemplateFiles("ManagerPage/ManagerLayout","ManagerPage/ManagerBar","ManagerFile")
		files:=sqlOperate.Files()
		t.Execute(writer, files)
	}
}


//有待修改
func UploadFileM(writer http.ResponseWriter, request *http.Request)  {
	if(CheckService.CheckManager(writer,request)){

		//上限为100G的文件
		err := request.ParseMultipartForm(MaxMemory)

		if err != nil {
			DealWrongs.DealWrongs(ErrorCodeForSystem,&writer,request)
		}

		//文件的验证
		fileHeader := request.MultipartForm.File["upload"][0]

		title :=request.PostFormValue("title")

		tag   :=request.PostFormValue("tag")

		if(!ManagerOperator.CheckUploadTag(tag,fileHeader.Filename)){
			DealWrongs.DealWrongs(ErrorCodeForTag,&writer,request)
		}

		ManagerOperator.UploadFile(fileHeader,tag,title)

		http.Redirect(writer,request,"/Manager_File",302)
	}
}
