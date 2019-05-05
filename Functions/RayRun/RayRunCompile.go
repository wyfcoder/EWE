package RayRun

import (
	"errors"
	"github.com/WebForEME/AMethod/Compile"
	"github.com/WebForEME/AMethod/TextDeal"
	"math"
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
func RayRunCompile(instructs []Compile.Instruct) (error, string) {

	err := errors.New("CW")

	for i := 0; i < len(instructs); i++ {
		head := instructs[i].Head
		key := CommandSet[head]
		isOk := false
		if key == 0 {
			return err, CommandNotExit + string('\n') + "At " + head
		}
		switch key {
		case E_IRI:
			isOk = CheckEIRI(instructs[i].Body)
			break
		case E_PRO:
			isOk = CheckEPRO(instructs[i].Body)
			break
		case R_AH1:
			isOk = CheckRAH(instructs[i].Body)
			break
		case R_AH0:
			isOk = CheckRAH(instructs[i].Body)
			break
		}
		if !isOk {
			return err, CommandParamWrong + string('\n') + "At " + Compile.ReturnString(instructs[i])
		}
	}

	return nil, CommandSuccess
}

// 指令如下:

// E_IRI()
// E_IRI(PM)
// E_IRI(2012-01-12,AM,20,20)
func CheckEIRI(body []string) bool {
	isOk :=false
	size := len(body)
	switch size {
	case 0:
		return true
	case 1:
		return CheckTime(body[0])
	case 4:
		param := body[0]
		_, isOk = TextDeal.StringToDate(param)
		if !isOk {
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
		_, isOk := TextDeal.StringToDate(param)
		return isOk
	default:
		return false
	}
}

//R_AH0(f,angle)            f:频率 MHz angle: -90,90)  n不考虑磁场
//R_AH1(f,angle)            f:频率 MHz angle: [0,90)  n不考虑磁场  出两条曲线
func CheckRAH(body []string) bool{

	if len(body) == 2{
		param :=body[0]

		number ,isOk :=TextDeal.StringToFloat(param)

		if !isOk || number <= 0{
			return false
		}

		number,isOk =TextDeal.StringToFloat(param)

		if !isOk || number >= 90 || number <= -90{
			return false
		}

		return true
	}

	return false
}

func CheckTime(time string) bool {

	if time != AM && time != PM {
		return false
	}
	return true
}
