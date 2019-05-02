package main

import (
	"github.com/WebForEME/Functions"
	"github.com/WebForEME/Functions/AboutLab"
	"github.com/WebForEME/Functions/ContactUs"
	"github.com/WebForEME/Functions/DeleteService"
	"github.com/WebForEME/Functions/Manager_File"
	"github.com/WebForEME/Functions/Manager_Notice"
	"github.com/WebForEME/Functions/Manager_Page"
	"github.com/WebForEME/Functions/User_Data"
	"github.com/WebForEME/Functions/User_File"
	"github.com/WebForEME/Functions/User_Market"
	"net/http"
)

func main() {
	p("EME", "started at", config.Address)

	mux := http.NewServeMux()
	files := http.FileServer(http.Dir(config.Static))
	mux.Handle("/static/", http.StripPrefix("/static/", files))


	//登陆
	mux.HandleFunc("/login", logIn)

	mux.HandleFunc("/dealLogin",dealLogin)


	mux.HandleFunc("/home",home)

	mux.HandleFunc("/download", User_File.Download)

	mux.HandleFunc("/downloadFile",Functions.Download)

    mux.HandleFunc("/deleteMFile",DeleteService.DeleteMFile)

	//注册
	mux.HandleFunc("/register", register)

	mux.HandleFunc("/sign_up_account", signUpAccount)

	mux.HandleFunc("/successfulSignUp", successfulSignUp)

	mux.HandleFunc("/deal_wrong", dealWrong)

	mux.HandleFunc("/forget", forget)

	mux.HandleFunc("/deal_forget", dealForget)

	mux.HandleFunc("/show_information", showInformation)

	mux.HandleFunc("/updatePassword", updatePassword)

	//绘图
	mux.HandleFunc("/draw",draw)

	//删除文件
	mux.HandleFunc("/deleteData",User_Data.DeleteFile)



	//User    集市功能 数据功能
	mux.HandleFunc("/Market",User_Market.Market)

	mux.HandleFunc("/Data",User_Data.Data)

	mux.HandleFunc("/checkUploadFile",User_Data.CheckUploadFile)

	//User Lab
	mux.HandleFunc("/Lab",AboutLab.Lab)

	//User Contact Us
	mux.HandleFunc("/ContactUs",ContactUs.ContactUs)
	//Manager 模式登陆,Notice 服务 和 File 服务

	mux.HandleFunc("/Manager_Notice",Manager_Page.ManagerPage)

	mux.HandleFunc("/Delete_Notice",Manager_Notice.DeleteNotice)

	mux.HandleFunc("/Manager_File",Manager_File.ManagerFile)

	mux.HandleFunc("/uploadFileM",Manager_File.UploadFileM)

	mux.HandleFunc("/addNotice",Manager_Notice.AddNotice)

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}

	server.ListenAndServe()
}