package InitWebsite

import (
	"github.com/WebForEME/Functions/DealWrongs"
	"net/http"
)

//纠错系统的初始化
func InitDealWrong(mux *http.ServeMux){
	mux.HandleFunc("/deal_wrong", DealWrongs.SolveWrongs)
}


