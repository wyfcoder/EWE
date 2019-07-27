package Plot

import (
	"github.com/WebForEME/AMethod/HtmlTool"
	"github.com/WebForEME/makeTemplate"
	"github.com/WebForEME/sqlOperate"
)

const Vector2DScriptPre  = `
        <script src="https://code.highcharts.com.cn/highcharts/highcharts.js"></script>
        <script src="https://code.highcharts.com.cn/highcharts/modules/exporting.js"></script>
        <script src="https://code.highcharts.com.cn/highcharts/modules/oldie.js"></script>
        <script src="https://code.highcharts.com.cn/highcharts/modules/vector.js"></script>
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
	    title: {
		    text: title
	    },
 xAxis: {
            reversed: false,
            title: {
                enabled: true,
                text: x
            },
            labels: {
                formatter: function () {
                    return this.value;
                }
            },
            maxPadding: 0.05,
            showLastLabel: true
        },
        yAxis: {
            title: {
                text: y
            },
            labels: {
                formatter: function () {
                    return this.value;
                }
            },
            lineWidth: 2
        },
        legend: {
				layout: 'vertical',
				align: 'right',
				verticalAlign: 'middle'
		},
	    series: [{
`

const Vector2DScriptTail  = `}]
})
}
</script>`

func VectorPlot(data *[]sqlOperate.DataPlot) string{
	return Head + Navbar + BodyPre + makeTemplate.MakeSelectComponent(data)+ BodyCenter + Vector2DScriptPre + getDataForVector(data) + Vector2DScriptTail + BodyTail
}

func getDataForVector(plot *[]sqlOperate.DataPlot) string{
	html:=""
	for i:=0 ;i<len(*plot) ;i++{
		if i!=0{
			html+="{"
		}
		html+="name: '"+(*plot)[i].GetLineName()+"',"
		html+="type: 'vector',"
		html+="data :["
		html+=HtmlTool.DrawData4D((*plot)[i].GetLineData())
		html+="]"
		if i!=len(*plot)-1{
			html+="},"
		}
	}
	return html
}