package User_Market

import (
	"github.com/WebForEME/Functions/CheckService"
	Market2 "github.com/WebForEME/makeTemplate/Market"
	"html/template"
	"net/http"
)

//打开集市功能
func Market(writer http.ResponseWriter, r *http.Request) {
	if(CheckService.CheckUser(writer,r)){
	t := template.New("tmpl.html")
	t, _ = t.Parse(Market2.MakeMarketTemplate())
	t.Execute(writer, nil)
	}
}
