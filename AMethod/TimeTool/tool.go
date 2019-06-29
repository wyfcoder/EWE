package TimeTool

import (
	"strconv"
	"time"
)

//比较时间与现在的时间是否一致
func IsSameTime(year string,month string ,day string) bool{
	timeString :=TimeStringNow()
	if timeString.Year != year || timeString.Month!=month || timeString.Day !=day{
		return false
	}
	return true
}



//string 转为 日期。 接收一个 string 返回 日期 和 是否有误

func StringToDate(text string)  (time.Time,bool){


	timeN,err:= time.Parse("2006-01-02 15:04:05",text+" 00:00:00")
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
func TextToTime(text string) TimeString{

	time , _ :=StringToDate(text)

	timeM :=TimeString{}

	timeM.Day=strconv.Itoa(time.Day())

	timeM.Year=strconv.Itoa(time.Year())

	timeM.Month=strconv.Itoa(MonthToInt(time.Month()))

	return timeM
}

//把现在的时间转换为一个TimeString对象
func TimeStringNow() TimeString{
	now :=time.Now()
	yearN :=strconv.Itoa(now.Year())
	monthN :=strconv.Itoa(MonthToInt(now.Month()))
	dayN :=strconv.Itoa(now.Day())

	timeString :=TimeString{yearN,monthN,dayN}
	return timeString
}

func NowString() string{
	timeString :=time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05")
	return timeString
}

