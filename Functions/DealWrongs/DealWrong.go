package DealWrongs

import (
	"fmt"
	"github.com/WebForEME/Functions"
	"net/http"
)

//处理错误 并返回帮助信息 TODO :构建这个错误处理系统
func DealWrongs(errorCode int,writer *http.ResponseWriter, request *http.Request){
	DealErrorCode(errorCode,writer,request)
	http.Redirect(*writer,request,"/deal_wrong",302)
}

//显示模块
func SolveWrongs(writer http.ResponseWriter, request *http.Request){

	wrong :=Functions.GetCookieValue("ErrorInformation",request)
	solve :=Functions.GetCookieValue("ErrorSolve",request)

	dealWrong := DealWrong{
		"",
		"",
	}
	if   wrong == "" {
		dealWrong.Wrong=NullErrorInformation
		dealWrong.Solve=NullErrorSolve
	}else{
		dealWrong.Wrong=wrong
		dealWrong.Solve=solve
		clearError(writer)
	}

	t := Functions.ParseTemplateFiles("DealWrong/DealWrongLayout", "Tools/public.navbar", "DealWrong/DealWrong")
	err:=t.Execute(writer, dealWrong)
	fmt.Println(err)
}

//处理错误码 ,借助cooke来显示解决方法和错误 统一从cookie中取出错误值
func DealErrorCode(errorCode int, writer *http.ResponseWriter,request * http.Request){
	//获取存储的前错误
	pre_information :=Functions.GetPreWrongCookie(request)

	switch errorCode {
	case ErrorCodeForSystem:    //系统错误
		writerError(pre_information+SystemErrorInformation,SystemErrorSolve,writer)
	case ErrorCodeForLogin:     //登陆错误
		writerError(pre_information+LoginErrorInformation,LoginErrorSolve,writer)
	case ErrorCodeForRegister: //注册错误
		writerError(pre_information+RegisterErrorInformation,RegisterErrorSolve,writer)
	case ErrorCodeForForget:     //找回密码错误
		writerError(pre_information+ForgetErrorInformation,ForgetErrorSolve,writer)
	case ErrorCodeForTimeOut:    //在线超时
		writerError(pre_information+TimeOutErrorInformation,TimeOutErrorSolve,writer)
	case ErrorCodeForLoginAgain: //重复登陆
		writerError(pre_information+LoginAgainInformation,LoginAgainErrorSolve,writer)
	case ErrorCodeForTag: //标签错误
		writerError(pre_information+TagErrorInformation,TagErrorSolve,writer)
	case ErrorCodeForFileName: //文件名重复
		writerError(pre_information+FileNameErrorInformation,FileNameErrorSolve,writer)
	case ErrorCodeForFileSize: //文件大小过大
		writerError(pre_information+FileSizeErrorInformation,FileSizeErrorSolve,writer)
	case ErrorCodeForFileContent:
		writerError(pre_information+FileContentErrorInformation,FileContentErrorSolve,writer)
	case ErrorCodeForFileNumber:
		writerError(pre_information+FileNumberErrorInformation,FileNumberErrorSolve,writer)
	case ErrorCodeForFileNull:
		writerError(pre_information+FileNullErrorInformation,FileNullErrorSolve,writer)
	case ErrorCodeForDrawTag:
		writerError(pre_information+DrawTagErrorInformation,DrawTagErrorSolve,writer)
	}
	Functions.DeletePreWrongCookie(writer)
}

func writerError(information string,solve string,w *http.ResponseWriter){
	Functions.AddCookie("ErrorInformation",information,w,TIMEFORCOOKIE)
	Functions.AddCookie("ErrorSolve",solve,w,TIMEFORCOOKIE)
}

func clearError(w http.ResponseWriter){
	Functions.DeleteCookie(w,"ErrorInformation")
	Functions.DeleteCookie(w,"ErrorSolve")
}