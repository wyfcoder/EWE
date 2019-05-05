package RayRun

import (
	"github.com/WebForEME/AMethod/Compile"
	"github.com/WebForEME/Functions/RayRun/CountELine"
	"github.com/WebForEME/Functions/RayRun/RayRunDataStruct"
	"sync"
)

//需要对指令进行计算，一类指令为环境，另一类为射线,返回值 环境组合 射线组合
func RayRunCount(instructs []Compile.Instruct) ([]RayRunDataStruct.RayRunData, []RayRunDataStruct.RayRunData,error) {
	eLines := []RayRunDataStruct.RayRunData{} //电子浓度数据集
	rLines := []RayRunDataStruct.RayRunData{} //射线数据集

	eLine := RayRunDataStruct.RayRunData{}
	hasE := false

	for i := 0; i < len(instructs); i++ {
		instruct := instructs[i]
		char := instruct.Head[0]
		//判断是否为电子密度曲线
		if char == 'E' {
			err:=countELine(instruct, &eLine)
			if err!= nil{
				return nil,nil,err
			}
			eLines = append(eLines, eLine)
			if !hasE {
				hasE = true
			}
		} else {
			if !hasE {
				addNullELine(&eLine)
				eLines = append(eLines, eLine)
				hasE = true
			}
			end := findRayCount(&instructs, i)
			group := sync.WaitGroup{}
			group.Add(end - i) //设计出等待队列
			s := addRLine(&rLines, end-i)
			for s < len(rLines) {
				go countRLine(instructs[i], &eLine, &(rLines[s]), &group)
				s++
				i++
			}
			group.Wait()
			i--
		}
	}
	return eLines, rLines,nil
}

//计算电子密度随高度变化曲线
func countELine(instruct Compile.Instruct, eLine *RayRunDataStruct.RayRunData) error{
	mode :=CommandSet[instruct.Head]
	switch mode {
	case E_IRI:
		err:=CountELine.CountEIRI(instruct,eLine)
		return err
	case E_PRO:
		CountELine.CountEPRO(instruct,eLine)
		break
	}
	return nil
}

//计算射线轨迹
func countRLine(instruct Compile.Instruct, eLine *RayRunDataStruct.RayRunData, rLine *RayRunDataStruct.RayRunData, group *sync.WaitGroup) {
	defer group.Done()
}



// 找出射线的条数搜素范围内所有的射线，并为其安排空间,在相同的电子环境下计算出结果
func findRayCount(instructs *[]Compile.Instruct, start int) int {

	for ; start < len(*instructs); start++ {

		char := (*instructs)[start].Head[0]
		if char == 'E' {
			break
		}
	}

	return start
}

//添加空白射线
func addRLine(rLine *[]RayRunDataStruct.RayRunData, count int) int {
	start := len(*rLine)
	line := RayRunDataStruct.RayRunData{}
	for i := 0; i < count; i++ {
		(*rLine) = append((*rLine), line)
	}
	return start
}

//添加一个默认的模型
func addNullELine(eLine *RayRunDataStruct.RayRunData){
	instruct:= Compile.Instruct{}
	instruct.Head="E_IRI"
	countELine(instruct,eLine)
}
