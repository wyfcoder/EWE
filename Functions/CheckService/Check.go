package CheckService

import (
	"github.com/WebForEME/Functions"
	"github.com/WebForEME/sqlOperate"
	"net/http"
)

//检验用户状态的有效性
func CheckUser(writer http.ResponseWriter,request *http.Request) bool{
	if !Functions.IsCookieExit(request, "account") {
		Functions.DealWrongCookie(request, writer, "The time is out.", "Login again.", "/login")
		http.Redirect(writer, request, "/deal_wrong", 302)
	} else if !CheckCode(request){
		t :=Functions.ParseTemplateFiles("layout", "close.navbar", "closeWindow")
		t.Execute(writer, "The password has expired.")
	}else {
		return true
	}
	return false
}

//检验Manager状态的有效性
func CheckManager(writer http.ResponseWriter,request *http.Request) bool{
	if !Functions.IsCookieExit(request, "account") {
		Functions.DealWrongCookie(request, writer, "The time is out.", "Login again.", "/login")
		http.Redirect(writer, request, "/deal_wrong", 302)
	} else if !CheckMCode(request){
		t := Functions.ParseTemplateFiles("layout", "close.navbar", "closeWindow")
		t.Execute(writer, "The password has expired.")
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