package makeTemplate

import (
	"strconv"
)

func MakeHeatMapTemplate(heatMap []float64)  string{
	html:=`<!DOCTYPE HTML>
<html>
<head>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">

    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=9">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    
    <title>EWE</title>

    <link href="/static/css/bootstrap.min.css" rel="stylesheet">
    <link href="/static/css/font-awesome.min.css" rel="stylesheet">
    <script type="text/javascript" src="http://api.map.baidu.com/api?v=2.0&ak=znos8mGRdphs0uYwfsffj4NOLofXIMmT"></script>
    <script type="text/javascript" src="http://api.map.baidu.com/library/Heatmap/2.0/src/Heatmap_min.js"></script>

    <style type="text/css">
        ul,li{list-style: none;margin:0;padding:0;float:left;}
        html{height:100%}
        body{height:100%;padding:0px;font-family:"微软雅黑";}
        #container{height:80%;width:100%;}
        #r-result{height:100%;}
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
<div id="container"></div>
<div id="r-result">
<input type="button"  onclick="openHeatmap();" value="Show"/>
<input type="button"  onclick="closeHeatmap();" value="Close"/>
</div>
</div>
<script type="text/javascript">
    var map = new BMap.Map("container");          // 创建地图实例
    map.centerAndZoom(new BMap.Point(116.404, 39.915), 11);
    map.addControl(new BMap.NavigationControl());
    map.addControl(new BMap.ScaleControl());
    map.addControl(new BMap.OverviewMapControl());
    map.addControl(new BMap.MapTypeControl());
    var point = new BMap.Point(`+strconv.FormatFloat(heatMap[0], 'f', -1, 64)+`,`+strconv.FormatFloat(heatMap[1], 'f', -1, 64)+`);`

	html+=`map.centerAndZoom(point, 15);             // 初始化地图，设置中心点坐标和地图级别
    map.enableScrollWheelZoom(); // 允许滚轮缩放
    var points =[`

	for i:=0;i< len(heatMap);i+=3{
		ln:=strconv.FormatFloat(heatMap[i], 'f', -1, 64)
		la:=strconv.FormatFloat(heatMap[i+1], 'f', -1, 64)
		count:=strconv.FormatFloat(heatMap[i+2], 'f', -1, 64)
		html+=`{"lng":`+ln+`,"lat":`+la+`,"count":`+count+`}`
		if(i!=len(heatMap)-3) {
			html+=","
			}else {
				html+="];"
		}
	}
	html+=` if(!isSupportCanvas()){
        alert('热力图目前只支持有canvas支持的浏览器,您所使用的浏览器不能使用热力图功能~')
    }
    //详细的参数,可以查看heatmap.js的文档 https://github.com/pa7/heatmap.js/blob/master/README.md
    //参数说明如下:
    /* visible 热力图是否显示,默认为true
     * opacity 热力的透明度,1-100
     * radius 势力图的每个点的半径大小
     * gradient  {JSON} 热力图的渐变区间 . gradient如下所示
     *	{
            .2:'rgb(0, 255, 255)',
            .5:'rgb(0, 110, 255)',
            .8:'rgb(100, 0, 255)'
        }
        其中 key 表示插值的位置, 0~1.
            value 为颜色值.
     */
    heatmapOverlay = new BMapLib.HeatmapOverlay({"radius":20});
    map.addOverlay(heatmapOverlay);
    heatmapOverlay.setDataSet({data:points,max:100});
    //是否显示热力图
    function openHeatmap(){
        heatmapOverlay.show();
    }
    function closeHeatmap(){
        heatmapOverlay.hide();
    }
    closeHeatmap();
    function setGradient(){
        /*格式如下所示:
       {
             0:'rgb(102, 255, 0)',
             .5:'rgb(255, 170, 0)',
             1:'rgb(255, 0, 0)'
       }*/
        var gradient = {};
        var colors = document.querySelectorAll("input[type='color']");
        colors = [].slice.call(colors,0);
        colors.forEach(function(ele){
            gradient[ele.getAttribute("data-key")] = ele.value;
        });
        heatmapOverlay.setOptions({"gradient":gradient});
    }
    //判断浏览区是否支持canvas
    function isSupportCanvas(){
        var elem = document.createElement('canvas');
        return !!(elem.getContext && elem.getContext('2d'));
    }
</script>
</body>
</html>`
	return html
}