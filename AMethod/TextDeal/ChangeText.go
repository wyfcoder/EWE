package TextDeal

import (
	"strconv"
	"time"
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

//string 转为 日期。 接收一个 string 返回 日期 和 是否有误

func StringToDate(text string)  (time.Time,bool){
	timeN,err:= time.Parse("2006-02-01",text)
	if err !=nil{
		return timeN,false
	}
	return timeN,true
}


//把月份转换为 int
func MonthToInt(month time.Month) int{
	m:=month.String()
	switch m {
	case "January":return 1
	case "February":return 2
	case "March": return 3
	case "April":return 4
	case "May": return 5
	case "June":return 6
	case "July":return 7
	case "August":return 8
	case "September":return 9
	case "October":return 10
	case "November":return 11
	case "December":return 12
	}
	return -1
}

//把时间转换为 Time Time为内部数据 用于其他接口 默认输入为正确
func TextToTime(text string) Time{

	time , _ :=StringToDate(text)

	timeM :=Time{}

	timeM.Day=strconv.Itoa(time.Day())

	timeM.Year=strconv.Itoa(time.Year())

	timeM.Month=strconv.Itoa(MonthToInt(time.Month()))

	return timeM
}