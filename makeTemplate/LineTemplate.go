package makeTemplate

import "github.com/WebForEME/AMethod/HtmlTool"

func MakeLineTemplate(datas []float64) string {
	html := `
<html>
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <meta charset="utf-8">
    <title>EWE</title>

    <script src="/static/js/jquery.min.js" type="text/javascript"></script>
    <script src="/static/js/colpick.js" type="text/javascript"></script>
    <link rel="stylesheet" href="/static/css/colpick.css" type="text/css"/>
    <link href="/static/css/bootstrap.min.css" rel="stylesheet">
    <link href="/static/css/font-awesome.min.css" rel="stylesheet">
    <link rel="icon" href="https://static.jianshukeji.com/highcharts/images/favicon.ico">


    <style type="text/css">
        #container{margin-right:10%;height:70%;width:90%;float: left}
        #edit{margin-right:30%;margin-left:30%}
#picker {
	margin:0;
	padding:0;
	border:0;
	width:150px;
	height:20px;
	border-right:20px solid green;
	line-height:20px;
}
    </style>
  
</head>
<body>`

	html += Navbar

	html += ` 
<p class="lead" >Graph</p>
<div id="container"></div>
<p class="lead">Edit Graph</p>
<div id="edit">
<input class="form-control" placeholder="Name" id="title">
<br/>
<input class="form-control" placeholder="x-label" id="xLabel">
<br/>
<input class="form-control" placeholder="y-label" id="yLabel">
<br/>
<input type="text"  id="picker"  placeholder="select color"></input>
<br/>
<br/>
<button class="btn btn btn-primary btn-small" id="button1" onclick="redraw()">change</button>
</div>

<script src="/static/js/bootstrap.min.js"></script>
<script src="https://img.highcharts.com.cn/highcharts/highcharts.js"></script>
<script src="https://img.highcharts.com.cn/highcharts/modules/exporting.js"></script>
<script src="https://img.highcharts.com.cn/highcharts/modules/oldie.js"></script>
<script src="https://img.highcharts.com.cn/highcharts-plugins/highcharts-zh_CN.js"></script>
<script>$('#picker').colpick({

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
var m2=[`

    html+=HtmlTool.DrawData(&datas)
	html += `];
    drawP("2D Graph","x","y");

function redraw() {
    let color =document.getElementById("picker");
    let name  =document.getElementById("title");
    let xName =document.getElementById("xLabel");
    let yName =document.getElementById("yLabel");
Highcharts.setOptions({
		colors: ["#"+color.value]
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
            enabled: false
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
        series: [{
            name: 'line',
            data:m2 }]
    });
}
</script>
</body>
</html>`
	return html
}
