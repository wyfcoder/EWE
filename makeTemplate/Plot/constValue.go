package Plot

const Head  =`<html>
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
        #container{margin-left:2%;margin-right:10%;height:70%;width:90%;float: left}
        #edit{margin-left:5%;margin-right:10%}
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

const Navbar  =`<div class="navbar navbar-default navbar-static-top" role="navigation" style="background: white">
    <div class="container">
        <div class="navbar-header">
            <button type="button" class="navbar-toggle collapsed" data-toggle="collapse" data-target=".navbar-collapse">
                <span class="sr-only">Toggle navigation</span>
                <span class="icon-bar"></span>
                <span class="icon-bar"></span>
                <span class="icon-bar"></span>
            </button>
            <a class="navbar-brand" >
                <i class="fa fa-line-chart"></i>
                Plot
            </a>
        </div>
        <div class="navbar-collapse collapse">

            <ul class="nav navbar-nav">
                <li><a href="/PlotCenter"  style="font-size: 20px;color: #00bf00">Console</a></li>
            </ul>

            <ul class="nav navbar-nav">
                <li><a href="/DocumentForPlot"  target="_blank" style="font-size: 20px;color: #0b99bc">Document</a></li>
            </ul>

            <ul class="nav navbar-nav">
                <li><a href="/EPrograms"  style="font-size: 20px;color: #ff76b5">back</a></li>
            </ul>
        </div>
    </div>
</div>`

const Body  =`
<div id="container"></div>
<HR style="FILTER: progid:DXImageTransform.Microsoft.Shadow(color:#987cb9,direction:145,strength:15)" width="80%" color=#987cb9 SIZE=1>
<div id="edit">
<p class="lead">Edit Canvas<button class="btn btn btn-default" onclick="redraw()" style="float: right" >Change</button></p>
<HR style="FILTER: progid:DXImageTransform.Microsoft.Shadow(color:#987cb9,direction:145,strength:15)" width="80%" color=#987cb9 SIZE=1>
<input type="text"  id="picker"  placeholder="select color">
<select>
<option>assss</option>
<option>bsss</option>
</select>
</input>
<br/>
<br/>
<input class="form-control" placeholder="title" id="title">
<br/>
<input class="form-control" placeholder="subTitle" id="subTitle">
<br/>
<input class="form-control" placeholder="x-label" id="xLabel">
<br/>
<input class="form-control" placeholder="y-label" id="yLabel">
<br/>
<br/>
</div>

<script src="https://code.highcharts.com.cn/highcharts/highcharts.js"></script>
        <script src="https://code.highcharts.com.cn/highcharts/modules/exporting.js"></script>
        <script src="https://code.highcharts.com.cn/highcharts/modules/series-label.js"></script>
        <script src="https://code.highcharts.com.cn/highcharts/modules/oldie.js"></script>
        <script src="https://code.highcharts.com.cn/highcharts-plugins/highcharts-zh_CN.js"></script>
        <script src="https://code.highcharts.com.cn/highcharts/themes/grid-light.js"></script>
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
    drawP("Graph","x","y");

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

const  Tail = `
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
</body>
</html>`