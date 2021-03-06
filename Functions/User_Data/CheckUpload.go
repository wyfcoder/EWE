package User_Data

import (
	"github.com/WebForEME/Functions"
	"github.com/WebForEME/Functions/CheckService"
	"github.com/WebForEME/sqlOperate"
	"io/ioutil"
	"net/http"
	"strconv"
)

func CheckUploadFile(writer http.ResponseWriter, request *http.Request) {
	if CheckService.CheckUser(writer, request) {
		err := request.ParseMultipartForm(1024 * 1024 * 12) //读取12M的数据

		if err != nil {
			Functions.Danger(err, "Cannot parse form.")
			sysW := Functions.SystemWrong()
			Functions.DealWrongCookie(request, writer, sysW.W, sysW.S, sysW.Wa)
			http.Redirect(writer, request, "/deal_wrong", 302)
			return
		}

		fileHeader := request.MultipartForm.File["file"][0]
		name := request.PostFormValue("name")
		tag := request.PostFormValue("selectType")
		describe := request.PostFormValue("describe")
		time := request.PostFormValue("date")
		account:=Functions.GetCookieValue("account",request)

		//对数据的类型与数据内容直接匹配
		suffixName :=Functions.GetSuffixName(fileHeader.Filename)

		if(checkSuffixName(suffixName)){

			//把文件内容->    byte[] -> string
			file,_:=fileHeader.Open()

			byte,_:=ioutil.ReadAll(file)

			text :=string(byte)

			if(checkText(tag,&text)){
				//保存数据，并刷新
				i:=sqlOperate.AddDataFile(name,tag,time,&describe,&text,account)
				if(i!=0){
					dataWrong :=Functions.DataUniqueWrong()
					Functions.DealWrongCookie(request, writer, dataWrong.W, dataWrong.S, dataWrong.Wa)
					http.Redirect(writer, request, "/deal_wrong", 302)
					return
				}
				http.Redirect(writer,request,"/Data",302)
				return
			}
		}
	}
	//错误提示
	dataWrong :=Functions.DataWrong()
	Functions.DealWrongCookie(request, writer, dataWrong.W, dataWrong.S, dataWrong.Wa)
	http.Redirect(writer, request, "/deal_wrong", 302)
}

func checkText(tag string, text *string) bool {
	ok := false
	count := 0
	ok, count, (*text) = dealText(text)
	if !ok {
		return false
	}
	if(count==0) {
		return false
	}
	switch tag {
	case "2D Scatter Diagram":
		if count%2 == 0 {
			return true
		}
	case "2D Graph":
		if count%2 == 0 {
			return true
		}
	case "2D Polar Graph":
		if count%2 == 0 {
			return true
		}
	case "Heat Map Graph":
		if count%3 == 0 {
			return true
		}
	}
	return false
}

func dealText(text *string) (bool, int, string) {
	m := 0
	s := []byte{}
	s2 := ""
	datas := ""
	for i := 0; i < len(*text); i++ {
		if (*text)[i] == '\r' {
			if (*text)[i+1] == '\n' { // Windows换行
				i++
				continue
			} else {
				return false, 0, ""
			}
		}else if (*text)[i] != ' ' {
			s = append(s, (*text)[i])
			if i != len(*text)-1 {
				continue
			}
		}
		s2 = string(s)
		datas += s2 + " "
		_, err := strconv.ParseFloat(s2, 64)
		if err == nil {
			m += 1
			s = s[0:0]
		} else {
			return false, 0, ""
		}
	}
	return true, m, datas
}

func checkSuffixName(name string) bool{
	return name=="txt" || name=="tad"
}