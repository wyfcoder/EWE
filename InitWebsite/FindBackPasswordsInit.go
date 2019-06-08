package InitWebsite

import (
	"github.com/WebForEME/Functions/FindBackPasswords"
	"net/http"
)

func InitFindBackPasswords(mux *http.ServeMux){

	mux.HandleFunc("/forget", FindBackPasswords.Forget) //主页

	mux.HandleFunc("/deal_forget", FindBackPasswords.DealForget) //解决

	mux.HandleFunc("/show_information", FindBackPasswords.ShowInformation) //成功解决

	mux.HandleFunc("/updatePassword", FindBackPasswords.UpdatePassword) //修改密码

}
