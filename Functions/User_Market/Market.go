package User_Market

import (
	"github.com/WebForEME/AMethod/TimeTool"
	"github.com/WebForEME/Functions"
	"github.com/WebForEME/Functions/CheckService"
	"github.com/WebForEME/Functions/DealWrongs"
	Market2 "github.com/WebForEME/makeTemplate/Market"
	"github.com/WebForEME/sqlOperate"
	"html/template"
	"net/http"
)


//打开集市功能
func Market(writer http.ResponseWriter, r *http.Request) {
	if(CheckService.CheckUser(writer,r)){
	t := template.New("tmpl.html")
	t, _ = t.Parse(Market2.MakeMarketTemplate())
	t.Execute(writer, nil)
	}
}

func FeedbackDeal(writer http.ResponseWriter, request *http.Request){


	if(CheckService.CheckUser(writer,request)) {

		err := request.ParseForm()

		if err != nil {
			Functions.Danger(err, "Cannot parse form.")
			DealWrongs.DealWrongs(DealWrongs.ErrorCodeForSystem,&writer,request)
		}
		message := request.PostFormValue("content") //首先获取消息的值
		account :=Functions.GetCookieValue("accountOfUser",request)
		time    :=TimeTool.NowString()
		err =sqlOperate.InsertIntoFeedbackTable(account,message,time)      //存入数据库

		if err!= nil{
			DealWrongs.DealWrongs(DealWrongs.ErrorCodeForSystem,&writer,request)
		}else{
			successFeedback(writer,request)
		}
		                                               //显示成功操作
	}
}

func successFeedback(writer http.ResponseWriter, request *http.Request){
	err := request.ParseForm()

	if err != nil {
		Functions.Danger(err, "Cannot parse form.")
		DealWrongs.DealWrongs(DealWrongs.ErrorCodeForSystem,&writer,request)
	}

	t := Functions.ParseTemplateFiles("Tools/layout", "Tools/public.navbar", "Success/Success")

	t.Execute(writer, "Your feedback is committed.")
}
