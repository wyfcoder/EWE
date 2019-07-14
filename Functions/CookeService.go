package Functions

import "net/http"
//TODO

//cookie 服务 获取用户的账号

//获取 userAccount
func GetUserAccountCookieValue(request *http.Request) string{
	return GetCookieValue(UserAccountCookieName,request)
}

//设置UserAccount 的值
func SetUserAccountCookieValue(request *http.Request,value string) {

}

//获取 UserCode

func GetUserCodeCookieValue(request *http.Request) string{
	return GetCookieValue(UserCodeCookieName,request)
}
func SetUserCodeValue(request *http.Request,value string)  {
	
}

//获取 ManagerAccount
func GetManagerAccountCookieValue(request *http.Request)string{
	return  GetCookieValue(ManagerAccountCookieName,request)
}
func SetManagerAccountCookieValue(request *http.Request,value string){

}

//获取 ManagerCode
func GetManagerCodeCookieValue(request *http.Request) string{
	return GetCookieValue(ManagerCodeCookieName,request)
}
func SetManagerCodeCoolieValue(request *http.Request,value string){

}



