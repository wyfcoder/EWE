package CountELine

import (
	"fmt"
	"github.com/WebForEME/AMethod/Compile"
	"github.com/WebForEME/AMethod/TextDeal"
	"github.com/WebForEME/AMethod/TimeTool"
	"github.com/WebForEME/Functions/RayRun/RayRunDataStruct"
	"github.com/WebForEME/sqlOperate/programDB/RayRun"
	"math"
	"strconv"
	"time"
)



//计算 E_PRO 指令 接收 指令 数据容器
//把所有的参数编码为 一个列 采用一个初始化函数进行值读取和编号 给每一个数据编号列调用时
//从数据库直接调用
func CountEPRO(instruct Compile.Instruct, eLine *RayRunDataStruct.RayRunData) {
	if len(instruct.Body) == 0 {
		CountEPROOne(eLine, TimeTool.MonthToInt(time.Now().Month()), time.Now().Day())
		eLine.Name=PROJECTNAME+"Auto."
	} else {
		timeN, _ := TimeTool.StringToDate(instruct.Body[0])
		CountEPROOne(eLine, TimeTool.MonthToInt(timeN.Month()), timeN.Day())
		eLine.Name=PROJECTNAME+instruct.Body[0]
	}
}
func CountEPROOne(eLine *RayRunDataStruct.RayRunData, month int, day int) {
	id := GetId(month, day)
	text:=""
	RayRun.SelectFromRayRun(strconv.Itoa(id),&text)
	_,_,text=TextDeal.DealText(&text)
	data :=[]float64{}
	TextDeal.DealText2(&text,&data)
	EPROMode(data[0],data[1],data[2],data[3],data[4],data[5],eLine) //调用
}
//编码时间，获得一个唯一的标识符
func GetId(month int, day int) int {
	month = (month-1) * 3
	if day < 10 {
		month += 1
	} else if day < 20 {
		month += 2
	} else {
		month += 3
	}
	return month
}
//直接主函数 By XiaoMing
func EPROMode( foE float64, foF2 float64, hmF1 float64, hmF2 float64, NmE float64, NmF2 float64, eLine *RayRunDataStruct.RayRunData) {
	hmE:=110.0
	h, eds_e:=0.0,0.0;      //高度，E层电子密度
	ymE := 20.0       //E层的厚度
	ymF2 := hmF2-hmF1;      //F2半厚度，ymF2=hmf2-hmf1
	f23000, eds_f1, eds_f2, eds_up:=0.0,0.0,0.0,0.0;   //M（3000）F2，f1,f2,上电离层电子密度
	X, fj, Sj, r12, dm, hj, HH, fp2:=0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0;
	dh := STEP;      //步长

	fmt.Println(ymE+f23000+dm+fp2)

	//计算E层电子密度分布

	h=START
	for  h<hmE{
		eds_e=(1.0-math.Pow((hmE-h)*0.05,2.0))*NmE
		(*eLine).Data = append((*eLine).Data, h)
		(*eLine).Data = append((*eLine).Data, eds_e)
		h+=dh
	}

	X = foF2 / foE;
	fj = 1.7 * foE;
	Sj = 1.24 * math.Pow(10.0, 10.0) * math.Pow(fj, 2.0);
	r12 = 200;     //太阳月平均黑子数
	f23000 = 3.169;     //M(3000)F2   (需要输入)
	dm = 0.18 / (X - 1.4) + 0.096 * (r12 - 25.0) / 150.0;
	hj = hmF2 - ymF2 * math.Sqrt(1.0 - math.Pow(fj / foF2, 2.0));   //F1层的顶部高度

	//计算F1层电子密度分布
	for h<hj{
		eds_f1 = (Sj - NmE) / (hj - hmE) * h + (NmE * hj - Sj * hmE) / (hj - hmE);    //F1层电子密度计算
		(*eLine).Data = append((*eLine).Data, h)
		(*eLine).Data = append((*eLine).Data, eds_f1)
		h+=dh
	}

	//计算F2电子密度分布
	for  h<hmF2{
		eds_f2 = NmF2 * (1.0 -math.Pow((hmF2 - h) / ymF2, 2.0));      //抛物线模型计算法f2层电子密度
		(*eLine).Data = append((*eLine).Data, h)
		(*eLine).Data = append((*eLine).Data, eds_f2)
		h+=dh
	}


	for h< END {
		HH = 0.85 * math.Pow((1.0 + h / 6378.0), 2.0) * 200;     //一个参数
		eds_up = NmF2 * math.Exp(0.5 * (1.0 - (h - hmF2) / HH - math.Exp((h - hmF2) / HH)));     //卡尔曼模型计算上电离层电子密度
		(*eLine).Data = append((*eLine).Data, h)
		(*eLine).Data = append((*eLine).Data, eds_up)
		h+=dh
	}


}


