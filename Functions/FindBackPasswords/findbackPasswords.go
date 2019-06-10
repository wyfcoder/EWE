package FindBackPasswords

import (
	"github.com/WebForEME/Functions"
	"github.com/WebForEME/Functions/DealWrongs"
	"github.com/WebForEME/randomOperator"
	"github.com/WebForEME/sqlOperate"
	"net/http"
)

const (
	ERROCODEFORSYSTEM = -1 //系统错误码
	ERRPCDEFORFORGET  = 2  //找回密码错误
)

//忘记密码界面
func Forget(writer http.ResponseWriter, request *http.Request){
	t :=Functions.ParseTemplateFiles("Login/login.layout", "Forget/forget")
	t.Execute(writer, nil)
}

//获取输入查找用户信息
func DealForget(writer http.ResponseWriter, request *http.Request){

	err := request.ParseForm()

	if err != nil {
		Functions.Danger(err, "Cannot parse form.")
		DealWrongs.DealWrongs(ERROCODEFORSYSTEM,&writer,request)
	}

	user := request.PostFormValue("accountOrName")
	findPassword := request.PostFormValue("findPassword")

	//解析输入
	u, err2 := sqlOperate.FindBackPassword(user, findPassword)

	if err2 != nil {
		DealWrongs.DealWrongs(ERRPCDEFORFORGET,&writer,request)
		return
	}

	//添加一个Cookie的内容，保存的时间有限。
	Functions.MakeUserInformationCookie(writer,u)

	//成果修改
	http.Redirect(writer, request, "/show_information", 302)
}

//显示用户信息
func ShowInformation(writer http.ResponseWriter, request *http.Request)  {
	//从cookie中获取对应的用户的信息
	user :=Functions.GetUserCookie(request)
	Functions.GenerateHTML(writer, &user, "Login/login.layout","Forget/showInformationOfUsers")
}

//修改用户密码
func UpdatePassword(writer http.ResponseWriter, request *http.Request){

	err := request.ParseForm()

	if err != nil {
		Functions.Danger(err, "Cannot parse form.")
		DealWrongs.DealWrongs(ERROCODEFORSYSTEM,&writer,request)
	}

	account := Functions.GetCookieValue("account", request)

	password := request.PostFormValue("password")

	sqlOperate.ResetPassword(account, password) //插入数据

	sqlOperate.UpdateCode(account, randomOperator.CreateVerificationCode()) //更新验证码

	Functions.DeleteUserCookie(writer)


	t := Functions.ParseTemplateFiles("Tools/layout", "Tools/public.navbar", "Success/Success")

	t.Execute(writer, "Change password successfully.")
}