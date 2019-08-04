package PSR

import (
	"github.com/WebForEME/AMethod/CommandsTool"
	"github.com/WebForEME/AMethod/TextDeal"
	"github.com/WebForEME/Functions"
	"github.com/WebForEME/Functions/DealWrongs"
	"net/http"
	"strings"
)

func PSR(writer http.ResponseWriter, request *http.Request) {
	t := Functions.ParseTemplateFiles("PSR/PSRLayout","PSR/PSRToolBar","PSR/PSR")
	t.Execute(writer, nil)
}


//处理查询
func PSRDeal(writer http.ResponseWriter, request *http.Request){
	err := request.ParseForm()
	if err != nil {
		Functions.Danger(err, "Cannot parse form.")
		DealWrongs.DealWrongs(DealWrongs.ErrorCodeForSystem,&writer,request)
	}
	codes := request.PostFormValue("codes")
	commands :=strings.FieldsFunc(codes,TextDeal.SpiltSemicolon)
	isOk,index,searchCommands := CommandsTool.SearchCommands(commands)
	if isOk {
		//判断数据源是否存在 TODO:数据库实现
		PSRShowResult(codes,searchCommands,writer,request)
	}else{
		errorInformation := "The \"" + commands[index] +"\" is wrong"
		PSRShowError(codes,errorInformation,writer,request)
	}

}

func PSRShowResult(commandString string,commands []CommandsTool.SearchCommand,writer http.ResponseWriter, request *http.Request){

}

func PSRShowError(commandString string,error string,writer http.ResponseWriter, request *http.Request){
	infromation :=Information{}
	infromation.error =error
	infromation.commands = commandString
	t := Functions.ParseTemplateFiles("PSR/PSRLayout","PSR/PSRToolBar","PSR/PSR")
	t.Execute(writer,infromation)
}

