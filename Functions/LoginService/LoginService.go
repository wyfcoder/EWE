package LoginService

import (
	"github.com/WebForEME/Functions"
	"github.com/WebForEME/Functions/DealWrongs"
	"net/http"
)

const
(
	USERTIME   = 10000 //用户在线上的最长时间
	MANGERTIME = 10000 //管理者在线上的最长时间
	ERROCODEFORLOGIN = 0 //登陆错误码
	ERROCODEFORSYSTEM = -1 //系统错误码
)


func Login(writer http.ResponseWriter, request *http.Request){
	t := Functions.ParseTemplateFiles("Login/login.layout", "Login/login")
	t.Execute(writer, nil)
}



func DealLogin(writer http.ResponseWriter, request *http.Request){
	err := request.ParseForm()

	if err != nil {
		Functions.Danger(err, "Cannot parse form.")
		DealWrongs.DealWrongs(ERROCODEFORSYSTEM,writer,request)
	}

	people := request.PostFormValue("people")

	if people == "User" {
		userMode(writer, request)
	} else if people == "Manager" {
		managerMode(writer, request)
	}

}