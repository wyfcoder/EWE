package LoginService

import (
	"github.com/WebForEME/Functions"
	"github.com/WebForEME/Functions/DealWrongs"
	"github.com/WebForEME/randomOperator"
	"github.com/WebForEME/sqlOperate"
	"net/http"
)

//客户模式的登陆
//验证账户，并把用户账号写入cookie，一段时间
//写入一个新的随机码，用于后续的验证操作，保证线上只有一个用户
//设置Cookie的时间有限，保证在一段时间之后，用户自动下线
func userMode(writer http.ResponseWriter, request *http.Request) {
	//获取页面的用户名和密码
	user := request.PostFormValue("accountOrName")
	password := request.PostFormValue("password")

	//检查
	u, err2 := sqlOperate.Login(user, password)

	if err2 != nil {
		/*fmt.Println(err2)
		//使用cookie去完成信息的存储

		Functions.DealWrongCookie(request, writer, err2.Error(), solutionToLogin, wayToLogin)
		http.Redirect(writer, request, "/deal_wrong", 302)*/
		// TODO:错误处理系统的构建
		DealWrongs.DealWrongs(ERROCODEFORLOGIN,writer,request);
		return
	}

	//获取验证码
	code := randomOperator.CreateVerificationCode()
	sqlOperate.UpdateCode(u.Account, code)

	Functions.AddCookie("account", u.Account, writer, USERTIME) //有限时间的链接
	Functions.AddCookie("code", code, writer, 1000)

	http.Redirect(writer, request, "/Market", 302)
}
