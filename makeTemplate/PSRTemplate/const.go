package PSRTemplate

const(
	HtmlHead = `<!DOCTYPE html>
<html>
<head>
    <link href="/static/css/bootstrap.min.css" rel="stylesheet">
    <link href="/static/css/font-awesome.min.css" rel="stylesheet">
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
    <meta name="viewport" content="initial-scale=1.0, user-scalable=no" />
    <script type="text/javascript" src="//api.map.baidu.com/api?v=2.0&ak=znos8mGRdphs0uYwfsffj4NOLofXIMmT"></script>
    <script type="text/javascript" src="//api.map.baidu.com/library/Heatmap/2.0/src/Heatmap_min.js"></script>
    <title>PSR</title>
    <style type="text/css">
        ul,li{list-style: none;margin:0;padding:0;float:left;}
        html{height:100%}
        body{height:100%;margin:0px;padding:0px}
        #edit{margin-left:5%;margin-right:10%}
        #container{height: 60%;width: 94%;margin: 3%}
    </style>
</head>
<body>

<div class="navbar navbar-default navbar-static-top" role="navigation" style="background: white">
    <div class="container">
        <div class="navbar-header">
            <button type="button" class="navbar-toggle collapsed" data-toggle="collapse" data-target=".navbar-collapse">
                <span class="sr-only">Toggle navigation</span>
                <span class="icon-bar"></span>
                <span class="icon-bar"></span>
                <span class="icon-bar"></span>
            </button>
            <a class="navbar-brand" >
                <i class="fa fa-rss"></i>
                PSR Heat Map
            </a>
        </div>
    </div>
</div>



<div id="container"></div>

<HR style="FILTER:progid:DXImageTransform.Microsoft.Shadow(color:#987cb9,direction:145,strength:15)" width="80%" color=#987cb9 SIZE=1>

<div id="edit">
    <p class="lead" ><a>Data List </a><button class="btn btn-default" onclick="return openHeatmap()" style="float: right">Draw</button></p>
`

	HtmlMapPre = `
</div>
<script type="text/javascript">
    var map = new BMap.Map("container");          // 创建地图实例
    map.enableScrollWheelZoom(); // 允许滚轮缩放
`
	htmlMapScriptPre =`
    if(!isSupportCanvas()){
        alert('热力图目前只支持有canvas支持的浏览器,您所使用的浏览器不能使用热力图功能~')
    }
    heatmapOverlay = new BMapLib.HeatmapOverlay({"radius":20});
    map.addOverlay(heatmapOverlay);`

	htmlMapScriptNext =`
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
`
	HtmlTail =`<script src="/static/js/bootstrap.min.js"></script>
</body>
</html>`

)
