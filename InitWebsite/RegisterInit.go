package InitWebsite

import (
	"github.com/WebForEME/Functions/Register"
	"net/http"
)

func InitRegister(mux *http.ServeMux){
	//注册
	mux.HandleFunc("/register", Register.Register)
	//检查，通过后即可
	mux.HandleFunc("/sign_up_account", Register.SignUpAccount)
	//成功注册
	mux.HandleFunc("/successfulSignUp", Register.SuccessfulSignUp)
}



