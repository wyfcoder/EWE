package Register

import (
	"github.com/WebForEME/Functions"
	"github.com/WebForEME/Functions/DealWrongs"
	"github.com/WebForEME/sqlOperate"
	"net/http"
)

const (
	ERRORCODEFORREGISTER = 1;
	ERROCODEFORSYSTEM=-1;
)

//注册页面 把格式放在JS里进行检查
func Register(writer http.ResponseWriter, request *http.Request){
	t := Functions.ParseTemplateFiles("Login/login.layout", "SignUp/signUp")
	t.Execute(writer, nil)
}

//检查账号的有效性，是否已存在账号。
func SignUpAccount(writer http.ResponseWriter, request *http.Request){
	err := request.ParseForm()

	if err != nil {
		Functions.Danger(err, "Cannot parse form.")
		DealWrongs.DealWrongs(ERROCODEFORSYSTEM,writer,request)
	}

	name := request.PostFormValue("name")
	account := request.PostFormValue("account")
	findPassword := request.PostFormValue("findPassword")
	password := request.PostFormValue("password")

	err, message:= sqlOperate.AddUsers(name, account, password, findPassword)

	//处理重复错误
	if err != nil {
		Functions.MakePreWrongCookie(message,writer)
		DealWrongs.DealWrongs(ERRORCODEFORREGISTER,writer,request) //注册错误
		return
	}


	http.Redirect(writer, request, "/successfulSignUp?name="+name, 302)
}

func SuccessfulSignUp(writer http.ResponseWriter, request *http.Request){

	err := request.ParseForm()

	if err != nil {
		Functions.Danger(err, "Cannot parse form.")
		DealWrongs.DealWrongs(ERROCODEFORSYSTEM,writer,request)
	}

	name := request.Form["name"][0]

	t := Functions.ParseTemplateFiles("Tools/layout", "Tools/public.navbar", "Success/Success")

	t.Execute(writer, "Dear "+name+", your account is active.")

}
