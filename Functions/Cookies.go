package Functions

import (
	"net/http"
)

const WrongName = "WrongInformation"
const SolutionName  ="SolutionInformation"
const WayName  ="WayInformation"

func DealWrongCookie(req *http.Request, writer http.ResponseWriter, w string,s string,way string) {

	if IsCookieExit(req, WrongName) {
		DeleteCookie(writer, WrongName)
		DeleteCookie(writer,SolutionName)
		DeleteCookie(writer,WayName)
	}
	AddCookie(WrongName, w, writer, 25)
	AddCookie(SolutionName,s,writer,25)
	AddCookie(WayName,way,writer,25)
}

func DealInformationCookie(req *http.Request, w http.ResponseWriter, s string) {

}

//检查cookie是否存在
func IsCookieExit(req *http.Request, s string) bool {
	_, err := req.Cookie(s)
	if err == nil {
		return true
	}
	return false
}

//删除cookie
func DeleteCookie(w http.ResponseWriter, s string) {
	cookie := http.Cookie{Name: s, Path: "/", MaxAge: -1}
	http.SetCookie(w, &cookie)
}

//添加cookie,time的单位是秒
func AddCookie(name string, information string, w http.ResponseWriter, time int) {
	cookie := http.Cookie{
		Name:     name,
		Value:    information,
		HttpOnly: true,
		MaxAge:   time,
	}
	http.SetCookie(w, &cookie)
}

//返回指定的cookie的值
func GetCookieValue(name string, req *http.Request) string {
	cookie, _ := req.Cookie(name)
	return cookie.Value
}
