package Functions

import (
	"github.com/WebForEME/sqlOperate"
	"net/http"
)

//制作特定的Cookie


//PreWrong

const PreWrongTime  =1000
const PreWrong  ="preWrong"
//PreWrong pre错误存储 用于进一步的处理
func MakePreWrongCookie(message string, write *http.ResponseWriter){
	AddCookie(PreWrong,message,write,PreWrongTime)
}

//获取 PreWrong
func GetPreWrongCookie(req *http.Request) string{
	return GetCookieValue(PreWrong,req)
}
//删除 PreWrongCookie
func DeletePreWrongCookie(write *http.ResponseWriter)  {
	DeleteCookie(*write,PreWrong)
}

//WrongShow


//获取存储信息 - User
const InformationTime  = 500
const UserName  ="nameOfUser"
const UserAccount  ="accountOfUser"
const UserSFP  ="sfpOfUser"
func MakeUserInformationCookie(writer http.ResponseWriter,user sqlOperate.User) {
	AddCookie(UserName, user.Name, &writer, InformationTime)
	AddCookie(UserAccount, user.Account,&writer, InformationTime)
	AddCookie(UserSFP, user.Sfp,&writer, InformationTime)
}

func GetUserCookie(request *http.Request) sqlOperate.User{
	return  sqlOperate.User{
		Name:     GetCookieValue(UserName, request),
		Password: "",
		Account:  GetCookieValue(UserAccount, request),
		Sfp:      GetCookieValue(UserSFP, request),
	}
}

func DeleteUserCookie(writer http.ResponseWriter){
	DeleteCookie(writer,UserName)
	DeleteCookie(writer,UserAccount)
	DeleteCookie(writer,UserSFP)
}


//Make Cookie for Manager
const ManagerCode  = "codeOfManager"
const MangerAccount  = "accountOfManager"
const ManagerTime  = 10000
func MakeManagerCookie(code string,account string,writer *http.ResponseWriter){
	AddCookie(ManagerCode, code, writer, ManagerTime)
	AddCookie(MangerAccount, account, writer, ManagerTime) //有限时间的链接
}

//Make Cookie for user
const UserCode  ="codeOfUser"
const UserTime  = 10000

func MakeUserCookie(code string,account string,writer *http.ResponseWriter){
	AddCookie(UserCode, code, writer, UserTime)
	AddCookie(UserAccount, account, writer, UserTime) //有限时间的链接
}



