package User_Plot

import (
	"github.com/WebForEME/AMethod/TextDeal"
	"github.com/WebForEME/Functions"
	"github.com/WebForEME/Functions/CheckService"
	"github.com/WebForEME/Functions/DealWrongs"
	"github.com/WebForEME/sqlOperate"
	"io/ioutil"
	"net/http"
)

func CheckUploadFile(writer http.ResponseWriter, request *http.Request) {

	if !CheckService.CheckUser(writer, request) { //连接超时
		DealWrongs.DealWrongs(CheckService.ErrorCodeForTimeOut,&writer,request)
	}

	err := request.ParseMultipartForm(1024 * 1024 * 120) //读取120M的数据

	if err != nil { //文件大小大于12M
		DealWrongs.DealWrongs(DealWrongs.ErrorCodeForFileSize,&writer,request)
	}

	//一次只读取单个文件
	if len(request.MultipartForm.File["file"]) !=1{
		DealWrongs.DealWrongs(DealWrongs.ErrorCodeForFileNumber,&writer,request)
	}

	//获取参数
	fileHeader := request.MultipartForm.File["file"][0]
	name := request.PostFormValue("name")
	tag := request.PostFormValue("selectType")
	describe := request.PostFormValue("describe")
	time := request.PostFormValue("date")
	account:=Functions.GetUserAccountCookieValue(request)

	//对数据的类型与数据内容直接匹配
	suffixName :=Functions.GetSuffixName(fileHeader.Filename)

	if(checkSuffixName(suffixName)){
		//把文件内容->    byte[] -> string
		file,_:=fileHeader.Open()

		byte,_:=ioutil.ReadAll(file)

		text :=string(byte)

		if(checkText(tag,&text)){
			//保存数据，并刷新
			i:=sqlOperate.AddDataFile(name,tag,time,&describe,&text,account)
			if(i!=0){ //文件重复
				DealWrongs.DealWrongs(DealWrongs.ErrorCodeForFileName,&writer,request)
			}
			http.Redirect(writer,request,"/PlotCenter",302)
		}else{
			DealWrongs.DealWrongs(DealWrongs.ErrorCodeForFileContent,&writer,request)
		}
	}else{ //文件的标签错误
		DealWrongs.DealWrongs(DealWrongs.ErrorCodeForTag,&writer,request)
	}

}

func checkText(tag string, text *string) bool {
	ok := false
	count := 0
	ok, count, (*text) = TextDeal.DealText(text)
	if !ok {
		return false
	}
	if(count==0) {
		return false
	}
	switch tag {
	case "2D Curve":
		if count%2 == 0 {
			return true
		}
	case "2D Polar":
		if count%2 == 0 {
			return true
		}
	case "2D Vector":
		if count%4==0{
			return true
		}
	}
	return false
}

func checkSuffixName(name string) bool{
	return name=="txt" || name=="tad"
}