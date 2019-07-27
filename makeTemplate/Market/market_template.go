package Market

import (
	"github.com/WebForEME/makeTemplate"
	"github.com/WebForEME/sqlOperate"
)

func MakeMarketTemplate() string {

	//前部分
	html := PreBody

	//添加Navbar
	html += makeTemplate.Navbar

	html += Body_News


	html += Body_Feedback

	html += Body_Support

	html += Body_Github

	html += DynamicMode()  //动态的获取通知数据

	html += Body_End

	return html
}

func DynamicMode() string{
	//获得所有的Notices -把其赋值给一个javascript的数组
	//监听按钮，上滑下滑事件
	notices := sqlOperate.Notices()
	//通知数据转化为javascript数组
	datas := `<script>
    var   position=1
    var   datas=[`
	if len(notices) == 0{
		datas += "];"
	} else{
	for i := 0; i < len(notices); i++ {
		datas += `[`
		datas += "`" + notices[i].Title + "`"
		datas += `,`
		datas += "`" + notices[i].Date + "`"
		datas += `,`
		datas += "`" + notices[i].Content + "`"
		datas += `]`
		if len(notices)-1 != i {
			datas += `,`
		} else {
			datas += `]; 
    var   size=datas.length;
    showInformation();
`
		}
	}
	}

	datas +=`function next(){
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
	return datas
}

const PreBody  = `<!DOCTYPE HTML>
<html>
<head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=9">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>EWE</title>

    <link href="/static/css/bootstrap.min.css" rel="stylesheet">
    <link href="/static/css/font-awesome.min.css" rel="stylesheet">


    <link href="/static/css/bootstrap.min.css" rel="stylesheet">
    <link href="/static/css/font-awesome.min.css" rel="stylesheet">
    <style type="text/css">
        html{height:100%}
        body{height:100%;padding:0px;font-family:"微软雅黑";}
        #news{height:35%;width:80%;margin-left:2%;margin-right:2%}
        #plot{height:70%;width:80%;margin-left:2%;margin-left:2%;margin-top:5%}
        #container2{height:100%;width:100%;margin-left:2%;margin:5%}
        .next{margin-left:40%}
        .pre {float:right;margin-right:15%}
    </style>
   

    <style type="text/css">
            .text{ width:90%;height:80%;margin-left:5%;}
            #feedBack{width:100%; height:30%;margin-left:2%;margin-right:2%}
            #support{width:100%; height:10%;margin-left:2%;margin-right:2%;margin-bottom:5%}
            #sourceCode{width:100%;height:10%;margin-left:15%}
            #left{float:left;}
            #left_button{float:left;margin-left:2%}
            #news{width:100%; height:50%;margin-left:2%;}
     </style>

      <style type="text/css">
            .img-responsive {
                display: inline-block;
                height: auto;
                max-width: 100%;
            }
        </style>
</head>
<body>`

const Body_News  =`
<div id="news">
    <a><p class="lead" id="title"></p></a>
   <textarea rows="3" class="text"  id="content" style="font-size: 20px" readonly>
</textarea>
<br/><br/>
 <button class="btn btn-default next"onclick="return previous()">previous</button>
 <button class="btn btn-primary" onclick="return next()">next</button>
</div>
<br/>`

const Body_Github  =`
<div id="sourceCode">
<div id="left"><img src="/static/img/Support/github.jpg"></div>
<div id="left_button"><button class="btn btn-success"><a href="https://github.com/wyfcoder/EWE"  target="_blank" style="font-size: 15px;color: #FFFFFF">Source codes for website.</a></button></div>
<div id="left_button"><button class="btn btn-primary"><a href="https://github.com/ElephantWithCode/MapApplication" target="_blank"style="font-size: 15px;color: #FFFFFF">Source codes for android.</a></button></div>
</div>
<br/>
`

const Body_Feedback  = `
    <form action="/feedbackDeal" method="post">
        <div id="feedBack">
   <p class="lead" style="font-size: 20px;margin-left:7%;">Feedback<button class="btn btn-default pre">Submit</button></p>
   <textarea rows="5" class="text" style="font-size: 20px" id="content" name="content" placeholder="   Any suggests for our andorid appliaction, webside or questions can be edited.
   They will make us better.
   If you want to ask us some questions, please leave your contact information.
   Thank you." required></textarea>
        </div>
    </form>
<br/>`

const Body_Support  = `
<div id="support" style = "margin-left:10%;">
    <p class="lead">Technical Support
    <a class="lead" href="https://golang.google.cn/" target="_blank">Go  </a>
    <a class="lead"  href="https://developer.android.google.cn/" target="_blank">Android  </a>
    <a class="lead" href="http://lbsyun.baidu.com/" target="_blank">Baidu map  </a>
    <a class="lead" href="https://www.highcharts.com.cn/" target="_blank">HighCharts</a>
    </p>
`

const Body_End =`
    </div>
    <script src="/static/js/jquery-2.1.1.min.js"></script>
    <script src="/static/js/bootstrap.min.js"></script>
    <script src="/static/js/jquery-1.11.0.min.js" type="text/javascript"></script>
</body>
</html>`