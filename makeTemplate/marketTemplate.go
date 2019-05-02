package makeTemplate

import "github.com/WebForEME/sqlOperate"

func MakeMarketTemplate() string {
	//获得所有的Notices -把其赋值给一个javascript的数组
	//监听按钮，上滑下滑事件
	notices :=sqlOperate.Notices()

	html:=`<!DOCTYPE HTML>
<html>
<head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=9">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>EWE</title>

    <link href="/static/css/bootstrap.min.css" rel="stylesheet">
    <link href="/static/css/font-awesome.min.css" rel="stylesheet">
    <script type="text/javascript" src="http://api.map.baidu.com/api?v=2.0&ak=znos8mGRdphs0uYwfsffj4NOLofXIMmT"></script>
    <script type="text/javascript" src="http://api.map.baidu.com/library/Heatmap/2.0/src/Heatmap_min.js"></script>

    <style type="text/css">
        html{height:100%}
        body{height:100%;padding:0px;font-family:"微软雅黑";}
        #news{height:45%;width:80%;margin-left:2%;margin-right:2%}
        #plot{height:70%;width:80%;margin-left:2%;margin-left:2%;margin-top:5%}
        #container2{height:100%;width:100%;margin-left:2%;margin:5%}
        .text{width:100%;height:70%}
        .next{margin-left:40%}
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
          <li><a href="/Data">Data</a></li>
        </ul>

        <ul class="nav navbar-nav">
          <li><a href="/Market">Market</a></li>
        </ul>

        <ul class="nav navbar-nav">
          <li><a href="/download">Download</a></li>
        </ul>

        <ul class="nav navbar-nav">
          <li><a href="/Lab">About Lab</a></li>
        </ul>

        <ul class="nav navbar-nav">
          <li><a>Contact Us</a></li>
        </ul>

      </div>
    </div>
  </div>

<div id="news">
    <a><p class="lead" id="title"></p></a>
   <textarea rows="3" class="text"  id="content" readonly>
</textarea>
<br/>
 <button class="btn btn-default next"onclick="return previous()">previous</button>
 <button class="btn btn-primary" onclick="return next()">next</button>
</div>
<div id="plot">
<a><p class="lead">Signal Foresee</p></a>
<div id="container2">
</div>
</div>

<script type="text/javascript">
 var map = new BMap.Map("container2");
            map.centerAndZoom(new BMap.Point(116.404, 39.915), 11);
            map.addControl(new BMap.NavigationControl());
            map.addControl(new BMap.ScaleControl());
            map.addControl(new BMap.OverviewMapControl());
            map.addControl(new BMap.MapTypeControl());
            map.centerAndZoom(new BMap.Point(108.840053,34.129522), 15.8);
            map.enableScrollWheelZoom();
    if(!isSupportCanvas()){
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
</script>`

	//数据转化为javascript数组
   datas:=`<script>
   var   position=1
   var   datas=[`
   for i:=0;i<len(notices);i++{
   	datas+=`[`
   	datas+="`"+notices[i].Title+"`"
   	datas+=`,`
   	datas+="`"+notices[i].Date+"`"
   	datas+=`,`
   	datas+="`"+notices[i].Content+"`"
   	datas+=`]`
   	if(len(notices)-1!=i){
   	datas+=`,`
	}else{
		datas+=`] 
    var   size=datas.length;
    showInformation();
`
	}
   }
   html+=datas

   //书写监听函数
   functions :=`
function next(){
if(position==size){
position=1;
}
else{
position=position+1;
}
showInformation();
}

function previous(){
if(position==1){
position=size;
}

else{
position=position-1;
}
showInformation();
}

function showInformation(){
   var t="Notice "+position.toString()+"/"+size.toString()+" "+datas[position-1][1]+" "+datas[position-1][0];
   document.getElementById("title").innerText=t;
   document.getElementById("content").value=datas[position-1][2];
}
</script>`

   html+=functions

    end :=`<script src="/static/js/jquery-2.1.1.min.js"></script>
    <script src="/static/js/bootstrap.min.js"></script>

</body>
</html>`

    html+=end
	return html
}