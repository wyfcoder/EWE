package InitWebsite

import (
	"github.com/WebForEME/Functions/LoginService"
	"net/http"
)

//登陆初始化 TODO:把登陆的逻辑功能进行模块化
func InitLogin(mux *http.ServeMux) {

	mux.HandleFunc("/login", LoginService.Login)

	mux.HandleFunc("/dealLogin", LoginService.DealLogin)

}
