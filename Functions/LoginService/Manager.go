package LoginService

import (
	"github.com/WebForEME/AMethod/MathTool"
	"github.com/WebForEME/Functions"
	"github.com/WebForEME/Functions/DealWrongs"
	"github.com/WebForEME/sqlOperate"
	"net/http"
)

//管理者模式的登陆
func managerMode(writer http.ResponseWriter, request *http.Request) {

	user := request.PostFormValue("accountOrName")
	password := request.PostFormValue("password")

	u, err2 := sqlOperate.ManagerLogin(user, password)

	if err2 != nil {
		DealWrongs.DealWrongs(ERROCODEFORLOGIN,&writer,request);
		return
	}

	code := MathTool.CreateVerificationCode()

	sqlOperate.UpdateCode2(u.Account, code)

	Functions.MakeManagerCookie(code,u.Account,&writer)

	http.Redirect(writer, request, "/Manager_Notice", 302) //跳转管理界面
}
