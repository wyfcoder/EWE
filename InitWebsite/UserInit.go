package InitWebsite

import (
	"github.com/WebForEME/Functions"
	"github.com/WebForEME/Functions/AboutLab"
	"github.com/WebForEME/Functions/Draw"
	"github.com/WebForEME/Functions/EProgram"
	"github.com/WebForEME/Functions/PSR"
	"github.com/WebForEME/Functions/RayRun"
	"github.com/WebForEME/Functions/User_File"
	"github.com/WebForEME/Functions/User_Market"
	"github.com/WebForEME/Functions/User_Plot"
	"net/http"
)

//用户端的一组访问初始化

//对Market涉及的功能进行初始化 wyf
func InitMarket(mux *http.ServeMux){
	//主页
	mux.HandleFunc("/Market", User_Market.Market)
	//涉及反馈的一块
	mux.HandleFunc("/feedbackDeal",User_Market.FeedbackDeal)
}

//对下载模块的初始化 wyf
func InitDownload(mux *http.ServeMux){
	//主页 下载逻辑
	mux.HandleFunc("/download", User_File.Download)

	mux.HandleFunc("/downloadFile", Functions.Download)
}

//对lab模块进行初始化
func InitLab(mux *http.ServeMux){
	mux.HandleFunc("/Lab", AboutLab.Lab)
}


func InitEProgram(mux *http.ServeMux){

	mux.HandleFunc("/EPrograms",EProgram.EProgram)

	//1 绘图模块    TODO 修改使其模块化 draw
	mux.HandleFunc("/PlotCenter", User_Plot.PlotCenter)
	//绘图中心 可以进行文件管理 绘图选择
	mux.HandleFunc("/PlotCheckUploadFile", User_Plot.CheckUploadFile)
	//处理绘制的图像上传功能
	mux.HandleFunc("/PlotDeleteData", User_Plot.DeleteFile)
	//删除数据

	//mux.HandleFunc("/PlotCheck",Draw.DrawCheck) //TODO 需要解决
	//处理绘制图像的合法性

	mux.HandleFunc("/Plot",Draw.DrawCharts) //TODO 困难之处
	//绘制图像

	//2 RayRun 小程序,射线追踪 TODO 还有一些函数可以实现
	mux.HandleFunc("/RayRun", RayRun.RayRun)
	mux.HandleFunc("/RayRunDeal", RayRun.RayRunDeal)
	mux.HandleFunc("/DocumentForRayRunScript", RayRun.RayRunHelpScript)
	mux.HandleFunc("/DocumentForRayRunIRI", RayRun.RayRunHelpIRI)
	mux.HandleFunc("/DocumentForRayRunPRO", RayRun.RayRunHelpPRO)
	mux.HandleFunc("/DocumentForRayRunAH0", RayRun.RayRunHelpAH0)

	//3 PSR   查询绘制数据程序 TODO 有待实现
	mux.HandleFunc("/PSR",PSR.PSR)
}

