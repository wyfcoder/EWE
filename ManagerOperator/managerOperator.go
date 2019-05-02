package ManagerOperator

import (
	"github.com/WebForEME/sqlOperate"
	"io/ioutil"
	"mime/multipart"
	"os"
)

//检查标签是否一致
func CheckUploadTag(tag string,fileName string) bool {
	selectTag := ""
	for i := len(fileName) - 1; i >= 0; i-- {
		if (fileName[i] == '.') {
			break
		}
		selectTag += string(fileName[i])
	}
	//比较
	switch tag {

	case ".text":
		if (selectTag != Text) {
			return false
		}
	case ".ppt":
		if (selectTag!=PPT) {
			return false
		}
	case ".pptx":{
		if(selectTag!=PPTX){
			return false
		}
	}
	case ".xlsx":{
		if(selectTag!=XLSX){
			return false
		}
	}
	case ".docx":
		if (selectTag!=DOCX) {
			return false
		}
	case ".doc":{
		if(selectTag!=DOC){
			return false
		}
	}
	case ".dat":
		if (selectTag != DAT) {
			return false
		}
	case ".pdf":
		if (selectTag != PDF) {
			return false
		}
	case ".png":
		if (selectTag != PNG) {
			return false
		}
	case ".rar":
		if(selectTag !=Rar){
			return false
		}
	case ".apk":
		if(selectTag!=Apk){
			return false
		}
	case ".exe":{
		if(selectTag!= EXE){
			return false
		}
	}
	case ".xls":{
		if(selectTag!=XLS){
			return false
		}
	}

	}
	return true
}

func isExit(list []string,value string) bool{
	for i:=0;i<len(list);i++{
		if list[i]==value{
		   return true
	}
	}
	return false
}

//上传文件到指定的目录
func UploadFile(fileHeader *multipart.FileHeader,tag string ,title string){

	file, err := fileHeader.Open()

	road :=Dir+fileHeader.Filename
	file2, err := os.OpenFile(road, os.O_RDWR|os.O_CREATE, 0766)

	if err == nil {
		data,_:= ioutil.ReadAll(file)
		file2.Write(data)
	}

	file2.Close()

	saveInformationToDB(tag,title,road)
}

//将文件信息存入数据库
func saveInformationToDB(tag string,title string,road string){
	sqlOperate.AddFileM(title,tag,road)
}