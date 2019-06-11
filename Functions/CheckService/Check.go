package CheckService

import (
	"github.com/WebForEME/Functions"
	"github.com/WebForEME/Functions/DealWrongs"
	"github.com/WebForEME/sqlOperate"
	"net/http"
)

//连接超时
const ErrorCodeForTimeOut  = 3;

//被连续登陆
const ErrorCodeForLoginAgain  =4;

//检验用户状态的有效性
func CheckUser(writer http.ResponseWriter,request *http.Request) bool{
	if !Functions.IsCookieExit(request, "account") {
		DealWrongs.DealWrongs(ErrorCodeForTimeOut,&writer,request)
	} else if !CheckCode(request){
		DealWrongs.DealWrongs(ErrorCodeForLoginAgain,&writer,request)
	}else {
		return true
	}
	return false
}

//检验Manager状态的有效性
func CheckManager(writer http.ResponseWriter,request *http.Request) bool{
	if !Functions.IsCookieExit(request, "account") {
		DealWrongs.DealWrongs(ErrorCodeForTimeOut,&writer,request)
	} else if !CheckMCode(request){
		DealWrongs.DealWrongs(ErrorCodeForLoginAgain,&writer,request)
	}else {
		return true
	}
	return false
}



func CheckCode(req *http.Request) bool {
	code :=Functions.GetCookieValue("code", req)
	account := Functions.GetCookieValue("account", req)
	return sqlOperate.CheckCode(account, code)
}

func CheckMCode(req *http.Request) bool {
	code := Functions.GetCookieValue("code", req)
	account := Functions.GetCookieValue("account", req)
	return sqlOperate.CheckMCode(account, code)
}