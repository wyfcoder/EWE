package makeTemplate

import "strconv"

func MakePlorGraph(datas []float64) string {
	m:=`<html>
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">

    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=9">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>EWE</title>

    <script src="/static/js/jquery.min.js" type="text/javascript"></script>
    <script src="/static/js/colpick.js" type="text/javascript"></script>
    <link rel="stylesheet" href="static/css/colpick.css" type="text/css"/>
    <link href="/static/css/bootstrap.min.css" rel="stylesheet">
    <link href="/static/css/font-awesome.min.css" rel="stylesheet">
    <meta charset="utf-8"><link rel="icon" href="https://static.jianshukeji.com/highcharts/images/favicon.ico">
    <meta name="viewport" content="width=device-width, initial-scale=1">
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
<body>
<div class="navbar navbar-default navbar-static-top" role="navigation">
    <div class="container">
      <div class="navbar-header">
        <button type="button" class="navbar-toggle collapsed" data-toggle="collapse" data-target=".navbar-collapse">
          <span class="sr-only">Toggle navigation</span>
          <span class="icon-bar"></span>
          <span class="icon-bar"></span>
          <span class="icon-bar"></span>
        </button>
        <a class="navbar-brand">
          <i class="fa fa-signal"></i>
          Ewe
        </a>
      </div>
      <div class="navbar-collapse collapse">

        <ul class="nav navbar-nav">
          <li><a href="/market">Market</a></li>
        </ul>

        <ul class="nav navbar-nav">
          <li><a href="/home">MyData</a></li>
        </ul>

        <ul class="nav navbar-nav">
          <li><a href="/download">Download</a></li>
        </ul>

        <ul class="nav navbar-nav">
          <li><a href="/upload">Upload Data</a></li>
        </ul>

        <ul class="nav navbar-nav">
          <li><a>About Lab</a></li>
        </ul>

        <ul class="nav navbar-nav">
          <li><a>Contact Us</a></li>
        </ul>

      </div>
    </div>
  </div>
<p class="lead">Graph</p>
<div id="container"></div>
<p class="lead">Edit Graph</p>
<div id="edit">
<br/>
    <input class="form-control" placeholder="title" id="title">
<br/>
    <select class="form-control" name="selectType" id="fen">
        <option>4</option>
        <option>8</option>
        <option>16</option>
        <option>32</option>
        <option>64</option>
    </select>
<br/>
    <select class="form-control" name="selectType" id="type">
        <option>line</option>
        <option>area</option>
    </select>
<br/>
    <input type="text"  id="picker"  placeholder="select color"></input>
<br/>
<br/>
    <button class="btn btn-primary btn-sm" id="button1" onclick="redraw()">change</button>
</div>
<script src="/static/js/bootstrap.min.js"></script>
<script src="https://img.highcharts.com.cn/highcharts/highcharts.js"></script>
<script src="https://img.highcharts.com.cn/highcharts/highcharts-more.js"></script>
<script src="https://img.highcharts.com.cn/highcharts/modules/exporting.js"></script>
<script src="https://img.hcharts.cn/highcharts-plugins/highcharts-zh_CN.js"></script>
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

	for i:=0;i<len(datas);i+=2{
		x:=strconv.FormatFloat(datas[i], 'f', -1, 64)
		y:=strconv.FormatFloat(datas[i+1], 'f', -1, 64)
		m+="["+x+","+y+"]"
		if(i!=len(datas)-2){
			m+=","
		}else{
			m+="]"
		}
	}
	m+=`;
    drawP("Plor Graph","line",8);
function redraw() {
    let color =document.getElementById("picker");
    let name  =document.getElementById("title");
    let type  =document.getElementById("type");
    let fen   =document.getElementById("fen")
Highcharts.setOptions({
		colors: ["#"+color.value+"72"]
	});
    chart=drawP( name.value.toString(),type.value.toString(),fen.value);
}

function drawP(title,type,fen){
return  Highcharts.chart('container', {
		chart: {
				polar: true
		},
		title: {
				text: title
		},
		pane: {
				startAngle: 0,
				endAngle: 360
		},
		xAxis: {
				tickInterval: 360/fen,
				min: 0,
				max: 360,
				labels: {
						formatter: function () {
								return this.value + '°';
						}
				}
		},
		yAxis: {
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
		series: [
           {
				type: type,
                name: type,
				data: m2
		}]
});
}
</script>
</body>
</html>`
	return m
}
