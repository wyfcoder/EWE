package EProgram

import (
	"github.com/WebForEME/Functions"
	"github.com/WebForEME/Functions/CheckService"
	"net/http"
)

//应用模块入口
func EProgram(writer http.ResponseWriter, request *http.Request){
	if(CheckService.CheckUser(writer,request)){
		t := Functions.ParseTemplateFiles("Programs/programsLayout", "Tools/private.navbar", "Programs/programs")
		t.Execute(writer,nil)
	}
}
