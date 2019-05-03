package makeTemplate

func MakeTest() string {
	return `
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml"><head>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
    <link rel="shortcut icon" href="http://simbyone.com/wp-content/uploads/2014/04/3455e6f65c33232a060c28829a49b1cb.png">
    <title>Page Preloading Effects</title>

    <link href='http://fonts.googleapis.com/css?family=Source+Sans+Pro:200,300,400,600,700,900,200italic,300italic,400italic,600italic,700italic,900italic' rel='stylesheet' type='text/css'>
    <link rel="stylesheet" href="http://maxcdn.bootstrapcdn.com/font-awesome/4.3.0/css/font-awesome.min.css">
    <link href="/static/css/Icomoon/style.css" rel="stylesheet" type="text/css" />
    <script type="text/javascript" src="/static/js/jquery-2.0.2.min.js"></script>
    <style>
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
</head>
<body>


<div id="loading">
    <div id="loading-center">
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
</div>
</body>
</html>

`
}