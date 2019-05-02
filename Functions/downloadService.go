package Functions

import (
	"fmt"
	"net/http"
)

//下载文件的操作
func Download(writer http.ResponseWriter, request *http.Request){
	if(!DealContent(writer,request)) {
		return
	}

	path := request.Form["path"][0]
	name:=request.Form["name"][0]
	writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", name))
	writer.Header().Add("Content-Type", "application/octet-stream")
	http.ServeFile(writer, request, path)
}