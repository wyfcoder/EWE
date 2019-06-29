package InitWebsite



import (
	"github.com/WebForEME/Functions/DeleteService"
	"github.com/WebForEME/Functions/Manager_Feedback"
	"github.com/WebForEME/Functions/Manager_File"
	"github.com/WebForEME/Functions/Manager_Notice"
	"github.com/WebForEME/Functions/Manager_Page"
	"net/http"
)

//管理者的一组 访问

//实现通知模块
func InitNotice(mux *http.ServeMux){

	mux.HandleFunc("/Manager_Notice", Manager_Page.ManagerPage)

	mux.HandleFunc("/Delete_Notice", Manager_Notice.DeleteNotice)

	mux.HandleFunc("/addNotice", Manager_Notice.AddNotice)
}

//实现数据上传模块 TODO:有 BUG
func InitUploadFile(mux *http.ServeMux){

	mux.HandleFunc("/Manager_File", Manager_File.ManagerFile)

	mux.HandleFunc("/uploadFileM", Manager_File.UploadFileM)

	mux.HandleFunc("/deleteMFile", DeleteService.DeleteMFile)

}

//实现查看反馈信息模块 TODO
func InitFeedBack(mux *http.ServeMux){
	mux.HandleFunc("/Manager_Feedback",Manager_Feedback.ManagerFeedback)

	mux.HandleFunc("/deleteFeedback",Manager_Feedback.DeleteFeedback)
}


//实现综合数据查看模块 TODO
func InitComprehensiveControl() {

}
