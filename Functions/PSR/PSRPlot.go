package PSR

import (
	"github.com/WebForEME/AMethod/TextDeal"
	"github.com/WebForEME/Functions/DealWrongs"
	"github.com/WebForEME/makeTemplate/PSRTemplate"
	PSRDB "github.com/WebForEME/sqlOperate/programDB/PSR"
	"github.com/WebForEME/sqlOperate/programDB/PSR/PSRData"
	"html/template"
	"net/http"
)

func PSRPlot(writer http.ResponseWriter, request *http.Request) {

	//获取请求列表
	listString    :=request.FormValue("list")
	if listString ==  ""{
		DealWrongs.DealWrongs(DealWrongs.ErrorCodeForFileNull,&writer,request)
	}

	//解析 1|2019-12-11|test2   ... 获取选择列表
	list := []string{}
	TextDeal.DecodeListString2s(listString,&list)

	//如果为空，表示用户没有选择
	if len(list) == 0{
		DealWrongs.DealWrongs(DealWrongs.ErrorCodeForFileNull,&writer,request)
		return
	}
	//从数据库里查询数据并保持
	data := [] PSRData.PSRDataR{}

	PSRDB.SetValue(list,&data)

	t := template.New("tmpl.html")

	t, _ = t.Parse(PSRTemplate.MakePSRTemplate(&data))

	t.Execute(writer, nil)

	/*t := Functions.ParseTemplateFiles("test")
	t.Execute(writer,nil)*/
}
