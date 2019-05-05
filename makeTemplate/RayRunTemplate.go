package makeTemplate

//射线追踪的模板程序,先架构一个静态的网页
const (
	RayRunHead = `<head>
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

    <link rel="shortcut icon" href="http://simbyone.com/wp-content/uploads/2014/04/3455e6f65c33232a060c28829a49b1cb.png">
    <link href='http://fonts.googleapis.com/css?family=Source+Sans+Pro:200,300,400,600,700,900,200italic,300italic,400italic,600italic,700italic,900italic' rel='stylesheet' type='text/css'>
    <link rel="stylesheet" href="http://maxcdn.bootstrapcdn.com/font-awesome/4.3.0/css/font-awesome.min.css">
    <link href="/static/css/Icomoon/style.css" rel="stylesheet" type="text/css" />
    <script type="text/javascript" src="/static/js/jquery-2.0.2.min.js"></script>

    <style type="text/css">
            #box {
            width: 90%;
            height: 60%;
            margin-left:10%;
        }
        #tab_nav {
            margin: 0;
            padding: 0;
            height: 25px;
            line-height: 24px;
        }
        #tab_nav li {
            float: left;
            margin: 0% 20%;
            list-style: none;
            border: 1px solid #999;
            border-bottom: none;
            height: 24px;
            width: 110px;
            text-align: center;
            background: #FFF;
        }
        a {
            font: bold 14px/24px "微软雅黑", Verdana, Arial,
            Helvetica,sans-serif;
            color: green;
            text-decoration: none;
        }
        a:hover {
            color: red;
        }
        #tab_content {
            width: 80%;
            height: 80%;
            border: 1px solid #999;
            font: bold 4em/273px "微软雅黑", Verdana, Arial,
            Helvetica,sans-serif;
            text-align: center;
            background: #FFF;
            overflow: hidden;
        }
        #t_1,#t_2,#t_3 {
            width:  100%;
            height: 100%;
        }
        #leadText{float: left;margin-left:5%}
        .text{width:50%;height:30%;margin-left:9%;border-width: 1px 1px 1px 5px; border-style: solid; border-color:lightblue}
        .compile{width:25%;height:30%;border-width: 1px 1px 1px 5px; border-style: solid; border-color:black}
        .buttonP{width:20%;height:10%;margin-left:20%;border-radius: 10px 20px 30px 40px;} 
        #loading{
            background-color: #FFFFFF;
            display:none;
            height: 100%;
            width: 100%;
            position: fixed;
            z-index: 1;
            margin-top: 0px;
            top: 0px;
        }
        #loading-center{
            width: 100%;
            height: 100%;
            position: relative;
        }
        #loading-center-absolute {
            position: absolute;
            left: 50%;
            top: 50%;
            height: 50px;
            width: 150px;
            margin-top: -25px;
            margin-left: -75px;

        }
        .object{
            width: 8px;
            height: 50px;
            margin-right:5px;
            background-color: #45b29d;
            -webkit-animation: animate 1s infinite;
            animation: animate 1s infinite;
            float: left;
        }

        .object:last-child {
            margin-right: 0px;
        }

        .object:nth-child(10){
            -webkit-animation-delay: 0.9s;
            animation-delay: 0.9s;
        }
        .object:nth-child(9){
            -webkit-animation-delay: 0.8s;
            animation-delay: 0.8s;
        }
        .object:nth-child(8){
            -webkit-animation-delay: 0.7s;
            animation-delay: 0.7s;
        }
        .object:nth-child(7){
            -webkit-animation-delay: 0.6s;
            animation-delay: 0.6s;
        }
        .object:nth-child(6){
            -webkit-animation-delay: 0.5s;
            animation-delay: 0.5s;
        }
        .object:nth-child(5){
            -webkit-animation-delay: 0.4s;
            animation-delay: 0.4s;
        }
        .object:nth-child(4){
            -webkit-animation-delay: 0.3s;
            animation-delay: 0.3s;
        }
        .object:nth-child(3){
            -webkit-animation-delay: 0.2s;
            animation-delay: 0.2s;
        }
        .object:nth-child(2){
            -webkit-animation-delay: 0.1s;
            animation-delay: 0.1s;
        }



        @-webkit-keyframes animate {

            50% {
                -ms-transform: scaleY(0);
                -webkit-transform: scaleY(0);
                transform: scaleY(0);

            }

        }

        @keyframes animate {
            50% {
                -ms-transform: scaleY(0);
                -webkit-transform: scaleY(0);
                transform: scaleY(0);
            }
        }
    </style>
  

        <script src="https://img.highcharts.com.cn/highcharts/highcharts.js"></script>
        <script src="https://img.highcharts.com.cn/highcharts/modules/exporting.js"></script>
        <script src="https://img.highcharts.com.cn/highcharts/modules/oldie.js"></script>
        <script src="https://img.highcharts.com.cn/highcharts-plugins/highcharts-zh_CN.js"></script>
</head>`

	RayRunBodyPre = `<body>
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
          <li><a href="/RayRun">Console</a></li>
        </ul>

        <ul class="nav navbar-nav">
          <li><a href="/DocumentForRayRun">Document</a></li>
        </ul>

        <ul class="nav navbar-nav">
          <li><a href="/Data">back</a></li>
        </ul>
      </div>
    </div>
  </div>`

	RayRunBodyGraph = `
<br/>
<br/>
<div id="box">
    <ul id="tab_nav">
        <li><a href="#t_1">电子密度变化图</a></li>
        <li><a href="#t_2">射线轨迹图</a></li>
    </ul>

    <div id="tab_content">
        <div id="t_1"></div>
        <div id="t_2"></div>
    </div>
</div>
<br/>
`

	RayRunBodyConsolePre = `
 <form action="/RayRunDeal" method="post">
 <ul id="tab_nav">
        <li>编辑器</li>
        <li>检查 && 输出</li>
 </ul>
  <div id="text">
   <textarea rows="5" class="text"  id="content" name="content" required>`
	RayRunBodyConsoleNex = `</textarea>`
	RayRunBodyCompilePre=`
     <textarea rows="5" class="compile" id="compile" name="compile" readonly>`
	RayRunBodyCompileNext=`
	</textarea>
  </div>
 <button class="btn btn-primary buttonP"  onclick="showLoading()">Go</button>
 </form>
    <div id="loading">
        <div id="loading-center-absolute">
            <div class="object"></div>
            <div class="object"></div>
            <div class="object"></div>
            <div class="object"></div>
            <div class="object"></div>
            <div class="object"></div>
            <div class="object"></div>
            <div class="object"></div>
            <div class="object"></div>
            <div class="object"></div>
        </div>
</div>
`
	RayRunBodyELine = `
<script>
var chart = Highcharts.chart('t_1', {
    chart: {
        type: 'spline',
        inverted: true
    },
    title: {
        text: '电子密度变化曲线'
    },
    xAxis: {
        reversed: false,
        title: {
            enabled: true,
            text: '高度'
        },
        labels: {
            formatter: function () {
                return this.value + 'km';
            }
        },
        maxPadding: 0.05,
        showLastLabel: true
    },
    yAxis: {
        title: {
            text: '距离'
        },
        labels: {
            formatter: function () {
                return this.value + 'km';
            }
        },
        lineWidth: 2
    },
    legend: {
        enabled: false
    },
    tooltip: {
        headerFormat: '<b>{series.name}</b><br/>',
        pointFormat: '{point.x} km: {point.y} km'
    },
    plotOptions: {
        spline: {
            marker: {
                enable: false
            }
        }
    },
    series: [{
`
	RayRunBodyRLine = `
}]
});
</script>
<script>
var chart2 = Highcharts.chart('t_2', {
    chart2: {
        type: 'spline',
        inverted: false
    },
    title: {
        text: '射线轨迹'
    },
    xAxis: {
        reversed: false,
        title: {
            enabled: true,
            text: '距离'
        },
        labels: {
            formatter: function () {
                return this.value + 'km';
            }
        },
        maxPadding: 0.05,
        showLastLabel: true
    },
    yAxis: {
        title: {
            text: '高度'
        },
        labels: {
            formatter: function () {
                return this.value + 'km';
            }
        },
        lineWidth: 2
    },
    legend: {
        enabled: false
    },
    tooltip: {
        headerFormat: '<b>{series.name}</b><br/>',
        pointFormat: '{point.x} km: {point.y}km'
    },
    plotOptions: {
        spline: {
            marker: {
                enable: false
            }
        }
    },
    series: [{
`
	RayRunTail = `
    }]
    });
</script>
    <script src="/static/js/jquery-2.1.1.min.js"></script>
    <script src="/static/js/bootstrap.min.js"></script>
    <script>
function showLoading(){
          document.getElementById("loading").style.display="block"
}
   window.onunload = onunload_handler;
   function onunload_handler() {
        load=document.getElementById ("loading")
        load.style.display="none"
        }     
    </script>

<script type="text/javascript">

function setCss(opt){
var sr=document.getElementById("textarea-1");

var len=sr.value.length;
setSelectionRange(sr,len,len); //将光标定位到文本最后 
}

function setSelectionRange(input, selectionStart, selectionEnd) {
 if (input.setSelectionRange) {  
   input.focus();  
   input.setSelectionRange(selectionStart, selectionEnd);  
 }  
 else if (input.createTextRange) {  
   var range = input.createTextRange();  
   range.collapse(true);  
   range.moveEnd('character', selectionEnd);  
   range.moveStart('character', selectionStart);  
   range.select();  
 }  
}
</script>
</body>`
	RayRunTextMain = "@main:\nE_"
	RayRunTextCompile ="Hello , I' will check your input :)\n"

)

func MakeRayRun() string {
	//第一部分：网页导航栏
	html := RayRunHead

	html += RayRunBodyPre + RayRunBodyGraph + RayRunBodyConsolePre + RayRunTextMain + RayRunBodyConsoleNex+RayRunBodyCompilePre+RayRunTextCompile+RayRunBodyCompileNext+ RayRunBodyELine + RayRunBodyRLine + RayRunTail

	return html
}

//导入信息 拆解出两部分 一部分为 body 一部分为 wrong
func MakeRayRunWrite(text string) string{
	size:=len(text)-1
	i:=size
	k:=0

	for i >= 0{
		if text[i] == '*'{
			break
		}
		i--
	}

	body := ""
	for k < i{
		body += string(text[k])
		k++
	}

	compile := ""

	for  i <= size{
		compile +=  string(text[i])
		i++
	}

	html := RayRunHead
	html += RayRunBodyPre + RayRunBodyGraph + writeEdit(body)+writeCompile(compile)+ RayRunBodyELine + RayRunBodyRLine + RayRunTail

	return html
}

//把 内容写入到 edit 里面
func writeEdit(text string) string{
	return RayRunBodyConsolePre + text +RayRunBodyConsoleNex
}

//把内容写到 compile 里面
func writeCompile(text string) string{
	return  RayRunBodyCompilePre+RayRunTextCompile+text+RayRunBodyCompileNext
}
