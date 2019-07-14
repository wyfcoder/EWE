package User_Plot

import (
	"github.com/WebForEME/Functions"
	"github.com/WebForEME/Functions/CheckService"
	"github.com/WebForEME/Functions/DealWrongs"
	"github.com/WebForEME/sqlOperate"
	"net/http"
)

//绘图中心界面 包括 绘制的数据集 添加 删除数据集等功能
func PlotCenter(writer http.ResponseWriter,request *http.Request){
	if(CheckService.CheckUser(writer,request)){
		t := Functions.ParseTemplateFiles("Tools/layout", "Plot/plotToolBar", "Plot/plotCenter")
		datas, err:= sqlOperate.Datas(Functions.GetUserAccountCookieValue(request))
		if err != nil{
			DealWrongs.DealWrongs(DealWrongs.ErrorCodeForSystem,&writer,request)
			return
		}
		t.Execute(writer, datas)
	}
}
