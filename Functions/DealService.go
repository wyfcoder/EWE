package Functions

import (
	"fmt"
	"html/template"
	"net/http"
)

//处理html内容
func ParseTemplateFiles(filenames ...string) (t *template.Template) {
	var files []string
	t = template.New("layout")
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("template/%s.html", file))
	}
	t = template.Must(t.ParseFiles(files...))
	return
}

func DealContent(writer http.ResponseWriter,request *http.Request) bool{
	err := request.ParseForm()
	if err != nil {
		Danger(err, "Cannot parse form.")
		sysW := SystemWrong()
		DealWrongCookie(request, &writer, sysW.W, sysW.S, sysW.Wa)
		http.Redirect(writer, request, "/deal_wrong", 302)
		return false
	}
	return true
}

func Danger(args ...interface{}) {
	logger.SetPrefix("ERROR ")
	logger.Println(args...)
}

func GenerateHTML(writer http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("template/%s.html", file))
	}
	templates := template.Must(template.ParseFiles(files...))

	templates.ExecuteTemplate(writer, "layout", data)
}

//获取文件的后缀名

func GetSuffixName(name string) string{
	suffixName :=""
	for i:=len(name)-1;i>0;i--{
		if name[i]=='.'{
			break
		}
		suffixName+=string(name[i])
	}
	return suffixName
}