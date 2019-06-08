package User_Market

import (
	"github.com/WebForEME/Functions/CheckService"
	"github.com/WebForEME/makeTemplate"
	"html/template"
	"net/http"
)

//打开集市功能
func Market(w http.ResponseWriter, r *http.Request) {
	if(CheckService.CheckUser(w,r)){
	t := template.New("tmpl.html")
	t, _ = t.Parse(makeTemplate.MakeMarketTemplate())
	t.Execute(w, nil)
	}
}
