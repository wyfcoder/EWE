package Draw

//一般曲线的绘制

import (
	"github.com/WebForEME/AMethod/TextDeal"
	"github.com/WebForEME/Functions"
	"github.com/WebForEME/Functions/CheckService"
	"github.com/WebForEME/Functions/DealWrongs"
	"github.com/WebForEME/makeTemplate/Plot"
	"github.com/WebForEME/sqlOperate"
	"html/template"
	"net/http"
)

const (
	ERRORCODEFORSYSTEM=-1;

)

//绘图逻辑
func DrawCharts(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		DealWrongs.DealWrongs(ERRORCODEFORSYSTEM,&writer,request)
	}

	//安全检查
	if CheckService.CheckUser(writer,request){
		plot(writer,request)
	}
}

//对应的绘图类型 TODO : 绘图模块混乱 需要修改
func plot(writer http.ResponseWriter,request *http.Request){
	listString    :=request.FormValue("list")
	if listString ==  ""{
		DealWrongs.DealWrongs(DealWrongs.ErrorCodeForFileNull,&writer,request)
	}
	//解析 0 1 2 ... 获取选择列表
	list := []int{}
	TextDeal.DecodeListString(listString,&list)

	//如果为空，表示用户没有选择
	if len(list) == 0{
		DealWrongs.DealWrongs(DealWrongs.ErrorCodeForFileNull,&writer,request)
		return
	}
	//调用数据库进行数据比对，并且生成一个数据表格，可用于绘制曲线
	data := []sqlOperate.DataPlot{}
	err  := sqlOperate.DataTables(Functions.GetUserAccountCookieValue(request),&data)
	if err != nil{
		DealWrongs.DealWrongs(DealWrongs.ErrorCodeForSystem,&writer,request)
		return
	}
	//开始比对tag，如果不同就报错
	dataDraw := []sqlOperate.DataPlot{}
	tag := data[list[0]].Tag

	for i:=0;i<len(list);i++{
		if tag != data[list[i]].Tag {
			DealWrongs.DealWrongs(DealWrongs.ErrorCodeForDrawTag,&writer,request)
			return
		}
		dataDraw=append(dataDraw,data[list[i]])
	}
	t :=plotMain(&dataDraw)

	t.Execute(writer, nil)
}

//绘图主要逻辑实现 使用动态模板去实现界面和界面的绘制逻辑
func plotMain(data *[]sqlOperate.DataPlot) *template.Template{
	t := template.New("line.html")
	tag :=(*data)[0].Tag
	switch tag {

	case Curve2D:
		t,_=t.Parse(Plot.Curve2D(data))
	case Polar2D:

	case Vector2D:

	}
	return t
}