package ContactUs

import (
	"github.com/WebForEME/Functions"
	"github.com/WebForEME/Functions/CheckService"
	"net/http"
)

//包括用户反馈信息，技术支内容

func ContactUs(writer http.ResponseWriter, request *http.Request){
	//检查用户有效性
	//加载页面
	if(CheckService.CheckUser(writer,request)){
		t := Functions.ParseTemplateFiles("ContactUsLay", "private.navbar","ContactUs")
		t.Execute(writer,nil)
	}

}