package RayRun

import (
	"errors"
	"github.com/WebForEME/AMethod/Compile"
	"github.com/WebForEME/AMethod/TextDeal"
	"github.com/WebForEME/AMethod/TimeTool"
	"math"
	"time"
)

//射线函数编译，输入的是一个instructs，进行检查，出错的话输出错误信息
//函数集合规定

const (
	WRONG = 0
	E_IRI = 1
	E_PRO = 2
	R_AH0 = 3
	R_AH1 = 4

	//错误信息
	CommandNotExit    = "Command  is not exit . :< "
	CommandParamWrong = "Command's param is wrong. :< "
	//固定参数
	AM = "AM"
	PM = "PM"
	//正确信息
	CommandSuccess = "Your code is right :>"
	WEBWRONG       ="Sorry , our web service is broken :<"
)

//CommandSet   命令字典
var CommandSet map[string]int = map[string]int{
	"Wrong": WRONG,
	"E_IRI": E_IRI,
	"E_PRO": E_PRO,
	"R_AH0": R_AH0,
	"R_AH1": R_AH1,
}

//检查编译的错误
// 第一种 指令不存在
// 第二种 参数有误，确定好参数的形式
func RayRunCompile(instructs *[]Compile.Instruct) (error, string) {

	err := errors.New("CW")

	for i := 0; i < len(*instructs); i++ {
		head := (*instructs)[i].Head
		key := CommandSet[head]
		isOk := false
		if key == 0 {
			return err, CommandNotExit + string('\n') + "At " + head
		}
		switch key {
		case E_IRI:
			isOk = CheckEIRI((*instructs)[i].Body)
			break
		case E_PRO:
			isOk = CheckEPRO((*instructs)[i].Body)
			break
		case R_AH1:
			isOk = CheckRAH((*instructs)[i].Body,instructs,i)
			break
		case R_AH0:
			isOk = CheckRAH((*instructs)[i].Body,instructs,i)
			break
		}
		if !isOk {
			return err, CommandParamWrong + string('\n') + "At " + Compile.ReturnString((*instructs)[i])
		}
	}


	return nil, CommandSuccess
}

// 指令如下:

// E_IRI()
// E_IRI(PM)
// E_IRI(2012-01-12,AM,20,20)
func CheckEIRI(body []string) bool {
	size := len(body)
	switch size {
	case 0:
		return true
	case 1:
		return CheckTime(body[0])
	case 4:
		param := body[0]
		timec, isOk := TimeTool.StringToDate(param)
		if !isOk || timec.Year() < 1959|| timec.Year() > time.Now().Year(){
			return false
		}
		isOk = CheckTime(body[1])
		if !isOk {
			return false
		}
		param = body[2]
		latitude, isOk := TextDeal.StringToFloat(param) //纬度
		if  !isOk || math.Abs(latitude) > 90{
			return false
		}
		param=body[3]
		longitude,isOk:=TextDeal.StringToFloat(param)
		if  !isOk  || math.Abs(longitude) >180{           //经度
			return false
		}
		return true
	}
	return false
}

//E_IRI 指令 函数如下:
// E_PRO()
// E_PRO(2016-01-02)
func CheckEPRO(body []string) bool {
	size := len(body)
	switch size {
	case 0:
		return true
	case 1:
		param := body[0]
		_, isOk := TimeTool.StringToDate(param)
		return isOk
	default:
		return false
	}
}

//R_AH0(f,angle)            f:频率 MHz angle: -90,90)  n不考虑磁场
//R_AH1(f,angle)            f:频率 MHz angle: [0,90)  n不考虑磁场  出两条曲线
func CheckRAH(body []string,instructs *[]Compile.Instruct,position int) bool{
	if len(body) == 2{
		param :=body[0]
		number ,isOk :=TextDeal.StringToFloat(param)
		if !isOk || !CheckF(number){
			return false
		}
		number,isOk =TextDeal.StringToFloat(param)

		if !isOk || !CheckA(number){
			return false
		}
		return true
	}else if len(body)== 3{   //循环队列
	     //分析  (f1-f2,step,angle) (f,step,angle1-angle2)
		step:=body[1]
		number ,isOk :=TextDeal.StringToFloat(step)
		if !CheckStep(number){
			return false
		}
		param :=body[0]
		number ,isOk =TextDeal.StringToFloat(param)
		if isOk{
			start , end , isOk :=CheckContinueA(body)
			if !isOk{
				return false
			}
			MakeListForFA(start,end,step,body[0],instructs,position,false)
		}else{
			start , end , isOk :=CheckContinueF(body)
			if !isOk{
				return false
			}
			MakeListForFA(start,end,step,body[2],instructs,position,true)
		}
		return true
	  }
	return false
}


//检查连续的 f
func CheckContinueF(body []string) (string,string ,bool){
	//对应 (10-20,1,100)
	a ,isOk :=TextDeal.StringToFloat(body[2])
	if !isOk || !CheckA(a){
		return "","",false
	}
	params,isOk :=Compile.GetParams(body[0])
	if !isOk  || len(params) !=2 {
		return "","",false
	}
	if !CheckParams(params){
		return "","",false
	}
	return params[0],params[1],true
}

//检查连续的 a 对应 (10,1,10-100)
func CheckContinueA(body []string) (string,string ,bool){

	f , isOk :=TextDeal.StringToFloat(body[0])
	if  !isOk || !CheckF(f){
		return  "","",false
	}
	params,isOk :=Compile.GetParams(body[2])

	if !isOk  || len(params) !=2 {
		return "","",false
	}
	if !CheckParams(params){
		return "","",false
	}
	return params[0],params[1],true
}
//检查 F 的合法性
func CheckF(f float64)bool{
	if f <= 0{
		return false
	}
	return true
}
//检查 a 的合法性
func CheckA(a float64)bool{
	if math.Abs(a) >= 90{
		return false
	}
	return true
}
//检查步长
func CheckStep(step float64) bool{
	return CheckF(step)
}
//检查PM 与 AM
func CheckTime(time string) bool {

	if time != AM && time != PM {
		return false
	}
	return true
}

//检查连续的变量 保证 a>b
func CheckParams(list []string) bool{
	start,isOk :=TextDeal.StringToFloat(list[0])
	if !isOk{
		return false
	}
	end ,isOk :=TextDeal.StringToFloat(list[1])
	if !isOk{
		return false
	}
	if start >= end{
		return false
	}
	return true
}

//制作f 或 a 队列
func MakeListForFA(start string,end string,step string,con string,instructs *[]Compile.Instruct,i int,isF bool){
	numberStart,_:=TextDeal.StringToFloat(start)
	numberEnd,_:=TextDeal.StringToFloat(end)
	numberStep,_:=TextDeal.StringToFloat(step)
	instructsNew :=[]Compile.Instruct{}
	head :=(*instructs)[i].Head
	for numberStart < numberEnd{
		a:=TextDeal.FloatToString(numberStart)
		if isF {
			instructsNew=append(instructsNew,MakeInstruct(a,con,head))
		}else {
			instructsNew = append(instructsNew,MakeInstruct(con,a,head))
			}
		numberStart += numberStep
	}
	if isF {
		instructsNew=append(instructsNew,MakeInstruct(end,con,head))
	}else{
		instructsNew=append(instructsNew,MakeInstruct(con,end,head))
	}
	Compile.Combine(instructs,&instructsNew,i)
}

func MakeInstruct(f string ,a string,head string )Compile.Instruct{
	instruct :=Compile.Instruct{}
	instruct.Head=head
	instruct.Body=append(instruct.Body, "")
	instruct.Body=append(instruct.Body, "")
	instruct.Body[0]=f
	instruct.Body[1]=a
	return instruct
}
