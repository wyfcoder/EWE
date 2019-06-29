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

	html += Body_Github

	html += Body_Support

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
			datas += `] 
    var   size=datas.length;
    showInformation();
`
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
    <script type="text/javascript" src="http://api.map.baidu.com/api?v=2.0&ak=znos8mGRdphs0uYwfsffj4NOLofXIMmT"></script>
    <script type="text/javascript" src="http://api.map.baidu.com/library/Heatmap/2.0/src/Heatmap_min.js"></script>


    <link rel="stylesheet" href="http://jrain.oscitas.netdna-cdn.com/tutorial/css/fontawesome-all.min.css">
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/owl-carousel/1.3.3/owl.carousel.min.css"/>

    <link href="/static/css/bootstrap.min.css" rel="stylesheet">
    <link href="/static/css/font-awesome.min.css" rel="stylesheet">
    <style type="text/css">
        html{height:100%}
        body{height:100%;padding:0px;font-family:"微软雅黑";}
        #news{height:35%;width:80%;margin-left:2%;margin-right:2%}
        #plot{height:70%;width:80%;margin-left:2%;margin-left:2%;margin-top:5%}
        #container2{height:100%;width:100%;margin-left:2%;margin:5%}
        .next{margin-left:40%}
    </style>
   

    <style type="text/css">
            .text{width:80%;height:80%;margin_left:10%}
            #feedBack{width:100%; height:30%;margin-left:2%;margin-right:2%}
            #support{width:100%; height:30%;margin-left:2%;margin-right:2%}
            #sourceCode{width:100%;height:10%;margin-left:15%}
            #left{float:left;}
            #left_button{float:left;margin-left:2%}
     </style>

      <style type="text/css">
            .img-responsive {
                display: inline-block;
                height: auto;
                max-width: 100%;
            }
            #news{width:100%; height:50%;margin-left:2%;margin-right:2%}

            .post-slide{
                background: #fff;
                margin: 0 15px;
                padding:27px 30px;
                border-bottom: 1px solid #dedde1;
            }
            .post-slide .post-header{
                margin-bottom: 10px;
            }
            .post-slide .subtitle{
                color:#0b99bc;
                font-size:14px;
                display: inline-block;
                margin-bottom:5px;
                text-transform: uppercase;
                transition: all 0.4s ease 0s;
            }
            .post-slide .subtitle:hover{
                color:#333;
                text-decoration: none;
            }
            .post-slide .post-title{
                margin: 0;
            }
            .post-slide .post-title a{
                color:#333;
                font-size:18px;
                font-weight: bold;
                text-transform:capitalize;
                transition: all 0.4s ease 0s;
            }
            .post-slide .post-title a:hover{
                color:#0b99bc;
                text-decoration: none;
            }
            .post-slide .pic{
                overflow: hidden;
                position: relative;
            }
            .post-slide .pic img{
                width: 100%;
                height: auto;
                transform: rotate(0deg) scale(1,1);
                transition: all 0.9s ease 0s;
            }
            .post-slide:hover img{
                transform: rotate(-2deg) scale(1.1,1.1);
            }
            .post-slide .pic:after{
                content: "";
                position: absolute;
                top:0;
                left:0;
                width: 100%;
                height: 100%;
                background: rgba(255,255,255,0);
                transition: all 0.9s ease 0s;

            }
            .post-slide:hover .pic:after{
                background: rgba(255,255,255,0.2);
            }
            .post-slide .post-bar{
                list-style: none;
                padding:12px 0;
                margin: 0;
            }
            .post-slide .post-bar li{
                display: inline-block;
                margin-right:3px;
                color:#aaa;
            }
            .post-slide .post-bar li:last-child{
                margin-right: 0;
            }
            .post-slide .post-bar li a{
                color:#aaa;
                transition:0.3s ease;
            }
            .post-slide .post-bar li a:hover{
                text-decoration: none;
                color:#0b99bc;
            }
            .post-slide .post-description{
                font-size: 14px;
                line-height: 24px;
                margin-bottom:15px;
                color:#767676;
            }
            .post-slide .read-more{
                color:#0b99bc;
                font-size: 14px;
                font-style: italic;
                text-transform: capitalize;
            }
            .post-slide .read-more:hover{
                color:#333;
                text-decoration:none;
            }
            @media only screen and (max-width: 479px) {
                .post-slide{
                    padding: 15px;
                }
                .post-slide .post-bar li{
                    margin-bottom: 5px;
                }
                .post-slide .post-bar li:last-child{
                    margin-bottom: 0;
                }
            }
        </style>
</head>
<body>`

const Body_News  =`
<div id="news">
    <a><p class="lead" id="title"></p></a>
   <textarea rows="3" class="text"  id="content" style="font-size: 20px" readonly>
</textarea>
<br/>
 <button class="btn btn-default next"onclick="return previous()">previous</button>
 <button class="btn btn-primary" onclick="return next()">next</button>
</div>
<br/>`

const Body_Github  =`
<div id="sourceCode">
<div id="left"><img src="/static/img/Support/github.jpg"></div>
<div id="left_button"><button class="btn btn-success"><a href="https://github.com/wyfcoder/EWE"  target="_blank" style="font-size: 15px;color: #FFFFFF">Source codes for website.</a></button></div>
<div id="left_button"><button class="btn btn-primary"><a href=""  target="_blank"style="font-size: 15px;color: #FFFFFF">Source codes for android.</a></button></div>
</div>
<br/>
`

const Body_Feedback  = `
    <form action="/feedbackDeal" method="post">
        <div id="feedBack">
   <p class="lead" ><a style="font-size: 20px;color: #00bf00">Feedback</a></p>
   <textarea rows="5" class="text" style="font-size: 20px" id="content" name="content" placeholder="   Any suggests for our andorid appliaction, webside or questions can be edited.
   They will make us better.
   If you want to ask us some questions, please leave your contact information.
   Thank you." required></textarea>
        <br/>
        <button class="btn btn-success">Go</button>
        </div>
    </form>
<br/>`

const Body_Support  = `
<div id="support">
   <a><p class="lead">Technical Support</p></a>
<br/>
 <div class="htmleaf-container">
  <div class="demo" style="background: #f7f8fa;">
   <div class="container">
    <div class="row">
     <div class="col-md-12">
      <div id="news-slider" class="owl-carousel">

       <div class="post-slide">
        <div class="post-header">
         <a class=lead>Go</a>
        </div>

        <div class="pic">
         <img src="static/img/Support/Go.jpg" alt=""/>
        </div>

        <ul class="post-bar">
         <li><i class="fa fa-users"></i> <a>Support</a></li>
        </ul>
        <p class="post-description">
         Go is an open source programming language that makes it easy to build simple, reliable, and efficient software.
         <br/>
         Our website is builded by Go.
        </p>
        <a href="https://golang.google.cn/" class="read-more" target="_blank">Official website</a>
       </div>

       <div class="post-slide">
        <div class="post-header">
         <a class="lead">Android</a>
        </div>
        <div class="pic">
         <img src="static/img/Support/android.jpg." alt=""/>
        </div>
        <ul class="post-bar">
         <li><i class="fa fa-users"></i> <a >Support</a></li>
        </ul>
        <p class="post-description">
            Android is a free and open source operating system based on Linux. Mainly used in mobile devices, such as smartphones and tablets, led and developed by Google (Google) and the Open Mobile Alliance.
        </p>
        <a href="https://developer.android.google.cn/" class="read-more" target="_blank">Official website</a>
       </div>

       <div class="post-slide">
        <div class="post-header">
         <a class="lead">Baidu map</a>
        </div>
        <div class="pic">
         <img src="static/img/Support/baidumap.jpg." alt=""/>
        </div>
        <ul class="post-bar">
         <li><i class="fa fa-users"></i> <a >Support</a></li>
        </ul>
        <p class="post-description">
            Baidu Map is a platform for users to provide travel-related services including intelligent route planning, intelligent navigation (driving, walking, cycling), real-time road conditions and so on.
        </p>
        <a href="http://lbsyun.baidu.com/" class="read-more" target="_blank">Official website</a>
       </div>

       <div class="post-slide">
        <div class="post-header">
         <a class="lead">highcharts</a>
        </div>
        <div class="pic">
         <img src="static/img/Support/highcharts.png" alt=""/>
        </div>
        <ul class="post-bar">
         <li><i class="fa fa-users"></i> <a >Support</a></li>
        </ul>
        <p class="post-description">
            Highcharts is a graphics library written in pure JavaScript.
            It can easily add interactive graphics to web sites or web applications, and it is free for personal learning, personal websites and non-commercial use.
        </p>
        <a href="https://www.highcharts.com.cn/" class="read-more" target="_blank">Official website</a>
       </div>
      </div>
     </div>
    </div>
   </div>
  </div>
 </div>
</div>
`

const Body_End =`<script src="/static/js/jquery-2.1.1.min.js"></script>
    <script src="/static/js/bootstrap.min.js"></script>
        <script src="/static/js/jquery-1.11.0.min.js" type="text/javascript"></script>
    <script type="text/javascript" src="https://cdnjs.cloudflare.com/ajax/libs/owl-carousel/1.3.3/owl.carousel.min.js"></script>
    <script>
        $(function(){
            $("#news-slider").owlCarousel({
                items : 3,
                itemsDesktop:[1199,3],
                itemsDesktopSmall:[1000,2],
                itemsMobile : [650,1],
                navigationText:false,
                autoPlay:true
            });
        })
    </script>
</body>
</html>`