package main

import (
	"fmt"
	"github.com/WebForEME/AMethod/MathTool"
	"github.com/WebForEME/makeTemplate"
	"github.com/WebForEME/sqlOperate"
	"html/template"
	"net/http"
)

//解析HTML文件函数
func parseTemplateFiles(filenames ...string) (t *template.Template) {
	var files []string
	t = template.New("layout")
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("template/%s.html", file))
	}
	t = template.Must(t.ParseFiles(files...))
	return
}

//登陆逻辑
func logIn(w http.ResponseWriter, r *http.Request) {
	t := parseTemplateFiles("login.layout", "public.navbar", "login")
	t.Execute(w, nil)
}

func dealLogin(writer http.ResponseWriter, request *http.Request) {

	err := request.ParseForm()
	if err != nil {
		danger(err, "Cannot parse form.")
		sysW := SystemWrong()
		DealWrongCookie(request, writer, sysW.W, sysW.S, sysW.Wa)
		http.Redirect(writer, request, "/deal_wrong", 302)
	}

	people := request.PostFormValue("people")

	if people == "User" {
		userMode(writer, request)
	} else if people == "Manager" {
		managerMode(writer, request)
	}

}

func userMode(writer http.ResponseWriter, request *http.Request) {
	//验证账户，并把用户账号写入cookie，一段时间
	//写入一个新的随机码，用于后续的验证操作，保证线上只有一个用户
	//设置Cookie的时间有限，保证在一段时间之后，用户自动下线
	user := request.PostFormValue("accountOrName")
	password := request.PostFormValue("password")

	u, err2 := sqlOperate.Login(user, password)
	if err2 != nil {
		fmt.Println(err2)
		//使用cookie去完成信息的存储
		DealWrongCookie(request, writer, err2.Error(), solutionToLogin, wayToLogin)
		http.Redirect(writer, request, "/deal_wrong", 302)
		return
	}
	addCookie("account", u.Account, writer, 1000) //有限时间的链接

	code := MathTool.CreateVerificationCode()
	sqlOperate.UpdateCode(u.Account, code)
	addCookie("code", code, writer, 1000)
	http.Redirect(writer, request, "/Market", 302)
}

func managerMode(writer http.ResponseWriter, request *http.Request) {
	user := request.PostFormValue("accountOrName")
	password := request.PostFormValue("password")

	u, err2 := sqlOperate.ManagerLogin(user, password)

	if err2 != nil {
		fmt.Println(err2)
		//使用cookie去完成信息的存储
		DealWrongCookie(request, writer, err2.Error(), solutionToLogin, wayToLogin)
		http.Redirect(writer, request, "/deal_wrong", 302)
		return
	}

	code := MathTool.CreateVerificationCode()
	sqlOperate.UpdateCode2(u.Account, code)
	addCookie("code", code, writer, 1500)
	addCookie("account", u.Account, writer, 1500) //有限时间的链接

	http.Redirect(writer, request, "/Manager_Notice", 302) //跳转管理界面
}

//检验用户状态的有效性
func checkUser(writer http.ResponseWriter, request *http.Request) bool {
	if !isCookieExit(request, "account") {
		DealWrongCookie(request, writer, "The time is out.", "Login again.", "/login")
		http.Redirect(writer, request, "/deal_wrong", 302)
	} else if !checkCode(request) {
		t := parseTemplateFiles("layout", "close.navbar", "closeWindow")
		t.Execute(writer, "The password has expired.")
	} else {
		return true
	}
	return false
}

/*func home(writer http.ResponseWriter, request *http.Request) {
	if checkUser(writer, request) {
		datas, _ := sqlOperate.Datas(getCookieValue("account", request))
		generateHTML(writer, datas, "layout", "private.navbar", "home")
	}
}*/

//注册逻辑
func register(w http.ResponseWriter, r *http.Request) {
	t := parseTemplateFiles("login.layout", "public.navbar", "signup")
	t.Execute(w, nil)
}

// POST /signup
// Create the user account
//处理可能出现的错误
func signUpAccount(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	message := "" //错误信息收集器
	if err != nil {
		danger(err, "Cannot parse form.")
		sysW := SystemWrong()
		DealWrongCookie(request, writer, sysW.W, sysW.S, sysW.Wa)
		http.Redirect(writer, request, "/deal_wrong", 302)
	}
	name := request.PostFormValue("name")
	account := request.PostFormValue("account")
	findPassword := request.PostFormValue("findPassword")
	password := request.PostFormValue("password")

	err, message = sqlOperate.AddUsers(name, account, password, findPassword)
	if err != nil {
		fmt.Println(err)
		//使用cookie去完成信息的存储
		DealWrongCookie(request, writer, message, solutionToSignUp, wayToSignUp)
		http.Redirect(writer, request, "/deal_wrong", 302)
		return
	}
	http.Redirect(writer, request, "/successfulSignUp?name="+name, 302)
}

func successfulSignUp(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		danger(err, "Cannot parse form.")
		sysW := SystemWrong()
		DealWrongCookie(request, writer, sysW.W, sysW.S, sysW.Wa)
		http.Redirect(writer, request, "/deal_wrong", 302)
	}
	name := request.Form["name"][0]
	t := parseTemplateFiles("layout", "close.navbar", "closeWindow")
	t.Execute(writer, "Dear "+name+", your account is active.")
}

//处理错误
func dealWrong(writer http.ResponseWriter, request *http.Request) {
	//获取信息
	wrong := sqlOperate.Wrongs{
		W:  getCookieValue(WrongName, request),
		S:  getCookieValue(SolutionName, request),
		Wa: getCookieValue(WayName, request),
	}
	generateHTML(writer, &wrong, "layout", "wrong.navbar", "dealWrong")
}

//加载找回密码的界面
func forget(writer http.ResponseWriter, request *http.Request) {
	t := parseTemplateFiles("login.layout", "public.navbar", "forget")
	t.Execute(writer, nil)
}

//处理找回密码的提交
func dealForget(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		danger(err, "Cannot parse form.")
		sysW := SystemWrong()
		DealWrongCookie(request, writer, sysW.W, sysW.S, sysW.Wa)
		http.Redirect(writer, request, "/deal_wrong", 302)
	}
	user := request.PostFormValue("accountOrName")
	findPassword := request.PostFormValue("findPassword")
	//解析输入
	u, err2 := sqlOperate.FindBackPassword(user, findPassword)
	if err2 != nil {
		fmt.Println(err2)
		//使用cookie去完成信息的存储
		DealWrongCookie(request, writer, err2.Error(), solutionToFindPassword, wayToFindback)
		http.Redirect(writer, request, "/deal_wrong", 302)
		return
	}
	//添加一个Cookie的内容，保存的时间有限。
	addCookie("name", u.Name, writer, 5)
	addCookie("account", u.Account, writer, 60) //为后续使用
	addCookie("sfp", u.Sfp, writer, 5)
	http.Redirect(writer, request, "/show_information", 302)
}

func showInformation(writer http.ResponseWriter, request *http.Request) {
	user := sqlOperate.User{
		Name:     getCookieValue("name", request),
		Password: "",
		Account:  getCookieValue("account", request),
		Sfp:      getCookieValue("sfp", request),
	}
	//t := parseTemplateFiles("login.layout", "public.navbar", "dealWrong")
	generateHTML(writer, &user, "login.layout", "public.navbar", "showInformationOfUsers")
}

//更新用户密码
func updatePassword(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		danger(err, "Cannot parse form.")
		sysW := SystemWrong()
		DealWrongCookie(request, writer, sysW.W, sysW.S, sysW.Wa)
		http.Redirect(writer, request, "/deal_wrong", 302)
	}
	account := getCookieValue("account", request)
	password := request.PostFormValue("password")

	sqlOperate.ResetPassword(account, password)
	sqlOperate.UpdateCode(account, MathTool.CreateVerificationCode())
	t := parseTemplateFiles("layout", "close.navbar", "closeWindow")
	t.Execute(writer, changePassword)
	deleteCookie(writer, "account")
}

func checkCode(req *http.Request) bool {
	code := getCookieValue("code", req)
	account := getCookieValue("account", req)
	return sqlOperate.CheckCode(account, code)
}

//绘图逻辑
func draw(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		danger(err, "Cannot parse form.")
		sysW := SystemWrong()
		DealWrongCookie(request, writer, sysW.W, sysW.S, sysW.Wa)
		http.Redirect(writer, request, "/deal_wrong", 302)
	}
	//安全检查
	if !checkCode(request) {
		t := parseTemplateFiles("layout", "close.navbar", "closeWindow")
		t.Execute(writer, "The password has expired.")
		return
	}
	id := request.Form["id"]
	name := request.Form["name"]
	tag := request.Form["tag"]
	t := template.New("tmpl.html")
	switch tag[0] {
	case "Heat Map Graph":
		t, _ = t.Parse(makeTemplate.MakeHeatMapTemplate(sqlOperate.GetValue(id[0], name[0])))
	case "2D Graph":
		t, _ = t.Parse(makeTemplate.MakeLineTemplate(sqlOperate.GetValue(id[0], name[0])))
	case "2D Scatter Diagram":
		t, _ = t.Parse(makeTemplate.MakeScatterTemplate(sqlOperate.GetValue(id[0], name[0])))
	case "2D Polar Graph":
		t, _ = t.Parse(makeTemplate.MakePlorGraph(sqlOperate.GetValue(id[0], name[0])))
	}
	t.Execute(writer, nil)
}

func danger(args ...interface{}) {
	logger.SetPrefix("ERROR ")
	logger.Println(args...)
}

func test(writer http.ResponseWriter, request *http.Request) {
	t := template.New("tmpl.html")
	t, _ = t.Parse(makeTemplate.MakeTest())
	t.Execute(writer, nil)
	return
}
