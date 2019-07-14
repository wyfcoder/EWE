package RayRun

import (
	"github.com/WebForEME/AMethod/Compile"
	"github.com/WebForEME/Functions"
	RayRun2 "github.com/WebForEME/makeTemplate/RayRun"
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

	err,instructs,text:= Compile.Compile(&text) //编译初级文本

	t := template.New("tmpl.html")

	if err!= nil{
		t,_=t.Parse(RayRun2.MakeRayRunWrite(text))
		t.Execute(writer,nil)
		return
	}

	err,text2 :=RayRunCompile(&instructs)       //编译对应程序的指令
	Compile.CombineText(&text,&text2)
	if err!= nil{
		t,_=t.Parse(RayRun2.MakeRayRunWrite(text))
		t.Execute(writer,nil)
		return
	}
	eLines,rLines,err:= RayRunCount(instructs)
	if err !=nil{                                                                  //网址访问问题
	    text +=string('\n')+WEBWRONG
		t,_=t.Parse(RayRun2.MakeRayRunWrite(text))
		t.Execute(writer,nil)
		return
	}

	t, _ = t.Parse(RayRun2.DrawData(&eLines,&rLines,text))
	t.Execute(writer, nil)
}
