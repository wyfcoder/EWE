package TextDeal

import (
	"strconv"
)

//完成一些数据转换的函数

//string to float64 接收一个 string 返回一个 float64 和 是否有误


func StringToFloat(text string) (float64,bool){
	 number ,err:=strconv.ParseFloat(text,64)
	 if err!=nil{
	 	return number,false
	 }
	 return number,true
}


//float64 to String
func FloatToString(f float64) string{
	return strconv.FormatFloat(f,'f',-1,64)
}