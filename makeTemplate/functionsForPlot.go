package makeTemplate

import (
	"github.com/WebForEME/sqlOperate"
)

//绘图组件颜色选择组件
func MakeSelectComponent(data *[]sqlOperate.DataPlot) string{
	selectScript := `<select id="selectLine">`
	for i:=0 ; i < len(*data) ;i++{
		selectScript += `<option>`+(*data)[i].GetLineName()+`</option>`
	}
	selectScript +=`</select>`
	return selectScript
}