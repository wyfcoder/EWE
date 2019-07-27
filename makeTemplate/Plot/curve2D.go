package Plot

import (
	"github.com/WebForEME/AMethod/HtmlTool"
	"github.com/WebForEME/makeTemplate"
	"github.com/WebForEME/sqlOperate"
)

const LineScriptPre  =` 
<script src="https://code.highcharts.com.cn/highcharts/highcharts.js"></script>
<script src="https://code.highcharts.com.cn/highcharts/modules/exporting.js"></script>
<script src="https://code.highcharts.com.cn/highcharts/modules/series-label.js"></script>
<script src="https://code.highcharts.com.cn/highcharts/modules/oldie.js"></script>
<script src="https://code.highcharts.com.cn/highcharts-plugins/highcharts-zh_CN.js"></script>
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
            type: 'spline'
        },
        title: {
            text: title
        },
        subtitle: {
            text: ''
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
        tooltip: {
            headerFormat: '<b>{series.name}</b><br/>',
            pointFormat: '{point.x},{point.y}'
        },
        plotOptions: {
            spline: {
                marker: {
                    enable: false
                }
            }
        },
        series: [{`

const  LineScriptTail = `
}],
		responsive: {
				rules: [{
						condition: {
								maxWidth: 500
						},
						chartOptions: {
								legend: {
										layout: 'horizontal',
										align: 'center',
										verticalAlign: 'bottom'
								}
						}
				}]
		}
})
}
</script>
<script>
$('#picker').colpick({

	layout:'hex',

	submit:0,

	colorScheme:'dark',

	onChange:function(hsb,hex,rgb,el,bySetColor) {

		$(el).css('border-color','#'+hex);

		// Fill the text box just if the color was set using the picker, and not the colpickSetColor function.

		if(!bySetColor) $(el).val(hex);

	}

}).keyup(function(){

	$(this).colpickSetColor(this.value);

});
</script>
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
</script>
`

//把数据转换到script脚本里去
func Curve2D(data *[]sqlOperate.DataPlot) string{
	return Head + Navbar + BodyPre + makeTemplate.MakeSelectComponent(data)+ BodyCenter + LineScriptPre + getDataForLine(data) + LineScriptTail + BodyTail
}

func getDataForLine(plot *[]sqlOperate.DataPlot) string{
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
