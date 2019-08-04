package Plot

import (
	"github.com/WebForEME/AMethod/HtmlTool"
	"github.com/WebForEME/makeTemplate"
	"github.com/WebForEME/sqlOperate"
)

const PolorScriptPre  = ` 
        <script src="https://code.highcharts.com.cn/highcharts/highcharts.js"></script>
        <script src="https://code.highcharts.com.cn/highcharts/highcharts-more.js"></script>
        <script src="https://code.highcharts.com.cn/highcharts/modules/exporting.js"></script>
        <script src="https://img.hcharts.cn/highcharts-plugins/highcharts-zh_CN.js"></script>
        <script src="https://code.highcharts.com.cn/highcharts/modules/series-label.js"></script>
        <script src="https://code.highcharts.com.cn/highcharts/modules/oldie.js"></script>
        <script src="https://code.highcharts.com.cn/highcharts/themes/grid-light.js"></script>
<script>
drawP("Graph","x","y");
function redraw() {
    let color =document.getElementById("picker");
    let name  =document.getElementById("title");
    let xName =document.getElementById("xLabel");
    let yName =document.getElementById("yLabel");
    var selectLine = document.getElementById("selectLine");
    var index = selectLine.selectedIndex;
    let myColors =Highcharts.getOptions().colors;
    myColors[index] = '#'+color.value;
    Highcharts.setOptions({
		colors: myColors
	});
    chart=drawP( name.value.toString(),xName.value.toString(), yName.value.toString());
}
function drawP(title,x,y){
return Highcharts.chart('container', {
        	chart: {
		polar: true
	},
	title: {
		text: title
	},
	xAxis: {
          title: {
                enabled: true,
                text: x
            },
		tickInterval: 45,
		labels: {
			formatter: function () {
				return this.value + '°';
			}
		}
	},
	yAxis: {
         title: {
                text: y
            },
		min: 0
	},
	plotOptions: {
		series: {
			pointStart: 0,
			pointInterval: 45
		},
		column: {
			pointPadding: 0,
			groupPadding: 0
		}
	},
	series: [{`

const PolarScriptTai  =`}]
})
}
</script>`

func Polar2D(data *[]sqlOperate.DataPlot)string{
	return Head + Navbar + BodyPre+makeTemplate.MakeSelectComponent(data) + BodyCenter + PolorScriptPre + getDataForPolar(data)+ PolarScriptTai + BodyTail
}

func getDataForPolar(plot *[]sqlOperate.DataPlot) string{
	html:=""
	for i:=0 ;i<len(*plot) ;i++{
		if i!=0{
			html+="{"
		}
		html+="name: '"+(*plot)[i].GetLineName()+"',"
		html+= "type: 'column',"
		html+="data :["
		html+=HtmlTool.DrawData((*plot)[i].GetLineData())
		html+="]"
		if i!=len(*plot)-1{
			html+="},"
		}
	}
	return html
}