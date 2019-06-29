package Draw

import (
	"github.com/WebForEME/Functions"
	"github.com/WebForEME/Functions/CheckService"
	"github.com/WebForEME/Functions/DealWrongs"
	"github.com/WebForEME/makeTemplate"
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
		Functions.Danger(err, "Cannot parse form.")
		DealWrongs.DealWrongs(ERRORCODEFORSYSTEM,&writer,request)
	}

	//安全检查
	if CheckService.CheckUCode(request){
		plot(writer,request)
	}
}

//对应的绘图类型 TODO : 绘图模块混乱 需要修改
func plot(writer http.ResponseWriter,request *http.Request){
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
