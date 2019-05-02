package Functions

import "github.com/WebForEME/sqlOperate"

const wrong1 = "/deal_wrong/?information=Wrong at registering an account.The Length of account is not 11.You need to set the length of it is 11.Such as 15050430001."

const wrong2 = "/deal_wrong/?information=Wrong at registering an account.The Length of password is less than 6.Set your password's length longer."

const wrong3 = "/deal_wrong/?information=Wrong at registering an account.The password you first edit is not the same with the password you edit secondly.Make sure they are same."


func SystemWrong()sqlOperate.Wrongs{
	return sqlOperate.Wrongs{
		W:"System wrong.",
		S:"Close window.",
		Wa:"/login",
	}
}

func DataUniqueWrong()  sqlOperate.Wrongs{
	return sqlOperate.Wrongs{
		W:"The name is used.",
		S:"Rename the file you upload.",
		Wa:"/Data",
	}
}

func DataWrong() sqlOperate.Wrongs{
	return sqlOperate.Wrongs{
		W:"Data format is wrong.",
		S:"Check your data file.Upload your file again.",
		Wa:"/Data",
	}
}