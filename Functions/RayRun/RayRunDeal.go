package RayRun

import (
	"fmt"
	"github.com/WebForEME/Functions"
	"github.com/WebForEME/makeTemplate"
	"html/template"
	"net/http"
)

//对输入的文本进行处理,翻译
func RayRunDeal(writer http.ResponseWriter, request *http.Request) {

	//获取文本
	if !Functions.DealContent(writer, request) {
		return
	}

	text := request.PostFormValue("content")

	fmt.Print(text)

	for i := 0; i < 80000000; i++ {
		for j := 0; j < 100; j++ {

		}
	}
	t := template.New("tmpl.html")
	t, _ = t.Parse(makeTemplate.MakeRayRun())
	t.Execute(writer, nil)
}
