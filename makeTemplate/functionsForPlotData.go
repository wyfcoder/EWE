package makeTemplate

import "github.com/WebForEME/AMethod/HtmlTool"

//提供把 [] float 转化为 string 并保存模板的服务 三维数据
//适配：highCharts
func DataDraw3D(plot *[]DataPlot) string{
	html:=""
	for i:=0 ;i<len(*plot) ;i++{
		if i!=0{
			html+="{"
		}
		html+="name: '"+(*plot)[i].GetLineName()+"',"
		html+="data :["
		html+=HtmlTool.DrawData3D((*plot)[i].GetLineData())
		html+="]"
		if i!=len(*plot)-1{
			html+="},"
		}
	}
	return html
}

//提供把 [] float 转化为 string 并保存模板的服务 二维数据
//适配：highCharts
func DataDraw2D(plot *[]DataPlot) string{
	html:=""
	for i:=0 ;i<len(*plot) ;i++{
		if i!=0{
			html+="{"
		}
		html+="name: '"+(*plot)[i].GetLineName()+"',"
		html+="data :["
		html+=HtmlTool.DrawData((*plot)[i].GetLineData())
		html+="]"
		if i!=len(*plot)-1{
			html+="},"
		}
	}
	return html
}