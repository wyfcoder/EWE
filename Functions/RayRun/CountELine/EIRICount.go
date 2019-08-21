package CountELine

import (
	"github.com/WebForEME/AMethod/Compile"
	"github.com/WebForEME/AMethod/TextDeal"
	"github.com/WebForEME/AMethod/TimeTool"
	"github.com/WebForEME/AMethod/WebTool"
	"github.com/WebForEME/Functions/RayRun/RayRunDataStruct"
	"github.com/WebForEME/sqlOperate/programDB/RayRun"
	time2 "time"
)


//计算 E_IRI指令 统一调用 一个指令 要么从数据库里读数据 要么在Web上直接访问其数据
// E_IRI()
// E_IRI(PM)
// E_IRI(2012-01-12,AM,20,20)
func CountEIRI(instruct Compile.Instruct, eLine *RayRunDataStruct.RayRunData) error{
	if len(instruct.Body) != 4{                     //调用一个更新程序 调用数据库文件 解析即可得到数据
	time := AM
		       if len(instruct.Body) == 0{
		       	now:=time2.Now()
		       	if now.Hour() >= 8 || now.Hour()<= 18{
		       		time=PM  //做一个判断 给出时间点
				}
			   }else{
			   	 if instruct.Body[0] == "PM"{
			   	 	time =PM
				 }
			   }
		   eLine.Data=GetDatabaseData(time)
	       eLine.Name=IRIName + "Auto:"+time
	       return nil
	}else{
	     err:=GetWebData(instruct,& eLine.Data)
		 eLine.Name =IRIName + instruct.Body[0]+" "+instruct.Body[1]+" ("+instruct.Body[2]+","+instruct.Body[3]+")"
		 return err
	}
}

// 一次完整的访问记录
func CombineWebData(year string , month string ,day string , time string, latitude string , longitude string) (string,error){
	dataText:=""
	webside:=""
	text:=""
	text,err := PostWebData(year,month,day,time,latitude,longitude,"95","595")
	if err!= nil{
		return "",err
	}
	//需要提取网址 获取IRI的更直接的地址
	webside=WebTool.GetWebside(IRIDIRWEBSIDE,&text)

	text,err= WebTool.GetWebGet(webside)
	if err!= nil{
		return "",err
	}
	dataText +=text
	dataText += string('\n')


	text ,err=PostWebData(year,month,day,time,latitude,longitude,"595.5","1000")
	if err !=nil{
		return "" ,err
	}
	webside=WebTool.GetWebside(IRIDIRWEBSIDE,&text)

	text,err= WebTool.GetWebGet(webside)

	if err!= nil{
		return "",err
	}
	dataText +=text

	return dataText,err
}

//访问接口 返回一个数据文件
//输入 年 月 日 一天的时间点  开始点 截至点 做一次 POST
func PostWebData(year string , month string ,day string , time string, latitude string , longitude string ,start string, end string )(string,error){
	post :=IRIPOSTHead+IRIPOSTYEAR+year+IRIPOSTMONTH+month+IRIPOSTDAY+day+IRIPOSTHOUR+time+IRIPOSTLATITUDE+latitude+IRIPOSTLONGITUDE+longitude+IRIPOSTSTART+start+IRIPOSTSTOP+end
	return WebTool.GetWebPost(IRIWEBSIDE,post)
}

func GetWebData(instruct Compile.Instruct,data *[]float64)error{
    date :=TimeTool.TextToTime(instruct.Body[0])
    time :=AM  //0
    if instruct.Body[1] == "PM"{
    	time=PM
	}
	dataText,err := CombineWebData(date.Year,date.Month,date.Day,time,instruct.Body[2],instruct.Body[3])
	if err!= nil {
		return err
	}
	_, _, dataText = TextDeal.DealText(&dataText)
	TextDeal.DealText2(&dataText, data)
	return nil
	}

//数据库操作
func GetDatabaseData(time string) []float64{
	//给用户回馈数据里的信息
	year,month,day,txt := RayRun.GetIRIData(time)
	_,_,txt=TextDeal.DealText(&txt)
	data :=[]float64{}
	TextDeal.DealText2(&txt,&data)
	if !TimeTool.IsSameTime(year,month,day){
		go UpdateDatabase()
	}
	return data
}

func UpdateDatabase(){
	//获取PM AM 的信息 执行四次的POST
	timeString :=TimeTool.TimeStringNow()
	am,err :=CombineWebData(timeString.Year,timeString.Month,timeString.Day,AM,XiAnLatitude,XiAnLongitude)
	if err!=nil{
		return
	}
	pm,err :=CombineWebData(timeString.Year,timeString.Month,timeString.Day,PM,XiAnLatitude,XiAnLongitude)
	RayRun.UpDateIRIData(am,pm)
}
