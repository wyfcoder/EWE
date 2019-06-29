package main

import (
	"github.com/WebForEME/Functions/User_Plot"
	"github.com/WebForEME/InitWebsite"
	"net/http"
)

func main() {

	//sqlOperate.InitRayRun()  //运行一次之后就可以注释掉
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir(config.Static))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	//登陆
	InitWebsite.InitLogin(mux)
	//注册
	InitWebsite.InitRegister(mux)
	//纠错系统
	InitWebsite.InitDealWrong(mux)
	//密码找回
	InitWebsite.InitFindBackPasswords(mux)

	//User 集市功能 下载功能 实验室介绍 应用程序模块
	InitWebsite.InitMarket(mux)
	InitWebsite.InitDownload(mux)
	InitWebsite.InitEProgram(mux)
	InitWebsite.InitLab(mux)


	mux.HandleFunc("/Data", User_Plot.Data)
	mux.HandleFunc("/checkUploadFile", User_Plot.CheckUploadFile)

	//Manager 模式登陆,Notice 服务 和 File 服务
	InitWebsite.InitNotice(mux)
	InitWebsite.InitUploadFile(mux);
	InitWebsite.InitFeedBack(mux);

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}

	server.ListenAndServe()
}