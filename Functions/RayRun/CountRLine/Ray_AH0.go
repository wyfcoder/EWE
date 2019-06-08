package CountRLine

import (
	"fmt"
	"github.com/WebForEME/AMethod/Compile"
	"github.com/WebForEME/AMethod/Physics"
	"github.com/WebForEME/AMethod/TextDeal"
	"github.com/WebForEME/Functions/RayRun/RayRunDataStruct"
	"math"
)

//不考虑磁场的轨迹
//直接计算
func CountAH0(instruct Compile.Instruct, eLine *RayRunDataStruct.RayRunData, rLine *RayRunDataStruct.RayRunData ){

	f,_:=TextDeal.StringToFloat(instruct.Body[0])
	angle,_:=TextDeal.StringToFloat(instruct.Body[1])
	h:=0.0  //高度
	x:=0.0  //距离
	n2:=0.0  //折射率
	theta:=angleToTheta(angle)
	sinTheta:=math.Sin(theta)
	isOut :=false
	//射线第一段 在大气中自由传播

	for h< START{
	x=math.Tan(theta)*h;
	(*rLine).Data=append((*rLine).Data, x)
	(*rLine).Data=append((*rLine).Data, h)
	h+=STEP
	}
	fmt.Println(x)
	i:=1

	for h< END{
		n2=neToN2((*eLine).Data[i],f)
		if n2 <= 0 {  //消散情况
			isOut=true
			break
		}
		if n2 < sinTheta*sinTheta { //遇到反射点
 			isOut=true
			break
		}
		x+=1.0/math.Sqrt(n2-sinTheta*sinTheta)*STEP*sinTheta
		(*rLine).Data=append((*rLine).Data, x)
		(*rLine).Data=append((*rLine).Data, h)
		i+=2
		h+=STEP
	}
	if isOut{
		//对称法直接处理
		size :=len((*rLine).Data)
		x=(*rLine).Data[size-2]
		for i:=size-3 ; i>=0 ;i-=2{
			x+=(*rLine).Data[i+1]-(*rLine).Data[i-1]
			(*rLine).Data=append((*rLine).Data,x)
			(*rLine).Data=append((*rLine).Data,(*rLine).Data[i])
		}
	}

	rLine.Name=eLine.Name +": "+instruct.Body[0]+"MHz "+instruct.Body[1]+"度"
}

//电子密度转化为n
func neToN2(ne float64,f float64) float64{
	C:=Physics.E*Physics.E/(Physics.ME*Physics.PEM)/(math.Pi*math.Pi*4)
	fp2:=ne*C
	f2:=f*1E6*f*1E6
	return 1-fp2/f2
}

func angleToTheta(angle float64)float64  {
	return math.Pi*angle/180.0
}


